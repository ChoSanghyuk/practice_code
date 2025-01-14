package solana

import (
	"context"
	"sync"

	"github.com/gagliardetto/solana-go"
)

type WalletManager struct {
	mintAcQ   WalletCircleQueue
	initAcQ   WalletCircleQueue
	targetAcQ WalletCircleQueue
}

// todo. account manager config로 바꿔서 작성
func NewWalletManager(sm *SolanaManager, conf WalletManagerConfig) (*WalletManager, error) {

	mintN := int(conf.N)
	initN := int(conf.N)
	trgtN := int(conf.M)

	mintAc := make([]*solana.Wallet, mintN)
	for i := 0; i < mintN; i++ {
		mintAc[i] = sm.CreateAccount(context.Background())
	}

	initAc := make([]*solana.Wallet, initN)
	for i := 0; i < initN; i++ {
		account, err := sm.CreateAccountWithFaucet(context.Background(), 1000)
		if err != nil {
			return nil, err
		}
		initAc[i] = account
	}

	trgtAc := make([]*solana.Wallet, trgtN)
	for i := 0; i < trgtN; i++ {
		trgtAc[i] = sm.CreateAccount(context.Background())
	}

	return &WalletManager{
		mintAcQ: WalletCircleQueue{
			wallets: mintAc,
			idx:     0,
		},
		initAcQ: WalletCircleQueue{
			wallets: initAc,
			idx:     0,
		},
		targetAcQ: WalletCircleQueue{
			wallets: trgtAc,
			idx:     0,
		},
	}, nil
}

// func (am *WalletManager) NextMintWallet() *solana.Wallet {
// 	return am.mintAcQ.pop()
// }

func (am *WalletManager) NextUserWallet() *solana.Wallet {
	return am.initAcQ.pop()
}

func (am *WalletManager) NextTrgtWallet() *solana.Wallet {
	return am.targetAcQ.pop()
}

func (wm *WalletManager) NextMintInitWallet() (mintAc *solana.Wallet, initAc *solana.Wallet) {

	mutex.Lock()
	i := &wm.mintAcQ.idx
	mintAc = wm.mintAcQ.wallets[*i]
	initAc = wm.initAcQ.wallets[*i]
	*i++
	if *i == len(wm.mintAcQ.wallets) {
		*i = 0
	}
	mutex.Unlock()
	return mintAc, initAc
}

func (wm *WalletManager) MintN() int {
	return len(wm.mintAcQ.wallets)
}

func (am *WalletManager) NewWallet() *solana.Wallet {
	return solana.NewWallet()
}

func (am *WalletManager) AllWallets() map[string][]string {

	m := make(map[string][]string)

	mintAc := make([]string, len(am.mintAcQ.wallets))
	for i := 0; i < len(mintAc); i++ {
		mintAc[i] = am.mintAcQ.wallets[i].PrivateKey.String()
	}
	m["mintAc"] = mintAc

	ownerAc := make([]string, len(am.initAcQ.wallets))
	for i := 0; i < len(ownerAc); i++ {
		ownerAc[i] = am.initAcQ.wallets[i].PrivateKey.String()
	}
	m["initHolderAc"] = ownerAc

	trgtAc := make([]string, len(am.targetAcQ.wallets))
	for i := 0; i < len(trgtAc); i++ {
		trgtAc[i] = am.targetAcQ.wallets[i].PrivateKey.String()
	}
	m["targetAc"] = trgtAc

	return m
}

type WalletCircleQueue struct {
	wallets []*solana.Wallet
	idx     int
}

var mutex sync.Mutex

func (wq *WalletCircleQueue) pop() *solana.Wallet {

	mutex.Lock()
	rtn := wq.wallets[wq.idx]
	wq.idx++
	if wq.idx == len(wq.wallets) {
		wq.idx = 0
	}
	mutex.Unlock()
	return rtn
}
