package avalanche_go

import (
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// 0.006967999366309176
func TestTraderJoeClient(t *testing.T) {
	traderJoeClient, err := NewTraderJoeClient("")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("MyBalanceOf", func(t *testing.T) {
		_, err := traderJoeClient.activeId()
		if err != nil {
			t.Fatal(err)
		}

		myAvax, myUsdc, err := traderJoeClient.MyBalanceOf(big.NewInt(8363921))
		if err != nil {
			t.Fatal(err)
		}
		t.Log(myAvax)
		t.Log(myUsdc)
	})

	t.Run("MyBalanceOfBatch", func(t *testing.T) {
		ids, amounts, amountX, amountY, err := traderJoeClient.MyBalanceOfBatch(big.NewInt(8363922), 200)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(ids)
		t.Log(amounts)
		t.Log(amountX)
		t.Log(amountY)
	})

	t.Run("balance", func(t *testing.T) {
		// bin, err := traderJoeClient.activeId()
		// if err != nil {
		// 	t.Fatal(err)
		// }
		// binId := bin.Int64()

		balance, err := traderJoeClient.balance(big.NewInt(8364164))
		if err != nil {
			t.Fatal(err)
		}
		t.Log(balance)
	})

	t.Run("balanceOfBatch", func(t *testing.T) {
		_, err := traderJoeClient.activeId()
		if err != nil {
			t.Fatal(err)
		}

		balanceOfBatch, err := traderJoeClient.balanceOfBatch(big.NewInt(8363922), 200)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(balanceOfBatch)
	})

	t.Run("activeId", func(t *testing.T) {
		activeId, err := traderJoeClient.activeId()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(activeId)
	})

	t.Run("getTokenX", func(t *testing.T) {
		tokenX, err := traderJoeClient.getTokenX()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(tokenX) //WAVAX
	})

	t.Run("getTokenY", func(t *testing.T) {
		tokenY, err := traderJoeClient.getTokenY()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(tokenY) //USDC
	})

	t.Run("isApprovedForAll", func(t *testing.T) {
		isApprovedForAll, err := traderJoeClient.isApprovedForAll()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(isApprovedForAll)
	})

	t.Run("getTargetId", func(t *testing.T) {
		targetId, err := traderJoeClient.getTargetId(16.8512, 24.51376935)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(targetId)
	})

	t.Run("getBinStep", func(t *testing.T) {
		binStep, err := traderJoeClient.getBinStep()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(binStep)
	})

	t.Run("getPriceFromId", func(t *testing.T) {
		price, err := traderJoeClient.getPriceFromId(big.NewInt(8363922))
		if err != nil {
			t.Fatal(err)
		}
		t.Log(price)
	})

	t.Run("GetSwapIn", func(t *testing.T) {
		amountIn, amountOutLeft, fee, err := traderJoeClient.GetSwapIn(common.HexToAddress("0x864d4e5Ee7318e97483DB7EB0912E09F161516EA"), big.NewInt(17522755))
		if err != nil {
			t.Fatal(err)
		}
		t.Log(amountIn)
		t.Log(amountOutLeft)
		t.Log(fee)
	})
}

func TestTemp(t *testing.T) {
	pk := os.Getenv("PK")
	fmt.Println(pk)
}

func TestRemoveLiquidity(t *testing.T) {

	pk := os.Getenv("PK")
	traderJoeClient, err := NewTraderJoeClient(pk)
	if err != nil {
		t.Fatal(err)
	}

	// t.Run("getAddress", func(t *testing.T) {
	// 	// pk to address
	// 	pkBytes, err := crypto.HexToECDSA(pk[2:])
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	address := crypto.PubkeyToAddress(pkBytes.PublicKey)
	// 	t.Log(address.Hex())
	// })

	var ids, amounts []*big.Int
	var amountX, amountY *big.Int
	t.Run("MyBalanceOfBatch", func(t *testing.T) {
		ids, amounts, amountX, amountY, err = traderJoeClient.MyBalanceOfBatch(big.NewInt(8364164), 200)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(ids)
		// t.Log(amounts)
		t.Log(amountX)
		t.Log(amountY)
	})

	t.Run("RemoveLiquidity", func(t *testing.T) {

		hash, err := traderJoeClient.RemoveLiquidity(ids, amounts, amountX, amountY, true)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(hash)
	})
}

func TestSwapNATIVEForExactTokens(t *testing.T) {

	pk := os.Getenv("PK")
	traderJoeClient, err := NewTraderJoeClient(pk)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("SwapNATIVEForExactTokens", func(t *testing.T) {
		hash, err := traderJoeClient.SwapNATIVEForExactTokens(
			big.NewInt(17663755),
			common.HexToAddress("0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E"),
			big.NewInt(time.Now().Add(time.Hour*1).Unix()),
		)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(hash)
	})
}

func TestSwapExactTokensForNATIVE(t *testing.T) {

	pk := os.Getenv("PK")
	traderJoeClient, err := NewTraderJoeClient(pk)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("SwapExactTokensForNATIVE", func(t *testing.T) {
		hash, err := traderJoeClient.SwapExactTokensForNATIVE(
			big.NewInt(100),
			big.NewInt(0),
			common.HexToAddress("0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E"),
			big.NewInt(time.Now().Add(time.Hour*1).Unix()),
		)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(hash)
	})
}

func TestSwapExactTokensForTokens(t *testing.T) {

	pk := os.Getenv("PK")
	traderJoeClient, err := NewTraderJoeClient(pk)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("SwapExactTokensForTokens", func(t *testing.T) {
		hash, err := traderJoeClient.SwapExactTokensForTokens(
			big.NewInt(100),
			big.NewInt(0),
			common.HexToAddress("0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E"),
			big.NewInt(time.Now().Add(time.Hour*1).Unix()),
		)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(hash)
	})
}
