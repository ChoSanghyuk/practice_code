// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transaction

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MyImplementV1MetaData contains all meta data concerning the MyImplementV1 contract.
var MyImplementV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801562000043575f80fd5b505f62000055620001cf60201b60201c565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f808267ffffffffffffffff161480156200009e5750825b90505f60018367ffffffffffffffff16148015620000d257505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015620000e1575080155b1562000119576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550831562000167576001855f0160086101000a81548160ff0219169083151502179055505b8315620001c4575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051620001bb919062000253565b60405180910390a15b50505050506200026e565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f819050919050565b5f6200023b620002356200022f84620001f6565b62000212565b620001ff565b9050919050565b6200024d816200021b565b82525050565b5f602082019050620002685f83018462000242565b92915050565b608051611288620002955f395f81816104f60152818161054b015261070501526112885ff3fe608060405260043610610090575f3560e01c8063715018a611610058578063715018a6146101565780638129fc1c1461016c5780638da5cb5b14610182578063ad3cb1cc146101ac578063f2fde38b146101d657610090565b80630d8e6e2c1461009457806320965255146100be5780634f1ef286146100e857806352d1902d14610104578063552410771461012e575b5f80fd5b34801561009f575f80fd5b506100a86101fe565b6040516100b59190610d77565b60405180910390f35b3480156100c9575f80fd5b506100d2610206565b6040516100df9190610da8565b60405180910390f35b61010260048036038101906100fd9190610f68565b61020e565b005b34801561010f575f80fd5b5061011861022d565b6040516101259190610fda565b60405180910390f35b348015610139575f80fd5b50610154600480360381019061014f919061101d565b61025e565b005b348015610161575f80fd5b5061016a610267565b005b348015610177575f80fd5b5061018061027a565b005b34801561018d575f80fd5b50610196610402565b6040516101a39190611057565b60405180910390f35b3480156101b7575f80fd5b506101c0610437565b6040516101cd91906110ea565b60405180910390f35b3480156101e1575f80fd5b506101fc60048036038101906101f7919061110a565b610470565b005b5f6001905090565b5f8054905090565b6102166104f4565b61021f826105da565b61022982826105e5565b5050565b5f610236610703565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b905090565b805f8190555050565b61026f61078a565b6102785f610811565b565b5f6102836108e2565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f808267ffffffffffffffff161480156102cb5750825b90505f60018367ffffffffffffffff161480156102fe57505f3073ffffffffffffffffffffffffffffffffffffffff163b145b90508115801561030c575080155b15610343576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315610390576001855f0160086101000a81548160ff0219169083151502179055505b61039933610909565b6103a161091d565b83156103fb575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040516103f2919061118a565b60405180910390a15b5050505050565b5f8061040c610927565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b61047861078a565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036104e8575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016104df9190611057565b60405180910390fd5b6104f181610811565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614806105a157507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661058861094e565b73ffffffffffffffffffffffffffffffffffffffff1614155b156105d8576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6105e261078a565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561064d57506040513d601f19601f8201168201806040525081019061064a91906111cd565b60015b61068e57816040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016106859190611057565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b81146106f457806040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004016106eb9190610fda565b60405180910390fd5b6106fe83836109a1565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614610788576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610792610a13565b73ffffffffffffffffffffffffffffffffffffffff166107b0610402565b73ffffffffffffffffffffffffffffffffffffffff161461080f576107d3610a13565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016108069190611057565b60405180910390fd5b565b5f61081a610927565b90505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b610911610a1a565b61091a81610a5a565b50565b610925610a1a565b565b5f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b5f61097a7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b610ade565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6109aa82610ae7565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f81511115610a0657610a008282610bb0565b50610a0f565b610a0e610c30565b5b5050565b5f33905090565b610a22610c6c565b610a58576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610a62610a1a565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610ad2575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610ac99190611057565b60405180910390fd5b610adb81610811565b50565b5f819050919050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b03610b4257806040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401610b399190611057565b60405180910390fd5b80610b6e7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b610ade565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f808473ffffffffffffffffffffffffffffffffffffffff1684604051610bd9919061123c565b5f60405180830381855af49150503d805f8114610c11576040519150601f19603f3d011682016040523d82523d5f602084013e610c16565b606091505b5091509150610c26858383610c8a565b9250505092915050565b5f341115610c6a576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f610c756108e2565b5f0160089054906101000a900460ff16905090565b606082610c9f57610c9a82610d17565b610d0f565b5f8251148015610cc557505f8473ffffffffffffffffffffffffffffffffffffffff163b145b15610d0757836040517f9996b315000000000000000000000000000000000000000000000000000000008152600401610cfe9190611057565b60405180910390fd5b819050610d10565b5b9392505050565b5f81511115610d295780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61ffff82169050919050565b610d7181610d5b565b82525050565b5f602082019050610d8a5f830184610d68565b92915050565b5f819050919050565b610da281610d90565b82525050565b5f602082019050610dbb5f830184610d99565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610dfb82610dd2565b9050919050565b610e0b81610df1565b8114610e15575f80fd5b50565b5f81359050610e2681610e02565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610e7a82610e34565b810181811067ffffffffffffffff82111715610e9957610e98610e44565b5b80604052505050565b5f610eab610dc1565b9050610eb78282610e71565b919050565b5f67ffffffffffffffff821115610ed657610ed5610e44565b5b610edf82610e34565b9050602081019050919050565b828183375f83830152505050565b5f610f0c610f0784610ebc565b610ea2565b905082815260208101848484011115610f2857610f27610e30565b5b610f33848285610eec565b509392505050565b5f82601f830112610f4f57610f4e610e2c565b5b8135610f5f848260208601610efa565b91505092915050565b5f8060408385031215610f7e57610f7d610dca565b5b5f610f8b85828601610e18565b925050602083013567ffffffffffffffff811115610fac57610fab610dce565b5b610fb885828601610f3b565b9150509250929050565b5f819050919050565b610fd481610fc2565b82525050565b5f602082019050610fed5f830184610fcb565b92915050565b610ffc81610d90565b8114611006575f80fd5b50565b5f8135905061101781610ff3565b92915050565b5f6020828403121561103257611031610dca565b5b5f61103f84828501611009565b91505092915050565b61105181610df1565b82525050565b5f60208201905061106a5f830184611048565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156110a757808201518184015260208101905061108c565b5f8484015250505050565b5f6110bc82611070565b6110c6818561107a565b93506110d681856020860161108a565b6110df81610e34565b840191505092915050565b5f6020820190508181035f83015261110281846110b2565b905092915050565b5f6020828403121561111f5761111e610dca565b5b5f61112c84828501610e18565b91505092915050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f819050919050565b5f61117461116f61116a84611135565b611151565b61113e565b9050919050565b6111848161115a565b82525050565b5f60208201905061119d5f83018461117b565b92915050565b6111ac81610fc2565b81146111b6575f80fd5b50565b5f815190506111c7816111a3565b92915050565b5f602082840312156111e2576111e1610dca565b5b5f6111ef848285016111b9565b91505092915050565b5f81519050919050565b5f81905092915050565b5f611216826111f8565b6112208185611202565b935061123081856020860161108a565b80840191505092915050565b5f611247828461120c565b91508190509291505056fea2646970667358221220b521dfb7e5336f57f16ee4055fec30e1f6fdb06315e5c2b2ab97cdbbe809f34064736f6c63430008140033",
}

