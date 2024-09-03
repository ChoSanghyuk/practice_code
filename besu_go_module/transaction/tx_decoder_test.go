package transaction

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestTxDecode(t *testing.T) {

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
