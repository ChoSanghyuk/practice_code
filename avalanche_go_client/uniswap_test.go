package avalanche_go

import (
	"avalanche_go_client/evmtxbroker"
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/stretchr/testify/assert"
)

func TestUniswapClient(t *testing.T) {

	pk := os.Getenv("PK")
	uniswapClient, err := NewUniswapClient(&UniswapClientConfig{
		url:      "https://api.avax.network/ext/bc/C/rpc",
		pk:       pk,
		urAddr:   "0x94b75331AE8d42C1b61065089B7d48FE14aA73b7",
		pmAddr:   "0x000000000022D473030F116dDEE9F6B43aC78BA3",
		gasLimit: big.NewInt(300000),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Swap_WavaxToUsdc", func(t *testing.T) {

		tokenIn := common.HexToAddress("0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7")
		tokenOut := common.HexToAddress("0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e")
		amountIn := big.NewInt(1e17)
		amountOutMin := big.NewInt(0)

		tx, err := uniswapClient.Swap(tokenIn, tokenOut, amountIn, amountOutMin)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(tx)
	})

	t.Run("Swap_UsdcToWavax", func(t *testing.T) {

		tokenIn := common.HexToAddress("0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e")
		tokenOut := common.HexToAddress("0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7")
		amountIn := big.NewInt(1e6)
		amountOutMin := big.NewInt(0)

		tx, err := uniswapClient.Swap(tokenIn, tokenOut, amountIn, amountOutMin)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(tx)
	})

	t.Run("NativeSwap", func(t *testing.T) {

		// tokenIn := common.HexToAddress("0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7")
		tokenOut := common.HexToAddress("0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e")
		amountIn := big.NewInt(1e17)
		amountOutMin := big.NewInt(0)

		tx, err := uniswapClient.SwapNativeForToken(tokenOut, amountIn, amountOutMin)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(tx)
	})

	t.Run("UsdtUsdcMutualSwap", func(t *testing.T) {

		usdc := common.HexToAddress("0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e")
		usdt := common.HexToAddress("0x9702230A8Ea53601f5cD2dc00fDBc13d4dF4A8c7")
		amountIn := big.NewInt(1e6)
		amountOutMin := big.NewInt(1e5)

		tokenIn, tokenOut := usdc, usdt

		for i := 0; i < 9; i++ {
			tx, err := uniswapClient.Swap(tokenIn, tokenOut, amountIn, amountOutMin)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(tx)
			time.Sleep(10 * time.Second)
			tokenIn, tokenOut = tokenOut, tokenIn
		}

	})
}