// MyImplementV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use MyImplementV1MetaData.ABI instead.
var MyImplementV1ABI = MyImplementV1MetaData.ABI

// MyImplementV1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyImplementV1MetaData.Bin instead.
var MyImplementV1Bin = MyImplementV1MetaData.Bin

// DeployMyImplementV1 deploys a new Ethereum contract, binding an instance of MyImplementV1 to it.
func DeployMyImplementV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MyImplementV1, error) {
	parsed, err := MyImplementV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyImplementV1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyImplementV1{MyImplementV1Caller: MyImplementV1Caller{contract: contract}, MyImplementV1Transactor: MyImplementV1Transactor{contract: contract}, MyImplementV1Filterer: MyImplementV1Filterer{contract: contract}}, nil
}

// MyImplementV1 is an auto generated Go binding around an Ethereum contract.
type MyImplementV1 struct {
	MyImplementV1Caller     // Read-only binding to the contract
	MyImplementV1Transactor // Write-only binding to the contract
	MyImplementV1Filterer   // Log filterer for contract events
}

// MyImplementV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type MyImplementV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyImplementV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MyImplementV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyImplementV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyImplementV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyImplementV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyImplementV1Session struct {
	Contract     *MyImplementV1    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyImplementV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyImplementV1CallerSession struct {
	Contract *MyImplementV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MyImplementV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyImplementV1TransactorSession struct {
	Contract     *MyImplementV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MyImplementV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type MyImplementV1Raw struct {
	Contract *MyImplementV1 // Generic contract binding to access the raw methods on
}

// MyImplementV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyImplementV1CallerRaw struct {
	Contract *MyImplementV1Caller // Generic read-only contract binding to access the raw methods on
}

// MyImplementV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyImplementV1TransactorRaw struct {
	Contract *MyImplementV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMyImplementV1 creates a new instance of MyImplementV1, bound to a specific deployed contract.
func NewMyImplementV1(address common.Address, backend bind.ContractBackend) (*MyImplementV1, error) {
	contract, err := bindMyImplementV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyImplementV1{MyImplementV1Caller: MyImplementV1Caller{contract: contract}, MyImplementV1Transactor: MyImplementV1Transactor{contract: contract}, MyImplementV1Filterer: MyImplementV1Filterer{contract: contract}}, nil
}

// NewMyImplementV1Caller creates a new read-only instance of MyImplementV1, bound to a specific deployed contract.
func NewMyImplementV1Caller(address common.Address, caller bind.ContractCaller) (*MyImplementV1Caller, error) {
	contract, err := bindMyImplementV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyImplementV1Caller{contract: contract}, nil
}

// NewMyImplementV1Transactor creates a new write-only instance of MyImplementV1, bound to a specific deployed contract.
func NewMyImplementV1Transactor(address common.Address, transactor bind.ContractTransactor) (*MyImplementV1Transactor, error) {
	contract, err := bindMyImplementV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyImplementV1Transactor{contract: contract}, nil
}

// NewMyImplementV1Filterer creates a new log filterer instance of MyImplementV1, bound to a specific deployed contract.
func NewMyImplementV1Filterer(address common.Address, filterer bind.ContractFilterer) (*MyImplementV1Filterer, error) {
	contract, err := bindMyImplementV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyImplementV1Filterer{contract: contract}, nil
}

// bindMyImplementV1 binds a generic wrapper to an already deployed contract.
func bindMyImplementV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MyImplementV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyImplementV1 *MyImplementV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyImplementV1.Contract.MyImplementV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyImplementV1 *MyImplementV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyImplementV1.Contract.MyImplementV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyImplementV1 *MyImplementV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyImplementV1.Contract.MyImplementV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyImplementV1 *MyImplementV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyImplementV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyImplementV1 *MyImplementV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyImplementV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyImplementV1 *MyImplementV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyImplementV1.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MyImplementV1 *MyImplementV1Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyImplementV1.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MyImplementV1 *MyImplementV1Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _MyImplementV1.Contract.UPGRADEINTERFACEVERSION(&_MyImplementV1.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MyImplementV1 *MyImplementV1CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MyImplementV1.Contract.UPGRADEINTERFACEVERSION(&_MyImplementV1.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() view returns(uint256)
func (_MyImplementV1 *MyImplementV1Caller) GetValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyImplementV1.contract.Call(opts, &out, "getValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() view returns(uint256)
func (_MyImplementV1 *MyImplementV1Session) GetValue() (*big.Int, error) {
	return _MyImplementV1.Contract.GetValue(&_MyImplementV1.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() view returns(uint256)
func (_MyImplementV1 *MyImplementV1CallerSession) GetValue() (*big.Int, error) {
	return _MyImplementV1.Contract.GetValue(&_MyImplementV1.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(uint16)
func (_MyImplementV1 *MyImplementV1Caller) GetVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _MyImplementV1.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(uint16)
func (_MyImplementV1 *MyImplementV1Session) GetVersion() (uint16, error) {
	return _MyImplementV1.Contract.GetVersion(&_MyImplementV1.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(uint16)
func (_MyImplementV1 *MyImplementV1CallerSession) GetVersion() (uint16, error) {
	return _MyImplementV1.Contract.GetVersion(&_MyImplementV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyImplementV1 *MyImplementV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyImplementV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyImplementV1 *MyImplementV1Session) Owner() (common.Address, error) {
	return _MyImplementV1.Contract.Owner(&_MyImplementV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyImplementV1 *MyImplementV1CallerSession) Owner() (common.Address, error) {
	return _MyImplementV1.Contract.Owner(&_MyImplementV1.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MyImplementV1 *MyImplementV1Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MyImplementV1.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MyImplementV1 *MyImplementV1Session) ProxiableUUID() ([32]byte, error) {
	return _MyImplementV1.Contract.ProxiableUUID(&_MyImplementV1.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MyImplementV1 *MyImplementV1CallerSession) ProxiableUUID() ([32]byte, error) {
	return _MyImplementV1.Contract.ProxiableUUID(&_MyImplementV1.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_MyImplementV1 *MyImplementV1Transactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyImplementV1.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_MyImplementV1 *MyImplementV1Session) Initialize() (*types.Transaction, error) {
	return _MyImplementV1.Contract.Initialize(&_MyImplementV1.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_MyImplementV1 *MyImplementV1TransactorSession) Initialize() (*types.Transaction, error) {
	return _MyImplementV1.Contract.Initialize(&_MyImplementV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyImplementV1 *MyImplementV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyImplementV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyImplementV1 *MyImplementV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _MyImplementV1.Contract.RenounceOwnership(&_MyImplementV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyImplementV1 *MyImplementV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MyImplementV1.Contract.RenounceOwnership(&_MyImplementV1.TransactOpts)
}

// SetValue is a paid mutator transaction binding the contract method 0x55241077.
//
// Solidity: function setValue(uint256 newValue) returns()
func (_MyImplementV1 *MyImplementV1Transactor) SetValue(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _MyImplementV1.contract.Transact(opts, "setValue", newValue)
}

// SetValue is a paid mutator transaction binding the contract method 0x55241077.
//
// Solidity: function setValue(uint256 newValue) returns()
func (_MyImplementV1 *MyImplementV1Session) SetValue(newValue *big.Int) (*types.Transaction, error) {
	return _MyImplementV1.Contract.SetValue(&_MyImplementV1.TransactOpts, newValue)
}

// SetValue is a paid mutator transaction binding the contract method 0x55241077.
//
// Solidity: function setValue(uint256 newValue) returns()
func (_MyImplementV1 *MyImplementV1TransactorSession) SetValue(newValue *big.Int) (*types.Transaction, error) {
	return _MyImplementV1.Contract.SetValue(&_MyImplementV1.TransactOpts, newValue)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyImplementV1 *MyImplementV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MyImplementV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyImplementV1 *MyImplementV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MyImplementV1.Contract.TransferOwnership(&_MyImplementV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyImplementV1 *MyImplementV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MyImplementV1.Contract.TransferOwnership(&_MyImplementV1.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MyImplementV1 *MyImplementV1Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MyImplementV1.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MyImplementV1 *MyImplementV1Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MyImplementV1.Contract.UpgradeToAndCall(&_MyImplementV1.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MyImplementV1 *MyImplementV1TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MyImplementV1.Contract.UpgradeToAndCall(&_MyImplementV1.TransactOpts, newImplementation, data)
}

// MyImplementV1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MyImplementV1 contract.
type MyImplementV1InitializedIterator struct {
	Event *MyImplementV1Initialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyImplementV1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyImplementV1Initialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyImplementV1Initialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyImplementV1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyImplementV1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyImplementV1Initialized represents a Initialized event raised by the MyImplementV1 contract.
type MyImplementV1Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MyImplementV1 *MyImplementV1Filterer) FilterInitialized(opts *bind.FilterOpts) (*MyImplementV1InitializedIterator, error) {

	logs, sub, err := _MyImplementV1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MyImplementV1InitializedIterator{contract: _MyImplementV1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MyImplementV1 *MyImplementV1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MyImplementV1Initialized) (event.Subscription, error) {

	logs, sub, err := _MyImplementV1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyImplementV1Initialized)
				if err := _MyImplementV1.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MyImplementV1 *MyImplementV1Filterer) ParseInitialized(log types.Log) (*MyImplementV1Initialized, error) {
	event := new(MyImplementV1Initialized)
	if err := _MyImplementV1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyImplementV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MyImplementV1 contract.
type MyImplementV1OwnershipTransferredIterator struct {
	Event *MyImplementV1OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyImplementV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyImplementV1OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyImplementV1OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyImplementV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyImplementV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyImplementV1OwnershipTransferred represents a OwnershipTransferred event raised by the MyImplementV1 contract.
type MyImplementV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyImplementV1 *MyImplementV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MyImplementV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MyImplementV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MyImplementV1OwnershipTransferredIterator{contract: _MyImplementV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyImplementV1 *MyImplementV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MyImplementV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MyImplementV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyImplementV1OwnershipTransferred)
				if err := _MyImplementV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyImplementV1 *MyImplementV1Filterer) ParseOwnershipTransferred(log types.Log) (*MyImplementV1OwnershipTransferred, error) {
	event := new(MyImplementV1OwnershipTransferred)
	if err := _MyImplementV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyImplementV1UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the MyImplementV1 contract.
type MyImplementV1UpgradedIterator struct {
	Event *MyImplementV1Upgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyImplementV1UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyImplementV1Upgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyImplementV1Upgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyImplementV1UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyImplementV1UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyImplementV1Upgraded represents a Upgraded event raised by the MyImplementV1 contract.
type MyImplementV1Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MyImplementV1 *MyImplementV1Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MyImplementV1UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MyImplementV1.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MyImplementV1UpgradedIterator{contract: _MyImplementV1.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MyImplementV1 *MyImplementV1Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MyImplementV1Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MyImplementV1.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyImplementV1Upgraded)
				if err := _MyImplementV1.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MyImplementV1 *MyImplementV1Filterer) ParseUpgraded(log types.Log) (*MyImplementV1Upgraded, error) {
	event := new(MyImplementV1Upgraded)
	if err := _MyImplementV1.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
