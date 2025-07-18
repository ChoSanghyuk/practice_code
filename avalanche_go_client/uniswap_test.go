package avalanche_go

import (
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestUniswapClient(t *testing.T) {

	pk := os.Getenv("PK")
	uniswapClient, err := NewUniswapClient(pk)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Swap_WavaxToUsdc", func(t *testing.T) {

		tokenIn := common.HexToAddress("0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7")
		tokenOut := common.HexToAddress("0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e")
		amountIn := big.NewInt(1e17)
		amountOutMin := big.NewInt(0)

		tx, err := uniswapClient.Swap(tokenIn, tokenOut, amountIn, amountOutMin)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(tx)
	})

	t.Run("Swap_UsdcToWavax", func(t *testing.T) {

		tokenIn := common.HexToAddress("0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e")
		tokenOut := common.HexToAddress("0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7")
		amountIn := big.NewInt(1e6)
		amountOutMin := big.NewInt(0)

		tx, err := uniswapClient.Swap(tokenIn, tokenOut, amountIn, amountOutMin)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(tx)
	})

	t.Run("NativeSwap", func(t *testing.T) {

		// tokenIn := common.HexToAddress("0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7")
		tokenOut := common.HexToAddress("0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e")
		amountIn := big.NewInt(1e17)
		amountOutMin := big.NewInt(0)

		tx, err := uniswapClient.SwapNativeForToken(tokenOut, amountIn, amountOutMin)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(tx)
	})
}
