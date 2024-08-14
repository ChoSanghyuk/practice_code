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

// ClientReceiptMetaData contains all meta data concerning the ClientReceipt contract.
var ClientReceiptMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610163806100206000396000f3fe60806040526004361061001e5760003560e01c8063b214faa514610023575b600080fd5b61003d600480360381019061003891906100a6565b61003f565b005b803373ffffffffffffffffffffffffffffffffffffffff167f19dacbf83c5de6658e14cbf7bcae5c15eca2eedecf1c66fbca928e4d351bea0f3460405161008691906100e2565b60405180910390a350565b6000813590506100a081610116565b92915050565b6000602082840312156100bc576100bb610111565b5b60006100ca84828501610091565b91505092915050565b6100dc81610107565b82525050565b60006020820190506100f760008301846100d3565b92915050565b6000819050919050565b6000819050919050565b600080fd5b61011f816100fd565b811461012a57600080fd5b5056fea26469706673582212205d22107b9d8ba4d652cf18fe6bec34810bbdd096d89305c4de620156580b525c64736f6c63430008060033",
}

// ClientReceiptABI is the input ABI used to generate the binding from.
// Deprecated: Use ClientReceiptMetaData.ABI instead.
var ClientReceiptABI = ClientReceiptMetaData.ABI

// ClientReceiptBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ClientReceiptMetaData.Bin instead.
var ClientReceiptBin = ClientReceiptMetaData.Bin

// DeployClientReceipt deploys a new Ethereum contract, binding an instance of ClientReceipt to it.
func DeployClientReceipt(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ClientReceipt, error) {
	parsed, err := ClientReceiptMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ClientReceiptBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClientReceipt{ClientReceiptCaller: ClientReceiptCaller{contract: contract}, ClientReceiptTransactor: ClientReceiptTransactor{contract: contract}, ClientReceiptFilterer: ClientReceiptFilterer{contract: contract}}, nil
}

// ClientReceipt is an auto generated Go binding around an Ethereum contract.
type ClientReceipt struct {
	ClientReceiptCaller     // Read-only binding to the contract
	ClientReceiptTransactor // Write-only binding to the contract
	ClientReceiptFilterer   // Log filterer for contract events
}

// ClientReceiptCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClientReceiptCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientReceiptTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClientReceiptTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientReceiptFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClientReceiptFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientReceiptSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClientReceiptSession struct {
	Contract     *ClientReceipt    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClientReceiptCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClientReceiptCallerSession struct {
	Contract *ClientReceiptCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ClientReceiptTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClientReceiptTransactorSession struct {
	Contract     *ClientReceiptTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ClientReceiptRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClientReceiptRaw struct {
	Contract *ClientReceipt // Generic contract binding to access the raw methods on
}

// ClientReceiptCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClientReceiptCallerRaw struct {
	Contract *ClientReceiptCaller // Generic read-only contract binding to access the raw methods on
}

// ClientReceiptTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClientReceiptTransactorRaw struct {
	Contract *ClientReceiptTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClientReceipt creates a new instance of ClientReceipt, bound to a specific deployed contract.
func NewClientReceipt(address common.Address, backend bind.ContractBackend) (*ClientReceipt, error) {
	contract, err := bindClientReceipt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClientReceipt{ClientReceiptCaller: ClientReceiptCaller{contract: contract}, ClientReceiptTransactor: ClientReceiptTransactor{contract: contract}, ClientReceiptFilterer: ClientReceiptFilterer{contract: contract}}, nil
}

// NewClientReceiptCaller creates a new read-only instance of ClientReceipt, bound to a specific deployed contract.
func NewClientReceiptCaller(address common.Address, caller bind.ContractCaller) (*ClientReceiptCaller, error) {
	contract, err := bindClientReceipt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClientReceiptCaller{contract: contract}, nil
}

// NewClientReceiptTransactor creates a new write-only instance of ClientReceipt, bound to a specific deployed contract.
func NewClientReceiptTransactor(address common.Address, transactor bind.ContractTransactor) (*ClientReceiptTransactor, error) {
	contract, err := bindClientReceipt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClientReceiptTransactor{contract: contract}, nil
}

// NewClientReceiptFilterer creates a new log filterer instance of ClientReceipt, bound to a specific deployed contract.
func NewClientReceiptFilterer(address common.Address, filterer bind.ContractFilterer) (*ClientReceiptFilterer, error) {
	contract, err := bindClientReceipt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClientReceiptFilterer{contract: contract}, nil
}

// bindClientReceipt binds a generic wrapper to an already deployed contract.
func bindClientReceipt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClientReceiptMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClientReceipt *ClientReceiptRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClientReceipt.Contract.ClientReceiptCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClientReceipt *ClientReceiptRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClientReceipt.Contract.ClientReceiptTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClientReceipt *ClientReceiptRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClientReceipt.Contract.ClientReceiptTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClientReceipt *ClientReceiptCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClientReceipt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClientReceipt *ClientReceiptTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClientReceipt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClientReceipt *ClientReceiptTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClientReceipt.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xb214faa5.
//
// Solidity: function deposit(bytes32 id) payable returns()
func (_ClientReceipt *ClientReceiptTransactor) Deposit(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _ClientReceipt.contract.Transact(opts, "deposit", id)
}

// Deposit is a paid mutator transaction binding the contract method 0xb214faa5.
//
// Solidity: function deposit(bytes32 id) payable returns()
func (_ClientReceipt *ClientReceiptSession) Deposit(id [32]byte) (*types.Transaction, error) {
	return _ClientReceipt.Contract.Deposit(&_ClientReceipt.TransactOpts, id)
}

// Deposit is a paid mutator transaction binding the contract method 0xb214faa5.
//
// Solidity: function deposit(bytes32 id) payable returns()
func (_ClientReceipt *ClientReceiptTransactorSession) Deposit(id [32]byte) (*types.Transaction, error) {
	return _ClientReceipt.Contract.Deposit(&_ClientReceipt.TransactOpts, id)
}

// ClientReceiptDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ClientReceipt contract.
type ClientReceiptDepositIterator struct {
	Event *ClientReceiptDeposit // Event containing the contract specifics and raw log

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
func (it *ClientReceiptDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClientReceiptDeposit)
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
		it.Event = new(ClientReceiptDeposit)
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
func (it *ClientReceiptDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClientReceiptDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClientReceiptDeposit represents a Deposit event raised by the ClientReceipt contract.
type ClientReceiptDeposit struct {
	From  common.Address
	Id    [32]byte
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x19dacbf83c5de6658e14cbf7bcae5c15eca2eedecf1c66fbca928e4d351bea0f.
//
// Solidity: event Deposit(address indexed from, bytes32 indexed id, uint256 value)
func (_ClientReceipt *ClientReceiptFilterer) FilterDeposit(opts *bind.FilterOpts, from []common.Address, id [][32]byte) (*ClientReceiptDepositIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ClientReceipt.contract.FilterLogs(opts, "Deposit", fromRule, idRule)
	if err != nil {
		return nil, err
	}
	return &ClientReceiptDepositIterator{contract: _ClientReceipt.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x19dacbf83c5de6658e14cbf7bcae5c15eca2eedecf1c66fbca928e4d351bea0f.
//
// Solidity: event Deposit(address indexed from, bytes32 indexed id, uint256 value)
func (_ClientReceipt *ClientReceiptFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ClientReceiptDeposit, from []common.Address, id [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ClientReceipt.contract.WatchLogs(opts, "Deposit", fromRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClientReceiptDeposit)
				if err := _ClientReceipt.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x19dacbf83c5de6658e14cbf7bcae5c15eca2eedecf1c66fbca928e4d351bea0f.
//
// Solidity: event Deposit(address indexed from, bytes32 indexed id, uint256 value)
func (_ClientReceipt *ClientReceiptFilterer) ParseDeposit(log types.Log) (*ClientReceiptDeposit, error) {
	event := new(ClientReceiptDeposit)
	if err := _ClientReceipt.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
