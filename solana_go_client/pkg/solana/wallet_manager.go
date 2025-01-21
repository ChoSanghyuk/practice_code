package solana

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gagliardetto/solana-go"
)

type WalletManager struct {
	mintAcQ   WalletCircleQueue
	initAcQ   WalletCircleQueue
	targetAcQ WalletCircleQueue
}

type SolMI interface {
	CreateAccount(ctx context.Context) *solana.Wallet
	CreateAccountWithFaucet(ctx context.Context, solAmount uint64) (*solana.Wallet, error)
	WalletFromPK(pk string) (*solana.Wallet, error)
}

// todo. account manager config로 바꿔서 작성
func NewWalletManager(path string, sm SolMI, conf WalletManagerConfig) (*WalletManager, error) {

	n := int(conf.N)
	m := int(conf.M)

	if !fileExists(path) {
		err := genWallets(path, sm)
		if err != nil {
			return nil, fmt.Errorf("generate wallet 중 오류 %w", err)
		}
	}

	wm, err := loadWallet(path, n, m, sm)
	if err != nil {
		return nil, fmt.Errorf("load wallet 중 오류 %w", err)
	}

	return wm, nil
}

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

func (am *WalletManager) AllAddress() map[string][]string {

	m := make(map[string][]string)

	mintAc := make([]string, len(am.mintAcQ.wallets))
	for i := 0; i < len(mintAc); i++ {
		mintAc[i] = am.mintAcQ.wallets[i].PublicKey().String()
	}
	m["mintAc"] = mintAc

	ownerAc := make([]string, len(am.initAcQ.wallets))
	for i := 0; i < len(ownerAc); i++ {
		ownerAc[i] = am.initAcQ.wallets[i].PublicKey().String()
	}
	m["initHolderAc"] = ownerAc

	trgtAc := make([]string, len(am.targetAcQ.wallets))
	for i := 0; i < len(trgtAc); i++ {
		trgtAc[i] = am.targetAcQ.wallets[i].PublicKey().String()
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

/************************************************** load **************************************************/

const totalNumber = 1000

const mint = "mint"
const initHolder = "init_holder"
const target = "target"

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func genWallets(path string, solm SolMI) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	write := func(buf string) {
		if err != nil {
			return
		}
		_, err = file.Write([]byte(buf + "\n"))
	}

	var wallet *solana.Wallet
	for i := 0; i < totalNumber; i++ {
		wallet = solm.CreateAccount(context.Background())
		write(fmt.Sprintf("%s:%s:%s", mint, wallet.PrivateKey, wallet.PublicKey())) // todo key값 변수화
	}
	for i := 0; i < totalNumber; i++ {
		wallet, err = solm.CreateAccountWithFaucet(context.Background(), 100)
		write(fmt.Sprintf("%s:%s:%s", initHolder, wallet.PrivateKey, wallet.PublicKey()))
	}
	if err != nil {
		return fmt.Errorf("wallet 초기화 중 실패. %w", err)
	}

	for i := 0; i < totalNumber; i++ {
		wallet := solm.CreateAccount(context.Background())
		write(fmt.Sprintf("%s:%s:%s", target, wallet.PrivateKey, wallet.PublicKey()))
	}

	time.Sleep(10 * time.Second)
	return nil
}

func loadWallet(path string, n, m int, solm SolMI) (*WalletManager, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	mintAcs := make([]*solana.Wallet, 0, n)
	initAcs := make([]*solana.Wallet, 0, n)
	trgtAcs := make([]*solana.Wallet, 0, m)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, ":")
		switch words[0] {
		case mint:
			if len(mintAcs) < n {
				wallet, _ := solm.WalletFromPK(words[1])
				mintAcs = append(mintAcs, wallet)
			}
		case initHolder:
			if len(initAcs) < n {
				wallet, _ := solm.WalletFromPK(words[1])
				initAcs = append(initAcs, wallet)
			}
		case target:
			if len(trgtAcs) < m {
				wallet, _ := solm.WalletFromPK(words[1])
				trgtAcs = append(trgtAcs, wallet)
			}
		}
	}
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &WalletManager{
		mintAcQ: WalletCircleQueue{
			wallets: mintAcs,
			idx:     0,
		},
		initAcQ: WalletCircleQueue{
			wallets: initAcs,
			idx:     0,
		},
		targetAcQ: WalletCircleQueue{
			wallets: trgtAcs,
			idx:     0,
		},
	}, nil
}
