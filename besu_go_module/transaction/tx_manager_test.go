package transaction

import (
	"fmt"
	"go_module/config"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var addr = "0x42699A7612A82f1d9C36148af9C77354759b210b"
var contractAbi, _ = AgeInfoStrageMetaData.GetAbi()
var contractBin = AgeInfoStrageMetaData.Bin

func TestDeploy(t *testing.T) {
	// pk := info.BesuKey["account1"]["privateKey"][2:]

	account := config.Config.Accounts["account1"]

	addr, txHash, _, err := Deploy(account.PrivateKey[2:], contractAbi, contractBin)

	if err != nil {
		t.Error(err)
	}
	fmt.Println("생성된 컨트랙트 주소", addr)
	fmt.Println("수행된 트랜잭션 해시", txHash)
	/*
		생성된 컨트랙트 주소 0x42699A7612A82f1d9C36148af9C77354759b210b
		수행된 트랜잭션 해시 0x125a67bd36de9a0ab91bf86df5feb08ccf125166bfd318c6471314bb61aefa57
	*/
}

func TestCall(t *testing.T) {

	rtn, err := Call(nil, addr, contractAbi, "getAge", big.NewInt(1))
	if err != nil {
		fmt.Println(err, "Call 실패")
	}

	fmt.Printf("%v", rtn)

	assert.NotNil(t, rtn)

}

const n = 1000000

// 단일 스레드 1000개 소요 시간 : 492
func TestMultiCall(t *testing.T) {
	var rtns []interface{} = make([]interface{}, n)

	address := common.HexToAddress(addr)

	input, err := contractAbi.Pack("getAge", big.NewInt(1))
	if err != nil {
		fmt.Println(err)
	}

	msg := ethereum.CallMsg{
		From: common.Address{},
		To:   &address,
		Data: input,
	}

	s := time.Now().UnixMilli()
	for i := 0; i < n; i++ {
		rtn, err := CallByMsg(nil, msg)
		if err != nil {
			fmt.Println(err, "Call 실패")
		}
		rtns[i] = rtn
	}
	e := time.Now().UnixMilli()

	// fmt.Println(rtns)
	_ = rtns
	fmt.Printf("단일 스레드 %d개 소요 시간 : %d", n, e-s)

}

// Goroutine 사용 시, 1000개 소요 시간 : 296
func TestMultiWithGoGroutineCall(t *testing.T) {
	var rtns [][]byte = make([][]byte, n)

	address := common.HexToAddress(addr)

	input, err := contractAbi.Pack("getAge", big.NewInt(1))
	if err != nil {
		fmt.Println(err)
	}

	msg := ethereum.CallMsg{
		From: common.Address{},
		To:   &address,
		Data: input,
	}
	c := make(chan []byte)

	s := time.Now().UnixMilli()
	for i := 0; i < n; i++ {
		go callByMsgWithGoRoutineForTest(c, nil, msg)
	}

	for i := 0; i < n; i++ {
		rtns[i] = <-c
	}

	e := time.Now().UnixMilli()

	// fmt.Println(rtns)
	_ = rtns
	fmt.Printf("Goroutine 사용 시, %d개 소요 시간 : %d", n, e-s)

}

func TestWrite(t *testing.T) {

	pk := config.Config.Accounts["account1"].PrivateKey[2:]

	rtn, err := Write(pk, addr, contractAbi, "setAgeList", big.NewInt(1), big.NewInt(30))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%v", rtn.Hex())
}
