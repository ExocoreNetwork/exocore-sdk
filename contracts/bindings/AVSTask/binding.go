// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractAVSTask

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

// AVSTaskTask is an auto generated low-level Go binding around an user-defined struct.
type AVSTaskTask struct {
	NumberToBeSquared     *big.Int
	Name                  string
	TaskResponsePeriod    uint64
	TaskChallengePeriod   uint64
	ThresholdPercentage   uint64
	TaskStatisticalPeriod uint64
}

// AVSTaskTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type AVSTaskTaskResponse struct {
	TaskID        uint64
	NumberSquared *big.Int
}

// ContractAVSTaskMetaData contains all meta data concerning the ContractAVSTask contract.
var ContractAVSTaskMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"taskIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"TaskChallengedSuccessfully\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numberToBeSquared\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"taskResponsePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskChallengePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"thresholdPercentage\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\"}],\"name\":\"createNewTask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"numberToBeSquared\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"taskResponsePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskChallengePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"thresholdPercentage\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\"}],\"internalType\":\"structAVSTask.Task\",\"name\":\"task\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"taskID\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberSquared\",\"type\":\"uint256\"}],\"internalType\":\"structAVSTask.TaskResponse\",\"name\":\"taskResponse\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"operatorAddress\",\"type\":\"string\"}],\"name\":\"raiseAndResolveChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506107e08061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c8063157cf04414610038578063a99d83471461005f575b5f5ffd5b61004b6100463660046103a2565b610072565b604051901515815260200160405180910390f35b61004b61006d366004610424565b61016c565b6040805160c08101825287815260208082018890526001600160401b0387811683850152868116606084015284811660a08401528516608083015291515f92839161090191638bf30a699133918c916100cd918891016104d3565b60408051601f19818403018152828252805160209182012090830152016040516020818303038152906040528b8b8b8b6040518863ffffffff1660e01b815260040161011f9796959493929190610550565b6020604051808303815f875af115801561013b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061015f91906105ba565b9998505050505050505050565b5f8061017b60208501856105d9565b905084355f61018a82806105f2565b90506020860135811460018190036101a8575f9450505050506102e3565b5f6109016001600160a01b031663e4465980338b6040516020016101cc9190610643565b60408051601f198184030181528282528051602091820120908301520160408051601f1981840301815291905261020660208d018d6105d9565b8c604051602001610217919061071d565b60408051601f19818403018152828252805160209182012090830152016040516020818303038152906040528c6040518663ffffffff1660e01b8152600401610264959493929190610746565b6020604051808303815f875af1158015610280573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102a491906105ba565b60405190915033906001600160401b038716907f586973008e6ed96c5552e4527b5f525f2863837da7b1f3f591e32b0a183a9354905f90a39450505050505b9392505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261030d575f5ffd5b81356001600160401b03811115610326576103266102ea565b604051601f8201601f19908116603f011681016001600160401b0381118282101715610354576103546102ea565b60405281815283820160200185101561036b575f5ffd5b816020850160208301375f918101602001919091529392505050565b80356001600160401b038116811461039d575f5ffd5b919050565b5f5f5f5f5f5f60c087890312156103b7575f5ffd5b8635955060208701356001600160401b038111156103d3575f5ffd5b6103df89828a016102fe565b9550506103ee60408801610387565b93506103fc60608801610387565b925061040a60808801610387565b915061041860a08801610387565b90509295509295509295565b5f5f5f8385036080811215610437575f5ffd5b84356001600160401b0381111561044c575f5ffd5b850160c0818803121561045d575f5ffd5b93506040601f1982011215610470575f5ffd5b5060208401915060608401356001600160401b0381111561048f575f5ffd5b61049b868287016102fe565b9150509250925092565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b60208152815160208201525f602083015160c060408401526104f860e08401826104a5565b90506001600160401b0360408501511660608401526001600160401b0360608501511660808401526001600160401b0360808501511660a08401526001600160401b0360a08501511660c08401528091505092915050565b6001600160a01b038816815260e0602082018190525f90610573908301896104a5565b828103604084015261058581896104a5565b6001600160401b0397881660608501529587166080840152505091841660a083015290921660c0909201919091529392505050565b5f602082840312156105ca575f5ffd5b815180151581146102e3575f5ffd5b5f602082840312156105e9575f5ffd5b6102e382610387565b808202811582820484141761061557634e487b7160e01b5f52601160045260245ffd5b92915050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60208082528235828201525f9083013536849003601e19018112610665575f5ffd5b83016020810190356001600160401b03811115610680575f5ffd5b80360382131561068e575f5ffd5b60c060408501526106a360e08501828461061b565b9150506106b260408501610387565b6001600160401b0381166060850152506106ce60608501610387565b6001600160401b0381166080850152506106ea60808501610387565b6001600160401b03811660a08501525061070660a08501610387565b6001600160401b03811660c0850152509392505050565b604081016001600160401b0361073284610387565b168252602092830135929091019190915290565b6001600160a01b038616815260a0602082018190525f90610769908301876104a5565b6001600160401b0386166040840152828103606084015261078a81866104a5565b9050828103608084015261079e81856104a5565b9897505050505050505056fea264697066735822122020350e19118aada8c90599277b640745388a769f7c57a1292e346e4796a0ca1764736f6c634300081c0033",
}

// ContractAVSTaskABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAVSTaskMetaData.ABI instead.
var ContractAVSTaskABI = ContractAVSTaskMetaData.ABI

// ContractAVSTaskBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAVSTaskMetaData.Bin instead.
var ContractAVSTaskBin = ContractAVSTaskMetaData.Bin

// DeployContractAVSTask deploys a new Ethereum contract, binding an instance of ContractAVSTask to it.
func DeployContractAVSTask(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractAVSTask, error) {
	parsed, err := ContractAVSTaskMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAVSTaskBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractAVSTask{ContractAVSTaskCaller: ContractAVSTaskCaller{contract: contract}, ContractAVSTaskTransactor: ContractAVSTaskTransactor{contract: contract}, ContractAVSTaskFilterer: ContractAVSTaskFilterer{contract: contract}}, nil
}

// ContractAVSTask is an auto generated Go binding around an Ethereum contract.
type ContractAVSTask struct {
	ContractAVSTaskCaller     // Read-only binding to the contract
	ContractAVSTaskTransactor // Write-only binding to the contract
	ContractAVSTaskFilterer   // Log filterer for contract events
}

// ContractAVSTaskCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAVSTaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAVSTaskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAVSTaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAVSTaskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAVSTaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAVSTaskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAVSTaskSession struct {
	Contract     *ContractAVSTask  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractAVSTaskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAVSTaskCallerSession struct {
	Contract *ContractAVSTaskCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ContractAVSTaskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAVSTaskTransactorSession struct {
	Contract     *ContractAVSTaskTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractAVSTaskRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAVSTaskRaw struct {
	Contract *ContractAVSTask // Generic contract binding to access the raw methods on
}

// ContractAVSTaskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAVSTaskCallerRaw struct {
	Contract *ContractAVSTaskCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAVSTaskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAVSTaskTransactorRaw struct {
	Contract *ContractAVSTaskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAVSTask creates a new instance of ContractAVSTask, bound to a specific deployed contract.
func NewContractAVSTask(address common.Address, backend bind.ContractBackend) (*ContractAVSTask, error) {
	contract, err := bindContractAVSTask(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAVSTask{ContractAVSTaskCaller: ContractAVSTaskCaller{contract: contract}, ContractAVSTaskTransactor: ContractAVSTaskTransactor{contract: contract}, ContractAVSTaskFilterer: ContractAVSTaskFilterer{contract: contract}}, nil
}

// NewContractAVSTaskCaller creates a new read-only instance of ContractAVSTask, bound to a specific deployed contract.
func NewContractAVSTaskCaller(address common.Address, caller bind.ContractCaller) (*ContractAVSTaskCaller, error) {
	contract, err := bindContractAVSTask(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAVSTaskCaller{contract: contract}, nil
}

// NewContractAVSTaskTransactor creates a new write-only instance of ContractAVSTask, bound to a specific deployed contract.
func NewContractAVSTaskTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAVSTaskTransactor, error) {
	contract, err := bindContractAVSTask(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAVSTaskTransactor{contract: contract}, nil
}

// NewContractAVSTaskFilterer creates a new log filterer instance of ContractAVSTask, bound to a specific deployed contract.
func NewContractAVSTaskFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAVSTaskFilterer, error) {
	contract, err := bindContractAVSTask(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAVSTaskFilterer{contract: contract}, nil
}

// bindContractAVSTask binds a generic wrapper to an already deployed contract.
func bindContractAVSTask(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractAVSTaskMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAVSTask *ContractAVSTaskRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAVSTask.Contract.ContractAVSTaskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAVSTask *ContractAVSTaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAVSTask.Contract.ContractAVSTaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAVSTask *ContractAVSTaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAVSTask.Contract.ContractAVSTaskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAVSTask *ContractAVSTaskCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAVSTask.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAVSTask *ContractAVSTaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAVSTask.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAVSTask *ContractAVSTaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAVSTask.Contract.contract.Transact(opts, method, params...)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x157cf044.
//
// Solidity: function createNewTask(uint256 numberToBeSquared, string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod) returns(bool)
func (_ContractAVSTask *ContractAVSTaskTransactor) CreateNewTask(opts *bind.TransactOpts, numberToBeSquared *big.Int, name string, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint64, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _ContractAVSTask.contract.Transact(opts, "createNewTask", numberToBeSquared, name, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x157cf044.
//
// Solidity: function createNewTask(uint256 numberToBeSquared, string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod) returns(bool)
func (_ContractAVSTask *ContractAVSTaskSession) CreateNewTask(numberToBeSquared *big.Int, name string, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint64, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _ContractAVSTask.Contract.CreateNewTask(&_ContractAVSTask.TransactOpts, numberToBeSquared, name, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x157cf044.
//
// Solidity: function createNewTask(uint256 numberToBeSquared, string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod) returns(bool)
func (_ContractAVSTask *ContractAVSTaskTransactorSession) CreateNewTask(numberToBeSquared *big.Int, name string, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint64, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _ContractAVSTask.Contract.CreateNewTask(&_ContractAVSTask.TransactOpts, numberToBeSquared, name, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xa99d8347.
//
// Solidity: function raiseAndResolveChallenge((uint256,string,uint64,uint64,uint64,uint64) task, (uint64,uint256) taskResponse, string operatorAddress) returns(bool)
func (_ContractAVSTask *ContractAVSTaskTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task AVSTaskTask, taskResponse AVSTaskTaskResponse, operatorAddress string) (*types.Transaction, error) {
	return _ContractAVSTask.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, operatorAddress)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xa99d8347.
//
// Solidity: function raiseAndResolveChallenge((uint256,string,uint64,uint64,uint64,uint64) task, (uint64,uint256) taskResponse, string operatorAddress) returns(bool)
func (_ContractAVSTask *ContractAVSTaskSession) RaiseAndResolveChallenge(task AVSTaskTask, taskResponse AVSTaskTaskResponse, operatorAddress string) (*types.Transaction, error) {
	return _ContractAVSTask.Contract.RaiseAndResolveChallenge(&_ContractAVSTask.TransactOpts, task, taskResponse, operatorAddress)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xa99d8347.
//
// Solidity: function raiseAndResolveChallenge((uint256,string,uint64,uint64,uint64,uint64) task, (uint64,uint256) taskResponse, string operatorAddress) returns(bool)
func (_ContractAVSTask *ContractAVSTaskTransactorSession) RaiseAndResolveChallenge(task AVSTaskTask, taskResponse AVSTaskTaskResponse, operatorAddress string) (*types.Transaction, error) {
	return _ContractAVSTask.Contract.RaiseAndResolveChallenge(&_ContractAVSTask.TransactOpts, task, taskResponse, operatorAddress)
}

// ContractAVSTaskTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractAVSTask contract.
type ContractAVSTaskTaskChallengedSuccessfullyIterator struct {
	Event *ContractAVSTaskTaskChallengedSuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractAVSTaskTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSTaskTaskChallengedSuccessfully)
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
		it.Event = new(ContractAVSTaskTaskChallengedSuccessfully)
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
func (it *ContractAVSTaskTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSTaskTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSTaskTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractAVSTask contract.
type ContractAVSTaskTaskChallengedSuccessfully struct {
	TaskIndex  uint64
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0x586973008e6ed96c5552e4527b5f525f2863837da7b1f3f591e32b0a183a9354.
//
// Solidity: event TaskChallengedSuccessfully(uint64 indexed taskIndex, address indexed challenger)
func (_ContractAVSTask *ContractAVSTaskFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint64, challenger []common.Address) (*ContractAVSTaskTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractAVSTask.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSTaskTaskChallengedSuccessfullyIterator{contract: _ContractAVSTask.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0x586973008e6ed96c5552e4527b5f525f2863837da7b1f3f591e32b0a183a9354.
//
// Solidity: event TaskChallengedSuccessfully(uint64 indexed taskIndex, address indexed challenger)
func (_ContractAVSTask *ContractAVSTaskFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractAVSTaskTaskChallengedSuccessfully, taskIndex []uint64, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractAVSTask.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSTaskTaskChallengedSuccessfully)
				if err := _ContractAVSTask.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
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

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0x586973008e6ed96c5552e4527b5f525f2863837da7b1f3f591e32b0a183a9354.
//
// Solidity: event TaskChallengedSuccessfully(uint64 indexed taskIndex, address indexed challenger)
func (_ContractAVSTask *ContractAVSTaskFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractAVSTaskTaskChallengedSuccessfully, error) {
	event := new(ContractAVSTaskTaskChallengedSuccessfully)
	if err := _ContractAVSTask.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
