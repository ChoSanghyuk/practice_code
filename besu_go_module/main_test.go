package main

import (
	"encoding/hex"
	"fmt"
	"go_module/config"
	g "go_module/graphql"
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

var cntr_addr = "0x77CAB8c457F7fDE65546C5C3845120Cdf84B813f"

func TestBesuCall(t *testing.T) {

	abi, err := transaction.AgeInfoStrageMetaData.GetAbi()
	if err != nil {
		t.Error(err)
	}

	data, err := abi.Pack("getAge", "name20")
	if err != nil {
		t.Error(err)
	}

	rtn, err := g.BesuCall(nil, g.Call{
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

	data, err := abi.Pack("setAge", "Jo", big.NewInt(29))
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

	rtn, err := g.BesuWrite("0x" + common.Bytes2Hex(encoded))
	if err != nil {
		t.Error(err)
	}

	t.Log(rtn)
}

func TestMultiBesuCall(t *testing.T) {

	n := 1000

	abi, err := transaction.AgeInfoStrageMetaData.GetAbi()
	if err != nil {
		t.Error(err)
	}

	callDatas := make([]g.Call, n)

	for i := range n {
		data, err := abi.Pack("getAge", fmt.Sprintf("name%d", i))
		if err != nil {
			t.Error(err)
		}
		callDatas[i] = g.Call{
			To:   cntr_addr,
			Data: "0x" + hex.EncodeToString(data),
		}
	}

	rtn, err := g.BesuMultiCall(nil, callDatas)
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

	n := 1000
	var gas uint64 = 100000

	abi, err := transaction.AgeInfoStrageMetaData.GetAbi()
	if err != nil {
		t.Error(err)
	}

	txs := make([]string, n)

	address := common.HexToAddress(cntr_addr)
	pk := config.Config.Accounts["account1"].PrivateKey[2:]

	for i := range n {

		data, err := abi.Pack("setAge", fmt.Sprintf("name%d", i), big.NewInt(int64(i)))
		if err != nil {
			t.Error(err)
		}

		tx, err := transaction.CreateSignedTx2(pk, &address, nil, data, gas, uint64(i))
		if err != nil {
			t.Error(err)
		}

		encoded, err := rlp.EncodeToBytes(tx)
		if err != nil {
			t.Error(err)
		}
		txs[i] = "0x" + common.Bytes2Hex(encoded)
	}

	rtn, err := g.BesuMultiWrite(txs)
	if err != nil {
		t.Error(err)
	}

	t.Log(rtn)

}

func TestCompareRpcGraphqlInOneThread(t *testing.T) {

	n := 100000

	msgList, callList, err := setDataForCompareCallTest(n)
	if err != nil {
		t.Error(err)
	}

	// /* RPC WRITE */
	// s := time.Now().UnixMilli()
	// for _, tx := range txList {
	// 	_, err := transaction.SendTx(tx)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// }
	// e := time.Now().UnixMilli()
	// t.Logf("rpc write %d개 소요시간 %dms\n", n, e-s)

	// /* GRAPHQL WRITE */
	// s = time.Now().UnixMilli()
	// _, err = g.BesuMultiWrite(encodedTxList)
	// if err != nil {
	// 	t.Error(err)
	// }
	// e = time.Now().UnixMilli()
	// t.Logf("graphql write %d개 소요시간 %dms\n", n, e-s)

	/************************ CALL ******************************/

	/* RPC Call */
	s := time.Now().UnixMilli()
	for _, msg := range msgList {
		_, err := transaction.CallByMsg(nil, msg)
		if err != nil {
			t.Error(err)
		}
	}
	e := time.Now().UnixMilli()
	t.Logf("rpc call %d개 소요시간 %dms\n", n, e-s)

	/* GRAPHQL CALL */
	s = time.Now().UnixMilli()
	_, err = g.BesuMultiCall(nil, callList)
	if err != nil {
		t.Error(err)
	}
	e = time.Now().UnixMilli()
	t.Logf("graphql call %d개 소요시간 %dms\n", n, e-s)
}

func TestCompareRpcGraphqlInMultiThread(t *testing.T) {

	const maxGoroutines = 70
	n := 100000

	msgList, callList, err := setDataForCompareCallTest(n)
	if err != nil {
		t.Error(err)
	}

	semaphore := make(chan struct{}, maxGoroutines)

	// /* RPC WRITE */
	// var wg1 sync.WaitGroup
	// wg1.Add(n)
	// s := time.Now().UnixMilli()
	// for _, tx := range txList {
	// 	go func(wg1 *sync.WaitGroup, tx *types.Transaction) {
	// 		_, err := transaction.SendTx(tx)
	// 		if err != nil {
	// 			t.Error(err)
	// 		}
	// 		wg1.Done()
	// 	}(&wg1, tx)
	// }
	// wg1.Wait()
	// e := time.Now().UnixMilli()
	// t.Logf("rpc write %d개 소요시간 %dms\n", n, e-s)

	// /* GRAPHQL WRITE */
	// var wg2 sync.WaitGroup
	// wg2.Add(n)
	// s = time.Now().UnixMilli()
	// for i := 0; i < n; i += g.BesuLimit {
	// 	upto := min(n, i+g.BesuLimit)
	// 	go func(wg2 *sync.WaitGroup, encodedTxList []string) {
	// 		_, err = g.BesuMultiWrite(encodedTxList)
	// 		if err != nil {
	// 			t.Error(err)
	// 		}
	// 		wg2.Done()
	// 	}(&wg2, encodedTxList[i:upto])
	// }
	// wg2.Wait()
	// e = time.Now().UnixMilli()
	// t.Logf("graphql write %d개 소요시간 %dms\n", n, e-s)

	/************************ CALL ******************************/

	/* RPC Call */
	var wg3 sync.WaitGroup
	wg3.Add(n)
	s := time.Now().UnixMilli()
	for _, msg := range msgList {
		semaphore <- struct{}{}
		go func(wg3 *sync.WaitGroup, msg ethereum.CallMsg) {
			_, err := transaction.CallByMsg(nil, msg)
			if err != nil {
				t.Error(err)
			}
			wg3.Done()
			<-semaphore
		}(&wg3, msg)
	}
	wg3.Wait()
	e := time.Now().UnixMilli()
	t.Logf("rpc call %d개 소요시간 %dms\n", n, e-s)

	/* GRAPHQL CALL */
	var wg4 sync.WaitGroup
	wg4.Add(n/g.BesuLimit + 1)
	s = time.Now().UnixMilli()
	for i := 0; i < n; i += g.BesuLimit {
		upto := min(n, i+g.BesuLimit)
		go func(wg4 *sync.WaitGroup, encodedTxList []g.Call) {
			_, err = g.BesuMultiCall(nil, encodedTxList)
			if err != nil {
				t.Error(err)
			}
			wg4.Done()
		}(&wg4, callList[i:upto])
	}
	wg4.Wait()
	e = time.Now().UnixMilli()
	t.Logf("graphql call %d개 소요시간 %dms\n", n, e-s)
}

func setDataForCompareCallTest(n int) (msgList []ethereum.CallMsg, callList []g.Call, err error) {

	abi, err := transaction.AgeInfoStrageMetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	address := common.HexToAddress(cntr_addr)

	/* RPC CALL 변수*/
	msgList = make([]ethereum.CallMsg, n)
	/* GRAPHQL CALL 변수*/
	callList = make([]g.Call, n)

	for i := range n {
		data, err := abi.Pack("getAge", fmt.Sprintf("person%d", i))
		if err != nil {
			return nil, nil, err
		}

		msgList[i] = ethereum.CallMsg{
			To:   &address,
			Data: data,
		}

		callList[i] = g.Call{
			To:   cntr_addr,
			Data: "0x" + hex.EncodeToString(data),
		}

	}

	return

}

func setDataForCompareWriteTest(n int) (txList []*types.Transaction, encodedTxList []string, err error) {

	abi, err := transaction.AgeInfoStrageMetaData.GetAbi()
	if err != nil {
		return nil, nil, err
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
			return nil, nil, err
		}
		tx, err := transaction.CreateSignedTx(pk, &address, nil, data)
		if err != nil {
			return nil, nil, err
		}
		txList[i] = tx

		encoded, err := rlp.EncodeToBytes(tx)
		if err != nil {
			return nil, nil, err
		}

		encodedTxList[i] = "0x" + common.Bytes2Hex(encoded)

	}
	return
}
