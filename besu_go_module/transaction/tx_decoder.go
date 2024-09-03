package transaction

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

func ParseTxData(tx *types.Transaction, abi *abi.ABI) (string, error) {

	// contractName
	txDataMap := make(map[string]interface{})

	data := tx.Data()
	method, err := abi.MethodById(data[:4])
	if err != nil {
		return "", fmt.Errorf("메소드 id 조회 실패 %w", err)
	}

	txDataMap["Method"] = method.Name

	params := make(map[string]interface{})
	err = method.Inputs.UnpackIntoMap(params, data[4:])
	if err != nil {
		return "", fmt.Errorf("파라미터 unpack 실패 %w", err)
	}

	txDataMap["Parameter"] = params

	j, err := json.MarshalIndent(txDataMap, "", "\t")
	if err != nil {
		return "", fmt.Errorf("마샬링 실패 %w", err)
	}

	return string(j), nil

}
