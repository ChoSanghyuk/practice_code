// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_key\",\"type\":\"uint256\"}],\"name\":\"getAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_key\",\"type\":\"uint256\"}],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getPriceList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_age\",\"type\":\"uint256\"}],\"name\":\"setAgeList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_key\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setNameList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_itemName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPriceList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610802806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630e37008a146100675780631bbefaf1146100975780635eac9098146100b35780636b8ff574146100e3578063a266f6c214610113578063fe093d3b1461012f575b600080fd5b610081600480360381019061007c9190610470565b61014b565b60405161008e91906105eb565b60405180910390f35b6100b160048036038101906100ac919061049d565b610168565b005b6100cd60048036038101906100c891906103cb565b610194565b6040516100da91906105eb565b60405180910390f35b6100fd60048036038101906100f89190610470565b6101bb565b60405161010a91906105c9565b60405180910390f35b61012d60048036038101906101289190610414565b610260565b005b610149600480360381019061014491906104f9565b610287565b005b600060026000838152602001908152602001600020549050919050565b8060016000848152602001908152602001600020908051906020019061018f9291906102a3565b505050565b600080826040516101a591906105b2565b9081526020016040518091039020549050919050565b60606001600083815260200190815260200160002080546101db906106cf565b80601f0160208091040260200160405190810160405280929190818152602001828054610207906106cf565b80156102545780601f1061022957610100808354040283529160200191610254565b820191906000526020600020905b81548152906001019060200180831161023757829003601f168201915b50505050509050919050565b8060008360405161027191906105b2565b9081526020016040518091039020819055505050565b8060026000848152602001908152602001600020819055505050565b8280546102af906106cf565b90600052602060002090601f0160209004810192826102d15760008555610318565b82601f106102ea57805160ff1916838001178555610318565b82800160010185558215610318579182015b828111156103175782518255916020019190600101906102fc565b5b5090506103259190610329565b5090565b5b8082111561034257600081600090555060010161032a565b5090565b60006103596103548461062b565b610606565b90508281526020810184848401111561037557610374610795565b5b61038084828561068d565b509392505050565b600082601f83011261039d5761039c610790565b5b81356103ad848260208601610346565b91505092915050565b6000813590506103c5816107b5565b92915050565b6000602082840312156103e1576103e061079f565b5b600082013567ffffffffffffffff8111156103ff576103fe61079a565b5b61040b84828501610388565b91505092915050565b6000806040838503121561042b5761042a61079f565b5b600083013567ffffffffffffffff8111156104495761044861079a565b5b61045585828601610388565b9250506020610466858286016103b6565b9150509250929050565b6000602082840312156104865761048561079f565b5b6000610494848285016103b6565b91505092915050565b600080604083850312156104b4576104b361079f565b5b60006104c2858286016103b6565b925050602083013567ffffffffffffffff8111156104e3576104e261079a565b5b6104ef85828601610388565b9150509250929050565b600080604083850312156105105761050f61079f565b5b600061051e858286016103b6565b925050602061052f858286016103b6565b9150509250929050565b60006105448261065c565b61054e8185610667565b935061055e81856020860161069c565b610567816107a4565b840191505092915050565b600061057d8261065c565b6105878185610678565b935061059781856020860161069c565b80840191505092915050565b6105ac81610683565b82525050565b60006105be8284610572565b915081905092915050565b600060208201905081810360008301526105e38184610539565b905092915050565b600060208201905061060060008301846105a3565b92915050565b6000610610610621565b905061061c8282610701565b919050565b6000604051905090565b600067ffffffffffffffff82111561064657610645610761565b5b61064f826107a4565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b6000819050919050565b82818337600083830152505050565b60005b838110156106ba57808201518184015260208101905061069f565b838111156106c9576000848401525b50505050565b600060028204905060018216806106e757607f821691505b602082108114156106fb576106fa610732565b5b50919050565b61070a826107a4565b810181811067ffffffffffffffff8211171561072957610728610761565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b6107be81610683565b81146107c957600080fd5b5056fea26469706673582212201bc70563972a90016812d64bfbedd28a742834b2a99b41eeb0e881844129810964736f6c63430008060033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetAge is a free data retrieval call binding the contract method 0x0e37008a.
//
// Solidity: function getAge(uint256 _key) view returns(uint256)
func (_Contract *ContractCaller) GetAge(opts *bind.CallOpts, _key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAge", _key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAge is a free data retrieval call binding the contract method 0x0e37008a.
//
// Solidity: function getAge(uint256 _key) view returns(uint256)
func (_Contract *ContractSession) GetAge(_key *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetAge(&_Contract.CallOpts, _key)
}

// GetAge is a free data retrieval call binding the contract method 0x0e37008a.
//
// Solidity: function getAge(uint256 _key) view returns(uint256)
func (_Contract *ContractCallerSession) GetAge(_key *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetAge(&_Contract.CallOpts, _key)
}

// GetName is a free data retrieval call binding the contract method 0x6b8ff574.
//
// Solidity: function getName(uint256 _key) view returns(string)
func (_Contract *ContractCaller) GetName(opts *bind.CallOpts, _key *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getName", _key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x6b8ff574.
//
// Solidity: function getName(uint256 _key) view returns(string)
func (_Contract *ContractSession) GetName(_key *big.Int) (string, error) {
	return _Contract.Contract.GetName(&_Contract.CallOpts, _key)
}

// GetName is a free data retrieval call binding the contract method 0x6b8ff574.
//
// Solidity: function getName(uint256 _key) view returns(string)
func (_Contract *ContractCallerSession) GetName(_key *big.Int) (string, error) {
	return _Contract.Contract.GetName(&_Contract.CallOpts, _key)
}

// GetPriceList is a free data retrieval call binding the contract method 0x5eac9098.
//
// Solidity: function getPriceList(string _key) view returns(uint256)
func (_Contract *ContractCaller) GetPriceList(opts *bind.CallOpts, _key string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPriceList", _key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceList is a free data retrieval call binding the contract method 0x5eac9098.
//
// Solidity: function getPriceList(string _key) view returns(uint256)
func (_Contract *ContractSession) GetPriceList(_key string) (*big.Int, error) {
	return _Contract.Contract.GetPriceList(&_Contract.CallOpts, _key)
}

// GetPriceList is a free data retrieval call binding the contract method 0x5eac9098.
//
// Solidity: function getPriceList(string _key) view returns(uint256)
func (_Contract *ContractCallerSession) GetPriceList(_key string) (*big.Int, error) {
	return _Contract.Contract.GetPriceList(&_Contract.CallOpts, _key)
}

// SetAgeList is a paid mutator transaction binding the contract method 0xfe093d3b.
//
// Solidity: function setAgeList(uint256 _key, uint256 _age) returns()
func (_Contract *ContractTransactor) SetAgeList(opts *bind.TransactOpts, _key *big.Int, _age *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setAgeList", _key, _age)
}

// SetAgeList is a paid mutator transaction binding the contract method 0xfe093d3b.
//
// Solidity: function setAgeList(uint256 _key, uint256 _age) returns()
func (_Contract *ContractSession) SetAgeList(_key *big.Int, _age *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetAgeList(&_Contract.TransactOpts, _key, _age)
}

// SetAgeList is a paid mutator transaction binding the contract method 0xfe093d3b.
//
// Solidity: function setAgeList(uint256 _key, uint256 _age) returns()
func (_Contract *ContractTransactorSession) SetAgeList(_key *big.Int, _age *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetAgeList(&_Contract.TransactOpts, _key, _age)
}

// SetNameList is a paid mutator transaction binding the contract method 0x1bbefaf1.
//
// Solidity: function setNameList(uint256 _key, string _name) returns()
func (_Contract *ContractTransactor) SetNameList(opts *bind.TransactOpts, _key *big.Int, _name string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setNameList", _key, _name)
}

// SetNameList is a paid mutator transaction binding the contract method 0x1bbefaf1.
//
// Solidity: function setNameList(uint256 _key, string _name) returns()
func (_Contract *ContractSession) SetNameList(_key *big.Int, _name string) (*types.Transaction, error) {
	return _Contract.Contract.SetNameList(&_Contract.TransactOpts, _key, _name)
}

// SetNameList is a paid mutator transaction binding the contract method 0x1bbefaf1.
//
// Solidity: function setNameList(uint256 _key, string _name) returns()
func (_Contract *ContractTransactorSession) SetNameList(_key *big.Int, _name string) (*types.Transaction, error) {
	return _Contract.Contract.SetNameList(&_Contract.TransactOpts, _key, _name)
}

// SetPriceList is a paid mutator transaction binding the contract method 0xa266f6c2.
//
// Solidity: function setPriceList(string _itemName, uint256 _price) returns()
func (_Contract *ContractTransactor) SetPriceList(opts *bind.TransactOpts, _itemName string, _price *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPriceList", _itemName, _price)
}

// SetPriceList is a paid mutator transaction binding the contract method 0xa266f6c2.
//
// Solidity: function setPriceList(string _itemName, uint256 _price) returns()
func (_Contract *ContractSession) SetPriceList(_itemName string, _price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetPriceList(&_Contract.TransactOpts, _itemName, _price)
}

// SetPriceList is a paid mutator transaction binding the contract method 0xa266f6c2.
//
// Solidity: function setPriceList(string _itemName, uint256 _price) returns()
func (_Contract *ContractTransactorSession) SetPriceList(_itemName string, _price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetPriceList(&_Contract.TransactOpts, _itemName, _price)
}
