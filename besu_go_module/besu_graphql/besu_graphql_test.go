package besu_graphql

import (
	"encoding/hex"
	"fmt"
	"go_module/config"
	"go_module/transaction"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

var cntr_addr = "0x42699A7612A82f1d9C36148af9C77354759b210b"

func TestBesuCall(t *testing.T) {

	abi, err := transaction.AgeInfoStrageMetaData.GetAbi()
	if err != nil {
		t.Error(err)
	}

	data, err := abi.Pack("getAge", "Jo")
	if err != nil {
		t.Error(err)
	}

	rtn, err := BesuCall(nil, Call{
		To:   cntr_addr,
		Data: "0x" + hex.EncodeToString(data),
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(rtn)
}

func TestBesuWrite(t *testing.T) {

	abi, err := transaction.AgeInfoStrageMetaData.GetAbi()
	if err != nil {
		t.Error(err)
	}

	data, err := abi.Pack("setAge", "min", big.NewInt(29))
	if err != nil {
		t.Error(err)
	}

	address := common.HexToAddress(cntr_addr)

	account := config.Config.Accounts["account1"]

	tx, err := transaction.CreateSignedTx(account.PrivateKey[2:], &address, nil, data)
	if err != nil {
		t.Error(err)
	}

	encoded, err := rlp.EncodeToBytes(tx)
	if err != nil {
		t.Error(err)
	}

	rtn, err := BesuWrite("0x" + common.Bytes2Hex(encoded))
	if err != nil {
		t.Error(err)
	}

	t.Log(rtn)
}

func TestMultiBesuCall(t *testing.T) {

	n := 200

	abi, err := transaction.AgeInfoStrageMetaData.GetAbi()
	if err != nil {
		t.Error(err)
	}

	callDatas := make([]Call, n)

	for i := range n {
		data, err := abi.Pack("getAge", fmt.Sprintf("name%d", i))
		if err != nil {
			t.Error(err)
		}
		callDatas[i] = Call{
			To:   cntr_addr,
			Data: "0x" + hex.EncodeToString(data),
		}
	}

	rtn, err := BesuMultiCall(nil, callDatas)
	if err != nil {
		t.Error(err)
	}

	t.Log(rtn)
}

/*
중간 트랜잭션 실패 시,
- 앞 트랜잭션에는 영향 X. 앞 트랜잭션은 정상적으로 반영
- 뒤 트랜잭션은 실행 X. 오류 발생 시, 즉시 오류 반환 후 뒤 트랜잭션 실행 X
*/
func TestMultiBesuWrite(t *testing.T) {

	abi, err := transaction.AgeInfoStrageMetaData.GetAbi()
	if err != nil {
		t.Error(err)
	}

	data1, err := abi.Pack("setAge", "Jo", big.NewInt(50))
	if err != nil {
		t.Error(err)
	}

	address := common.HexToAddress(cntr_addr)

	account := config.Config.Accounts["account1"]

	tx1, err := transaction.CreateSignedTx(account.PrivateKey[2:], &address, nil, data1)
	if err != nil {
		t.Error(err)
	}

	encoded1, err := rlp.EncodeToBytes(tx1)
	if err != nil {
		t.Error(err)
	}

	data2, err := abi.Pack("setAge", "min", big.NewInt(60))
	if err != nil {
		t.Error(err)
	}

	tx2, err := transaction.CreateSignedTx(account.PrivateKey[2:], &address, nil, data2)
	if err != nil {
		t.Error(err)
	}

	encoded2, err := rlp.EncodeToBytes(tx2)
	if err != nil {
		t.Error(err)
	}

	rtn, err := BesuMultiWrite([]string{
		"0x" + common.Bytes2Hex(encoded1),
		"0x" + common.Bytes2Hex(encoded2),
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(rtn)

}

func TestCompareRpcGraphqlInOneThread(t *testing.T) {

	n := 1000

	txList, encodedTxList, msgList, callList, err := setDataForCompareTest(n)
	if err != nil {
		t.Error(err)
	}

	/* RPC WRITE */
	s := time.Now().UnixMilli()
	for _, tx := range txList {
		_, err := transaction.SendTx(tx)
		if err != nil {
			t.Error(err)
		}
	}
	e := time.Now().UnixMilli()
	t.Logf("rpc write %d개 소요시간 %dms\n", n, e-s)

	/* GRAPHQL WRITE */
	s = time.Now().UnixMilli()
	_, err = BesuMultiWrite(encodedTxList)
	if err != nil {
		t.Error(err)
	}
	e = time.Now().UnixMilli()
	t.Logf("graphql write %d개 소요시간 %dms\n", n, e-s)

	/************************ CALL ******************************/

	/* RPC Call */
	s = time.Now().UnixMilli()
	for _, msg := range msgList {
		_, err := transaction.CallByMsg(nil, msg)
		if err != nil {
			t.Error(err)
		}
	}
	e = time.Now().UnixMilli()
	t.Logf("rpc call %d개 소요시간 %dms\n", n, e-s)

	/* GRAPHQL CALL */
	s = time.Now().UnixMilli()
	_, err = BesuMultiCall(nil, callList)
	if err != nil {
		t.Error(err)
	}
	e = time.Now().UnixMilli()
	t.Logf("graphql call %d개 소요시간 %dms\n", n, e-s)
}

func TestCompareRpcGraphqlInMultiThread(t *testing.T) {

	n := 1000

	txList, encodedTxList, msgList, callList, err := setDataForCompareTest(n)
	if err != nil {
		t.Error(err)
	}

	/* RPC WRITE */
	var wg1 sync.WaitGroup
	wg1.Add(n)
	s := time.Now().UnixMilli()
	for _, tx := range txList {
		go func(wg1 *sync.WaitGroup, tx *types.Transaction) {
			_, err := transaction.SendTx(tx)
			if err != nil {
				t.Error(err)
			}
			wg1.Done()
		}(&wg1, tx)
	}
	wg1.Wait()
	e := time.Now().UnixMilli()
	t.Logf("rpc write %d개 소요시간 %dms\n", n, e-s)

	/* GRAPHQL WRITE */
	var wg2 sync.WaitGroup
	wg2.Add(n)
	s = time.Now().UnixMilli()
	for i := 0; i < n; i += besuLimit {
		upto := min(n, i+besuLimit)
		go func(wg2 *sync.WaitGroup, encodedTxList []string) {
			_, err = BesuMultiWrite(encodedTxList)
			if err != nil {
				t.Error(err)
			}
			wg2.Done()
		}(&wg2, encodedTxList[i:upto])
	}
	wg2.Wait()
	e = time.Now().UnixMilli()
	t.Logf("graphql write %d개 소요시간 %dms\n", n, e-s)

	/************************ CALL ******************************/

	/* RPC Call */
	var wg3 sync.WaitGroup
	wg3.Add(n)
	s = time.Now().UnixMilli()
	for _, msg := range msgList {
		go func(wg3 *sync.WaitGroup, msg ethereum.CallMsg) {
			_, err := transaction.CallByMsg(nil, msg)
			if err != nil {
				t.Error(err)
			}
			wg3.Done()
		}(&wg3, msg)
	}
	wg3.Wait()
	e = time.Now().UnixMilli()
	t.Logf("rpc call %d개 소요시간 %dms\n", n, e-s)

	/* GRAPHQL CALL */
	var wg4 sync.WaitGroup
	wg4.Add(n)
	s = time.Now().UnixMilli()
	for i := 0; i < n; i += besuLimit {
		upto := min(n, i+besuLimit)
		go func(wg4 *sync.WaitGroup, encodedTxList []Call) {
			_, err = BesuMultiCall(nil, callList)
			if err != nil {
				t.Error(err)
			}
			wg4.Done()
		}(&wg4, callList[i:upto])
	}
	e = time.Now().UnixMilli()
	t.Logf("graphql call %d개 소요시간 %dms\n", n, e-s)
}

func setDataForCompareTest(n int) (txList []*types.Transaction, encodedTxList []string, msgList []ethereum.CallMsg, callList []Call, err error) {

	abi, err := transaction.AgeInfoStrageMetaData.GetAbi()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	address := common.HexToAddress(cntr_addr)

	pk := config.Config.Accounts["account1"].PrivateKey[2:]

	/* RPC WRITE 변수*/
	txList = make([]*types.Transaction, n)
	/* GRAPHQL WRITE 변수*/
	encodedTxList = make([]string, n)

	for i := range n {
		data, err := abi.Pack("setAge", fmt.Sprintf("person%d", i), big.NewInt(int64(i)))
		if err != nil {
			return nil, nil, nil, nil, err
		}
		tx, err := transaction.CreateSignedTx(pk, &address, nil, data)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		txList[i] = tx

		encoded, err := rlp.EncodeToBytes(tx)
		if err != nil {
			return nil, nil, nil, nil, err
		}

		encodedTxList[i] = "0x" + common.Bytes2Hex(encoded)

	}

	/* RPC CALL 변수*/
	msgList = make([]ethereum.CallMsg, n)
	/* GRAPHQL CALL 변수*/
	callList = make([]Call, n)

	for i := range n {
		data, err := abi.Pack("getAge", fmt.Sprintf("person%d", i))
		if err != nil {
			return nil, nil, nil, nil, err
		}

		msgList[i] = ethereum.CallMsg{
			To:   &address,
			Data: data,
		}

		callList[i] = Call{
			To:   cntr_addr,
			Data: "0x" + hex.EncodeToString(data),
		}

	}

	return

}
