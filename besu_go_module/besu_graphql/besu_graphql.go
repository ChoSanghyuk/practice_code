package besu_graphql

import (
	"context"
	"fmt"
	"go_module/config"
	"math/big"

	"github.com/machinebox/graphql"
)

type Call struct {
	To   string `json:"to"`
	Data string `json:"data"`
}

type BlockCallResp struct {
	Block struct {
		Call1 CallResp `json:"call1"`
	} `json:"block"`
}

type CallResp struct {
	Data    string `json:"data"`
	GasUsed string `json:"gasUsed"`
	Status  string `json:"status"`
}

var callQuery = `
	query getCall($blockNumber: Long, $callData: CallData!) {
		block(number: $blockNumber){
			call1 : call(data: $callData){
						data, 
						gasUsed, 
						status
					}
			}
		}
`
var mutQuery = `
	mutation($mutData: Bytes!) {
		t1: sendRawTransaction(data: $mutData)
		}
`

var client = graphql.NewClient(config.Config.Network.GraphqlUrl)

func BesuCall(bn *big.Int, callData Call) (BlockCallResp, error) {

	req := graphql.NewRequest(callQuery)

	req.Var("blockNumber", bn)
	req.Var("callData", callData)

	var res BlockCallResp

	err := client.Run(context.Background(), req, &res)
	if err != nil {
		return BlockCallResp{}, fmt.Errorf("client run 에러 %w", err)
	}

	return res, nil
}

func BesuMutliCall(bn *big.Int, callDatas []Call) (BlockCallResp, error) {

	req := graphql.NewRequest(callQuery)

	req.Var("blockNumber", bn)
	req.Var("callData", callDatas)

	var res BlockCallResp

	err := client.Run(context.Background(), req, &res)
	if err != nil {
		return BlockCallResp{}, fmt.Errorf("client run 에러 %w", err)
	}

	return res, nil
}

func BesuWrite(tx string) (string, error) {

	req := graphql.NewRequest(mutQuery)

	req.Var("mutData", tx)

	var txHash string

	err := client.Run(context.Background(), req, &txHash)
	if err != nil {
		return "", fmt.Errorf("client run 에러 %w", err)
	}

	return txHash, nil

}
