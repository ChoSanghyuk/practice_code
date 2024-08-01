package graphql

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/machinebox/graphql"
)

const BesuLimit = 190

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
	var s int = 0
	var e int = min(len(callDatas), BesuLimit)

loop:
	var varBuilder strings.Builder
	var callBuilder strings.Builder

	for i := s; i < e; i++ {

		varBuilder.WriteString(", ")
		varBuilder.WriteString(fmt.Sprintf(callVariableForm, i))

		callBuilder.WriteString(fmt.Sprintf(callForm, i, i))

	}

	query := fmt.Sprintf(multiCallQuery, varBuilder.String(), callBuilder.String())
	// fmt.Println(query)

	req := graphql.NewRequest(query)

	req.Var("blockNumber", bn)

	for i := s; i < e; i++ {
		req.Var(fmt.Sprintf("callData%d", i), callDatas[i])
	}

	err := client.Run(context.Background(), req, &res)
	if err != nil {
		return MultiCallsResp{}, fmt.Errorf("client run 에러 %w", err)
	}

	if e < len(callDatas) {
		s += BesuLimit
		e = min(len(callDatas), e+BesuLimit)
		goto loop
	}

	return res, nil
}

func BesuMultiWrite(txs []string) (map[string]string, error) {

	var res map[string]string
	var s int = 0
	var e int = min(len(txs), BesuLimit)

loop:
	var varBuilder strings.Builder
	var callBuilder strings.Builder

	for i := s; i < e; i++ {
		varBuilder.WriteString(fmt.Sprintf(mutVariableForm, i))
		callBuilder.WriteString(fmt.Sprintf(mutForm, i, i))

		if i < e-1 {
			varBuilder.WriteString(", ")
			callBuilder.WriteString(", ")
		}
	}

	query := fmt.Sprintf(multiMutQuery, varBuilder.String(), callBuilder.String())
	req := graphql.NewRequest(query)

	for i := s; i < e; i++ {
		req.Var(fmt.Sprintf("mutData%d", i), txs[i])
	}

	err := client.Run(context.Background(), req, &res)
	if err != nil {
		return nil, fmt.Errorf("client run 에러 %w", err)
	}

	if e < len(txs) {
		s += BesuLimit
		e = min(len(txs), e+BesuLimit)
		time.Sleep(3 * time.Second)
		goto loop
	}

	return res, nil
}
