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

// AgeInfoStrageMetaData contains all meta data concerning the AgeInfoStrage contract.
var AgeInfoStrageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNameList\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"age\",\"type\":\"uint256\"}],\"name\":\"setAge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610795806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806325b574aa14610051578063b0c8f9dc1461006d578063d83880f014610089578063db88acf4146100b9575b600080fd5b61006b600480360381019061006691906103ae565b6100d7565b005b61008760048036038101906100829190610365565b6100fe565b005b6100a3600480360381019061009e9190610365565b61013d565b6040516100b09190610545565b60405180910390f35b6100c1610164565b6040516100ce9190610523565b60405180910390f35b806000836040516100e8919061050c565b9081526020016040518091039020819055505050565b60018190806001815401808255809150506001900390600052602060002001600090919091909150908051906020019061013992919061023d565b5050565b6000808260405161014e919061050c565b9081526020016040518091039020549050919050565b60606001805480602002602001604051908101604052809291908181526020016000905b828210156102345783829060005260206000200180546101a790610662565b80601f01602080910402602001604051908101604052809291908181526020018280546101d390610662565b80156102205780601f106101f557610100808354040283529160200191610220565b820191906000526020600020905b81548152906001019060200180831161020357829003601f168201915b505050505081526020019060010190610188565b50505050905090565b82805461024990610662565b90600052602060002090601f01602090048101928261026b57600085556102b2565b82601f1061028457805160ff19168380011785556102b2565b828001600101855582156102b2579182015b828111156102b1578251825591602001919060010190610296565b5b5090506102bf91906102c3565b5090565b5b808211156102dc5760008160009055506001016102c4565b5090565b60006102f36102ee84610585565b610560565b90508281526020810184848401111561030f5761030e610728565b5b61031a848285610620565b509392505050565b600082601f83011261033757610336610723565b5b81356103478482602086016102e0565b91505092915050565b60008135905061035f81610748565b92915050565b60006020828403121561037b5761037a610732565b5b600082013567ffffffffffffffff8111156103995761039861072d565b5b6103a584828501610322565b91505092915050565b600080604083850312156103c5576103c4610732565b5b600083013567ffffffffffffffff8111156103e3576103e261072d565b5b6103ef85828601610322565b925050602061040085828601610350565b9150509250929050565b60006104168383610493565b905092915050565b6000610429826105c6565b61043381856105e9565b935083602082028501610445856105b6565b8060005b858110156104815784840389528151610462858261040a565b945061046d836105dc565b925060208a01995050600181019050610449565b50829750879550505050505092915050565b600061049e826105d1565b6104a881856105fa565b93506104b881856020860161062f565b6104c181610737565b840191505092915050565b60006104d7826105d1565b6104e1818561060b565b93506104f181856020860161062f565b80840191505092915050565b61050681610616565b82525050565b600061051882846104cc565b915081905092915050565b6000602082019050818103600083015261053d818461041e565b905092915050565b600060208201905061055a60008301846104fd565b92915050565b600061056a61057b565b90506105768282610694565b919050565b6000604051905090565b600067ffffffffffffffff8211156105a05761059f6106f4565b5b6105a982610737565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000819050919050565b82818337600083830152505050565b60005b8381101561064d578082015181840152602081019050610632565b8381111561065c576000848401525b50505050565b6000600282049050600182168061067a57607f821691505b6020821081141561068e5761068d6106c5565b5b50919050565b61069d82610737565b810181811067ffffffffffffffff821117156106bc576106bb6106f4565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b61075181610616565b811461075c57600080fd5b5056fea2646970667358221220a40077ec591de9fe88f4d33c3f235dae3e2fe3e91033fd9110ef53c5b5d0613564736f6c63430008060033",
}

// AgeInfoStrageABI is the input ABI used to generate the binding from.
// Deprecated: Use AgeInfoStrageMetaData.ABI instead.
var AgeInfoStrageABI = AgeInfoStrageMetaData.ABI

// AgeInfoStrageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AgeInfoStrageMetaData.Bin instead.
var AgeInfoStrageBin = AgeInfoStrageMetaData.Bin

