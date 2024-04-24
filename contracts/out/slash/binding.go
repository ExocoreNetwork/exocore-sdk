// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractslash

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

// ContractslashMetaData contains all meta data concerning the Contractslash contract.
var ContractslashMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"clientChainLzID\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"assetsAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"stakerAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"opAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"operatorAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"middlewareContractAddress\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"proportion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"}],\"name\":\"submitSlash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractslashABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractslashMetaData.ABI instead.
var ContractslashABI = ContractslashMetaData.ABI

// Contractslash is an auto generated Go binding around an Ethereum contract.
type Contractslash struct {
	ContractslashCaller     // Read-only binding to the contract
	ContractslashTransactor // Write-only binding to the contract
	ContractslashFilterer   // Log filterer for contract events
}

// ContractslashCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractslashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractslashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractslashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractslashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractslashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractslashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractslashSession struct {
	Contract     *Contractslash    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractslashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractslashCallerSession struct {
	Contract *ContractslashCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ContractslashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractslashTransactorSession struct {
	Contract     *ContractslashTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ContractslashRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractslashRaw struct {
	Contract *Contractslash // Generic contract binding to access the raw methods on
}

// ContractslashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractslashCallerRaw struct {
	Contract *ContractslashCaller // Generic read-only contract binding to access the raw methods on
}

// ContractslashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractslashTransactorRaw struct {
	Contract *ContractslashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractslash creates a new instance of Contractslash, bound to a specific deployed contract.
func NewContractslash(address common.Address, backend bind.ContractBackend) (*Contractslash, error) {
	contract, err := bindContractslash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contractslash{ContractslashCaller: ContractslashCaller{contract: contract}, ContractslashTransactor: ContractslashTransactor{contract: contract}, ContractslashFilterer: ContractslashFilterer{contract: contract}}, nil
}

// NewContractslashCaller creates a new read-only instance of Contractslash, bound to a specific deployed contract.
func NewContractslashCaller(address common.Address, caller bind.ContractCaller) (*ContractslashCaller, error) {
	contract, err := bindContractslash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractslashCaller{contract: contract}, nil
}

// NewContractslashTransactor creates a new write-only instance of Contractslash, bound to a specific deployed contract.
func NewContractslashTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractslashTransactor, error) {
	contract, err := bindContractslash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractslashTransactor{contract: contract}, nil
}

// NewContractslashFilterer creates a new log filterer instance of Contractslash, bound to a specific deployed contract.
func NewContractslashFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractslashFilterer, error) {
	contract, err := bindContractslash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractslashFilterer{contract: contract}, nil
}

// bindContractslash binds a generic wrapper to an already deployed contract.
func bindContractslash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractslashMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractslash *ContractslashRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractslash.Contract.ContractslashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractslash *ContractslashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractslash.Contract.ContractslashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractslash *ContractslashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractslash.Contract.ContractslashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractslash *ContractslashCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractslash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractslash *ContractslashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractslash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractslash *ContractslashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractslash.Contract.contract.Transact(opts, method, params...)
}

// SubmitSlash is a paid mutator transaction binding the contract method 0x8efb5427.
//
// Solidity: function submitSlash(uint16 clientChainLzID, bytes assetsAddress, bytes stakerAddress, uint256 opAmount, bytes operatorAddress, bytes middlewareContractAddress, string proportion, string proof) returns(bool success)
func (_Contractslash *ContractslashTransactor) SubmitSlash(opts *bind.TransactOpts, clientChainLzID uint16, assetsAddress []byte, stakerAddress []byte, opAmount *big.Int, operatorAddress []byte, middlewareContractAddress []byte, proportion string, proof string) (*types.Transaction, error) {
	return _Contractslash.contract.Transact(opts, "submitSlash", clientChainLzID, assetsAddress, stakerAddress, opAmount, operatorAddress, middlewareContractAddress, proportion, proof)
}

// SubmitSlash is a paid mutator transaction binding the contract method 0x8efb5427.
//
// Solidity: function submitSlash(uint16 clientChainLzID, bytes assetsAddress, bytes stakerAddress, uint256 opAmount, bytes operatorAddress, bytes middlewareContractAddress, string proportion, string proof) returns(bool success)
func (_Contractslash *ContractslashSession) SubmitSlash(clientChainLzID uint16, assetsAddress []byte, stakerAddress []byte, opAmount *big.Int, operatorAddress []byte, middlewareContractAddress []byte, proportion string, proof string) (*types.Transaction, error) {
	return _Contractslash.Contract.SubmitSlash(&_Contractslash.TransactOpts, clientChainLzID, assetsAddress, stakerAddress, opAmount, operatorAddress, middlewareContractAddress, proportion, proof)
}

// SubmitSlash is a paid mutator transaction binding the contract method 0x8efb5427.
//
// Solidity: function submitSlash(uint16 clientChainLzID, bytes assetsAddress, bytes stakerAddress, uint256 opAmount, bytes operatorAddress, bytes middlewareContractAddress, string proportion, string proof) returns(bool success)
func (_Contractslash *ContractslashTransactorSession) SubmitSlash(clientChainLzID uint16, assetsAddress []byte, stakerAddress []byte, opAmount *big.Int, operatorAddress []byte, middlewareContractAddress []byte, proportion string, proof string) (*types.Transaction, error) {
	return _Contractslash.Contract.SubmitSlash(&_Contractslash.TransactOpts, clientChainLzID, assetsAddress, stakerAddress, opAmount, operatorAddress, middlewareContractAddress, proportion, proof)
}
