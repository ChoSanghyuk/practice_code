package solana

import (
	"context"
	"log"
	"sync"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/test-go/testify/require"
)

func TestForPerformanceTest(t *testing.T) {

	ctx := context.Background()

	sm, err := newSolanaManager()
	require.NoError(t, err)

	pkBase58 := "Tx2y8ztcSXr2vx2jkEAYbtqKwDaceRFtwFRSFCD7cxEzPv4Zp6ZdUb9eNKLB9KwF8mkuSdd9rbZFPUmG9KPzYyH"
	payer, err := solana.WalletFromPrivateKeyBase58(pkBase58)
	require.NoError(t, err)
	log.Println("0. payer: ", payer.PublicKey())

	ownerCount := 1

	mintAccount := solana.NewWallet()
	sig, err := sm.SetMintAccount(ctx, mintAccount, payer, payer.PublicKey(), payer.PublicKey())
	require.NoError(t, err)
	log.Println("1. create spl token sig: ", sig)

	var wg sync.WaitGroup
	for i := 0; i < ownerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			owner := solana.NewWallet()
			amount := 100000000

			sig, ata, err := sm.Mint(ctx, payer, mintAccount.PublicKey(), owner.PublicKey(), payer.PublicKey(), uint64(amount))
			if err != nil {
				log.Println("민팅에러발생", err)
			}
			log.Println("2. ata :", ata)
			log.Println("3. 민팅 sig:", sig)
		}()
	}
	wg.Wait()
}

func TestTransferPerformance(t *testing.T) {
	// sm, err := newSolanaManager()
	// require.NoError(t, err)

	// pkBase58 := "5pJJmuBZcVZtECHpwncuDpJgVB7H4mD8JCTMTasvCcsrRLxSEQNGFwJwnmfEME5CKSDcUS2wGT9k73rKnyNcXBmp"
	// payer, err := solana.WalletFromPrivateKeyBase58(pkBase58)
	// require.NoError(t, err)
}

