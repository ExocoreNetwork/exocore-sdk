// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractavs

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

// ContractavsMetaData contains all meta data concerning the Contractavs contract.
var ContractavsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"avsAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"operatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"action\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"avsOwnerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"assetID\",\"type\":\"string\"}],\"name\":\"AVSAction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractavsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractavsMetaData.ABI instead.
var ContractavsABI = ContractavsMetaData.ABI

// Contractavs is an auto generated Go binding around an Ethereum contract.
type Contractavs struct {
	ContractavsCaller     // Read-only binding to the contract
	ContractavsTransactor // Write-only binding to the contract
	ContractavsFilterer   // Log filterer for contract events
}

// ContractavsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractavsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractavsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractavsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractavsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractavsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractavsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractavsSession struct {
	Contract     *Contractavs      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractavsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractavsCallerSession struct {
	Contract *ContractavsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ContractavsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractavsTransactorSession struct {
	Contract     *ContractavsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractavsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractavsRaw struct {
	Contract *Contractavs // Generic contract binding to access the raw methods on
}

// ContractavsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractavsCallerRaw struct {
	Contract *ContractavsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractavsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractavsTransactorRaw struct {
	Contract *ContractavsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractavs creates a new instance of Contractavs, bound to a specific deployed contract.
func NewContractavs(address common.Address, backend bind.ContractBackend) (*Contractavs, error) {
	contract, err := bindContractavs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contractavs{ContractavsCaller: ContractavsCaller{contract: contract}, ContractavsTransactor: ContractavsTransactor{contract: contract}, ContractavsFilterer: ContractavsFilterer{contract: contract}}, nil
}

// NewContractavsCaller creates a new read-only instance of Contractavs, bound to a specific deployed contract.
func NewContractavsCaller(address common.Address, caller bind.ContractCaller) (*ContractavsCaller, error) {
	contract, err := bindContractavs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractavsCaller{contract: contract}, nil
}

// NewContractavsTransactor creates a new write-only instance of Contractavs, bound to a specific deployed contract.
func NewContractavsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractavsTransactor, error) {
	contract, err := bindContractavs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractavsTransactor{contract: contract}, nil
}

// NewContractavsFilterer creates a new log filterer instance of Contractavs, bound to a specific deployed contract.
func NewContractavsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractavsFilterer, error) {
	contract, err := bindContractavs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractavsFilterer{contract: contract}, nil
}

// bindContractavs binds a generic wrapper to an already deployed contract.
func bindContractavs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractavsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractavs *ContractavsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractavs.Contract.ContractavsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractavs *ContractavsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractavs.Contract.ContractavsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractavs *ContractavsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractavs.Contract.ContractavsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractavs *ContractavsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractavs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractavs *ContractavsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractavs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractavs *ContractavsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractavs.Contract.contract.Transact(opts, method, params...)
}

// AVSAction is a paid mutator transaction binding the contract method 0xb6c1c1ca.
//
// Solidity: function AVSAction(string avsName, string avsAddress, string operatorAddress, uint64 action, string avsOwnerAddress, string assetID) returns(bool success)
func (_Contractavs *ContractavsTransactor) AVSAction(opts *bind.TransactOpts, avsName string, avsAddress string, operatorAddress string, action uint64, avsOwnerAddress string, assetID string) (*types.Transaction, error) {
	return _Contractavs.contract.Transact(opts, "AVSAction", avsName, avsAddress, operatorAddress, action, avsOwnerAddress, assetID)
}

// AVSAction is a paid mutator transaction binding the contract method 0xb6c1c1ca.
//
// Solidity: function AVSAction(string avsName, string avsAddress, string operatorAddress, uint64 action, string avsOwnerAddress, string assetID) returns(bool success)
func (_Contractavs *ContractavsSession) AVSAction(avsName string, avsAddress string, operatorAddress string, action uint64, avsOwnerAddress string, assetID string) (*types.Transaction, error) {
	return _Contractavs.Contract.AVSAction(&_Contractavs.TransactOpts, avsName, avsAddress, operatorAddress, action, avsOwnerAddress, assetID)
}

// AVSAction is a paid mutator transaction binding the contract method 0xb6c1c1ca.
//
// Solidity: function AVSAction(string avsName, string avsAddress, string operatorAddress, uint64 action, string avsOwnerAddress, string assetID) returns(bool success)
func (_Contractavs *ContractavsTransactorSession) AVSAction(avsName string, avsAddress string, operatorAddress string, action uint64, avsOwnerAddress string, assetID string) (*types.Transaction, error) {
	return _Contractavs.Contract.AVSAction(&_Contractavs.TransactOpts, avsName, avsAddress, operatorAddress, action, avsOwnerAddress, assetID)
}
