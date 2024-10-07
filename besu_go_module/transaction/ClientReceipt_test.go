package transaction

import (
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var clientReceiptContract *ClientReceipt

func init() {

	hashAddr := "0xE0323ECd69dD90DF1628Ce855D4f6299B54a5F5F"
	var addr common.Address
	var err error

	if hashAddr == "" {
		auth, _ := CreateTxOpts(privateKey1, nil)
		client, err := ethclient.Dial(url)
		if err != nil {
			log.Panic(err)
		}

		addr, _, clientReceiptContract, err = DeployClientReceipt(auth, client)
		if err != nil {
			log.Panic(err)
		}
	} else {
		addr := common.HexToAddress(hashAddr)
		clientReceiptContract, err = NewClientReceipt(addr, client)
		if err != nil {
			log.Panic(err)
		}
	}

	log.Printf("컨트랙트 주소 : %s", addr)
}

func TestEmitEvent(t *testing.T) {

	auth, err := CreateTxOpts(privateKey1, nil)
	if err != nil {
		t.Error(err)
	}

	tx, err := clientReceiptContract.Deposit(auth, [32]byte{32})
	if err != nil {
		t.Error(err)
	}
	t.Log(tx.Hash())

}
