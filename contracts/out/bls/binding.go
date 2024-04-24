// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractbls

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

// ContractblsMetaData contains all meta data concerning the Contractbls contract.
var ContractblsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubkey2\",\"type\":\"bytes\"}],\"name\":\"addTwoPubkeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"newPubkey\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"aggregatePubkeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"name\":\"aggregateSignatures\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msg_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"fastAggregateVerify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"generatePrivateKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"privkey\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"privkey\",\"type\":\"bytes\"}],\"name\":\"publicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"privkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"msg_\",\"type\":\"bytes32\"}],\"name\":\"sign\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msg_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// ContractblsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractblsMetaData.ABI instead.
var ContractblsABI = ContractblsMetaData.ABI

// Contractbls is an auto generated Go binding around an Ethereum contract.
type Contractbls struct {
	ContractblsCaller     // Read-only binding to the contract
	ContractblsTransactor // Write-only binding to the contract
	ContractblsFilterer   // Log filterer for contract events
}

// ContractblsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractblsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractblsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractblsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractblsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractblsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractblsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractblsSession struct {
	Contract     *Contractbls      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractblsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractblsCallerSession struct {
	Contract *ContractblsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ContractblsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractblsTransactorSession struct {
	Contract     *ContractblsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractblsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractblsRaw struct {
	Contract *Contractbls // Generic contract binding to access the raw methods on
}

// ContractblsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractblsCallerRaw struct {
	Contract *ContractblsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractblsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractblsTransactorRaw struct {
	Contract *ContractblsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractbls creates a new instance of Contractbls, bound to a specific deployed contract.
func NewContractbls(address common.Address, backend bind.ContractBackend) (*Contractbls, error) {
	contract, err := bindContractbls(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contractbls{ContractblsCaller: ContractblsCaller{contract: contract}, ContractblsTransactor: ContractblsTransactor{contract: contract}, ContractblsFilterer: ContractblsFilterer{contract: contract}}, nil
}

// NewContractblsCaller creates a new read-only instance of Contractbls, bound to a specific deployed contract.
func NewContractblsCaller(address common.Address, caller bind.ContractCaller) (*ContractblsCaller, error) {
	contract, err := bindContractbls(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractblsCaller{contract: contract}, nil
}

// NewContractblsTransactor creates a new write-only instance of Contractbls, bound to a specific deployed contract.
func NewContractblsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractblsTransactor, error) {
	contract, err := bindContractbls(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractblsTransactor{contract: contract}, nil
}

// NewContractblsFilterer creates a new log filterer instance of Contractbls, bound to a specific deployed contract.
func NewContractblsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractblsFilterer, error) {
	contract, err := bindContractbls(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractblsFilterer{contract: contract}, nil
}

// bindContractbls binds a generic wrapper to an already deployed contract.
func bindContractbls(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractblsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractbls *ContractblsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractbls.Contract.ContractblsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractbls *ContractblsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractbls.Contract.ContractblsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractbls *ContractblsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractbls.Contract.ContractblsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractbls *ContractblsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractbls.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractbls *ContractblsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractbls.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractbls *ContractblsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractbls.Contract.contract.Transact(opts, method, params...)
}

// AddTwoPubkeys is a free data retrieval call binding the contract method 0x594ba6de.
//
// Solidity: function addTwoPubkeys(bytes pubkey1, bytes pubkey2) pure returns(bytes newPubkey)
func (_Contractbls *ContractblsCaller) AddTwoPubkeys(opts *bind.CallOpts, pubkey1 []byte, pubkey2 []byte) ([]byte, error) {
	var out []interface{}
	err := _Contractbls.contract.Call(opts, &out, "addTwoPubkeys", pubkey1, pubkey2)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AddTwoPubkeys is a free data retrieval call binding the contract method 0x594ba6de.
//
// Solidity: function addTwoPubkeys(bytes pubkey1, bytes pubkey2) pure returns(bytes newPubkey)
func (_Contractbls *ContractblsSession) AddTwoPubkeys(pubkey1 []byte, pubkey2 []byte) ([]byte, error) {
	return _Contractbls.Contract.AddTwoPubkeys(&_Contractbls.CallOpts, pubkey1, pubkey2)
}

// AddTwoPubkeys is a free data retrieval call binding the contract method 0x594ba6de.
//
// Solidity: function addTwoPubkeys(bytes pubkey1, bytes pubkey2) pure returns(bytes newPubkey)
func (_Contractbls *ContractblsCallerSession) AddTwoPubkeys(pubkey1 []byte, pubkey2 []byte) ([]byte, error) {
	return _Contractbls.Contract.AddTwoPubkeys(&_Contractbls.CallOpts, pubkey1, pubkey2)
}

// AggregatePubkeys is a free data retrieval call binding the contract method 0xfa32cd08.
//
// Solidity: function aggregatePubkeys(bytes[] pubkeys) pure returns(bytes pubkey)
func (_Contractbls *ContractblsCaller) AggregatePubkeys(opts *bind.CallOpts, pubkeys [][]byte) ([]byte, error) {
	var out []interface{}
	err := _Contractbls.contract.Call(opts, &out, "aggregatePubkeys", pubkeys)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AggregatePubkeys is a free data retrieval call binding the contract method 0xfa32cd08.
//
// Solidity: function aggregatePubkeys(bytes[] pubkeys) pure returns(bytes pubkey)
func (_Contractbls *ContractblsSession) AggregatePubkeys(pubkeys [][]byte) ([]byte, error) {
	return _Contractbls.Contract.AggregatePubkeys(&_Contractbls.CallOpts, pubkeys)
}

// AggregatePubkeys is a free data retrieval call binding the contract method 0xfa32cd08.
//
// Solidity: function aggregatePubkeys(bytes[] pubkeys) pure returns(bytes pubkey)
func (_Contractbls *ContractblsCallerSession) AggregatePubkeys(pubkeys [][]byte) ([]byte, error) {
	return _Contractbls.Contract.AggregatePubkeys(&_Contractbls.CallOpts, pubkeys)
}

// AggregateSignatures is a free data retrieval call binding the contract method 0xa8309b9e.
//
// Solidity: function aggregateSignatures(bytes[] sigs) pure returns(bytes sig)
func (_Contractbls *ContractblsCaller) AggregateSignatures(opts *bind.CallOpts, sigs [][]byte) ([]byte, error) {
	var out []interface{}
	err := _Contractbls.contract.Call(opts, &out, "aggregateSignatures", sigs)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AggregateSignatures is a free data retrieval call binding the contract method 0xa8309b9e.
//
// Solidity: function aggregateSignatures(bytes[] sigs) pure returns(bytes sig)
func (_Contractbls *ContractblsSession) AggregateSignatures(sigs [][]byte) ([]byte, error) {
	return _Contractbls.Contract.AggregateSignatures(&_Contractbls.CallOpts, sigs)
}

// AggregateSignatures is a free data retrieval call binding the contract method 0xa8309b9e.
//
// Solidity: function aggregateSignatures(bytes[] sigs) pure returns(bytes sig)
func (_Contractbls *ContractblsCallerSession) AggregateSignatures(sigs [][]byte) ([]byte, error) {
	return _Contractbls.Contract.AggregateSignatures(&_Contractbls.CallOpts, sigs)
}

// FastAggregateVerify is a free data retrieval call binding the contract method 0x7ea5a4cf.
//
// Solidity: function fastAggregateVerify(bytes32 msg_, bytes signature, bytes[] pubkeys) pure returns(bool valid)
func (_Contractbls *ContractblsCaller) FastAggregateVerify(opts *bind.CallOpts, msg_ [32]byte, signature []byte, pubkeys [][]byte) (bool, error) {
	var out []interface{}
	err := _Contractbls.contract.Call(opts, &out, "fastAggregateVerify", msg_, signature, pubkeys)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FastAggregateVerify is a free data retrieval call binding the contract method 0x7ea5a4cf.
//
// Solidity: function fastAggregateVerify(bytes32 msg_, bytes signature, bytes[] pubkeys) pure returns(bool valid)
func (_Contractbls *ContractblsSession) FastAggregateVerify(msg_ [32]byte, signature []byte, pubkeys [][]byte) (bool, error) {
	return _Contractbls.Contract.FastAggregateVerify(&_Contractbls.CallOpts, msg_, signature, pubkeys)
}

// FastAggregateVerify is a free data retrieval call binding the contract method 0x7ea5a4cf.
//
// Solidity: function fastAggregateVerify(bytes32 msg_, bytes signature, bytes[] pubkeys) pure returns(bool valid)
func (_Contractbls *ContractblsCallerSession) FastAggregateVerify(msg_ [32]byte, signature []byte, pubkeys [][]byte) (bool, error) {
	return _Contractbls.Contract.FastAggregateVerify(&_Contractbls.CallOpts, msg_, signature, pubkeys)
}

// GeneratePrivateKey is a free data retrieval call binding the contract method 0x3504af4a.
//
// Solidity: function generatePrivateKey() pure returns(bytes privkey)
func (_Contractbls *ContractblsCaller) GeneratePrivateKey(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Contractbls.contract.Call(opts, &out, "generatePrivateKey")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GeneratePrivateKey is a free data retrieval call binding the contract method 0x3504af4a.
//
// Solidity: function generatePrivateKey() pure returns(bytes privkey)
func (_Contractbls *ContractblsSession) GeneratePrivateKey() ([]byte, error) {
	return _Contractbls.Contract.GeneratePrivateKey(&_Contractbls.CallOpts)
}

// GeneratePrivateKey is a free data retrieval call binding the contract method 0x3504af4a.
//
// Solidity: function generatePrivateKey() pure returns(bytes privkey)
func (_Contractbls *ContractblsCallerSession) GeneratePrivateKey() ([]byte, error) {
	return _Contractbls.Contract.GeneratePrivateKey(&_Contractbls.CallOpts)
}

// PublicKey is a free data retrieval call binding the contract method 0x7f00ac87.
//
// Solidity: function publicKey(bytes privkey) pure returns(bytes pubkey)
func (_Contractbls *ContractblsCaller) PublicKey(opts *bind.CallOpts, privkey []byte) ([]byte, error) {
	var out []interface{}
	err := _Contractbls.contract.Call(opts, &out, "publicKey", privkey)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PublicKey is a free data retrieval call binding the contract method 0x7f00ac87.
//
// Solidity: function publicKey(bytes privkey) pure returns(bytes pubkey)
func (_Contractbls *ContractblsSession) PublicKey(privkey []byte) ([]byte, error) {
	return _Contractbls.Contract.PublicKey(&_Contractbls.CallOpts, privkey)
}

// PublicKey is a free data retrieval call binding the contract method 0x7f00ac87.
//
// Solidity: function publicKey(bytes privkey) pure returns(bytes pubkey)
func (_Contractbls *ContractblsCallerSession) PublicKey(privkey []byte) ([]byte, error) {
	return _Contractbls.Contract.PublicKey(&_Contractbls.CallOpts, privkey)
}

// Sign is a free data retrieval call binding the contract method 0x0a1e2de8.
//
// Solidity: function sign(bytes privkey, bytes32 msg_) pure returns(bytes signature)
func (_Contractbls *ContractblsCaller) Sign(opts *bind.CallOpts, privkey []byte, msg_ [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Contractbls.contract.Call(opts, &out, "sign", privkey, msg_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Sign is a free data retrieval call binding the contract method 0x0a1e2de8.
//
// Solidity: function sign(bytes privkey, bytes32 msg_) pure returns(bytes signature)
func (_Contractbls *ContractblsSession) Sign(privkey []byte, msg_ [32]byte) ([]byte, error) {
	return _Contractbls.Contract.Sign(&_Contractbls.CallOpts, privkey, msg_)
}

// Sign is a free data retrieval call binding the contract method 0x0a1e2de8.
//
// Solidity: function sign(bytes privkey, bytes32 msg_) pure returns(bytes signature)
func (_Contractbls *ContractblsCallerSession) Sign(privkey []byte, msg_ [32]byte) ([]byte, error) {
	return _Contractbls.Contract.Sign(&_Contractbls.CallOpts, privkey, msg_)
}

// Verify is a free data retrieval call binding the contract method 0x5bf48e3a.
//
// Solidity: function verify(bytes32 msg_, bytes signature, bytes pubkey) pure returns(bool valid)
func (_Contractbls *ContractblsCaller) Verify(opts *bind.CallOpts, msg_ [32]byte, signature []byte, pubkey []byte) (bool, error) {
	var out []interface{}
	err := _Contractbls.contract.Call(opts, &out, "verify", msg_, signature, pubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x5bf48e3a.
//
// Solidity: function verify(bytes32 msg_, bytes signature, bytes pubkey) pure returns(bool valid)
func (_Contractbls *ContractblsSession) Verify(msg_ [32]byte, signature []byte, pubkey []byte) (bool, error) {
	return _Contractbls.Contract.Verify(&_Contractbls.CallOpts, msg_, signature, pubkey)
}

// Verify is a free data retrieval call binding the contract method 0x5bf48e3a.
//
// Solidity: function verify(bytes32 msg_, bytes signature, bytes pubkey) pure returns(bool valid)
func (_Contractbls *ContractblsCallerSession) Verify(msg_ [32]byte, signature []byte, pubkey []byte) (bool, error) {
	return _Contractbls.Contract.Verify(&_Contractbls.CallOpts, msg_, signature, pubkey)
}
