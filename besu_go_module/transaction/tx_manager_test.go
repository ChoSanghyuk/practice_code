package transaction

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go_module/config"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
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

func TestGetReceipt(t *testing.T) {

	txHash := "0xa81fd5469698659007209efdf32139cb241d3159c86bf24f5e43efe0ceb669b1"

	r, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

/*
생성된 컨트랙트 주소 0x87eED15a3230E1Fd6DcD0A1c619ec712C0196Ac0
수행된 트랜잭션 해시 0x6b2ea72c63439e89bce671eed453568a18a1b7bf8a510ed7be911b67127fec00

event
emit 트랜잭션 해시 0xa4ba3d49a84981185771faa2aa0f5ef2e5466676f1425fd3fbc2fedee3e345f5
*/

func TestReceipt(t *testing.T) {

	hash := common.HexToHash("0xa4ba3d49a84981185771faa2aa0f5ef2e5466676f1425fd3fbc2fedee3e345f5")
	r, _ := client.TransactionReceipt(context.Background(), hash)

	raw, _ := r.MarshalBinary()
	byte1, err := json.MarshalIndent(r, "", "\t")
	if err != nil {
		t.Log(err)
	}
	fmt.Println(string(byte1))
	fmt.Println(strings.Repeat("=", 20))

	// 	raw := common.Hex2Bytes(`7b22626c6f636b48617368223a22307830636232666365346336336230393437396230366331363566383431343036613435363562386636306139633162343534643331306432376463313634663636222c22626c6f636b4e756d626572223a22307832313061222c22636f6e747261637441646472657373223a22307838376565643135613332333065316664366463643061316336313965633731326330313936616330222c2263756d756c617469766547617355736564223a2230783733313930222c2266726f6d223a22307866653362353537653866623632623839663439313662373231626535356365623832386462643733222c2267617355736564223a2230783733313930222c226566666563746976654761735072696365223a223078336538222c226c6f6773223a5b5d2c226c6f6773426c6f6f6d223a223078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030
	// 3030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030222c22737461747573223a22307831222c22746f223a6e756c6c2c227472616e73616374696f6e48617368223a22307836623265613732633633343339653839626365363731656564343533353638613138613162376266386135313065643762653931316236373132376665633030222c227472616e73616374696f6e496e646578223a22307830222c2274797065223a22307830227`)

	new := types.Receipt{
		TxHash: common.HexToHash("0xa4ba3d49a84981185771faa2aa0f5ef2e5466676f1425fd3fbc2fedee3e345f5"),
	}
	err = new.UnmarshalBinary(raw)
	if err != nil {
		t.Log(err)
		return
	}
	byte2, err := json.MarshalIndent(new, "", "\t")
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(string(byte2))

}

func TestReceiptWriter(t *testing.T) {

	hash := common.HexToHash("0xa4ba3d49a84981185771faa2aa0f5ef2e5466676f1425fd3fbc2fedee3e345f5")
	r, _ := client.TransactionReceipt(context.Background(), hash)
	fmt.Println(strings.Repeat("=", 20))

	var buf bytes.Buffer

	r.EncodeRLP(&buf)
	fmt.Println("Binary Data :", hex.EncodeToString(buf.Bytes()))

	new := types.Receipt{}
	rlp.DecodeBytes(buf.Bytes(), new)

	jsonData, err := json.MarshalIndent(new, "", "\t")
	if err != nil {
		t.Log(err)
	}
	fmt.Println(string(jsonData))
}
