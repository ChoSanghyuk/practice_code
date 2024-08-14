package test

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go_module/config"
	g "go_module/graphql"
	"go_module/transaction"
	"math/big"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

/*
Graphql과 RPC call 간의 성능 비교 테스트
*/

var cntr_addr = "0x77CAB8c457F7fDE65546C5C3845120Cdf84B813f"

func TestBesuCall(t *testing.T) {

	abi, err := transaction.AgeInfoStrageMetaData.GetAbi()
	if err != nil {
		t.Error(err)
	}

	data, err := abi.Pack("getAge", "name6300")
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

	n := 100000
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
	n := 1000

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
	wg4.Add(n/g.BesuCallLimit + 1)
	s = time.Now().UnixMilli()
	for i := 0; i < n; i += g.BesuCallLimit {
		upto := min(n, i+g.BesuCallLimit)
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

/************************************************
Graphql Receipt 조회 성능 테스트
		rpc		graphql	rpc/graphql
1000	1551	509		3.04
10000	11588	3426	3.38
100000	105804	29240	3.61
************************************************/

func TestCompareRpcGraphqlReceiptCall(t *testing.T) {

	n := 10

	hashS := "0xa4ba3d49a84981185771faa2aa0f5ef2e5466676f1425fd3fbc2fedee3e345f5"
	hash := common.HexToHash(hashS)

	rpcLi := make([]*common.Hash, n)
	graphqlLi := make([]string, n)

	rpcRslt := make([]*types.Receipt, n)
	// graphqlRslt := make([]*types.Receipt, n)

	var err error

	for i := range n {
		rpcLi[i] = &hash
		graphqlLi[i] = hashS
	}

	s := time.Now().UnixMilli()
	for i := range n {
		rpcRslt[i], err = transaction.TransactionReceipt(context.Background(), *rpcLi[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	e := time.Now().UnixMilli()
	t.Logf("rpc receipt 조회 %d개 소요시간 %dms\n", n, e-s)

	s = time.Now().UnixMilli()
	graphqlRslt, err := g.GetMultiReceipt(graphqlLi)
	if err != nil {
		t.Fatal(err)
	}
	e = time.Now().UnixMilli()
	t.Logf("graphql receipt 조회 %d개 소요시간 %dms\n", n, e-s)

	if n <= 10 {
		j1, _ := json.MarshalIndent(rpcRslt, "", "\t")
		t.Log(string(j1))
		fmt.Println(strings.Repeat("=", 20))
		j2, _ := json.MarshalIndent(graphqlRslt, "", "\t")
		t.Log(string(j2))
	}

}

/************************************************
Graphql Receipt 멀티스레드 조회 성능 테스트
		rpc		graphql	rpc/graphql
1000	793		167		4.74
10000	3543	1453	2.43
100000	36684	10837	3.38
************************************************/

func TestCompareRpcGraphqlReceiptCallInMultiThread(t *testing.T) {

	n := 100000
	const maxGoroutines = 70

	hashS := "0xa4ba3d49a84981185771faa2aa0f5ef2e5466676f1425fd3fbc2fedee3e345f5"
	hash := common.HexToHash(hashS)

	rpcLi := make([]*common.Hash, n)
	graphqlLi := make([]string, n)
	for i := range n {
		rpcLi[i] = &hash
		graphqlLi[i] = hashS
	}

	var err error

	var wg sync.WaitGroup
	wg.Add(n)

	semaphore := make(chan struct{}, maxGoroutines)

	s := time.Now().UnixMilli()
	for i := range n {
		semaphore <- struct{}{}
		go func(wg *sync.WaitGroup, semaphore chan struct{}) {
			_, err := transaction.TransactionReceipt(context.Background(), *rpcLi[i])
			if err != nil {
				t.Error(err)
			}
			<-semaphore
			wg.Done()
		}(&wg, semaphore)

	}
	wg.Wait()
	e := time.Now().UnixMilli()
	t.Logf("rpc receipt 조회 %d개 소요시간 %dms\n", n, e-s)

	var wg2 sync.WaitGroup
	wg2.Add(n/g.BesuReceiptLimit + 1)
	s = time.Now().UnixMilli()
	for i := 0; i < n; i += g.BesuReceiptLimit {
		upto := min(n, i+g.BesuReceiptLimit)
		semaphore <- struct{}{}
		go func(wg2 *sync.WaitGroup, semaphore chan struct{}, hashLi []string) {
			_, err = g.GetMultiReceipt(hashLi)
			if err != nil {
				t.Error(err)
			}
			wg2.Done()
			<-semaphore
		}(&wg2, semaphore, graphqlLi[i:upto])
	}
	wg2.Wait()
	e = time.Now().UnixMilli()
	t.Logf("graphql receipt 조회 %d개 소요시간 %dms\n", n, e-s)

}
