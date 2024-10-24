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

// MyProxyMetaData contains all meta data concerning the MyProxy contract.
var MyProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161077e38038061077e83398181016040528101906100319190610502565b8181610043828261004c60201b60201c565b505050506105de565b61005b826100d060201b60201c565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f815111156100bd576100b7828261019f60201b60201c565b506100cc565b6100cb61022560201b60201c565b5b5050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b0361012b57806040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401610122919061056b565b60405180910390fd5b8061015d7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b61026160201b60201c565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f808473ffffffffffffffffffffffffffffffffffffffff16846040516101c891906105c8565b5f60405180830381855af49150503d805f8114610200576040519150601f19603f3d011682016040523d82523d5f602084013e610205565b606091505b509150915061021b85838361026a60201b60201c565b9250505092915050565b5f34111561025f576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f819050919050565b60608261028557610280826102fd60201b60201c565b6102f5565b5f82511480156102ab57505f8473ffffffffffffffffffffffffffffffffffffffff163b145b156102ed57836040517f9996b3150000000000000000000000000000000000000000000000000000000081526004016102e4919061056b565b60405180910390fd5b8190506102f6565b5b9392505050565b5f8151111561030f5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61037b82610352565b9050919050565b61038b81610371565b8114610395575f80fd5b50565b5f815190506103a681610382565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6103fa826103b4565b810181811067ffffffffffffffff82111715610419576104186103c4565b5b80604052505050565b5f61042b610341565b905061043782826103f1565b919050565b5f67ffffffffffffffff821115610456576104556103c4565b5b61045f826103b4565b9050602081019050919050565b5f5b8381101561048957808201518184015260208101905061046e565b5f8484015250505050565b5f6104a66104a18461043c565b610422565b9050828152602081018484840111156104c2576104c16103b0565b5b6104cd84828561046c565b509392505050565b5f82601f8301126104e9576104e86103ac565b5b81516104f9848260208601610494565b91505092915050565b5f80604083850312156105185761051761034a565b5b5f61052585828601610398565b925050602083015167ffffffffffffffff8111156105465761054561034e565b5b610552858286016104d5565b9150509250929050565b61056581610371565b82525050565b5f60208201905061057e5f83018461055c565b92915050565b5f81519050919050565b5f81905092915050565b5f6105a282610584565b6105ac818561058e565b93506105bc81856020860161046c565b80840191505092915050565b5f6105d38284610598565b915081905092915050565b610193806105eb5f395ff3fe608060405260043610610021575f3560e01c8063aaf10f421461003257610028565b3661002857005b61003061005c565b005b34801561003d575f80fd5b5061004661006e565b6040516100539190610144565b60405180910390f35b61006c61006761007c565b61008a565b565b5f61007761007c565b905090565b5f6100856100a9565b905090565b365f80375f80365f845af43d5f803e805f81146100a5573d5ff35b3d5ffd5b5f6100d57f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b6100fc565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f819050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61012e82610105565b9050919050565b61013e81610124565b82525050565b5f6020820190506101575f830184610135565b9291505056fea2646970667358221220bb83c8d142d3d1ae358335532fcf34349ebd3aee0c0c10758350def71514436a64736f6c63430008140033",
}

// MyProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use MyProxyMetaData.ABI instead.
var MyProxyABI = MyProxyMetaData.ABI

// MyProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyProxyMetaData.Bin instead.
var MyProxyBin = MyProxyMetaData.Bin

// DeployMyProxy deploys a new Ethereum contract, binding an instance of MyProxy to it.
func DeployMyProxy(auth *bind.TransactOpts, backend bind.ContractBackend, implementation common.Address, data []byte) (common.Address, *types.Transaction, *MyProxy, error) {
	parsed, err := MyProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyProxyBin), backend, implementation, data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyProxy{MyProxyCaller: MyProxyCaller{contract: contract}, MyProxyTransactor: MyProxyTransactor{contract: contract}, MyProxyFilterer: MyProxyFilterer{contract: contract}}, nil
}

