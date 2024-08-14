package graphql

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/machinebox/graphql"
)

/************************************************
Graphql Receipt 단일 스레드 환경 조회 성능 테스트 (단위:ms)
		rpc		graphql	rpc/graphql
1000	1551	509		3.04
10000	11588	3426	3.38
100000	105804	29240	3.61

Graphql Receipt 멀티스레드 환경 조회 성능 테스트 결과 (단위:ms)
(connection 제한으로 goroutine 수는 70로 유지)
		rpc		graphql	rpc/graphql
1000	793		167		4.74
10000	3543	1453	2.43
100000	36684	10837	3.38
************************************************/

const BesuReceiptLimit = 90

type GraphqlReceipt struct {
	TxHash          common.Hash `json:"hash"`
	CreatedContract struct {
		Address common.Address `json:"address"`
	} `json:"createdContract"`
	Gasused           string `json:"gasUsed"`
	EffectiveGasPrice string `json:"effectiveGasPrice"`
	Block             struct {
		BlockHash   common.Hash `json:"hash"`
		BlockNumber string      `json:"number"`
	} `json:"block"`
	TxIndex    string `json:"index"`
	RawReceipt string `json:"rawReceipt"`
}

var receiptQuery = `
	query getReceipt($txHash: Bytes32!) {
		transaction(hash: $txHash){
			hash,
			createdContract{
				address
			},
			gasUsed,
			effectiveGasPrice,
			block{
				hash,
				number
			},
			index,
			rawReceipt
		}
	}
`

var multiReceiptQuery = `
	query getReceipt(%s) {
		%s
	}

	fragment receiptFields on Transaction {
		hash,
		createdContract{
			address
		},
		gasUsed,
		effectiveGasPrice,
		block{
			hash,
			number
		},
		index,
		rawReceipt
	}
`

var receiptForm = `
	tx%d : transaction(hash: $txHash%d){
		...receiptFields
	}
`
var receiptVariableForm = "$txHash%d: Bytes32!"

func GetReceipt(hash string) (*types.Receipt, error) {

	req := graphql.NewRequest(receiptQuery)
	req.Var("txHash", hash)

	var res map[string]GraphqlReceipt
	err := client.Run(context.Background(), req, &res)
	if err != nil {
		return nil, fmt.Errorf("client run 에러 %w", err)
	}

	// j, err := json.MarshalIndent(res, "", "\t")
	// fmt.Println(string(j))

	r := res["transaction"]
	rtn, err := parseReceipt(&r)
	if err != nil {
		return nil, fmt.Errorf("receipt 파싱 에러 %w", err)
	}

	return rtn, nil
}

func GetMultiReceipt(hashList []string) ([]*types.Receipt, error) {

	var res map[string]GraphqlReceipt
	var s int = 0
	var e int = min(len(hashList), BesuReceiptLimit)

loop:
	var varBuilder strings.Builder
	var queryBuilder strings.Builder

	for i := s; i < e; i++ {

		if i != s {
			varBuilder.WriteString(", ")
		}
		varBuilder.WriteString(fmt.Sprintf(receiptVariableForm, i))

		queryBuilder.WriteString(fmt.Sprintf(receiptForm, i, i))

	}

	query := fmt.Sprintf(multiReceiptQuery, varBuilder.String(), queryBuilder.String())
	// fmt.Println(query)

	req := graphql.NewRequest(query)

	for i := s; i < e; i++ {
		req.Var(fmt.Sprintf("txHash%d", i), hashList[i])
	}

	err := client.Run(context.Background(), req, &res)
	if err != nil {
		return nil, fmt.Errorf("client run 에러 %w", err)
	}

	if e < len(hashList) {
		s += BesuReceiptLimit
		e = min(len(hashList), e+BesuReceiptLimit)
		goto loop
	}

	rtn := make([]*types.Receipt, len(hashList))
	var i int
	for _, v := range res {
		receipt, err := parseReceipt(&v)
		if err != nil {
			return nil, fmt.Errorf("receipt 파싱 에러 %w", err)
		}
		rtn[i] = receipt
		i++
	}

	return rtn, nil

}

func parseReceipt(res *GraphqlReceipt) (*types.Receipt, error) {

	gasUsed, err := strconv.ParseUint(strings.TrimLeft(res.Gasused, "0x"), 16, 64)
	if err != nil {
		return nil, fmt.Errorf("parse 에러 %w", err)
	}
	effGasPrice, err := strconv.ParseInt(strings.TrimLeft(res.EffectiveGasPrice, "0x"), 16, 64)
	if err != nil {
		return nil, fmt.Errorf("parse 에러 %w", err)
	}
	blockNum, err := strconv.ParseInt(strings.TrimLeft(res.Block.BlockNumber, "0x"), 16, 64)
	if err != nil {
		return nil, fmt.Errorf("parse 에러 %w", err)
	}
	txIndex, err := strconv.ParseUint(strings.TrimLeft(res.TxIndex, "0x"), 16, 64)
	if err != nil {
		return nil, fmt.Errorf("parse 에러 %w", err)
	}

	rtn := &types.Receipt{
		TxHash:            res.TxHash,
		ContractAddress:   res.CreatedContract.Address,
		GasUsed:           gasUsed,
		EffectiveGasPrice: big.NewInt(effGasPrice),
		BlockHash:         res.Block.BlockHash,
		BlockNumber:       big.NewInt(blockNum),
		TransactionIndex:  uint(txIndex),
	}

	err = rtn.UnmarshalBinary(common.Hex2Bytes(strings.TrimLeft(res.RawReceipt, "0x")))
	if err != nil {
		return nil, fmt.Errorf("unmarshal 실패 %w", err)
	}

	return rtn, nil
}
