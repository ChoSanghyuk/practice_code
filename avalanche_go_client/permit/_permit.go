package permit

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// Permit2 constants
const (
	PERMIT2_NAME            = "Permit2"
	PERMIT_DETAILS_TYPEHASH = "PermitDetails(address token,uint160 amount,uint48 expiration,uint48 nonce)"
	PERMIT_SINGLE_TYPEHASH  = "PermitSingle(PermitDetails details,address spender,uint256 sigDeadline)PermitDetails(address token,uint160 amount,uint48 expiration,uint48 nonce)"
)

// PermitDetails struct
type PermitDetails struct {
	Token      common.Address `json:"token"`
	Amount     *big.Int       `json:"amount"`
	Expiration uint64         `json:"expiration"`
	Nonce      uint64         `json:"nonce"`
}

// PermitSingle struct
type PermitSingle struct {
	Details     PermitDetails  `json:"details"`
	Spender     common.Address `json:"spender"`
	SigDeadline *big.Int       `json:"sigDeadline"`
}

// SignPermit2 generates the signature for Permit2 permit function
func SignPermit2(
	privateKey *ecdsa.PrivateKey,
	permit2Address common.Address,
	chainID *big.Int,
	permitSingle PermitSingle,
) ([]byte, error) {

	// Create EIP-712 domain
	domain := apitypes.TypedDataDomain{
		Name:              PERMIT2_NAME,
		ChainId:           (*math.HexOrDecimal256)(chainID),
		VerifyingContract: permit2Address.Hex(),
	}

	// Define the types
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

	// Create the message data
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

	// Create typed data
	typedData := apitypes.TypedData{
		Types:       types,
		PrimaryType: "PermitSingle",
		Domain:      domain,
		Message:     message,
	}

	/*
		AllowanceTransfer.sol의 37 line. permitSingle.hash() 하는 부분
		struct 데이터의 hash는 TYPEHASH와 같이 function sig를 가지고, 필드들을 encode하는 작업.

	*/
	hash, err := typedData.HashStruct("PermitSingle", message)
	if err != nil {
		return nil, fmt.Errorf("failed to hash struct: %v", err)
	}

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, fmt.Errorf("failed to hash domain: %v", err)
	}

	// Create the final hash (EIP-712)
	finalHash := crypto.Keccak256( // _hashTypedData 수행하는 부분
		[]byte{0x19, 0x01}, // EIP-712 prefix
		domainSeparator[:],
		hash[:],
	)

	// Sign the hash
	signature, err := crypto.Sign(finalHash, privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign: %v", err)
	}

	// Adjust v value for Ethereum (recovery ID + 27)
	if signature[64] < 27 {
		signature[64] += 27
	}

	return signature, nil
}

// Example usage
func main() {
	// Example private key (NEVER use this in production!)
	privateKeyHex := "your_private_key_here"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		panic(err)
	}

	// Example parameters
	permit2Address := common.HexToAddress("0x000000000022D473030F116dDEE9F6B43aC78BA3") // Permit2 mainnet
	chainID := big.NewInt(1)                                                            // Ethereum mainnet

	permitSingle := PermitSingle{
		Details: PermitDetails{
			Token:      common.HexToAddress("0xA0b86a33E6441A8b3278BA6E348b8b49d6B0334C"), // Example token
			Amount:     big.NewInt(1000000000000000000),                                   // 1 token (18 decimals)
			Expiration: 1735689600,                                                        // Unix timestamp
			Nonce:      0,
		},
		Spender:     common.HexToAddress("0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD"), // UniversalRouter
		SigDeadline: big.NewInt(1735689600),                                            // Unix timestamp
	}

	// Generate signature
	signature, err := SignPermit2(privateKey, permit2Address, chainID, permitSingle)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Signature: 0x%x\n", signature)
}
