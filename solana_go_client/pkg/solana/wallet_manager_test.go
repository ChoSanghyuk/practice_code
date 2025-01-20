package solana

import (
	"context"
	"log"
	"sync"
	"testing"

	"github.com/gagliardetto/solana-go"
)

type SolMock struct {
}

func (m *SolMock) CreateAccount(ctx context.Context) *solana.Wallet {
	return solana.NewWallet()
}

func (m *SolMock) CreateAccountWithFaucet(ctx context.Context, solAmount uint64) (*solana.Wallet, error) {
	wallet := solana.NewWallet()
	return wallet, nil
}

func (m *SolMock) WalletFromPK(pk string) (*solana.Wallet, error) {
	return solana.WalletFromPrivateKeyBase58(pk)
}

func TestInitAccountManager(t *testing.T) {

	sm := &SolMock{}

	wm, err := NewWalletManager("../../wallets.txt", sm, WalletManagerConfig{
		N: 10,
		M: 10,
	})
	if err != nil {
		log.Fatalf("Error creating account manager: %v", err)
	}
	log.Printf("%v", wm.AllAddress())

}
func TestAccountManager(t *testing.T) {

	sm := &SolMock{}

	am, err := NewWalletManager("", sm, WalletManagerConfig{
		N: 1000,
		M: 10,
	})
	if err != nil {
		log.Fatalf("Error creating account manager: %v", err)
	}
	t.Run("Lock test", func(t *testing.T) {

		var wg sync.WaitGroup
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				log.Printf("user: %s",
					// am.NextMintWallet().PublicKey(),
					am.NextUserWallet().PublicKey(),
				)
			}()
		}
		wg.Wait()
	})

	t.Run("Pair Query test", func(t *testing.T) {

		var wg sync.WaitGroup

		m := make(map[string]string)
		n := am.MintN()
		for i := 0; i < n; i++ {
			m[am.mintAcQ.wallets[i].PublicKey().String()] = am.initAcQ.wallets[i].PublicKey().String()
		}

		log.Println("Let's Start")
		for i := 0; i < 10000000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				w1, w2 := am.NextMintInitWallet()
				// log.Printf("mintAc: %s, initAc: %s", w1.PublicKey(), w2.PublicKey())
				if m[w1.PublicKey().String()] != w2.PublicKey().String() {
					log.Fatalf("비이상!!!! %s, %s", w1.PublicKey(), w2.PublicKey())
				}

			}()
		}
		wg.Wait()
	})

}
