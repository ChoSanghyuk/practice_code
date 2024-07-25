package transaction

import (
	"go_module/config"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

var privateKey1 = config.Config.Accounts["account1"].PrivateKey[2:]

func TestDeployAgeInfoStorage(t *testing.T) {

	auth, _ := CreateAuth(privateKey1)

	client, err := ethclient.Dial(url)
	if err != nil {
		t.Error(err)
	}

	addr, _, _, err := DeployAgeInfoStrage(auth, client)
	if err != nil {
		t.Error(err)
	}

	t.Logf("컨트랙트 주소 : %s", addr)
}

func TestWriteAgnInfoStorage(t *testing.T) {
	contract, err := NewAgeInfoStrage(common.HexToAddress("0x42699A7612A82f1d9C36148af9C77354759b210b"), client)
	if err != nil {
		t.Error(err)
	}

	auth, err := CreateAuth(privateKey1)
	if err != nil {
		t.Error(err)
	}

	tx, err := contract.AgeInfoStrageTransactor.SetAge(auth, "Jo", big.NewInt(50))
	if err != nil {
		t.Error(err)
	}
	t.Log(tx.Hash())
}

func TestCallAgnInfoStorage(t *testing.T) {
	contract, err := NewAgeInfoStrage(common.HexToAddress("0x42699A7612A82f1d9C36148af9C77354759b210b"), client)
	if err != nil {
		t.Error(err)
	}

	age, err := contract.AgeInfoStrageCaller.GetAge(nil, "Jo")
	if err != nil {
		t.Error(err)
	}

	t.Log(age)

	assert.Equal(t, big.NewInt(50), age)
}
