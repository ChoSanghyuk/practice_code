package besu_graphql

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/machinebox/graphql"
)

var multiCallQuery = `
	query getCall($blockNumber: Long%s) {
		block(number: $blockNumber){
			%s
		}
	
	fragment callFields on call {
		data, 
		gasUsed, 
		statu
	}
`

var callForm = `
	call%d : call(data: $callData%d){
		...callFields
	}
`
var callVariableForm = "$callData%d: CallData!"

type BlockCallsResp struct {
	Block struct {
		resps []CallResp
	} `json:"block"`
}

func BesuMultiCall(bn *big.Int, callDatas []Call) (BlockCallsResp, error) {

	var varBuilder strings.Builder
	var callBuilder strings.Builder

	for i := range len(callDatas) {

		varBuilder.WriteString(", ")
		varBuilder.WriteString(fmt.Sprintf(callVariableForm, i))

		callBuilder.WriteString(fmt.Sprintf(callForm, i, i))

	}

	query := fmt.Sprintf(multiCallQuery, varBuilder.String(), callBuilder.String())

	fmt.Println(query)

	req := graphql.NewRequest(query)

	req.Var("blockNumber", bn)

	for i, call := range callDatas {
		req.Var(fmt.Sprintf("callData%d", i), call)
	}

	var res BlockCallsResp

	err := client.Run(context.Background(), req, &res)
	if err != nil {
		return BlockCallsResp{}, fmt.Errorf("client run 에러 %w", err)
	}

	return res, nil
}
