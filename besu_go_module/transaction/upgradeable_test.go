package transaction

import (
	"context"
	"fmt"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
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

	implAddr, _, _, err := DeployMyImplementV1(auth1, client)
	if err != nil {
		t.Error(err)
	}

	// 2. 프록시 컨트랙트 배포 (프록시-구현 연결)
	auth2, _ := CreateTxOpts(privateKey1, nil)
	proxyAddr, _, _, err := DeployMyProxy(auth2, client, implAddr, nil)
	if err != nil {
		t.Error(err)
	}

	t.Logf("\n프록시 주소 : %s\n구현 컨트랙트 주소 : %s", proxyAddr, implAddr)
}

/*

프록시 주소 : 0x302811352bD2bD428AA2890e1136A0fa247e92dB
구현 컨트랙트 주소 : 0x4cfaD51A19141731Ff971ed16f18c72601BB8e56

*/

func TestCallProxy(t *testing.T) {

	proxyAddr := common.HexToAddress("0x4cfaD51A19141731Ff971ed16f18c72601BB8e56")

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
