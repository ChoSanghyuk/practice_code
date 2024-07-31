package besu_graphql

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/machinebox/graphql"
)

const besuLimit = 100

type MultiCallsResp struct {
	Block map[string]CallResp `json:"block"`
}

var multiCallQuery = `
	query getCall($blockNumber: Long%s) {
		block(number: $blockNumber){
			%s
		}
	}
	fragment callFields on CallResult {
		data, 
		gasUsed, 
		status
	}
`
var callForm = `
	call%d : call(data: $callData%d){
		...callFields
	}
`
var callVariableForm = "$callData%d: CallData!"

var multiMutQuery = `
	mutation(%s) {
		%s
	}
`
var mutForm = "tx%d: sendRawTransaction(data: $mutData%d) "
var mutVariableForm = "$mutData%d: Bytes!"

func BesuMultiCall(bn *big.Int, callDatas []Call) (MultiCallsResp, error) {

	var res MultiCallsResp
	var upto int = 0

loop:
	upto += besuLimit
	if len(callDatas) < upto {
		upto = len(callDatas)
	}

	var varBuilder strings.Builder
	var callBuilder strings.Builder

	for i := upto - besuLimit; i < upto; i++ {

		varBuilder.WriteString(", ")
		varBuilder.WriteString(fmt.Sprintf(callVariableForm, i))

		callBuilder.WriteString(fmt.Sprintf(callForm, i, i))

	}

	query := fmt.Sprintf(multiCallQuery, varBuilder.String(), callBuilder.String())
	// fmt.Println(query)

	req := graphql.NewRequest(query)

	req.Var("blockNumber", bn)

	for i := upto - besuLimit; i < upto; i++ {
		req.Var(fmt.Sprintf("callData%d", i), callDatas[i])
	}

	err := client.Run(context.Background(), req, &res)
	if err != nil {
		return MultiCallsResp{}, fmt.Errorf("client run 에러 %w", err)
	}

	if upto < len(callDatas) {
		goto loop
	}

	return res, nil
}

func BesuMultiWrite(txs []string) (map[string]string, error) {

	var varBuilder strings.Builder
	var callBuilder strings.Builder

	for i := range len(txs) {
		varBuilder.WriteString(fmt.Sprintf(mutVariableForm, i))
		callBuilder.WriteString(fmt.Sprintf(mutForm, i, i))

		if i < len(txs)-1 {
			varBuilder.WriteString(", ")
			callBuilder.WriteString(", ")
		}
	}

	query := fmt.Sprintf(multiMutQuery, varBuilder.String(), callBuilder.String())

	fmt.Println(query)
	req := graphql.NewRequest(query)

	for i, tx := range txs {
		req.Var(fmt.Sprintf("mutData%d", i), tx)
	}

	var res map[string]string

	err := client.Run(context.Background(), req, &res)
	if err != nil {
		return nil, fmt.Errorf("client run 에러 %w", err)
	}

	return res, nil
}
