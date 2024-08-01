package transaction

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"go_module/config"
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
	ci, err := strconv.ParseInt(config.Config.Network.ChainId, 10, 64)
	if err != nil {
		fmt.Println(err, "chain Id Int 변환 시 오류")
	}
	chainID = big.NewInt(ci)

	url = config.Config.Network.Url

	gl, err := strconv.ParseUint(config.Config.Network.GasLimit[2:], 16, 64)
	if err != nil {
		fmt.Println(err, "gasLimit Uint 변환 시 오류")
	}
	gasLimit = gl

	client, err = ethclient.Dial(url)
	if err != nil {
		fmt.Println(err, "ethclient.Dial 시 오류")
	}
}

func Deploy(pk string, abi *abi.ABI, bin string, params ...interface{}) (string, common.Hash, *bind.BoundContract, error) {

	auth, _ := CreateTxOpts(pk, nil)

	// Important! params... 처럼 Input을 풀지 않으면, [[]]와 같이 이중 구조로 감싸여짐
	address, tx, contract, err := bind.DeployContract(auth, *abi, common.FromHex(bin), client, params...)
	if err != nil {
		return "", common.Hash{}, nil, fmt.Errorf("Deploy 시 실패. %w", err)
	}

	return address.Hex(), tx.Hash(), contract, nil

}

func Write(pk string, addr string, abi *abi.ABI, method string, params ...interface{}) (common.Hash, error) {

	address := common.HexToAddress(addr)

	input, _ := abi.Pack(method, params...)

	signedTx, err := CreateSignedTx(pk, &address, nil, input)
	if err != nil {
		return common.Hash{}, fmt.Errorf("CreateTx시 오류. %w", err)
	}

	hash, err := SendTx(signedTx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("서명 및 전송 시 실패. %w", err)
	}

	return hash, nil
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

func CreateTxOpts(pk string, value *big.Int) (*bind.TransactOpts, error) {

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
	auth.Value = value
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	return auth, nil
}

func CreateSignedTx(pk string, to *common.Address, value *big.Int, data []byte) (*types.Transaction, error) {

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, fmt.Errorf("crypto.HexToECDSA 시 오류 %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		return nil, fmt.Errorf("crypto.PendingNonceAt 시 오류 %w", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("crypto.SuggestGasPrice 시 오류 %w", err)
	}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		To:       to,
		Value:    value,
		Data:     data,
		Gas:      gasLimit,
	})

	signedTx, err := SignTx(tx, privateKey)
	if err != nil {
		return nil, fmt.Errorf("sign 실패. %w", err)
	}

	return signedTx, nil
}
func CreateSignedTx2(pk string, to *common.Address, value *big.Int, data []byte, gas uint64, i uint64) (*types.Transaction, error) {

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, fmt.Errorf("crypto.HexToECDSA 시 오류 %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		return nil, fmt.Errorf("crypto.PendingNonceAt 시 오류 %w", err)
	}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce + i,
		GasPrice: big.NewInt(1),
		To:       to,
		Value:    value,
		Data:     data,
		Gas:      gas,
	})

	signedTx, err := SignTx(tx, privateKey)
	if err != nil {
		return nil, fmt.Errorf("sign 실패. %w", err)
	}

	return signedTx, nil
}

func SendTx(signedTx *types.Transaction) (common.Hash, error) {

	err := client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("send 실패. %w", err)
	}

	return signedTx.Hash(), nil

}

func SignTx(tx *types.Transaction, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {

	s := types.NewEIP155Signer(chainID)

	signedTx, err := types.SignTx(tx, s, privateKey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

// 성능 테스트를 위해, msg 완성 후 Call하는 로직
func CallByMsg(blockNumber *big.Int, msg ethereum.CallMsg) ([]byte, error) { // blockNumber : can be nil
	return client.CallContract(context.Background(), msg, blockNumber) // CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
}

func callByMsgWithGoRoutineForTest(c chan []byte, blockNumber *big.Int, msg ethereum.CallMsg) { // blockNumber : can be nil

	rtn, _ := client.CallContract(context.Background(), msg, blockNumber) // CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	c <- rtn

}
