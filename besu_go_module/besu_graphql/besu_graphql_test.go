package besu_graphql

import (
	"encoding/hex"
	"go_module/config"
	"go_module/transaction"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

	data, err := abi.Pack("setAge", "min", 29)
	if err != nil {
		t.Error(err)
	}

	address := common.HexToAddress(cntr_addr)

	account := config.Config.Accounts["account1"]

	tx, err := transaction.CreateTx(account.PrivateKey[2:], &address, nil, data)
	if err != nil {
		t.Error(err)
	}

	rtn, err := BesuWrite(tx.Hash().Hex())
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

	data, err := abi.Pack("getAge", "Jo")
	if err != nil {
		t.Error(err)
	}

	rtn, err := BesuMultiCall(nil, []Call{
		{
			To:   cntr_addr,
			Data: "0x" + hex.EncodeToString(data),
		},
		{
			To:   cntr_addr,
			Data: "0x" + hex.EncodeToString(data),
		},
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(rtn)
}
