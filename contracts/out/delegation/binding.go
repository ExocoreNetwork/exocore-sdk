// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractdelegation

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

// ContractdelegationMetaData contains all meta data concerning the Contractdelegation contract.
var ContractdelegationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"clientChainLzID\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"lzNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"assetsAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"stakerAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorAddr\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"opAmount\",\"type\":\"uint256\"}],\"name\":\"delegateToThroughClientChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"clientChainLzID\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"lzNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"assetsAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"stakerAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorAddr\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"opAmount\",\"type\":\"uint256\"}],\"name\":\"undelegateFromThroughClientChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractdelegationABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractdelegationMetaData.ABI instead.
var ContractdelegationABI = ContractdelegationMetaData.ABI

// Contractdelegation is an auto generated Go binding around an Ethereum contract.
type Contractdelegation struct {
	ContractdelegationCaller     // Read-only binding to the contract
	ContractdelegationTransactor // Write-only binding to the contract
	ContractdelegationFilterer   // Log filterer for contract events
}

// ContractdelegationCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractdelegationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractdelegationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractdelegationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractdelegationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractdelegationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractdelegationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractdelegationSession struct {
	Contract     *Contractdelegation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractdelegationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractdelegationCallerSession struct {
	Contract *ContractdelegationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContractdelegationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractdelegationTransactorSession struct {
	Contract     *ContractdelegationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContractdelegationRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractdelegationRaw struct {
	Contract *Contractdelegation // Generic contract binding to access the raw methods on
}

// ContractdelegationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractdelegationCallerRaw struct {
	Contract *ContractdelegationCaller // Generic read-only contract binding to access the raw methods on
}

// ContractdelegationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractdelegationTransactorRaw struct {
	Contract *ContractdelegationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractdelegation creates a new instance of Contractdelegation, bound to a specific deployed contract.
func NewContractdelegation(address common.Address, backend bind.ContractBackend) (*Contractdelegation, error) {
	contract, err := bindContractdelegation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contractdelegation{ContractdelegationCaller: ContractdelegationCaller{contract: contract}, ContractdelegationTransactor: ContractdelegationTransactor{contract: contract}, ContractdelegationFilterer: ContractdelegationFilterer{contract: contract}}, nil
}

// NewContractdelegationCaller creates a new read-only instance of Contractdelegation, bound to a specific deployed contract.
func NewContractdelegationCaller(address common.Address, caller bind.ContractCaller) (*ContractdelegationCaller, error) {
	contract, err := bindContractdelegation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractdelegationCaller{contract: contract}, nil
}

// NewContractdelegationTransactor creates a new write-only instance of Contractdelegation, bound to a specific deployed contract.
func NewContractdelegationTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractdelegationTransactor, error) {
	contract, err := bindContractdelegation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractdelegationTransactor{contract: contract}, nil
}

// NewContractdelegationFilterer creates a new log filterer instance of Contractdelegation, bound to a specific deployed contract.
func NewContractdelegationFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractdelegationFilterer, error) {
	contract, err := bindContractdelegation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractdelegationFilterer{contract: contract}, nil
}

// bindContractdelegation binds a generic wrapper to an already deployed contract.
func bindContractdelegation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractdelegationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractdelegation *ContractdelegationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractdelegation.Contract.ContractdelegationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractdelegation *ContractdelegationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractdelegation.Contract.ContractdelegationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractdelegation *ContractdelegationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractdelegation.Contract.ContractdelegationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractdelegation *ContractdelegationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractdelegation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractdelegation *ContractdelegationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractdelegation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractdelegation *ContractdelegationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractdelegation.Contract.contract.Transact(opts, method, params...)
}

// DelegateToThroughClientChain is a paid mutator transaction binding the contract method 0xedc32d0a.
//
// Solidity: function delegateToThroughClientChain(uint16 clientChainLzID, uint64 lzNonce, bytes assetsAddress, bytes stakerAddress, bytes operatorAddr, uint256 opAmount) returns(bool success)
func (_Contractdelegation *ContractdelegationTransactor) DelegateToThroughClientChain(opts *bind.TransactOpts, clientChainLzID uint16, lzNonce uint64, assetsAddress []byte, stakerAddress []byte, operatorAddr []byte, opAmount *big.Int) (*types.Transaction, error) {
	return _Contractdelegation.contract.Transact(opts, "delegateToThroughClientChain", clientChainLzID, lzNonce, assetsAddress, stakerAddress, operatorAddr, opAmount)
}

