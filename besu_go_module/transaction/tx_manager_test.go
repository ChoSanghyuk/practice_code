package transaction

import (
	"fmt"
	"go_module/config"
	contract "go_module/smart_contract"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

var addr = "0x42699A7612A82f1d9C36148af9C77354759b210b"

func TestDeploy(t *testing.T) {
	// pk := info.BesuKey["account1"]["privateKey"][2:]

	account := config.Config.Accounts["account1"]

	abi, _ := contract.ContractMetaData.GetAbi()
	bin := contract.ContractMetaData.Bin

	addr, txHash := Deploy(account.PrivateKey[2:], abi, bin)
	fmt.Println("생성된 컨트랙트 주소", addr)
	fmt.Println("수행된 트랜잭션 해시", txHash)
	/*
		생성된 컨트랙트 주소 0x42699A7612A82f1d9C36148af9C77354759b210b
		수행된 트랜잭션 해시 0x125a67bd36de9a0ab91bf86df5feb08ccf125166bfd318c6471314bb61aefa57
	*/
}

func TestCall(t *testing.T) {

	abi, _ := contract.ContractMetaData.GetAbi()

	rtn, err := Call(nil, addr, abi, "getAge", big.NewInt(1))
	if err != nil {
		fmt.Println(err, "Call 실패")
	}

	// a := fmt.Sprintf("%s", rtn)
	fmt.Println(rtn)

}

const n = 1000000

// 단일 스레드 1000개 소요 시간 : 492
func TestMultiCall(t *testing.T) {
	var rtns []interface{} = make([]interface{}, n)

	abi, _ := contract.ContractMetaData.GetAbi()

	address := common.HexToAddress(addr)

	input, err := abi.Pack("getAge", big.NewInt(1))
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
		rtn, err := callByMsgForTest(nil, msg)
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

	abi, _ := contract.ContractMetaData.GetAbi()

	address := common.HexToAddress(addr)

	input, err := abi.Pack("getAge", big.NewInt(1))
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

	abi, _ := contract.ContractMetaData.GetAbi()

	rtn := Write(pk, addr, abi, "setAgeList", big.NewInt(1), big.NewInt(30))
	fmt.Printf("%v", rtn.Hex())
}
