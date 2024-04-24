// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractdeposit

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

// ContractdepositMetaData contains all meta data concerning the Contractdeposit contract.
var ContractdepositMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"clientChainLzID\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"assetsAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"stakerAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"opAmount\",\"type\":\"uint256\"}],\"name\":\"depositTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"latestAssetState\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractdepositABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractdepositMetaData.ABI instead.
var ContractdepositABI = ContractdepositMetaData.ABI

// Contractdeposit is an auto generated Go binding around an Ethereum contract.
type Contractdeposit struct {
	ContractdepositCaller     // Read-only binding to the contract
	ContractdepositTransactor // Write-only binding to the contract
	ContractdepositFilterer   // Log filterer for contract events
}

// ContractdepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractdepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractdepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractdepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractdepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractdepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractdepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractdepositSession struct {
	Contract     *Contractdeposit  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractdepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractdepositCallerSession struct {
	Contract *ContractdepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ContractdepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractdepositTransactorSession struct {
	Contract     *ContractdepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractdepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractdepositRaw struct {
	Contract *Contractdeposit // Generic contract binding to access the raw methods on
}

// ContractdepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractdepositCallerRaw struct {
	Contract *ContractdepositCaller // Generic read-only contract binding to access the raw methods on
}

// ContractdepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractdepositTransactorRaw struct {
	Contract *ContractdepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractdeposit creates a new instance of Contractdeposit, bound to a specific deployed contract.
func NewContractdeposit(address common.Address, backend bind.ContractBackend) (*Contractdeposit, error) {
	contract, err := bindContractdeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contractdeposit{ContractdepositCaller: ContractdepositCaller{contract: contract}, ContractdepositTransactor: ContractdepositTransactor{contract: contract}, ContractdepositFilterer: ContractdepositFilterer{contract: contract}}, nil
}

// NewContractdepositCaller creates a new read-only instance of Contractdeposit, bound to a specific deployed contract.
func NewContractdepositCaller(address common.Address, caller bind.ContractCaller) (*ContractdepositCaller, error) {
	contract, err := bindContractdeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractdepositCaller{contract: contract}, nil
}

// NewContractdepositTransactor creates a new write-only instance of Contractdeposit, bound to a specific deployed contract.
func NewContractdepositTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractdepositTransactor, error) {
	contract, err := bindContractdeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractdepositTransactor{contract: contract}, nil
}

// NewContractdepositFilterer creates a new log filterer instance of Contractdeposit, bound to a specific deployed contract.
func NewContractdepositFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractdepositFilterer, error) {
	contract, err := bindContractdeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractdepositFilterer{contract: contract}, nil
}

// bindContractdeposit binds a generic wrapper to an already deployed contract.
func bindContractdeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractdepositMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractdeposit *ContractdepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractdeposit.Contract.ContractdepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractdeposit *ContractdepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractdeposit.Contract.ContractdepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractdeposit *ContractdepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractdeposit.Contract.ContractdepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractdeposit *ContractdepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractdeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractdeposit *ContractdepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractdeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractdeposit *ContractdepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractdeposit.Contract.contract.Transact(opts, method, params...)
}

// DepositTo is a paid mutator transaction binding the contract method 0x58bd9b81.
//
// Solidity: function depositTo(uint16 clientChainLzID, bytes assetsAddress, bytes stakerAddress, uint256 opAmount) returns(bool success, uint256 latestAssetState)
func (_Contractdeposit *ContractdepositTransactor) DepositTo(opts *bind.TransactOpts, clientChainLzID uint16, assetsAddress []byte, stakerAddress []byte, opAmount *big.Int) (*types.Transaction, error) {
	return _Contractdeposit.contract.Transact(opts, "depositTo", clientChainLzID, assetsAddress, stakerAddress, opAmount)
}

// DepositTo is a paid mutator transaction binding the contract method 0x58bd9b81.
//
// Solidity: function depositTo(uint16 clientChainLzID, bytes assetsAddress, bytes stakerAddress, uint256 opAmount) returns(bool success, uint256 latestAssetState)
func (_Contractdeposit *ContractdepositSession) DepositTo(clientChainLzID uint16, assetsAddress []byte, stakerAddress []byte, opAmount *big.Int) (*types.Transaction, error) {
	return _Contractdeposit.Contract.DepositTo(&_Contractdeposit.TransactOpts, clientChainLzID, assetsAddress, stakerAddress, opAmount)
}

// DepositTo is a paid mutator transaction binding the contract method 0x58bd9b81.
//
// Solidity: function depositTo(uint16 clientChainLzID, bytes assetsAddress, bytes stakerAddress, uint256 opAmount) returns(bool success, uint256 latestAssetState)
func (_Contractdeposit *ContractdepositTransactorSession) DepositTo(clientChainLzID uint16, assetsAddress []byte, stakerAddress []byte, opAmount *big.Int) (*types.Transaction, error) {
	return _Contractdeposit.Contract.DepositTo(&_Contractdeposit.TransactOpts, clientChainLzID, assetsAddress, stakerAddress, opAmount)
}