// DelegateToThroughClientChain is a paid mutator transaction binding the contract method 0xedc32d0a.
//
// Solidity: function delegateToThroughClientChain(uint16 clientChainLzID, uint64 lzNonce, bytes assetsAddress, bytes stakerAddress, bytes operatorAddr, uint256 opAmount) returns(bool success)
func (_Contractdelegation *ContractdelegationSession) DelegateToThroughClientChain(clientChainLzID uint16, lzNonce uint64, assetsAddress []byte, stakerAddress []byte, operatorAddr []byte, opAmount *big.Int) (*types.Transaction, error) {
	return _Contractdelegation.Contract.DelegateToThroughClientChain(&_Contractdelegation.TransactOpts, clientChainLzID, lzNonce, assetsAddress, stakerAddress, operatorAddr, opAmount)
}

// DelegateToThroughClientChain is a paid mutator transaction binding the contract method 0xedc32d0a.
//
// Solidity: function delegateToThroughClientChain(uint16 clientChainLzID, uint64 lzNonce, bytes assetsAddress, bytes stakerAddress, bytes operatorAddr, uint256 opAmount) returns(bool success)
func (_Contractdelegation *ContractdelegationTransactorSession) DelegateToThroughClientChain(clientChainLzID uint16, lzNonce uint64, assetsAddress []byte, stakerAddress []byte, operatorAddr []byte, opAmount *big.Int) (*types.Transaction, error) {
	return _Contractdelegation.Contract.DelegateToThroughClientChain(&_Contractdelegation.TransactOpts, clientChainLzID, lzNonce, assetsAddress, stakerAddress, operatorAddr, opAmount)
}

// UndelegateFromThroughClientChain is a paid mutator transaction binding the contract method 0x81d27842.
//
// Solidity: function undelegateFromThroughClientChain(uint16 clientChainLzID, uint64 lzNonce, bytes assetsAddress, bytes stakerAddress, bytes operatorAddr, uint256 opAmount) returns(bool success)
func (_Contractdelegation *ContractdelegationTransactor) UndelegateFromThroughClientChain(opts *bind.TransactOpts, clientChainLzID uint16, lzNonce uint64, assetsAddress []byte, stakerAddress []byte, operatorAddr []byte, opAmount *big.Int) (*types.Transaction, error) {
	return _Contractdelegation.contract.Transact(opts, "undelegateFromThroughClientChain", clientChainLzID, lzNonce, assetsAddress, stakerAddress, operatorAddr, opAmount)
}

// UndelegateFromThroughClientChain is a paid mutator transaction binding the contract method 0x81d27842.
//
// Solidity: function undelegateFromThroughClientChain(uint16 clientChainLzID, uint64 lzNonce, bytes assetsAddress, bytes stakerAddress, bytes operatorAddr, uint256 opAmount) returns(bool success)
func (_Contractdelegation *ContractdelegationSession) UndelegateFromThroughClientChain(clientChainLzID uint16, lzNonce uint64, assetsAddress []byte, stakerAddress []byte, operatorAddr []byte, opAmount *big.Int) (*types.Transaction, error) {
	return _Contractdelegation.Contract.UndelegateFromThroughClientChain(&_Contractdelegation.TransactOpts, clientChainLzID, lzNonce, assetsAddress, stakerAddress, operatorAddr, opAmount)
}

// UndelegateFromThroughClientChain is a paid mutator transaction binding the contract method 0x81d27842.
//
// Solidity: function undelegateFromThroughClientChain(uint16 clientChainLzID, uint64 lzNonce, bytes assetsAddress, bytes stakerAddress, bytes operatorAddr, uint256 opAmount) returns(bool success)
func (_Contractdelegation *ContractdelegationTransactorSession) UndelegateFromThroughClientChain(clientChainLzID uint16, lzNonce uint64, assetsAddress []byte, stakerAddress []byte, operatorAddr []byte, opAmount *big.Int) (*types.Transaction, error) {
	return _Contractdelegation.Contract.UndelegateFromThroughClientChain(&_Contractdelegation.TransactOpts, clientChainLzID, lzNonce, assetsAddress, stakerAddress, operatorAddr, opAmount)
}
