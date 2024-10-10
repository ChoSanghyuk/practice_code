package transaction

import (
	"context"
	"encoding/json"
	"fmt"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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

		// var abiEvent abi.Event // todo. topic과 대응되는 event 조회 필요
		eventInfo.EventName = abiEvent.Name

		paramMap := make(map[string]interface{})
		eventInfo.Parameter = paramMap

		err := abiEvent.Inputs.UnpackIntoMap(paramMap, log.Data)
		if err != nil {
			return "", nil
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
			return "", nil
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
		return "", nil
	}

	var _ types.Receipt

	return string(jsonData), nil

}

/*
revertReason 받을 수 있는 새 receipt.
대부분의 필드들은 우선 string으로 받아야 하며, 이후 형변환 필요
*/
type CustomReceipt struct {
	BlockHash         string       `json:"blockHash"`         // common.Hash
	BlockNumber       string       `json:"blockNumber"`       // *big.Int
	ContractAddress   string       `json:"contractAddress"`   // common.Address
	CumulativeGasUsed string       `json:"cumulativeGasUsed"` // uint64
	EffectiveGasPrice string       `json:"effectiveGasPrice"` // *big.Int
	From              string       `json:"from"`              // string
	GasUsed           string       `json:"gasUsed"`           // uint64
	Logs              []*types.Log `json:"logs"`              // []*types.Log
	Bloom             string       `json:"logsBloom"`         // types.Bloom
	RevertReason      string       `json:"revertReason"`      // string
	Status            string       `json:"status"`            // uint64
	To                string       `json:"to"`                // string
	TxHash            string       `json:"transactionHash"`   // common.Hash
	TransactionIndex  string       `json:"transactionIndex"`  // uint
	Type              string       `json:"type"`              // string
}

func TransactionCustomReceipt(client *ethclient.Client, ctx context.Context, txHash common.Hash) (*CustomReceipt, error) {
	var r *CustomReceipt
	err := client.Client().CallContext(ctx, &r, "eth_getTransactionReceipt", txHash)
	if err == nil && r == nil {
		return nil, ethereum.NotFound
	}

	return r, err

}
