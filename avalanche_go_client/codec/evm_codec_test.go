package codec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var txData = ""

func TestEvmContractCodec_parseTxData(t *testing.T) {

	json, err := os.ReadFile("../abi/LBrouter.json")
	if err != nil {
		t.Fatal(err)
	}
	lbRouterABIAbi, err := abi.JSON(bytes.NewReader(json))
	if err != nil {
		t.Fatal(err)
	}
	codec := NewEvmCodec(nil, common.Address{}, &lbRouterABIAbi)
	err = codec.unparseTxData(txData, "removeLiquidityNATIVE")
	if err != nil {
		t.Fatal(err)
	}
}

func TestMetaAggregationRouterV2(t *testing.T) {
	type SwapDescriptionV2 struct {
		SrcToken        common.Address   // address
		DstToken        common.Address   // address
		SrcReceivers    []common.Address // address[]
		SrcAmounts      []*big.Int       // uint256[]
		FeeReceivers    []common.Address // address[]
		FeeAmounts      []*big.Int       // uint256[]
		DstReceiver     common.Address   // address
		Amount          *big.Int         // uint256
		MinReturnAmount *big.Int         // uint256
		Flags           *big.Int         // uint256
		Permit          []byte           // bytes
	}

	// SwapExecutionParams is compatible with the Solidity struct
	type SwapExecutionParams struct {
		CallTarget    common.Address    // address
		ApproveTarget common.Address    // address
		TargetData    []byte            // bytes
		Desc          SwapDescriptionV2 // struct
		ClientData    []byte            // bytes
	}

	jsonAbi, err := os.ReadFile("../abi/MetaAggregationRouterV2.json")
	if err != nil {
		t.Fatal(err)
	}
	metaAggregationRouterV2ABIAbi, err := abi.JSON(bytes.NewReader(jsonAbi))
	if err != nil {
		t.Fatal(err)
	}

	t.Run("method_check", func(t *testing.T) {
		for k, v := range metaAggregationRouterV2ABIAbi.Methods {
			fmt.Println(k, common.Bytes2Hex(v.ID))
		}
	})

	t.Run("swap_packing", func(t *testing.T) {

		swapDescriptionV2 := SwapDescriptionV2{
			SrcToken:        common.HexToAddress("0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e"),
			DstToken:        common.HexToAddress("0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7"),
			SrcReceivers:    []common.Address{common.HexToAddress("0xc04f291347d21dc663f7646056db22bff8ce8430")},
			SrcAmounts:      []*big.Int{big.NewInt(109999002)},
			FeeReceivers:    []common.Address{common.HexToAddress("0xff53611968f1e5ca45cfca7918447e7f5776f6d3")},
			FeeAmounts:      []*big.Int{big.NewInt(1)},
			DstReceiver:     common.HexToAddress("0x6e4141d33021b52c91c28608403db4a0ffb50ec6"),
			Amount:          big.NewInt(10000000),
			MinReturnAmount: big.NewInt(40),
			Flags:           big.NewInt(16),
			Permit:          []byte{},
		}

		swapExecutionParams := SwapExecutionParams{
			CallTarget:    common.HexToAddress("0x6e4141d33021b52c91c28608403db4a0ffb50ec6"),
			ApproveTarget: common.HexToAddress("0x11476e10eb79ddffa6f2585be526d2bd840c3e20"),
			TargetData:    []byte{},
			Desc:          swapDescriptionV2,
			ClientData:    []byte{},
		}

		txData, err := metaAggregationRouterV2ABIAbi.Pack("swap", swapExecutionParams)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(common.Bytes2Hex(txData))
	})

	t.Run("swap_unpacking", func(t *testing.T) {
		txData := "e21fd0e900000000000000000000000000000000000000000000000000000000000000200000000000000000000000006e4141d33021b52c91c28608403db4a0ffb50ec6000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000003a000000000000000000000000000000000000000000000000000000000000005e000000000000000000000000000000000000000000000000000000000000002e0000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000b97ef9ef8734c71904d8002f8b6bc66dd9c48a6e000000000000000000000000b31f66aa3c1e785363f0875a1b74e27b85fd66c7000000000000000000000000c04f291347d21dc663f7646056db22bff8ce84300000000000000000000000000000000000000000000000000000000068635da000000000000000000000000000000000000000000000000000000000000002800000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000004063407a490000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000e00000000000000000000000006e4141d33021b52c91c28608403db4a0ffb50ec600000000000000000000000011476e10eb79ddffa6f2585be526d2bd840c3e20000000000000000000000000b97ef9ef8734c71904d8002f8b6bc66dd9c48a6e000000000000000000000000b31f66aa3c1e785363f0875a1b74e27b85fd66c70000000000000000000000000000000000000000000000000000000000989680000000000000000000000000ff53611968f1e5ca45cfca7918447e7f5776f6d300000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000082bc8cd63d000000000000000007cae18c00102a64000000000000000000000000b97ef9ef8734c71904d8002f8b6bc66dd9c48a6e000000000000000000000000b31f66aa3c1e785363f0875a1b74e27b85fd66c7000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000001e00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000c04f291347d21dc663f7646056db22bff8ce8430000000000000000000000000000000000000000000000000000000000098968000000000000000000000000000000000000000000000000007c0e8196157c3c70000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000000010000000000000000000000006e4141d33021b52c91c28608403db4a0ffb50ec60000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000098968000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002857b22536f75726365223a226c666a313334222c22416d6f756e74496e555344223a2231302e303135393133383037333336313932222c22416d6f756e744f7574555344223a2231302e303138353632333136313239373033222c22526566657272616c223a22222c22466c616773223a302c22416d6f756e744f7574223a22353631353039303934393733373738353332222c2254696d657374616d70223a313735313334313239362c22526f7574654944223a2231643765643133372d353231632d343363662d623130342d3133353561666531383638653a35373730373233352d356465622d346234332d396530312d356336303366383532653838222c22496e74656772697479496e666f223a7b224b65794944223a2231222c225369676e6174757265223a224a316d4e524330466b70793764456f6a716d764c327671674231585074724f4f67524f487a41507768777336777a3244736676307a3737323147556d4d32384f424c566c5a557862705670734c6763487a31612f39452f346d62724e76574c43537a794c6e643277476b4c5a44536d316172374e38504d484e316364494e755a62384e517a4333736770646d4b78332b594c5a4e643578544453317839465059306936643834432f717a7244374f6337466d54494d7937566869496f786d735275473141527639305468774d656b6a3473713974774f67527435633163586134326d42326b304e314c514177734a594b4c623137504666742b68314b6b483676597433414154644d31456d576e723459766f464c43796a6a59556f35616f4c387856714c35453455634b4f364246444d72716b624d314f594b653632456377706b77422b75597539724e524259546c7357584d4637673d3d227d7d000000000000000000000000000000000000000000000000000000"
		args := common.Hex2Bytes(txData)
		unpacked, err := metaAggregationRouterV2ABIAbi.Methods["swap"].Inputs.Unpack(args[4:])
		if err != nil {
			t.Fatal(err)
		}
		// t.Log(unpacked[0].(SwapExecutionParams))
		unpackedJSON, err := json.MarshalIndent(unpacked[0], "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(unpackedJSON))
	})
}

func TestUniversalRouter(t *testing.T) {

	jsonAbi, err := os.ReadFile("../abi/UniversalRouter.json")
	if err != nil {
		t.Fatal(err)
	}
	metaAggregationRouterV2ABIAbi, err := abi.JSON(bytes.NewReader(jsonAbi))
	if err != nil {
		t.Fatal(err)
	}

	t.Run("error", func(t *testing.T) {
		for k, v := range metaAggregationRouterV2ABIAbi.Errors {
			fmt.Println(k, v.ID)
		}
	})
}

func TestGetReceipt(t *testing.T) {

	client, err := ethclient.Dial("https://api.avax.network/ext/bc/C/rpc")
	if err != nil {
		t.Fatal(err)
	}

	codec := NewEvmCodec(client, common.Address{}, nil)
	// tx1. 0x879490912b85011eba12835ad9d14df373e751ca65dda5f600823ad3c5c8e768
	// tx2. 0x2bd517822ca76b5ba552d90879e4486b43ec753b133cded6f67b23316c9371be
	tx := common.HexToHash("0xbad4a71871ed80a0b48aef5b9a170a074431d6ad112ebb4e2b807c820111790d")
	receipt, err := codec.GetReceipt(tx)
	if err != nil {
		t.Fatal(err)
	}

	// parse receipt to json
	receiptJSON, err := json.MarshalIndent(receipt, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Receipt JSON:\n%s", string(receiptJSON))
}

func TestSubnets(t *testing.T) {
	var chain1Uri = "http://127.0.0.1:53896/ext/bc/2syE518fxK6R9XLLcamagg9CFiKB7Bf6WvMTh7KYfHZsrhVBEv/rpc"
	var chain2Uri = "http://127.0.0.1:54337/ext/bc/2oAyCEp6RKQVQvfgvfHjVywpULexzfhsXTsqKBQv9UgKGZwKVb/rpc"

	client1, err := ethclient.Dial(chain1Uri)
	if err != nil {
		t.Fatal(err)
	}

	client2, err := ethclient.Dial(chain2Uri)
	if err != nil {
		t.Fatal(err)
	}
	_, _ = client1, client2

	erc20Abi, err := abi.JSON(bytes.NewReader([]byte(erc20abi)))
	if err != nil {
		t.Fatal(err)
	}
	homeABIAbi, err := abi.JSON(bytes.NewReader([]byte(homeabi)))
	if err != nil {
		t.Fatal(err)
	}
	remoteABIAbi, err := abi.JSON(bytes.NewReader([]byte(remoteabi)))
	if err != nil {
		t.Fatal(err)
	}

	t.Run("blockchainId", func(t *testing.T) {

		codec := NewEvmCodec(client1, common.HexToAddress("0x789a5FDac2b37FCD290fb2924382297A6AE65860"), &homeABIAbi)
		rtn, err := codec.Call(nil, "getBlockchainID")
		if err != nil {
			t.Fatal(err)
		}
		t.Log(rtn)
	})

	t.Run("totalSupply", func(t *testing.T) {
		codec := NewEvmCodec(client1, common.HexToAddress("0x52C84043CD9c865236f11d9Fc9F56aa003c1f922"), &homeABIAbi)
		rtn, err := codec.Call(nil, "totalSupply")
		if err != nil {
			t.Fatal(err)
		}
		t.Log(rtn)
	})

	t.Run("Parsed Receipt", func(t *testing.T) {

		codec := NewEvmCodec(client1, common.Address{}, &homeABIAbi)
		// tx1. 0x706c68b47d380303f3c78399c685c2e7e173d07715bab58f478a82a0ffc46dab
		// tx2. 0x665f25b1cabc27c95391e47e18f78a0b986e7d7639359069e4463d1850a4b751
		tx := common.HexToHash("0x340e23b6292300180c0b9bd1684a032002f941dec0be1498b8b9ce2d4fa97ff2")
		receipt, err := codec.GetReceipt(tx)
		if err != nil {
			t.Fatal(err)
		}
		receiptJSON, err := json.MarshalIndent(receipt, "", "  ")
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("Receipt JSON:\n%s\n\n\n", string(receiptJSON))

		parsed, err := codec.ParseReceipt(receipt)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(parsed)
	})

	t.Run("Only receipt", func(t *testing.T) {
		codec := NewEvmCodec(client2, common.Address{}, &homeABIAbi)
		tx := common.HexToHash("0x8fb5a9ad6cde92fef43c21107f9ae242336a67d17ff6329a19e9b8f9bf26eb30")
		receipt, err := codec.GetReceipt(tx)
		if err != nil {
			t.Fatal(err)
		}
		receiptJSON, err := json.MarshalIndent(receipt, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Receipt JSON:\n%s\n\n\n", string(receiptJSON))
	})

	t.Run("mock supply", func(t *testing.T) {
		codec := NewEvmCodec(client1, common.HexToAddress("0xe17bDC68168d07c1776c362d596adaa5d52A1De7"), &erc20Abi)
		rtn, err := codec.Call(nil, "totalSupply")
		if err != nil {
			t.Fatal(err)
		}
		t.Log(rtn)
	})

	t.Run("recipient balance", func(t *testing.T) {
		codec := NewEvmCodec(client2, common.HexToAddress("0x32CaF0D54B0578a96A1aDc7269F19e7398358174"), &remoteABIAbi)
		rtn, err := codec.Call(nil, "balanceOf", common.HexToAddress("0x66291aa2EAa47d2F6b75A4c585437BDBf93D155d"))
		if err != nil {
			t.Fatal(err)
		}
		t.Log(rtn)
	})
}

func TestSendTemp(t *testing.T) {

	client, err := ethclient.Dial("https://api.avax.network/ext/bc/C/rpc")
	if err != nil {
		t.Fatal(err)
	}

	codec := NewEvmCodec(client, common.HexToAddress("0x94b75331AE8d42C1b61065089B7d48FE14aA73b7"), nil)
	pk := os.Getenv("PK")

	myAddr := common.HexToAddress("0xb4dd4fb3D4bCED984cce972991fB100488b59223")

	tx, err := codec.TestSend(High, &myAddr, "0x"+pk, "ateste")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tx)
}
