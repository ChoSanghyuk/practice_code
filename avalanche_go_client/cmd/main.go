package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
For C-Chain API, the URL is https://api.avax.network/ext/bc/C/rpc.
For X-Chain API, the URL is https://api.avax.network/ext/bc/X.
For P-Chain API, the URL is https://api.avax.network/ext/bc/P.
Note: on Fuji Testnet, use https://api.avax-test.network/ instead of https://api.avax.network/.
*/

func main() {
	// Public Avalanche C-Chain RPC endpoint (Mainnet)
	rpcURL := "https://api.avax.network/ext/bc/C/rpc"

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to Avalanche C-Chain: %v", err)
	}

	balance, _ := client.BalanceAt(context.Background(), common.HexToAddress("0xb4dd4fb3D4bCED984cce972991fB100488b59223"), nil)
	fmt.Printf("Balance: %s\n", balance)

	// Get latest block number
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("Failed to get block number: %v", err)
	}

	fmt.Printf("Latest block number: %d\n", blockNumber)
}