// DeployAgeInfoStrage deploys a new Ethereum contract, binding an instance of AgeInfoStrage to it.
func DeployAgeInfoStrage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AgeInfoStrage, error) {
	parsed, err := AgeInfoStrageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AgeInfoStrageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AgeInfoStrage{AgeInfoStrageCaller: AgeInfoStrageCaller{contract: contract}, AgeInfoStrageTransactor: AgeInfoStrageTransactor{contract: contract}, AgeInfoStrageFilterer: AgeInfoStrageFilterer{contract: contract}}, nil
}

// AgeInfoStrage is an auto generated Go binding around an Ethereum contract.
type AgeInfoStrage struct {
	AgeInfoStrageCaller     // Read-only binding to the contract
	AgeInfoStrageTransactor // Write-only binding to the contract
	AgeInfoStrageFilterer   // Log filterer for contract events
}

// AgeInfoStrageCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgeInfoStrageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgeInfoStrageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgeInfoStrageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgeInfoStrageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgeInfoStrageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgeInfoStrageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgeInfoStrageSession struct {
	Contract     *AgeInfoStrage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgeInfoStrageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgeInfoStrageCallerSession struct {
	Contract *AgeInfoStrageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AgeInfoStrageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgeInfoStrageTransactorSession struct {
	Contract     *AgeInfoStrageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AgeInfoStrageRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgeInfoStrageRaw struct {
	Contract *AgeInfoStrage // Generic contract binding to access the raw methods on
}

// AgeInfoStrageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgeInfoStrageCallerRaw struct {
	Contract *AgeInfoStrageCaller // Generic read-only contract binding to access the raw methods on
}

// AgeInfoStrageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgeInfoStrageTransactorRaw struct {
	Contract *AgeInfoStrageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgeInfoStrage creates a new instance of AgeInfoStrage, bound to a specific deployed contract.
func NewAgeInfoStrage(address common.Address, backend bind.ContractBackend) (*AgeInfoStrage, error) {
	contract, err := bindAgeInfoStrage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgeInfoStrage{AgeInfoStrageCaller: AgeInfoStrageCaller{contract: contract}, AgeInfoStrageTransactor: AgeInfoStrageTransactor{contract: contract}, AgeInfoStrageFilterer: AgeInfoStrageFilterer{contract: contract}}, nil
}

// NewAgeInfoStrageCaller creates a new read-only instance of AgeInfoStrage, bound to a specific deployed contract.
func NewAgeInfoStrageCaller(address common.Address, caller bind.ContractCaller) (*AgeInfoStrageCaller, error) {
	contract, err := bindAgeInfoStrage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgeInfoStrageCaller{contract: contract}, nil
}

// NewAgeInfoStrageTransactor creates a new write-only instance of AgeInfoStrage, bound to a specific deployed contract.
func NewAgeInfoStrageTransactor(address common.Address, transactor bind.ContractTransactor) (*AgeInfoStrageTransactor, error) {
	contract, err := bindAgeInfoStrage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgeInfoStrageTransactor{contract: contract}, nil
}

// NewAgeInfoStrageFilterer creates a new log filterer instance of AgeInfoStrage, bound to a specific deployed contract.
func NewAgeInfoStrageFilterer(address common.Address, filterer bind.ContractFilterer) (*AgeInfoStrageFilterer, error) {
	contract, err := bindAgeInfoStrage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgeInfoStrageFilterer{contract: contract}, nil
}

// bindAgeInfoStrage binds a generic wrapper to an already deployed contract.
func bindAgeInfoStrage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgeInfoStrageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgeInfoStrage *AgeInfoStrageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgeInfoStrage.Contract.AgeInfoStrageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgeInfoStrage *AgeInfoStrageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgeInfoStrage.Contract.AgeInfoStrageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgeInfoStrage *AgeInfoStrageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgeInfoStrage.Contract.AgeInfoStrageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgeInfoStrage *AgeInfoStrageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgeInfoStrage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgeInfoStrage *AgeInfoStrageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgeInfoStrage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgeInfoStrage *AgeInfoStrageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgeInfoStrage.Contract.contract.Transact(opts, method, params...)
}