func TestUnit(t *testing.T) {

	sm, err := newSolanaManager()
	if err != nil {
		log.Fatalf("Error creating solana manager: %v", err)
	}

	payerW, _ := solana.WalletFromPrivateKeyBase58("Tx2y8ztcSXr2vx2jkEAYbtqKwDaceRFtwFRSFCD7cxEzPv4Zp6ZdUb9eNKLB9KwF8mkuSdd9rbZFPUmG9KPzYyH")
	mintW, _ := solana.WalletFromPrivateKeyBase58("2Dv9x4xgj5LVaCSAVTEdLdZwFztQk3QBaq3yZ8GRpciHkrf58MneUmHNdFW7ChCpEFLniiknjY6JZuJU36FPJTBh") //("5DfHivrXhYga5Y1UbQafHuNzTWPzGn31F3Tv5t5h2bQn58cRT3ohxXQUoSxXKUssRuBrdCNJsks5cU3PQfCNrVHZ")
	owner1, _ := solana.WalletFromPrivateKeyBase58("5Zk5gLkpxumRaeprPFGV5x41DZXBgPARBLkBqgPaMbaFUFzhMwDSrGqBKvPcLTLQy1sDripVVvwmrTiDL1cNSHXN")
	owner2, _ := solana.WalletFromPrivateKeyBase58("52fhezV6r62Meszq8zQugSxPryzAs379aMTntv3jz9CcosmAjr6C2BFmY1hUNkAFgKfJDqVk7KsJaYUVF2zhgxj5")

	t.Run("account creation", func(t *testing.T) {
		// payer, err := sm.CreateAccountWithFaucet(context.Background(), 1000000)
		wallet, err := sm.CreateAccountWithFaucet(context.Background(), 1000)
		require.NoError(t, err)
		log.Printf("%v", wallet.PrivateKey)
	})
	// 54YCaNP5JfZuZb44qxPsjRjEghQDnXCMDMHvKR2484oaM2tgiufDrzRpNsHv7pR3HDrs3jCTWHBv5XeXCG3indn3

	t.Run("account fund", func(t *testing.T) {
		_, err := solana.WalletFromPrivateKeyBase58("32JsF4FWwjrwdiEGfoFvTCGpZDEckeqkTq7WmVF1AH2rDLtHy345KbmYZYvDvNmyaesgjJVRs8jXbMNhkUobk3Hr")
		require.NoError(t, err)

		tx, err := sm.RequestAirdrop(context.Background(), payerW, 100)
		require.NoError(t, err)
		log.Println(tx)
	})

	t.Run("account balance query", func(t *testing.T) {
		wallet, err := solana.WalletFromPrivateKeyBase58("5MBFFymun7cJoACtJ1gERkhAwcENiYsQT98rNDHj7XEEcSDbfMrzczUXkgx8EJZ7zZ4bK8o1NirKXPu4TkwX2MG5")
		require.NoError(t, err)

		pubKey := wallet.PublicKey()
		log.Printf("public key: %s", pubKey)

		balance, err := sm.Balance(context.Background(), pubKey)
		require.NoError(t, err)
		log.Println(balance, "balance")

	})

	t.Run("account balance query by address", func(t *testing.T) {

		pubKey := solana.MustPublicKeyFromBase58("7JTQ3nQzTwwW9F8NZZiP23PkK4SNXoNuGQAuC68WVVVp")
		log.Printf("public key: %s", pubKey)

		balance, err := sm.Balance(context.Background(), pubKey)
		require.NoError(t, err)
		log.Println(balance, "balance")

	})

	t.Run("tx query", func(t *testing.T) {
		txSig := solana.MustSignatureFromBase58("5wDEzcrvgFm7XQcZYp5PKCtmK5cNL55BJtdiVkf8sB551VvRwFn9dH9YCUxL9PS2Got9Kmiihx84UgQLFYXWzEzE")

		out, err := sm.rpc.GetTransaction(
			context.Background(),
			txSig,
			&rpc.GetTransactionOpts{
				Encoding: solana.EncodingBase64,
			},
		)
		require.NoError(t, err)
		log.Printf("%v", out)
	})

	t.Run("get token holder", func(t *testing.T) {
		account := "ASjUrtwV22BAVFL4w6bzXTiDFD9FEyMzmi7faKcEm1Jw" // "Goe4oBAW5tyioXYWohfKSFyiuyZ7nYLgbZ6wAYeD2VqW"
		pubKey := solana.MustPublicKeyFromBase58(account)

		holdersConfirmed, err := sm.GetTokenHolders(context.Background(), pubKey, rpc.CommitmentConfirmed)
		require.NoError(t, err)
		log.Println("unique holder count1:", len(holdersConfirmed))

		holdersFinalized, err := sm.GetTokenHolders(context.Background(), pubKey, rpc.CommitmentFinalized)
		require.NoError(t, err)
		log.Println("unique holder count:", len(holdersFinalized))
	})

	t.Run("token holder info", func(t *testing.T) {

		ata, _, _ := solana.FindAssociatedTokenAddress(owner1.PublicKey(), mintW.PublicKey())
		accountInfo, err := sm.rpc.GetAccountInfo(context.Background(), ata)
		if err == nil && accountInfo != nil {
			log.Println("ATA already exists:", ata)
			spew.Dump(accountInfo)
		}

	})

	t.Run("get recent blockhash", func(t *testing.T) {
		recent, err := sm.latestBlockHash(context.Background())
		require.NoError(t, err)
		log.Println(recent, "recent")
	})

	t.Run("airdrop", func(t *testing.T) {
		account := solana.NewWallet()

		_, err := sm.RequestAirdrop(context.Background(), account, 10)
		require.NoError(t, err)

		balance, err := sm.Balance(context.Background(), account.PublicKey())
		require.NoError(t, err)
		log.Println(account.PublicKey().String())
		log.Println(balance, "balance")
	})

	t.Run("balance", func(t *testing.T) {
		// account := "2GXNA7Vt1hRsre9HT7L3CRpiFznZwbc2mAapAbBVAkZf"
		// pubKey := solana.MustPublicKeyFromBase58(account)

		balance, err := sm.Balance(context.Background(), payerW.PublicKey())
		require.NoError(t, err)
		log.Println(balance, "balance")
	})

	t.Run("new token", func(t *testing.T) {

		// payer, err := solana.WalletFromPrivateKeyBase58("3KPVVBRV3YZMJ67pdg7QgB2kDzJ42Au4mKEShYNC8GP7HM7BdFNPzuSJ3GLKtv3D81NCj2KpZmJaPKcuTj2zUvGR")
		// require.NoError(t, err)
		token := solana.NewWallet()
		sig, err := sm.SetMintAccount(
			context.Background(),
			token,
			payerW,
			payerW.PublicKey(),
			payerW.PublicKey(),
		)

		require.NoError(t, err)
		log.Printf("%v\n", sig)
		log.Printf("%v\n", token)
	})

	t.Run("token query", func(t *testing.T) {

		mintW, _ = solana.WalletFromPrivateKeyBase58("2nat5jYGS6KDniyMj4NSdrijAEX2qXz9D3B8EevSo5ghph7qiCCVm9NTsM6i8NfANWcqysxGVjYANkHWuw7wfJrk")
		owner2, _ = solana.WalletFromPrivateKeyBase58("5Wd6UcGMQeX9LCjiWjCHzrDo15Mdmv25ZasbhKkG76geVC9hk7sJ2bvd2eoZvARrLD2uYS1RBJRFor8G9jVnQ6Cc")
		balance, err := sm.TokenBalance(
			context.Background(),
			mintW.PublicKey(),
			owner2.PublicKey(),
		)
		require.NoError(t, err)
		log.Printf("%s\n", *balance)
	})

	t.Run("minting", func(t *testing.T) {
		mintW, _ = solana.WalletFromPrivateKeyBase58("gn1KYWJBez7gsQ5J54BKqbtVz6D6AgYAzanbLXqfGAdskYmNphRzTWBsrPv9YebqcPUfBeV9TQtuD1JJYn1gK5m")
		owner := owner1
		var amount uint64 = 100000000

		log.Println("owner", owner.PrivateKey)
		sig, ata, err := sm.Mint(
			context.Background(),
			payerW,
			mintW.PublicKey(),
			owner.PublicKey(),
			payerW.PublicKey(),
			amount,
		)
		require.NoError(t, err)

		log.Printf("sig: %s\n", sig.String())
		log.Printf("ata: %s\n", ata.String())

	})

	t.Run("new token with minting", func(t *testing.T) {
		sig, ata, err := sm.SetMintAccountAndMint(
			context.Background(),
			payerW,
			solana.NewWallet(),
			owner1.PublicKey(),
			payerW.PublicKey(),
			payerW.PublicKey(),
			10000,
		)
		require.NoError(t, err)

		log.Printf("sig: %s\n", sig.String())
		log.Printf("ata: %s\n", ata.String())
	})

	t.Run("token account creation", func(t *testing.T) {

		sig, err := sm.CreateAta(context.Background(), payerW, owner1, mintW.PublicKey())
		require.NoError(t, err)

		log.Printf("sig: %s\n", sig.String())
	})

	t.Run("get token supply", func(t *testing.T) {

		out, err := sm.rpc.GetTokenSupply(
			context.Background(),
			mintW.PublicKey(),
			rpc.CommitmentFinalized,
		)
		require.NoError(t, err)
		spew.Dump(out)
	})

	t.Run("token transfer", func(t *testing.T) {

		owner2 := solana.NewWallet()
		log.Printf("pk: %s\n", owner2.PrivateKey.String())
		sig, err := sm.TransferToken(context.Background(), payerW, owner1, owner2, mintW.PublicKey(), 100)
		require.NoError(t, err)
		log.Printf("sig: %s\n", sig.String())
	})

	t.Run("sig verify - cache", func(t *testing.T) {
		sig := solana.MustSignatureFromBase58("2MtZD7v5sv8eruBxWjNu7bddfCVUiaWb8AgDMxNWCt6pPXcu87M4pM9rsoASBQsEbA5GPGd5hDwrDbUcknEiGboA")
		// sig, err := solana.SignatureFromBase58("2mpeu4DHc8jK4tJ7g7QxErS43pP8wo3wh4CaH3hDP2W2RTK5RF2mxsw1DkroVuuUSjs5Am7bPyuSTkB1sxUaktoB")
		require.NoError(t, err)

		out, err := sm.rpc.GetSignatureStatuses(
			context.Background(),
			true,
			sig,
		)
		require.NoError(t, err)
		spew.Dump(out)

		// d, err := json.MarshalIndent(out.Value, "", "\t")
		// if err != nil {
		// 	log.Println("error", err)
		// }
		// log.Println(string(d))
	})

	t.Run("sig verify - hist", func(t *testing.T) {
		// sig := solana.MustSignatureFromBase58("PQbSzj33m1Zh87pH8UAEqBGZvnSDwMZac6iZP87yv4pQ9RN4B3HW2kuUY5ESsh6quv67P8M7Bhmhefrs5AtUeDU")
		sig, err := solana.SignatureFromBase58("2Vst655kYZNQ9KyYWdXroe24YKtqMdrL98GKfcD96dM9B5R1s3t5FRezd1oZjnNs41LpBUG4gyr8BXRvmdRXV1Sb")
		require.NoError(t, err)

		out, err := sm.rpc.GetTransaction(
			context.Background(),
			sig,
			&rpc.GetTransactionOpts{
				Encoding:   solana.EncodingJSONParsed,
				Commitment: rpc.CommitmentFinalized,
			},
		)
		require.NoError(t, err)
		spew.Dump(out)

	})

}

func newSolanaManager() (*SolanaManager, error) {
	return NewSolanaManager(SolManagerConfig{
		// RPCURL: "https://bold-special-snow.solana-devnet.quiknode.pro/df682b672aa3f617aa87c8f3d80e16fe4f06e1ce",
		RPCURL: "http://124.50.46.159:50001", //"http://localhost:8899",
		// WSURL: "wss://bold-special-snow.solana-devnet.quiknode.pro/df682b672aa3f617aa87c8f3d80e16fe4f06e1ce",
		WSURL:      "ws://124.50.46.159:50002",
		Commitment: rpc.CommitmentFinalized,
		IsSync:     true,
		// RPCURL: "http://118.37.71.15:8899",
	})
}
