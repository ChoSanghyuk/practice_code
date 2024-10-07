package transaction

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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

type eventInfo struct {
	Address   common.Address         `json:"address"`
	EventName string                 `json:"event"`
	Index     uint                   `json:"index"`
	Parameter map[string]interface{} `json:"parameter"`
}

/*
types.Receipt에는 revertReason를 받는 필드가 없음
따라서, RevertReason string `json:"revertReason"` 와 같이 revertReason을 받을 필드를 가진 신규 struct가 필요함

```
var receipt 신규struct
err := client.Client().CallContext(ctx, &receipt, "", txHash)
```
*/
func ParseReceipt(abiEvent abi.Event, receipt *types.Receipt) (string, error) {

	events := make([]*eventInfo, len(receipt.Logs))
	for i, log := range receipt.Logs {

		eventInfo := eventInfo{}
		events[i] = &eventInfo

		eventInfo.Address = log.Address
		eventInfo.Index = log.Index

		// var abiEvent abi.Event // todo. topic과 대응되는 event 조회 필요. (log.Topics[0]이 Event ID의 Hex값)
		eventInfo.EventName = abiEvent.Name

		paramMap := make(map[string]interface{})
		eventInfo.Parameter = paramMap

		err := abiEvent.Inputs.UnpackIntoMap(paramMap, log.Data)
		if err != nil {
			// todo. error handle
		}

		indexed := make([]abi.Argument, len(log.Topics)-1)
		idx := 0
		for _, input := range abiEvent.Inputs {
			if input.Indexed {
				indexed[idx] = input
				idx++
			}
		}

		err = abi.ParseTopicsIntoMap(paramMap, indexed, log.Topics[1:])
		if err != nil {
			// todo. error handle
		}

		for i, input := range indexed {
			if input.Type.T == abi.FixedBytesTy || input.Type.T == abi.BytesTy {
				topic := log.Topics[i+1]
				paramMap[input.Name] = topic.Hex()
			}
		}
	}

	jsonData, err := json.Marshal(events)
	if err != nil {
		// todo. error handle
	}

	var _ types.Receipt

	return string(jsonData), nil

}

// revertReason 받을 수 있는 새 receipt. 필드를 string으로 받아야 하는지 테스트 필요
type CustomReceipt struct {
	BlockHash         common.Hash    `json:"blockHash"`
	BlockNumber       *big.Int       `json:"blockNumber"`
	ContractAddress   common.Address `json:"contractAddress"`
	CumulativeGasUsed uint64         `json:"cumulativeGasUsed"`
	EffectiveGasPrice *big.Int       `json:"effectiveGasPrice"`
	From              string         `json:"from"`
	GasUsed           uint64         `json:"gasUsed"`
	Logs              []*types.Log   `json:"logs"`
	Bloom             types.Bloom    `json:"logsBloom"`
	RevertReason      string         `json:"revertReason"`
	Status            uint64         `json:"status"`
	To                string         `json:"to"`
	TxHash            common.Hash    `json:"transactionHash"`
	TransactionIndex  uint           `json:"transactionIndex"`
	Type              string         `json:"type"`
}
