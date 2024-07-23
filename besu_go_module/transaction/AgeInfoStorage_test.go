package transaction

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestDeployAgeInfoStorage(t *testing.T) {

	auth := CreateAuth(privateKey1)

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

func TestWriteAndCallAgnInfoStorage(t *testing.T) {
	contract, err := NewAgeInfoStrage(common.HexToAddress("0x42699A7612A82f1d9C36148af9C77354759b210b"), client)
	if err != nil {
		t.Error(err)
	}

	key, err := crypto.HexToECDSA(privateKey1)
	if err != nil {
		t.Error(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		t.Error(err)
	}

	tx, err := contract.AgeInfoStrageTransactor.SetAge(auth, "Jo", big.NewInt(30))
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)
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

	assert.Equal(t, big.NewInt(30), age)
}
