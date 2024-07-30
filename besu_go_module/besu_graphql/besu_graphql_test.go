package besu_graphql

import (
	"encoding/hex"
	"go_module/config"
	"go_module/transaction"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

	abi, err := transaction.AgeInfoStrageMetaData.GetAbi()
	if err != nil {
		t.Error(err)
	}

	data1, err := abi.Pack("getAge", "Jo")
	if err != nil {
		t.Error(err)
	}
	data2, err := abi.Pack("getAge", "min")
	if err != nil {
		t.Error(err)
	}

	rtn, err := BesuMultiCall(nil, []Call{
		{
			To:   cntr_addr,
			Data: "0x" + hex.EncodeToString(data1),
		},
		{
			To:   cntr_addr,
			Data: "0x" + hex.EncodeToString(data2),
		},
	})
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
