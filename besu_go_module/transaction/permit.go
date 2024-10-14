package transaction

import (
	"context"
	"errors"
	"go_module/config"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type PermitSig struct {
	V        uint8
	R        []byte
	S        []byte
	Deadline *big.Int
}

func PermitSignature(tokenContract common.Address, owner, spender string, nonce *big.Int, value *big.Int) (*PermitSig, error) {

	deadline := big.NewInt(time.Now().Add(time.Second * 30).Unix())

	var contractAbi abi.ABI // Todo. ST, DT의 ABI

	/*
		todo. EIP-712
		DOMAIN_SEPARATOR : 함수 식별자가 동일한 타 컨트랙트에 사용하지 못 하도록 하기 위함
	*/
	packed, err := contractAbi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		return nil, err
	}
	data, err := client.CallContract(context.Background(),
		ethereum.CallMsg{
			From: common.Address{},
			To:   &tokenContract,
			Data: packed,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	unpacked, err := contractAbi.Unpack("DOMAIN_SEPARATOR", data)
	if err != nil {
		return nil, err
	}
	ds := unpacked[0].([32]byte)

	// Permit 트랜잭션 생성 및 서명
	uint256Type, _ := abi.NewType("uint256", "uint256", nil)
	addressType, _ := abi.NewType("address", "address", nil)

	permitSig := "Permit(address owner, address spender, uint256 value, uint nonce, uint256 deadline)"
	agruments := abi.Arguments{
		{Type: addressType},
		{Type: addressType},
		{Type: uint256Type},
		{Type: uint256Type},
		{Type: uint256Type},
	}

	packedPermit, err := agruments.Pack(common.HexToAddress(owner), common.HexToAddress(spender), value, nonce, deadline)
	if err != nil {
		return nil, err
	}

	permitStructHash := crypto.Keccak256(append(crypto.Keccak256([]byte(permitSig)), packedPermit...))

	/*
		EIP-191 header  : 0x19
		EIP-191 version : 0x01
	*/
	// rawMsgData := crypto.Keccak256(append([]byte{0x19, 0x01}, append(ds[:], permitStructHash...)...))
	/*
		todo. 아래 링크에서는 0x19부터 전부 keccak 처리. 아래처럼 permitHash만 keccak된 거로 사용해도 되는지 확인 필요
		https://ethereum.stackexchange.com/questions/152731/generate-r-s-v-permit-signature-variables-off-chain-using-golang
	*/
	rawMsgData := append([]byte{0x19, 0x01}, append(ds[:], permitStructHash...)...)

	pk := config.Config.Accounts["account1"].PrivateKey[2:]
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, err
	}

	sig, err := crypto.Sign(rawMsgData, privateKey)
	if err != nil {
		return nil, err
	}

	if len(sig) != 65 {
		return nil, errors.New("wrong signature length")
	}

	return &PermitSig{
		V:        sig[64] + 27,
		R:        sig[:32],
		S:        sig[32:64],
		Deadline: deadline,
	}, nil
}