// MyProxy is an auto generated Go binding around an Ethereum contract.
type MyProxy struct {
	MyProxyCaller     // Read-only binding to the contract
	MyProxyTransactor // Write-only binding to the contract
	MyProxyFilterer   // Log filterer for contract events
}

// MyProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyProxySession struct {
	Contract     *MyProxy   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyProxyCallerSession struct {
	Contract *MyProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MyProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyProxyTransactorSession struct {
	Contract     *MyProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MyProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyProxyRaw struct {
	Contract *MyProxy // Generic contract binding to access the raw methods on
}

// MyProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyProxyCallerRaw struct {
	Contract *MyProxyCaller // Generic read-only contract binding to access the raw methods on
}

// MyProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyProxyTransactorRaw struct {
	Contract *MyProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyProxy creates a new instance of MyProxy, bound to a specific deployed contract.
func NewMyProxy(address common.Address, backend bind.ContractBackend) (*MyProxy, error) {
	contract, err := bindMyProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyProxy{MyProxyCaller: MyProxyCaller{contract: contract}, MyProxyTransactor: MyProxyTransactor{contract: contract}, MyProxyFilterer: MyProxyFilterer{contract: contract}}, nil
}

// NewMyProxyCaller creates a new read-only instance of MyProxy, bound to a specific deployed contract.
func NewMyProxyCaller(address common.Address, caller bind.ContractCaller) (*MyProxyCaller, error) {
	contract, err := bindMyProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyProxyCaller{contract: contract}, nil
}

// NewMyProxyTransactor creates a new write-only instance of MyProxy, bound to a specific deployed contract.
func NewMyProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*MyProxyTransactor, error) {
	contract, err := bindMyProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyProxyTransactor{contract: contract}, nil
}

// NewMyProxyFilterer creates a new log filterer instance of MyProxy, bound to a specific deployed contract.
func NewMyProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*MyProxyFilterer, error) {
	contract, err := bindMyProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyProxyFilterer{contract: contract}, nil
}

// bindMyProxy binds a generic wrapper to an already deployed contract.
func bindMyProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MyProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyProxy *MyProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyProxy.Contract.MyProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyProxy *MyProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyProxy.Contract.MyProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyProxy *MyProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyProxy.Contract.MyProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyProxy *MyProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyProxy *MyProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyProxy *MyProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyProxy.Contract.contract.Transact(opts, method, params...)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_MyProxy *MyProxyCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyProxy.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_MyProxy *MyProxySession) GetImplementation() (common.Address, error) {
	return _MyProxy.Contract.GetImplementation(&_MyProxy.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_MyProxy *MyProxyCallerSession) GetImplementation() (common.Address, error) {
	return _MyProxy.Contract.GetImplementation(&_MyProxy.CallOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MyProxy *MyProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _MyProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MyProxy *MyProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MyProxy.Contract.Fallback(&_MyProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MyProxy *MyProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MyProxy.Contract.Fallback(&_MyProxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MyProxy *MyProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MyProxy *MyProxySession) Receive() (*types.Transaction, error) {
	return _MyProxy.Contract.Receive(&_MyProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MyProxy *MyProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _MyProxy.Contract.Receive(&_MyProxy.TransactOpts)
}

// MyProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the MyProxy contract.
type MyProxyUpgradedIterator struct {
	Event *MyProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *MyProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyProxyUpgraded)
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
		it.Event = new(MyProxyUpgraded)
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
func (it *MyProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyProxyUpgraded represents a Upgraded event raised by the MyProxy contract.
type MyProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MyProxy *MyProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MyProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MyProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MyProxyUpgradedIterator{contract: _MyProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MyProxy *MyProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MyProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MyProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyProxyUpgraded)
				if err := _MyProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_MyProxy *MyProxyFilterer) ParseUpgraded(log types.Log) (*MyProxyUpgraded, error) {
	event := new(MyProxyUpgraded)
	if err := _MyProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
