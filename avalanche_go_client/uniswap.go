package avalanche_go

import (
	"avalanche_go_client/codec"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UniswapClient struct {
	pk              string
	myAddr          common.Address
	senderAddr      common.Address
	thisAddr        common.Address
	universalRouter *codec.EvmContractCodec
}

func NewUniswapClient(pk string) (*UniswapClient, error) {
	client, err := ethclient.Dial("https://api.avax.network/ext/bc/C/rpc")
	if err != nil {
		return nil, err
	}

	urAbi, err := abi.JSON(strings.NewReader(`[{
		"name": "execute",
		"type": "function",
		"inputs": [
			{"name": "commands", "type": "bytes"},
			{"name": "inputs", "type": "bytes[]"},
			{"name": "deadline", "type": "uint256"}
		],
		"outputs": [],
		"stateMutability": "payable"
	}]`))
	if err != nil {
		return nil, err
	}

	universalRouter := codec.NewEvmCodec(client, common.HexToAddress("0x94b75331AE8d42C1b61065089B7d48FE14aA73b7"), &urAbi)
	return &UniswapClient{
		pk:              pk,
		myAddr:          common.HexToAddress("0xb4dd4fb3D4bCED984cce972991fB100488b59223"),
		senderAddr:      common.HexToAddress("0x0000000000000000000000000000000000000001"),
		thisAddr:        common.HexToAddress("0x0000000000000000000000000000000000000002"),
		universalRouter: universalRouter,
	}, nil
}

const (
	PERMIT2_PERMIT   byte = 0x10
	WRAP_ETH         byte = 0x0b
	UNWRAP_ETH       byte = 0x0c
	V3_SWAP_EXACT_IN byte = 0x00
	PAY_PORTION      byte = 0x06
	SWEEP            byte = 0x04
)

type WrapETH struct {
	Recipient common.Address
	Amount    *big.Int
}

type V3Path struct {
	TokenIn  common.Address
	Fee      *big.Int // 500
	TokenOut common.Address
}

type V3SwapExactIn struct {
	Recipient    common.Address // 0000000000000000000000000000000000000000000000000000000000000002
	AmountIn     *big.Int
	AmountOutMin *big.Int
	Path         []byte
	PayerIsUser  bool
}

type PayPortion struct {
	Token     common.Address
	Recipient common.Address
	Bips      *big.Int // 25
}

type Sweep struct {
	Token        common.Address
	Recipient    common.Address
	AmountOutMin *big.Int
}

const (
	wrapEthAbiJSON = `[{
		"name": "wrapETH",
		"type": "function",
		"inputs": [
			{ "type": "address" },
			{ "type": "uint256" }
		],
		"outputs": []
	}]`
	v3SwapExactInAbiJSON = `[{
		"name": "V3SwapExactIn",
		"type": "function",
		"inputs": [
			{ "type": "address" },
			{ "type": "uint256" },
			{ "type": "uint256" },
			{ "type": "bytes"	},
			{ "type": "bool" }
		],
		"outputs": []
	}]`
	payPortionAbiJSON = `[{
		"name": "PayPortion",
		"type": "function",
		"inputs": [
			{ "type": "address" },
			{ "type": "address" },
			{ "type": "uint256" }
		],
		"outputs": []
	}]`
	sweepAbiJSON = `[{
		"name": "Sweep",
		"type": "function",
		"inputs": [
			{ "type": "address" },
			{ "type": "address" },
			{ "type": "uint256" }
		],
		"outputs": []
	}]`
)

func (u *UniswapClient) SwapNativeForToken(tokenOut common.Address, amountIn *big.Int, amountOutMin *big.Int) (*common.Hash, error) {
	commands := []byte{}
	inputs := [][]byte{}

	wrapCmd, wrapInput, err := u.wrapETH(amountIn)
	if err != nil {
		return nil, err
	}
	commands = append(commands, wrapCmd)
	inputs = append(inputs, wrapInput)

	tokenIn := common.HexToAddress("0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7") // WAVAX
	//V3_SWAP_EXACT_IN
	swapCmd, swapInput, err := u.swapExactIn(tokenIn, tokenOut, amountIn, amountOutMin)
	if err != nil {
		return nil, err
	}
	commands = append(commands, swapCmd)
	inputs = append(inputs, swapInput)

	// SWEEP
	sweepCmd, sweepInput, err := u.sweep(tokenOut, amountOutMin)
	if err != nil {
		return nil, err
	}
	commands = append(commands, sweepCmd)
	inputs = append(inputs, sweepInput)

	tx, err := u.universalRouter.SendWithValue(codec.Standard, big.NewInt(300000), amountIn, &u.myAddr, u.pk, "execute", commands, inputs, big.NewInt(time.Now().Add(time.Hour*1).Unix()))
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

func (u *UniswapClient) Swap(tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOutMin *big.Int) (*common.Hash, error) {

	commands := []byte{} //V4_SWAP
	inputs := [][]byte{}

	// permit2 (omit)
	// commands = append(commands, PERMIT2_PERMIT)
	// type PermitDetails struct {
	// 	// ERC20 token address
	// 	token common.Address
	// 	// the maximum amount allowed to spend
	// 	amount *big.Int
	// 	// timestamp at which a spender's token allowances become invalid
	// 	expiration *big.Int
	// 	// an incrementing value indexed per owner,token,and spender for each signature
	// 	nonce *big.Int
	// }

	//V3_SWAP_EXACT_IN
	swapCmd, swapInput, err := u.swapExactIn(tokenIn, tokenOut, amountIn, amountOutMin)
	if err != nil {
		return nil, err
	}
	commands = append(commands, swapCmd)
	inputs = append(inputs, swapInput)

	// PAY_PORTION(한번만 테스트용으로 수행)
	// payPortionCmd, payPortionInput, err := u.payPortion(common.HexToAddress("0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e"), common.HexToAddress("7ffc3dbf3b2b50ff3a1d5523bc24bb5043837b14"))
	// if err != nil {
	// 	return nil, err
	// }
	// commands = append(commands, payPortionCmd)
	// inputs = append(inputs, payPortionInput)

	// SWEEP
	sweepCmd, sweepInput, err := u.sweep(tokenOut, amountOutMin)
	if err != nil {
		return nil, err
	}
	commands = append(commands, sweepCmd)
	inputs = append(inputs, sweepInput)

	tx, err := u.universalRouter.Send(codec.Standard, big.NewInt(300000), &u.myAddr, u.pk, "execute", commands, inputs, big.NewInt(time.Now().Add(time.Hour*1).Unix()))
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

func (u *UniswapClient) wrapETH(amountIn *big.Int) (command byte, input []byte, err error) {
	wrapETHABI, err := abi.JSON(strings.NewReader(wrapEthAbiJSON))
	if err != nil {
		return 0, nil, err
	}
	wrapETHInput, err := wrapETHABI.Pack("wrapETH", u.thisAddr, amountIn)
	if err != nil {
		return 0, nil, err
	}

	return WRAP_ETH, wrapETHInput[4:], nil // function sig omit
}

func (u *UniswapClient) swapExactIn(tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOutMin *big.Int) (command byte, input []byte, err error) {

	v3SwapExactInABI, err := abi.JSON(strings.NewReader(v3SwapExactInAbiJSON))
	if err != nil {
		return 0, nil, err
	}

	path := []byte{}
	// fee := 500
	// feeBytes := []byte{
	// 	byte(fee >> 16), // most significant byte
	// 	byte(fee >> 8),
	// 	byte(fee),
	// }
	feeBytes := []byte{0x00, 0x01, 0xf4} // 500

	path = append(path, tokenIn.Bytes()...)
	path = append(path, feeBytes...)
	path = append(path, tokenOut.Bytes()...)

	swap := V3SwapExactIn{
		Recipient:    u.thisAddr,
		AmountIn:     amountIn,
		AmountOutMin: amountOutMin,
		Path:         path,
		PayerIsUser:  true,
	}

	swapInput, err := v3SwapExactInABI.Pack("V3SwapExactIn", swap.Recipient, swap.AmountIn, swap.AmountOutMin, swap.Path, swap.PayerIsUser)
	if err != nil {
		return 0, nil, err
	}

	return V3_SWAP_EXACT_IN, swapInput[4:], nil // function sig omit
}

// Token:     common.HexToAddress("0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e"),
// Recipient: common.HexToAddress("7ffc3dbf3b2b50ff3a1d5523bc24bb5043837b14"),
func (u *UniswapClient) payPortion(token common.Address, recipient common.Address) (command byte, input []byte, err error) {

	// commands = append(commands, PAY_PORTION)
	payPortion := PayPortion{
		Token:     token,
		Recipient: recipient,
		Bips:      big.NewInt(25),
	}

	payPortionABI, err := abi.JSON(strings.NewReader(payPortionAbiJSON))
	if err != nil {
		return 0, nil, err
	}
	input, err = payPortionABI.Pack("PayPortion", payPortion.Token, payPortion.Recipient, payPortion.Bips)
	if err != nil {
		return 0, nil, err
	}

	return PAY_PORTION, input[4:], nil // function sig omit
}

func (u *UniswapClient) sweep(tokenOut common.Address, amountOutMin *big.Int) (command byte, input []byte, err error) {
	sweep := Sweep{
		Token:        tokenOut,
		Recipient:    u.myAddr,
		AmountOutMin: amountOutMin,
	}

	sweepABI, err := abi.JSON(strings.NewReader(sweepAbiJSON))
	if err != nil {
		return 0, nil, err
	}
	sweepInput, err := sweepABI.Pack("Sweep", sweep.Token, sweep.Recipient, sweep.AmountOutMin)
	if err != nil {
		return 0, nil, err
	}

	return SWEEP, sweepInput[4:], nil // function sig omit
}

/***********************************************deprecated****************************************************/
// deprecated. 단순 참고용
func (u *UniswapClient) Swap_GPT() (*common.Hash, error) {

	swapABI, _ := abi.JSON(strings.NewReader(`[{
		"inputs": [
			{"internalType": "address","name": "poolManager","type": "address"},
			{"internalType": "bytes","name": "poolKey","type": "bytes"},
			{"internalType": "uint128","name": "amountIn","type": "uint128"},
			{"internalType": "uint128","name": "minAmountOut","type": "uint128"},
			{"internalType": "address","name": "recipient","type": "address"}
		],
		"name": "V4SwapInput",
		"type": "function"
	}]`))

	commands := []byte{0x0a} // V4_SWAP_EXACT_IN_SINGLE

	poolManager := common.HexToAddress("0x06380c0e0912312b5150364b9dc4542ba0dbbc85") // actual PoolManager address
	recipient := u.myAddr
	amountIn := big.NewInt(1e18)  // 1 token
	minAmountOut := big.NewInt(0) // accept any output

	// Construct poolKey bytes manually
	token0 := common.HexToAddress("0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7") // WAVAX
	token1 := common.HexToAddress("0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e") // USDC
	fee := big.NewInt(3000)
	tickSpacing := big.NewInt(60)
	hasHooks := false

	// Solidity encoding: abi.encode(token0, token1, fee, tickSpacing, hasHooks)
	argTypes := abi.Arguments{
		{Type: mustNewType("address")},
		{Type: mustNewType("address")},
		{Type: mustNewType("uint24")},
		{Type: mustNewType("int24")},
		{Type: mustNewType("bool")},
	}

	poolKey, err := argTypes.Pack(token0, token1, fee, tickSpacing, hasHooks)
	if err != nil {
		return nil, err
	}

	// Pack input
	swapInput, _ := swapABI.Pack("V4SwapInput", poolManager, poolKey, amountIn, minAmountOut, recipient)
	inputs := [][]byte{swapInput}

	tx, err := u.universalRouter.Send(codec.High, nil, &u.myAddr, u.pk, "execute", commands, inputs)
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

func mustNewType(t string) abi.Type {
	typ, err := abi.NewType(t, "", nil)
	if err != nil {
		panic(err)
	}
	return typ
}
