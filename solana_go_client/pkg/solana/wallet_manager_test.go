package solana

import (
	"log"
	"sync"
	"testing"
)

func TestAccountManager(t *testing.T) {

	sm, err := newSolanaManager()
	if err != nil {
		log.Fatalf("Error creating solana manager: %v", err)
	}

	am, err := NewWalletManager(sm, WalletManagerConfig{
		N: 10,
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
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				w1, w2 := am.NextMintInitWallet()
				log.Printf("mintAc: %s, initAc: %s", w1.PublicKey(), w2.PublicKey())
			}()
		}
		wg.Wait()
	})

}