// GetAge is a free data retrieval call binding the contract method 0xd83880f0.
//
// Solidity: function getAge(string key) view returns(uint256)
func (_AgeInfoStrage *AgeInfoStrageCaller) GetAge(opts *bind.CallOpts, key string) (*big.Int, error) {
	var out []interface{}
	err := _AgeInfoStrage.contract.Call(opts, &out, "getAge", key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAge is a free data retrieval call binding the contract method 0xd83880f0.
//
// Solidity: function getAge(string key) view returns(uint256)
func (_AgeInfoStrage *AgeInfoStrageSession) GetAge(key string) (*big.Int, error) {
	return _AgeInfoStrage.Contract.GetAge(&_AgeInfoStrage.CallOpts, key)
}

// GetAge is a free data retrieval call binding the contract method 0xd83880f0.
//
// Solidity: function getAge(string key) view returns(uint256)
func (_AgeInfoStrage *AgeInfoStrageCallerSession) GetAge(key string) (*big.Int, error) {
	return _AgeInfoStrage.Contract.GetAge(&_AgeInfoStrage.CallOpts, key)
}

// GetNameList is a free data retrieval call binding the contract method 0xdb88acf4.
//
// Solidity: function getNameList() view returns(string[])
func (_AgeInfoStrage *AgeInfoStrageCaller) GetNameList(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _AgeInfoStrage.contract.Call(opts, &out, "getNameList")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetNameList is a free data retrieval call binding the contract method 0xdb88acf4.
//
// Solidity: function getNameList() view returns(string[])
func (_AgeInfoStrage *AgeInfoStrageSession) GetNameList() ([]string, error) {
	return _AgeInfoStrage.Contract.GetNameList(&_AgeInfoStrage.CallOpts)
}

// GetNameList is a free data retrieval call binding the contract method 0xdb88acf4.
//
// Solidity: function getNameList() view returns(string[])
func (_AgeInfoStrage *AgeInfoStrageCallerSession) GetNameList() ([]string, error) {
	return _AgeInfoStrage.Contract.GetNameList(&_AgeInfoStrage.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string name) returns()
func (_AgeInfoStrage *AgeInfoStrageTransactor) Add(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _AgeInfoStrage.contract.Transact(opts, "add", name)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string name) returns()
func (_AgeInfoStrage *AgeInfoStrageSession) Add(name string) (*types.Transaction, error) {
	return _AgeInfoStrage.Contract.Add(&_AgeInfoStrage.TransactOpts, name)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string name) returns()
func (_AgeInfoStrage *AgeInfoStrageTransactorSession) Add(name string) (*types.Transaction, error) {
	return _AgeInfoStrage.Contract.Add(&_AgeInfoStrage.TransactOpts, name)
}

// SetAge is a paid mutator transaction binding the contract method 0x25b574aa.
//
// Solidity: function setAge(string key, uint256 age) returns()
func (_AgeInfoStrage *AgeInfoStrageTransactor) SetAge(opts *bind.TransactOpts, key string, age *big.Int) (*types.Transaction, error) {
	return _AgeInfoStrage.contract.Transact(opts, "setAge", key, age)
}

// SetAge is a paid mutator transaction binding the contract method 0x25b574aa.
//
// Solidity: function setAge(string key, uint256 age) returns()
func (_AgeInfoStrage *AgeInfoStrageSession) SetAge(key string, age *big.Int) (*types.Transaction, error) {
	return _AgeInfoStrage.Contract.SetAge(&_AgeInfoStrage.TransactOpts, key, age)
}

// SetAge is a paid mutator transaction binding the contract method 0x25b574aa.
//
// Solidity: function setAge(string key, uint256 age) returns()
func (_AgeInfoStrage *AgeInfoStrageTransactorSession) SetAge(key string, age *big.Int) (*types.Transaction, error) {
	return _AgeInfoStrage.Contract.SetAge(&_AgeInfoStrage.TransactOpts, key, age)
}
