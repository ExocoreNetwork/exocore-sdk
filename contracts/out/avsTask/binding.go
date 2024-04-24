// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractavsTask

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

// ContractavsTaskMetaData contains all meta data concerning the ContractavsTask contract.
var ContractavsTaskMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkeyG1\",\"type\":\"bytes\"}],\"name\":\"NewPubkeyRegistration\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getRegisteredPubkey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"TaskContractAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"MetaInfo\",\"type\":\"string\"}],\"name\":\"registerAVSTask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"operator\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"}],\"name\":\"registerBLSPublicKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractavsTaskABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractavsTaskMetaData.ABI instead.
var ContractavsTaskABI = ContractavsTaskMetaData.ABI

// ContractavsTask is an auto generated Go binding around an Ethereum contract.
type ContractavsTask struct {
	ContractavsTaskCaller     // Read-only binding to the contract
	ContractavsTaskTransactor // Write-only binding to the contract
	ContractavsTaskFilterer   // Log filterer for contract events
}

// ContractavsTaskCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractavsTaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractavsTaskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractavsTaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractavsTaskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractavsTaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractavsTaskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractavsTaskSession struct {
	Contract     *ContractavsTask  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractavsTaskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractavsTaskCallerSession struct {
	Contract *ContractavsTaskCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ContractavsTaskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractavsTaskTransactorSession struct {
	Contract     *ContractavsTaskTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractavsTaskRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractavsTaskRaw struct {
	Contract *ContractavsTask // Generic contract binding to access the raw methods on
}

// ContractavsTaskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractavsTaskCallerRaw struct {
	Contract *ContractavsTaskCaller // Generic read-only contract binding to access the raw methods on
}

// ContractavsTaskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractavsTaskTransactorRaw struct {
	Contract *ContractavsTaskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractavsTask creates a new instance of ContractavsTask, bound to a specific deployed contract.
func NewContractavsTask(address common.Address, backend bind.ContractBackend) (*ContractavsTask, error) {
	contract, err := bindContractavsTask(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractavsTask{ContractavsTaskCaller: ContractavsTaskCaller{contract: contract}, ContractavsTaskTransactor: ContractavsTaskTransactor{contract: contract}, ContractavsTaskFilterer: ContractavsTaskFilterer{contract: contract}}, nil
}

// NewContractavsTaskCaller creates a new read-only instance of ContractavsTask, bound to a specific deployed contract.
func NewContractavsTaskCaller(address common.Address, caller bind.ContractCaller) (*ContractavsTaskCaller, error) {
	contract, err := bindContractavsTask(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractavsTaskCaller{contract: contract}, nil
}

// NewContractavsTaskTransactor creates a new write-only instance of ContractavsTask, bound to a specific deployed contract.
func NewContractavsTaskTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractavsTaskTransactor, error) {
	contract, err := bindContractavsTask(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractavsTaskTransactor{contract: contract}, nil
}

// NewContractavsTaskFilterer creates a new log filterer instance of ContractavsTask, bound to a specific deployed contract.
func NewContractavsTaskFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractavsTaskFilterer, error) {
	contract, err := bindContractavsTask(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractavsTaskFilterer{contract: contract}, nil
}

// bindContractavsTask binds a generic wrapper to an already deployed contract.
func bindContractavsTask(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractavsTaskMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractavsTask *ContractavsTaskRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractavsTask.Contract.ContractavsTaskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractavsTask *ContractavsTaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractavsTask.Contract.ContractavsTaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractavsTask *ContractavsTaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractavsTask.Contract.ContractavsTaskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractavsTask *ContractavsTaskCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractavsTask.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractavsTask *ContractavsTaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractavsTask.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractavsTask *ContractavsTaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractavsTask.Contract.contract.Transact(opts, method, params...)
}

// GetRegisteredPubkey is a paid mutator transaction binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) returns(bytes32)
func (_ContractavsTask *ContractavsTaskTransactor) GetRegisteredPubkey(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractavsTask.contract.Transact(opts, "getRegisteredPubkey", operator)
}

// GetRegisteredPubkey is a paid mutator transaction binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) returns(bytes32)
func (_ContractavsTask *ContractavsTaskSession) GetRegisteredPubkey(operator common.Address) (*types.Transaction, error) {
	return _ContractavsTask.Contract.GetRegisteredPubkey(&_ContractavsTask.TransactOpts, operator)
}

// GetRegisteredPubkey is a paid mutator transaction binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) returns(bytes32)
func (_ContractavsTask *ContractavsTaskTransactorSession) GetRegisteredPubkey(operator common.Address) (*types.Transaction, error) {
	return _ContractavsTask.Contract.GetRegisteredPubkey(&_ContractavsTask.TransactOpts, operator)
}

// RegisterAVSTask is a paid mutator transaction binding the contract method 0xfab952d6.
//
// Solidity: function registerAVSTask(string TaskContractAddress, string Name, string MetaInfo) returns(bool success)
func (_ContractavsTask *ContractavsTaskTransactor) RegisterAVSTask(opts *bind.TransactOpts, TaskContractAddress string, Name string, MetaInfo string) (*types.Transaction, error) {
	return _ContractavsTask.contract.Transact(opts, "registerAVSTask", TaskContractAddress, Name, MetaInfo)
}

// RegisterAVSTask is a paid mutator transaction binding the contract method 0xfab952d6.
//
// Solidity: function registerAVSTask(string TaskContractAddress, string Name, string MetaInfo) returns(bool success)
func (_ContractavsTask *ContractavsTaskSession) RegisterAVSTask(TaskContractAddress string, Name string, MetaInfo string) (*types.Transaction, error) {
	return _ContractavsTask.Contract.RegisterAVSTask(&_ContractavsTask.TransactOpts, TaskContractAddress, Name, MetaInfo)
}

// RegisterAVSTask is a paid mutator transaction binding the contract method 0xfab952d6.
//
// Solidity: function registerAVSTask(string TaskContractAddress, string Name, string MetaInfo) returns(bool success)
func (_ContractavsTask *ContractavsTaskTransactorSession) RegisterAVSTask(TaskContractAddress string, Name string, MetaInfo string) (*types.Transaction, error) {
	return _ContractavsTask.Contract.RegisterAVSTask(&_ContractavsTask.TransactOpts, TaskContractAddress, Name, MetaInfo)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xad1a2dbd.
//
// Solidity: function registerBLSPublicKey(string operator, bytes pubKey) returns(bool success)
func (_ContractavsTask *ContractavsTaskTransactor) RegisterBLSPublicKey(opts *bind.TransactOpts, operator string, pubKey []byte) (*types.Transaction, error) {
	return _ContractavsTask.contract.Transact(opts, "registerBLSPublicKey", operator, pubKey)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xad1a2dbd.
//
// Solidity: function registerBLSPublicKey(string operator, bytes pubKey) returns(bool success)
func (_ContractavsTask *ContractavsTaskSession) RegisterBLSPublicKey(operator string, pubKey []byte) (*types.Transaction, error) {
	return _ContractavsTask.Contract.RegisterBLSPublicKey(&_ContractavsTask.TransactOpts, operator, pubKey)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xad1a2dbd.
//
// Solidity: function registerBLSPublicKey(string operator, bytes pubKey) returns(bool success)
func (_ContractavsTask *ContractavsTaskTransactorSession) RegisterBLSPublicKey(operator string, pubKey []byte) (*types.Transaction, error) {
	return _ContractavsTask.Contract.RegisterBLSPublicKey(&_ContractavsTask.TransactOpts, operator, pubKey)
}

// ContractavsTaskNewPubkeyRegistrationIterator is returned from FilterNewPubkeyRegistration and is used to iterate over the raw logs and unpacked data for NewPubkeyRegistration events raised by the ContractavsTask contract.
type ContractavsTaskNewPubkeyRegistrationIterator struct {
	Event *ContractavsTaskNewPubkeyRegistration // Event containing the contract specifics and raw log

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
func (it *ContractavsTaskNewPubkeyRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractavsTaskNewPubkeyRegistration)
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
		it.Event = new(ContractavsTaskNewPubkeyRegistration)
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
func (it *ContractavsTaskNewPubkeyRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractavsTaskNewPubkeyRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractavsTaskNewPubkeyRegistration represents a NewPubkeyRegistration event raised by the ContractavsTask contract.
type ContractavsTaskNewPubkeyRegistration struct {
	Operator common.Address
	PubkeyG1 []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewPubkeyRegistration is a free log retrieval operation binding the contract event 0x3b2901eb6a9e6bb0aab2d26d8a660b938b499f70102f929cec7fb5a07d131879.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, bytes pubkeyG1)
func (_ContractavsTask *ContractavsTaskFilterer) FilterNewPubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*ContractavsTaskNewPubkeyRegistrationIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractavsTask.contract.FilterLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractavsTaskNewPubkeyRegistrationIterator{contract: _ContractavsTask.contract, event: "NewPubkeyRegistration", logs: logs, sub: sub}, nil
}

// WatchNewPubkeyRegistration is a free log subscription operation binding the contract event 0x3b2901eb6a9e6bb0aab2d26d8a660b938b499f70102f929cec7fb5a07d131879.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, bytes pubkeyG1)
func (_ContractavsTask *ContractavsTaskFilterer) WatchNewPubkeyRegistration(opts *bind.WatchOpts, sink chan<- *ContractavsTaskNewPubkeyRegistration, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractavsTask.contract.WatchLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractavsTaskNewPubkeyRegistration)
				if err := _ContractavsTask.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
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

// ParseNewPubkeyRegistration is a log parse operation binding the contract event 0x3b2901eb6a9e6bb0aab2d26d8a660b938b499f70102f929cec7fb5a07d131879.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, bytes pubkeyG1)
func (_ContractavsTask *ContractavsTaskFilterer) ParseNewPubkeyRegistration(log types.Log) (*ContractavsTaskNewPubkeyRegistration, error) {
	event := new(ContractavsTaskNewPubkeyRegistration)
	if err := _ContractavsTask.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
