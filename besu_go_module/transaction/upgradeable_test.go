package transaction

import (
	"context"
	"fmt"
	"strings"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
== 테스트 대상 파일 ==
MyProxy.go
MyImplementV1.go
MyImplementV2.go
*/

func TestInitialSetting(t *testing.T) {

	client, err := ethclient.Dial(url)
	if err != nil {
		t.Error(err)
	}

	// 1. 구현 컨트랙트 배포
	auth1, _ := CreateTxOpts(privateKey1, nil)

	implAddr, tx1, _, err := DeployMyImplementV1(auth1, client)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(tx1.Hash())

	// implAbi, err := MyImplementV1MetaData.GetAbi()
	// if err != nil {
	// 	t.Error(err)
	// }

	// input, err := implAbi.Pack("initialize")
	// if err != nil {
	// 	t.Error(err)
	// }

	// // t.Log(parsed)
	// // 2. 프록시 컨트랙트 배포 (프록시-구현 연결)
	// auth2, _ := CreateTxOpts(privateKey1, nil)
	// proxyAddr, _, _, err := DeployMyProxy(auth2, client, implAddr, input)
	// if err != nil {
	// 	t.Error(err)
	// }

	t.Logf("\n프록시 주소 : %s\n구현 컨트랙트 주소 : %s", "dfd", implAddr)
}

/*
프록시 주소 : 0xe135783649BfA7c9c4c6F8E528C7f56166efC8a6
구현 컨트랙트 주소 : 0xfeae27388A65eE984F452f86efFEd42AaBD438FD
*/

func TestCallProxy(t *testing.T) {

	proxyAddr := common.HexToAddress("0xfeae27388A65eE984F452f86efFEd42AaBD438FD")

	implAbi, err := MyImplementV1MetaData.GetAbi()
	if err != nil {
		t.Error(err)
	}

	input, err := implAbi.Pack("getVersion")
	if err != nil {
		t.Error(err)
	}

	msg := ethereum.CallMsg{
		From: common.Address{},
		To:   &proxyAddr,
		Data: input,
	}

	client, err := ethclient.Dial(url)
	if err != nil {
		t.Error(err)
	}
	raw, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(raw)

	data, err := implAbi.Unpack("gerVersion", raw)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(data)
}

func TestCallImpl(t *testing.T) {
	client, err := ethclient.Dial(url)
	if err != nil {
		t.Error(err)
	}

	implCtr, err := NewMyImplementV1(common.HexToAddress("0xfeae27388A65eE984F452f86efFEd42AaBD438FD"), client)
	if err != nil {
		t.Error(err)
	}

	rtn, err := implCtr.GetVersion(&bind.CallOpts{})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(rtn)

}

/*
0xdd1a2F168C288019E21148fc1097243823657C09
0x4405942d49b5962e16e57b8afa959c39b3d53b6f634c896c4bdd4b522a181cb9
*/
func TestDeplo(t *testing.T) {

	client, err := ethclient.Dial(url)
	if err != nil {
		t.Error(err)
	}

	// 1. 구현 컨트랙트 배포
	auth, _ := CreateTxOpts(privateKey1, nil)

	impleAbi, err := abi.JSON(strings.NewReader(MyImplementV1ABI))

	addr, tx, _, err := bind.DeployContract(auth, impleAbi, common.FromHex(MyImplementV1Bin), client)

	fmt.Println(addr)
	fmt.Println(tx.Hash())
}

func TestDeploy2(t *testing.T) {

	contractAbi, _ := MyImplementV1MetaData.GetAbi()
	contractBin := MyImplementV1MetaData.Bin

	addr, txHash, _, err := Deploy(privateKey1, contractAbi, contractBin)

	if err != nil {
		t.Error(err)
	}
	fmt.Println("생성된 컨트랙트 주소", addr)
	fmt.Println("수행된 트랜잭션 해시", txHash)
}
