// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractreward

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

// ContractrewardMetaData contains all meta data concerning the Contractreward contract.
var ContractrewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"clientChainLzID\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"assetsAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"withdrawRewardAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"opAmount\",\"type\":\"uint256\"}],\"name\":\"claimReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"latestAssetState\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractrewardABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractrewardMetaData.ABI instead.
var ContractrewardABI = ContractrewardMetaData.ABI

// Contractreward is an auto generated Go binding around an Ethereum contract.
type Contractreward struct {
	ContractrewardCaller     // Read-only binding to the contract
	ContractrewardTransactor // Write-only binding to the contract
	ContractrewardFilterer   // Log filterer for contract events
}

// ContractrewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractrewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractrewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractrewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractrewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractrewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractrewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractrewardSession struct {
	Contract     *Contractreward   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractrewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractrewardCallerSession struct {
	Contract *ContractrewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ContractrewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractrewardTransactorSession struct {
	Contract     *ContractrewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ContractrewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractrewardRaw struct {
	Contract *Contractreward // Generic contract binding to access the raw methods on
}

// ContractrewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractrewardCallerRaw struct {
	Contract *ContractrewardCaller // Generic read-only contract binding to access the raw methods on
}

// ContractrewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractrewardTransactorRaw struct {
	Contract *ContractrewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractreward creates a new instance of Contractreward, bound to a specific deployed contract.
func NewContractreward(address common.Address, backend bind.ContractBackend) (*Contractreward, error) {
	contract, err := bindContractreward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contractreward{ContractrewardCaller: ContractrewardCaller{contract: contract}, ContractrewardTransactor: ContractrewardTransactor{contract: contract}, ContractrewardFilterer: ContractrewardFilterer{contract: contract}}, nil
}

// NewContractrewardCaller creates a new read-only instance of Contractreward, bound to a specific deployed contract.
func NewContractrewardCaller(address common.Address, caller bind.ContractCaller) (*ContractrewardCaller, error) {
	contract, err := bindContractreward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractrewardCaller{contract: contract}, nil
}

// NewContractrewardTransactor creates a new write-only instance of Contractreward, bound to a specific deployed contract.
func NewContractrewardTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractrewardTransactor, error) {
	contract, err := bindContractreward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractrewardTransactor{contract: contract}, nil
}

// NewContractrewardFilterer creates a new log filterer instance of Contractreward, bound to a specific deployed contract.
func NewContractrewardFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractrewardFilterer, error) {
	contract, err := bindContractreward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractrewardFilterer{contract: contract}, nil
}

// bindContractreward binds a generic wrapper to an already deployed contract.
func bindContractreward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractrewardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractreward *ContractrewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractreward.Contract.ContractrewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractreward *ContractrewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractreward.Contract.ContractrewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractreward *ContractrewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractreward.Contract.ContractrewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractreward *ContractrewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractreward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractreward *ContractrewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractreward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractreward *ContractrewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractreward.Contract.contract.Transact(opts, method, params...)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb339d734.
//
// Solidity: function claimReward(uint16 clientChainLzID, bytes assetsAddress, bytes withdrawRewardAddress, uint256 opAmount) returns(bool success, uint256 latestAssetState)
func (_Contractreward *ContractrewardTransactor) ClaimReward(opts *bind.TransactOpts, clientChainLzID uint16, assetsAddress []byte, withdrawRewardAddress []byte, opAmount *big.Int) (*types.Transaction, error) {
	return _Contractreward.contract.Transact(opts, "claimReward", clientChainLzID, assetsAddress, withdrawRewardAddress, opAmount)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb339d734.
//
// Solidity: function claimReward(uint16 clientChainLzID, bytes assetsAddress, bytes withdrawRewardAddress, uint256 opAmount) returns(bool success, uint256 latestAssetState)
func (_Contractreward *ContractrewardSession) ClaimReward(clientChainLzID uint16, assetsAddress []byte, withdrawRewardAddress []byte, opAmount *big.Int) (*types.Transaction, error) {
	return _Contractreward.Contract.ClaimReward(&_Contractreward.TransactOpts, clientChainLzID, assetsAddress, withdrawRewardAddress, opAmount)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb339d734.
//
// Solidity: function claimReward(uint16 clientChainLzID, bytes assetsAddress, bytes withdrawRewardAddress, uint256 opAmount) returns(bool success, uint256 latestAssetState)
func (_Contractreward *ContractrewardTransactorSession) ClaimReward(clientChainLzID uint16, assetsAddress []byte, withdrawRewardAddress []byte, opAmount *big.Int) (*types.Transaction, error) {
	return _Contractreward.Contract.ClaimReward(&_Contractreward.TransactOpts, clientChainLzID, assetsAddress, withdrawRewardAddress, opAmount)
}
