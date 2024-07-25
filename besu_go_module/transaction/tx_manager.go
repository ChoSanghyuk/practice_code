package transaction

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"go_module/config"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var chainID *big.Int
var url string
var gasLimit uint64
var client *ethclient.Client

func init() {
	// t : temp
	t1, err := strconv.ParseInt(config.Config.Network.ChainId, 10, 64)
	if err != nil {
		fmt.Println(err, "chain Id Int 변환 시 오류")
	}
	chainID = big.NewInt(t1)

	url = config.Config.Network.Url

	t2, err := strconv.ParseUint(config.Config.Network.GasLimit[2:], 16, 64)
	if err != nil {
		fmt.Println(err, "gasLimit Uint 변환 시 오류")
	}
	gasLimit = t2

	client, err = ethclient.Dial(url)
	if err != nil {
		fmt.Println(err, "ethclient.Dial 시 오류")
	}
}

// fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
// fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0
func Deploy(pk string, abi *abi.ABI, bin string, params ...interface{}) (string, common.Hash) {

	auth, _ := CreateAuth(pk)

	// Important! params... 처럼 Input을 풀지 않으면, [[]]와 같이 이중 구조로 감싸여짐
	address, tx, contract, err := bind.DeployContract(auth, *abi, common.FromHex(bin), client, params...)
	if err != nil {
		fmt.Println(err, "Deploy 시 실패")
		return "", common.Hash{}
	}
	_ = contract

	return address.Hex(), tx.Hash()

}

func Write(pk string, addr string, abi *abi.ABI, method string, params ...interface{}) common.Hash {

	address := common.HexToAddress(addr)
	input, _ := abi.Pack(method, params...)
	return craftSignSendTx(pk, &address, nil, input)
}

func Call(blockNumber *big.Int, addr string, abi *abi.ABI, method string, params ...interface{}) ([]interface{}, error) { // blockNumber : can be nil

	address := common.HexToAddress(addr)
	input, err := abi.Pack(method, params...)
	if err != nil {
		return nil, errors.Join(errors.New(method+" Call 시, abi Pack Error"), err)
	}

	msg := ethereum.CallMsg{
		From: common.Address{},
		To:   &address,
		Data: input,
	}

	raw, err := client.CallContract(context.Background(), msg, blockNumber) // CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("%s Call 시, client Call 에러 발생", method), err)
	}

	data, err := abi.Unpack(method, raw)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("%s Call 시, abi Unpack Error", method), err)
	}

	return data, nil

}

func CreateAuth(pk string) (*bind.TransactOpts, error) {

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, fmt.Errorf("crypto.HexToECDSA 시 오류 %w", err)
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, fmt.Errorf("PendingNonceAt 시 오류 %w", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("SuggestGasPrice 시 오류 %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("NewKeyedTransactorWithChainID 시 오류 %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gasLimit   // in unitss
	auth.GasPrice = gasPrice

	return auth, nil
}

func craftSignSendTx(pk string, to *common.Address, value *big.Int, data []byte) common.Hash {

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		fmt.Println(err, "crypto.HexToECDSA 시 오류")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, _ := client.PendingNonceAt(context.Background(), from)
	fmt.Println("Nonce : ", nonce)

	gasPrice, _ := client.SuggestGasPrice(context.Background())

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		To:       to,
		Value:    value,
		Data:     data,
		Gas:      gasLimit,
	})

	signedTx, err := signTx(tx, privateKey)
	if err != nil {
		fmt.Println(err, "Sign 실패")
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err, "Send 실패")
	}

	return signedTx.Hash()
}

func signTx(tx *types.Transaction, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {

	s := types.NewEIP155Signer(chainID)

	signedTx, err := types.SignTx(tx, s, privateKey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

// 성능 테스트를 위해, msg 완성 후 Call하는 로직
func callByMsgForTest(blockNumber *big.Int, msg ethereum.CallMsg) ([]byte, error) { // blockNumber : can be nil
	return client.CallContract(context.Background(), msg, blockNumber) // CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
}

func callByMsgWithGoRoutineForTest(c chan []byte, blockNumber *big.Int, msg ethereum.CallMsg) { // blockNumber : can be nil

	rtn, _ := client.CallContract(context.Background(), msg, blockNumber) // CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	c <- rtn

}
