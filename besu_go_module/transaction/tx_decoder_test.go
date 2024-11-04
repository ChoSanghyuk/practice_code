package transaction

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestTxDecode(t *testing.T) {

	// AgeInfoStrage의 setAge Transaction Hash
	hash := common.HexToHash("0x4882aaa2fb03e326ecee675ebb3b7fc273d86b069265b630cae763b19c0609ef")
	tx, _, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		t.Error(err)
	}

	abi, err := AgeInfoStrageMetaData.GetAbi()
	if err != nil {
		t.Error(err)
	}

	parsed, err := ParseTxData(tx, abi)
	if err != nil {
		t.Error(err)
	}

	t.Log(parsed)

}

func TestParseReceipt(t *testing.T) {

	// ClientReceipt의 deposit Transaction Hash
	hash := common.HexToHash("0xf3e86a122a9a51e90c5bcf5fc5d85ccdb3274b981eb1b938db28365b851ce89f")
	abi, err := ClientReceiptMetaData.GetAbi()
	if err != nil {
		t.Error(err)
	}

	receipt, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		t.Error(err)
	}

	parsed, err := ParseReceipt(abi.Events["Deposit"], receipt)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(parsed)
}

func TestCustomReceipt(t *testing.T) {
	hash := common.HexToHash("0x3b2ab1d3c7820a084bd215cab13925814e3abbf0d3e0bf121d7f225dc1728c8c")

	receipt, err := TransactionCustomReceipt(client, context.Background(), hash)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", receipt)
}
