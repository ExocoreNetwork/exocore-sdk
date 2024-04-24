// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractwithdraw

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

// ContractwithdrawMetaData contains all meta data concerning the Contractwithdraw contract.
var ContractwithdrawMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"clientChainLzID\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"assetsAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"withdrawAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"opAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawPrinciple\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"latestAssetState\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractwithdrawABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractwithdrawMetaData.ABI instead.
var ContractwithdrawABI = ContractwithdrawMetaData.ABI

// Contractwithdraw is an auto generated Go binding around an Ethereum contract.
type Contractwithdraw struct {
	ContractwithdrawCaller     // Read-only binding to the contract
	ContractwithdrawTransactor // Write-only binding to the contract
	ContractwithdrawFilterer   // Log filterer for contract events
}

// ContractwithdrawCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractwithdrawCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractwithdrawTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractwithdrawTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractwithdrawFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractwithdrawFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractwithdrawSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractwithdrawSession struct {
	Contract     *Contractwithdraw // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractwithdrawCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractwithdrawCallerSession struct {
	Contract *ContractwithdrawCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ContractwithdrawTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractwithdrawTransactorSession struct {
	Contract     *ContractwithdrawTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractwithdrawRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractwithdrawRaw struct {
	Contract *Contractwithdraw // Generic contract binding to access the raw methods on
}

// ContractwithdrawCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractwithdrawCallerRaw struct {
	Contract *ContractwithdrawCaller // Generic read-only contract binding to access the raw methods on
}

// ContractwithdrawTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractwithdrawTransactorRaw struct {
	Contract *ContractwithdrawTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractwithdraw creates a new instance of Contractwithdraw, bound to a specific deployed contract.
func NewContractwithdraw(address common.Address, backend bind.ContractBackend) (*Contractwithdraw, error) {
	contract, err := bindContractwithdraw(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contractwithdraw{ContractwithdrawCaller: ContractwithdrawCaller{contract: contract}, ContractwithdrawTransactor: ContractwithdrawTransactor{contract: contract}, ContractwithdrawFilterer: ContractwithdrawFilterer{contract: contract}}, nil
}

// NewContractwithdrawCaller creates a new read-only instance of Contractwithdraw, bound to a specific deployed contract.
func NewContractwithdrawCaller(address common.Address, caller bind.ContractCaller) (*ContractwithdrawCaller, error) {
	contract, err := bindContractwithdraw(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractwithdrawCaller{contract: contract}, nil
}

// NewContractwithdrawTransactor creates a new write-only instance of Contractwithdraw, bound to a specific deployed contract.
func NewContractwithdrawTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractwithdrawTransactor, error) {
	contract, err := bindContractwithdraw(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractwithdrawTransactor{contract: contract}, nil
}

// NewContractwithdrawFilterer creates a new log filterer instance of Contractwithdraw, bound to a specific deployed contract.
func NewContractwithdrawFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractwithdrawFilterer, error) {
	contract, err := bindContractwithdraw(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractwithdrawFilterer{contract: contract}, nil
}

// bindContractwithdraw binds a generic wrapper to an already deployed contract.
func bindContractwithdraw(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractwithdrawMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractwithdraw *ContractwithdrawRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractwithdraw.Contract.ContractwithdrawCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractwithdraw *ContractwithdrawRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractwithdraw.Contract.ContractwithdrawTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractwithdraw *ContractwithdrawRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractwithdraw.Contract.ContractwithdrawTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractwithdraw *ContractwithdrawCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractwithdraw.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractwithdraw *ContractwithdrawTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractwithdraw.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractwithdraw *ContractwithdrawTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractwithdraw.Contract.contract.Transact(opts, method, params...)
}

// WithdrawPrinciple is a paid mutator transaction binding the contract method 0xcfcd2269.
//
// Solidity: function withdrawPrinciple(uint16 clientChainLzID, bytes assetsAddress, bytes withdrawAddress, uint256 opAmount) returns(bool success, uint256 latestAssetState)
func (_Contractwithdraw *ContractwithdrawTransactor) WithdrawPrinciple(opts *bind.TransactOpts, clientChainLzID uint16, assetsAddress []byte, withdrawAddress []byte, opAmount *big.Int) (*types.Transaction, error) {
	return _Contractwithdraw.contract.Transact(opts, "withdrawPrinciple", clientChainLzID, assetsAddress, withdrawAddress, opAmount)
}

// WithdrawPrinciple is a paid mutator transaction binding the contract method 0xcfcd2269.
//
// Solidity: function withdrawPrinciple(uint16 clientChainLzID, bytes assetsAddress, bytes withdrawAddress, uint256 opAmount) returns(bool success, uint256 latestAssetState)
func (_Contractwithdraw *ContractwithdrawSession) WithdrawPrinciple(clientChainLzID uint16, assetsAddress []byte, withdrawAddress []byte, opAmount *big.Int) (*types.Transaction, error) {
	return _Contractwithdraw.Contract.WithdrawPrinciple(&_Contractwithdraw.TransactOpts, clientChainLzID, assetsAddress, withdrawAddress, opAmount)
}

// WithdrawPrinciple is a paid mutator transaction binding the contract method 0xcfcd2269.
//
// Solidity: function withdrawPrinciple(uint16 clientChainLzID, bytes assetsAddress, bytes withdrawAddress, uint256 opAmount) returns(bool success, uint256 latestAssetState)
func (_Contractwithdraw *ContractwithdrawTransactorSession) WithdrawPrinciple(clientChainLzID uint16, assetsAddress []byte, withdrawAddress []byte, opAmount *big.Int) (*types.Transaction, error) {
	return _Contractwithdraw.Contract.WithdrawPrinciple(&_Contractwithdraw.TransactOpts, clientChainLzID, assetsAddress, withdrawAddress, opAmount)
}