func TestUniswapMethod(t *testing.T) {

	pk := os.Getenv("PK")
	if pk == "" {
		pk = "0000000000000000000000000000000000000000000000000000000000000001"
	}
	u, err := NewUniswapClient(&UniswapClientConfig{
		url:      "https://api.avax.network/ext/bc/C/rpc",
		pk:       pk,
		urAddr:   "0x94b75331AE8d42C1b61065089B7d48FE14aA73b7",
		pmAddr:   "0x000000000022D473030F116dDEE9F6B43aC78BA3",
		gasLimit: big.NewInt(300000),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Run("permit2Cmd", func(t *testing.T) {
		permitCmd, permitInput, err := u.permit2(PermitSingle{
			Details: PermitDetails{
				Token:      common.HexToAddress("0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e"),
				Amount:     big.NewInt(1e6),
				Expiration: big.NewInt(1759276800),
				Nonce:      big.NewInt(1),
			},
			Spender:     common.HexToAddress("0x94b75331AE8d42C1b61065089B7d48FE14aA73b7"),
			SigDeadline: big.NewInt(time.Now().Add(time.Hour * 1).Unix()),
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v", permitCmd)
		t.Logf("%v", common.Bytes2Hex(permitInput))
	})

	t.Run("permit2Nonce", func(t *testing.T) {
		permitNonce, err := u.permit2Nonce(common.HexToAddress("0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e"), common.HexToAddress("0x94b75331AE8d42C1b61065089B7d48FE14aA73b7"))
		if err != nil {
			t.Fatal(err)
		}
		t.Log(permitNonce)
	})

	t.Run("approve", func(t *testing.T) {
		// usdt , 2025년 10월 1일 수요일 오전 9:00:00 GMT+09:00
		tx, err := u.Approve(common.HexToAddress("0x9702230A8Ea53601f5cD2dc00fDBc13d4dF4A8c7"), common.HexToAddress("0x94b75331AE8d42C1b61065089B7d48FE14aA73b7"), big.NewInt(1e8), big.NewInt(1759276800))
		if err != nil {
			t.Fatal(err)
		}
		t.Log(tx)
	})
}

func TestPermit2Client(t *testing.T) {

	pk := os.Getenv("PK")
	if pk == "" {
		pk = "0000000000000000000000000000000000000000000000000000000000000001"
	}
	p, err := NewPermitClient(pk)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("permit2Tx", func(t *testing.T) {
		tx, err := p.permit2(PermitSingle{
			Details: PermitDetails{
				Token:      common.HexToAddress("0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e"),
				Amount:     big.NewInt(1e6),
				Expiration: big.NewInt(1759366800),
				Nonce:      big.NewInt(3),
			},
			Spender:     common.HexToAddress("0x94b75331AE8d42C1b61065089B7d48FE14aA73b7"),
			SigDeadline: big.NewInt(time.Now().Add(time.Hour * 1).Unix()),
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v", tx)
	})

	t.Run("permit2SingleHash", func(t *testing.T) {
		hashByte, err := p.permit2Hash(PermitSingle{
			Details: PermitDetails{
				Token:      common.HexToAddress("0x9702230A8Ea53601f5cD2dc00fDBc13d4dF4A8c7"),
				Amount:     big.NewInt(1e6),
				Expiration: big.NewInt(1759366800),
				Nonce:      big.NewInt(3),
			},
			Spender:     common.HexToAddress("0x94b75331AE8d42C1b61065089B7d48FE14aA73b7"),
			SigDeadline: big.NewInt(1759366800), //big.NewInt(time.Now().Add(time.Hour * 1).Unix()),
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v", common.Bytes2Hex(hashByte))

		assert.Equal(t, "0x"+common.Bytes2Hex(hashByte), "0x1a9387b519aea4038428ebee7865e669686604c70c129f802b9d8e2dbd0d3dad", "두 해시는 같아야 합니다")

	})

	t.Run("permit2SingleHashCompare", func(t *testing.T) {
		hashByte1, err := p.permit2Hash(PermitSingle{
			Details: PermitDetails{
				Token:      common.HexToAddress("0x9702230A8Ea53601f5cD2dc00fDBc13d4dF4A8c7"),
				Amount:     big.NewInt(1e6),
				Expiration: big.NewInt(1759366800),
				Nonce:      big.NewInt(3),
			},
			Spender:     common.HexToAddress("0x94b75331AE8d42C1b61065089B7d48FE14aA73b7"),
			SigDeadline: big.NewInt(1759366800), //big.NewInt(time.Now().Add(time.Hour * 1).Unix()),
		})
		if err != nil {
			t.Fatal(err)
		}

		hashByte2, err := p.permit2Hash2(PermitSingle{
			Details: PermitDetails{
				Token:      common.HexToAddress("0x9702230A8Ea53601f5cD2dc00fDBc13d4dF4A8c7"),
				Amount:     big.NewInt(1e6),
				Expiration: big.NewInt(1759366800),
				Nonce:      big.NewInt(3),
			},
			Spender:     common.HexToAddress("0x94b75331AE8d42C1b61065089B7d48FE14aA73b7"),
			SigDeadline: big.NewInt(1759366800), //big.NewInt(time.Now().Add(time.Hour * 1).Unix()),
		})
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, common.Bytes2Hex(hashByte1), common.Bytes2Hex(hashByte2), "두 해시는 같아야 합니다")

	})

}

/*
테스트용 PermitClient 정의
*/
type PermitClient struct {
	pk           *ecdsa.PrivateKey
	myAddr       common.Address
	permitRouter *evmtxbroker.EvmTxBroker
}

func NewPermitClient(pk string) (*PermitClient, error) {
	client, err := ethclient.Dial("https://api.avax.network/ext/bc/C/rpc")
	if err != nil {
		return nil, err
	}

	// permit2 abi 파일에서 permitBatch만 삭제. 기존 permit과 겹쳐서 파라미터 인식을 permitBatch로 하는 경우 발생.
	json, err := os.ReadFile("./abi/Permit2.json")
	if err != nil {
		return nil, err
	}
	permitAbi, err := abi.JSON(bytes.NewReader(json))
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, err
	}

	permitRouter := evmtxbroker.NewEvmTxBroker(client, common.HexToAddress("0x000000000022D473030F116dDEE9F6B43aC78BA3"), &permitAbi)
	return &PermitClient{
		pk:           privateKey,
		myAddr:       common.HexToAddress("0xb4dd4fb3D4bCED984cce972991fB100488b59223"),
		permitRouter: permitRouter,
	}, nil
}

func (p *PermitClient) permit2(permitSingle PermitSingle) (*common.Hash, error) {

	finalHash, err := p.permit2Hash(permitSingle)
	if err != nil {
		return nil, err
	}

	sig, err := crypto.Sign(finalHash, p.pk)
	if err != nil {
		return nil, err
	}

	// 이게 중요한가 벼..
	if sig[64] < 27 {
		sig[64] += 27
	}

	tx, err := p.permitRouter.Send(evmtxbroker.Standard, big.NewInt(300000), &p.myAddr, p.pk, "permit", p.myAddr, permitSingle, sig)
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

func (p *PermitClient) permit2Hash(permitSingle PermitSingle) ([]byte, error) {
	// Domain Serparator. evm_codec_test.go 테스트에 Permit2 컨트랙트에 직접 요청보내서 DOMAIN_SEPARATOR 가져오는 것 있음.
	typeHashSlice := crypto.Keccak256([]byte("EIP712Domain(string name,uint256 chainId,address verifyingContract)"))
	hashedNameSlice := crypto.Keccak256([]byte("Permit2"))

	var typeHash [32]byte
	var hashedName [32]byte
	copy(typeHash[:], typeHashSlice)
	copy(hashedName[:], hashedNameSlice)

	chainId := big.NewInt(43114)                                                     //big.NewInt(1)                                                         // big.NewInt(43114)
	permit2Addr := common.HexToAddress("0x000000000022D473030F116dDEE9F6B43aC78BA3") //common.HexToAddress("0x0000000000000000000000000000000000000000") //

	uint256Type, _ := abi.NewType("uint256", "", nil)
	bytes32Type, _ := abi.NewType("bytes32", "", nil)
	addressType, _ := abi.NewType("address", "", nil)

	arguments := abi.Arguments{
		{Type: bytes32Type}, // typeHash
		{Type: bytes32Type}, // nameHash
		{Type: uint256Type}, // chainID
		{Type: addressType}, // contract address
	}

	packed, err := arguments.Pack(typeHash, hashedName, chainId, permit2Addr)
	if err != nil {
		panic(fmt.Sprintf("Failed to pack data: %v", err))
	}

	// Hash the packed data with keccak256
	dsHash := crypto.Keccak256Hash(packed) // domainseparator hash
	/*             */

	domain := apitypes.TypedDataDomain{
		Name:              "Permit2",
		ChainId:           (*math.HexOrDecimal256)(chainId),
		VerifyingContract: permit2Addr.Hex(),
	}

	types := apitypes.Types{
		"EIP712Domain": {
			{Name: "name", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"PermitDetails": {
			{Name: "token", Type: "address"},
			{Name: "amount", Type: "uint160"},
			{Name: "expiration", Type: "uint48"},
			{Name: "nonce", Type: "uint48"},
		},
		"PermitSingle": {
			{Name: "details", Type: "PermitDetails"},
			{Name: "spender", Type: "address"},
			{Name: "sigDeadline", Type: "uint256"},
		},
	}

	message := apitypes.TypedDataMessage{
		"details": map[string]interface{}{
			"token":      permitSingle.Details.Token.Hex(),
			"amount":     permitSingle.Details.Amount.String(),
			"expiration": permitSingle.Details.Expiration,
			"nonce":      permitSingle.Details.Nonce,
		},
		"spender":     permitSingle.Spender.Hex(),
		"sigDeadline": permitSingle.SigDeadline.String(),
	}

	typedData := apitypes.TypedData{
		Types:       types,
		PrimaryType: "PermitSingle",
		Message:     message,
		Domain:      domain,
	}

	/*
		AllowanceTransfer.sol의 37 line. permitSingle.hash() 하는 부분
		struct 데이터의 hash는 TYPEHASH와 같이 function sig를 가지고, 필드들을 encode하는 작업.
	*/
	permitSingleHash, err := typedData.HashStruct("PermitSingle", message)
	if err != nil {
		return nil, fmt.Errorf("failed to hash struct: %v", err)
	}

	// _hashTypedData
	finalHash := crypto.Keccak256( // _hashTypedData 수행하는 부분
		[]byte{0x19, 0x01}, // EIP-712 prefix
		dsHash[:],
		permitSingleHash[:],
	)

	return finalHash, nil
}

func (p *PermitClient) permit2Hash2(permitSingle PermitSingle) ([]byte, error) {

	chainId := big.NewInt(43114)                                                     //big.NewInt(1)                                                         // big.NewInt(43114)
	permit2Addr := common.HexToAddress("0x000000000022D473030F116dDEE9F6B43aC78BA3") //common.HexToAddress("0x0000000000000000000000000000000000000000") //

	domain := apitypes.TypedDataDomain{
		Name:              "Permit2",
		ChainId:           (*math.HexOrDecimal256)(chainId),
		VerifyingContract: permit2Addr.Hex(),
	}

	types := apitypes.Types{
		"EIP712Domain": {
			{Name: "name", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"PermitDetails": {
			{Name: "token", Type: "address"},
			{Name: "amount", Type: "uint160"},
			{Name: "expiration", Type: "uint48"},
			{Name: "nonce", Type: "uint48"},
		},
		"PermitSingle": {
			{Name: "details", Type: "PermitDetails"},
			{Name: "spender", Type: "address"},
			{Name: "sigDeadline", Type: "uint256"},
		},
	}

	message := apitypes.TypedDataMessage{
		"details": map[string]interface{}{
			"token":      permitSingle.Details.Token.Hex(),
			"amount":     permitSingle.Details.Amount.String(),
			"expiration": permitSingle.Details.Expiration,
			"nonce":      permitSingle.Details.Nonce,
		},
		"spender":     permitSingle.Spender.Hex(),
		"sigDeadline": permitSingle.SigDeadline.String(),
	}

	typedData := apitypes.TypedData{
		Types:       types,
		PrimaryType: "PermitSingle",
		Message:     message,
		Domain:      domain,
	}

	/*
		AllowanceTransfer.sol의 37 line. permitSingle.hash() 하는 부분
		struct 데이터의 hash는 TYPEHASH와 같이 function sig를 가지고, 필드들을 encode하는 작업.
	*/
	finalHash, _, err := apitypes.TypedDataAndHash(typedData)
	if err != nil {
		return nil, fmt.Errorf("failed to hash struct: %v", err)
	}

	return finalHash, nil
}
