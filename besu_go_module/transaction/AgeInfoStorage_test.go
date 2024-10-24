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

	auth, _ := CreateTxOpts(privateKey1, nil)

	client, err := ethclient.Dial(url)
	if err != nil {
		t.Error(err)
	}

	addr, tx, _, err := DeployAgeInfoStrage(auth, client)
	if err != nil {
		t.Error(err)
	}

	t.Logf("트랜잭션 해시 : %s", tx.Hash())
	t.Logf("컨트랙트 주소 : %s", addr)
}

func TestWriteAgnInfoStorage(t *testing.T) {
	contract, err := NewAgeInfoStrage(common.HexToAddress("0x38DB5cc1954Cf38fe65849C9de7D38a2Eac4d79c"), client)
	if err != nil {
		t.Error(err)
	}

	auth, err := CreateTxOpts(privateKey1, nil)
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
	contract, err := NewAgeInfoStrage(common.HexToAddress("0x38DB5cc1954Cf38fe65849C9de7D38a2Eac4d79c"), client)
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
