// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractv3

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

// Contractv3MetaData contains all meta data concerning the Contractv3 contract.
var Contractv3MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"taskIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"TaskChallengedSuccessfully\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numberToBeSquared\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"taskResponsePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskChallengePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"thresholdPercentage\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\"}],\"name\":\"createNewTask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"numberToBeSquared\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"taskResponsePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskChallengePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"thresholdPercentage\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\"}],\"internalType\":\"structAVSTask.Task\",\"name\":\"task\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"taskID\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberSquared\",\"type\":\"uint256\"}],\"internalType\":\"structAVSTask.TaskResponse\",\"name\":\"taskResponse\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"operatorAddress\",\"type\":\"string\"}],\"name\":\"raiseAndResolveChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b505f80546001600160a01b031916331790556111a38061002d5f395ff3fe608060405234801561000f575f5ffd5b50600436106100a6575f3560e01c80636f48e1a21161006e5780636f48e1a21461013c5780637ea5a4cf1461014f5780638da5cb5b14610162578063c208dd991461018c578063c49bb52114610194578063de16bf46146101a7575f5ffd5b80631d4c8007146100aa57806333a057f5146100d35780634bdb04f9146100f657806354c77f711461010957806355b42cbe1461011c575b5f5ffd5b6100bd6100b8366004610691565b6101af565b6040516100ca9190610732565b60405180910390f35b6100e66100e1366004610802565b610226565b60405190151581526020016100ca565b6100e66101043660046108e0565b6102a2565b6100e6610117366004610919565b61030d565b61012f61012a3660046108e0565b610388565b6040516100ca91906109ce565b6100e661014a366004610b02565b610433565b6100e661015d366004610c8b565b6104ca565b5f54610174906001600160a01b031681565b6040516001600160a01b0390911681526020016100ca565b6100e661059c565b6100e66101a2366004610b02565b610607565b6100e661064d565b604051631d4c800760e01b81526001600160a01b03821660048201526060905f9061090190631d4c8007906024015f604051808303815f875af11580156101f8573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261021f9190810190610d66565b9392505050565b6040516333a057f560e01b81525f908190610901906333a057f590610257908a908a908a908a908a90600401610e28565b6020604051808303815f875af1158015610273573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102979190610e94565b979650505050505050565b6040516318cd2ab360e01b81525f908190610901906318cd2ab3906102cd9033908790600401610eb3565b6020604051808303815f875af11580156102e9573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061021f9190610e94565b604051632a95fddf60e21b81525f9081906109019063aa57f77c9061033e9033908a908a908a908a90600401610ed6565b6020604051808303815f875af115801561035a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061037e9190610e94565b9695505050505050565b604051632ada165f60e11b81526060905f90610901906355b42cbe906103b29086906004016109ce565b5f604051808303815f875af11580156103cd573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526103f49190810190610f0b565b90507ff8a35bbe260cb3c99c0febbe1d21b55e78ca43bdc6ed00e30007a9fadfe67e438160405161042591906109ce565b60405180910390a192915050565b5f5f6109016001600160a01b031663d9e5daa0338f8f8f8f8f8f8f8f8f8f8f6040518d63ffffffff1660e01b81526004016104799c9b9a99989796959493929190610f92565b6020604051808303815f875af1158015610495573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104b99190610e94565b9d9c50505050505050505050505050565b5f5f6108096001600160a01b03167f7ea5a4cf6ce6f6b705c2660435a26c37251a0c518b51dd1ef25344f05ece43a7888888888860405160240161051295949392919061109c565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516105509190611157565b5f604051808303815f865af19150503d805f8114610589576040519150601f19603f3d011682016040523d82523d5f602084013e61058e565b606091505b509098975050505050505050565b60405163d7a2398b60e01b81523360048201525f9081906109019063d7a2398b906024015b6020604051808303815f875af11580156105dd573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106019190610e94565b92915050565b5f5f6109016001600160a01b031663cde09950338f8f8f8f8f8f8f8f8f8f8f6040518d63ffffffff1660e01b81526004016104799c9b9a99989796959493929190610f92565b6040516351b27a6d60e11b81523360048201525f9081906109019063a364f4da906024016105c1565b80356001600160a01b038116811461068c575f5ffd5b919050565b5f602082840312156106a1575f5ffd5b61021f82610676565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f82825180855260208501945060208160051b830101602085015f5b8381101561072657601f198584030188526107108383516106aa565b60209889019890935091909101906001016106f4565b50909695505050505050565b602081525f61021f60208301846106d8565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561078057610780610744565b604052919050565b5f6001600160401b038211156107a0576107a0610744565b50601f01601f191660200190565b5f82601f8301126107bd575f5ffd5b8135602083015f6107d56107d084610788565b610758565b90508281528583830111156107e8575f5ffd5b828260208301375f92810160200192909252509392505050565b5f5f5f5f5f60a08688031215610816575f5ffd5b85356001600160401b0381111561082b575f5ffd5b610837888289016107ae565b95505060208601356001600160401b03811115610852575f5ffd5b61085e888289016107ae565b94505060408601356001600160401b03811115610879575f5ffd5b610885888289016107ae565b93505060608601356001600160401b038111156108a0575f5ffd5b6108ac888289016107ae565b92505060808601356001600160401b038111156108c7575f5ffd5b6108d3888289016107ae565b9150509295509295909350565b5f602082840312156108f0575f5ffd5b81356001600160401b03811115610905575f5ffd5b610911848285016107ae565b949350505050565b5f5f5f5f6080858703121561092c575f5ffd5b84356001600160401b03811115610941575f5ffd5b61094d878288016107ae565b94505060208501356001600160401b03811115610968575f5ffd5b610974878288016107ae565b93505060408501356001600160401b0381111561098f575f5ffd5b61099b878288016107ae565b92505060608501356001600160401b038111156109b6575f5ffd5b6109c2878288016107ae565b91505092959194509250565b602081525f61021f60208301846106aa565b80356001600160401b038116811461068c575f5ffd5b5f6001600160401b03821115610a0e57610a0e610744565b5060051b60200190565b5f82601f830112610a27575f5ffd5b8135610a356107d0826109f6565b8082825260208201915060208360051b860101925085831115610a56575f5ffd5b602085015b83811015610a965780356001600160401b03811115610a78575f5ffd5b610a87886020838a01016107ae565b84525060209283019201610a5b565b5095945050505050565b5f82601f830112610aaf575f5ffd5b8135610abd6107d0826109f6565b8082825260208201915060208360051b860101925085831115610ade575f5ffd5b602085015b83811015610a9657610af4816109e0565b835260209283019201610ae3565b5f5f5f5f5f5f5f5f5f5f5f6101608c8e031215610b1d575f5ffd5b8b356001600160401b03811115610b32575f5ffd5b610b3e8e828f016107ae565b9b5050610b4d60208d016109e0565b9950610b5b60408d01610676565b9850610b6960608d01610676565b9750610b7760808d01610676565b965060a08c01356001600160401b03811115610b91575f5ffd5b610b9d8e828f01610a18565b96505060c08c01356001600160401b03811115610bb8575f5ffd5b610bc48e828f01610a18565b955050610bd360e08d016109e0565b9350610be26101008d016109e0565b92506101208c01356001600160401b03811115610bfd575f5ffd5b610c098e828f016107ae565b9250506101408c01356001600160401b03811115610c25575f5ffd5b610c318e828f01610aa0565b9150509295989b509295989b9093969950565b5f5f83601f840112610c54575f5ffd5b5081356001600160401b03811115610c6a575f5ffd5b6020830191508360208260051b8501011115610c84575f5ffd5b9250929050565b5f5f5f5f5f60608688031215610c9f575f5ffd5b8535945060208601356001600160401b03811115610cbb575f5ffd5b8601601f81018813610ccb575f5ffd5b80356001600160401b03811115610ce0575f5ffd5b886020828401011115610cf1575f5ffd5b6020919091019450925060408601356001600160401b03811115610d13575f5ffd5b610d1f88828901610c44565b969995985093965092949392505050565b5f610d3d6107d084610788565b9050828152838383011115610d50575f5ffd5b8282602083015e5f602084830101529392505050565b5f60208284031215610d76575f5ffd5b81516001600160401b03811115610d8b575f5ffd5b8201601f81018413610d9b575f5ffd5b8051610da96107d0826109f6565b8082825260208201915060208360051b850101925086831115610dca575f5ffd5b602084015b83811015610e1d5780516001600160401b03811115610dec575f5ffd5b8501603f81018913610dfc575f5ffd5b610e0e89602083015160408401610d30565b84525060209283019201610dcf565b509695505050505050565b60a081525f610e3a60a08301886106aa565b8281036020840152610e4c81886106aa565b90508281036040840152610e6081876106aa565b90508281036060840152610e7481866106aa565b90508281036080840152610e8881856106aa565b98975050505050505050565b5f60208284031215610ea4575f5ffd5b8151801515811461021f575f5ffd5b6001600160a01b03831681526040602082018190525f90610911908301846106aa565b6001600160a01b038616815260a0602082018190525f90610ef9908301876106aa565b8281036040840152610e6081876106aa565b5f60208284031215610f1b575f5ffd5b81516001600160401b03811115610f30575f5ffd5b8201601f81018413610f40575f5ffd5b61091184825160208401610d30565b5f8151808452602084019350602083015f5b82811015610f885781516001600160401b0316865260209586019590910190600101610f61565b5093949350505050565b6001600160a01b038d16815261018060208201525f610fb561018083018e6106aa565b6001600160401b038d1660408401526001600160a01b038c1660608401526001600160a01b038b1660808401526001600160a01b038a1660a084015282810360c0840152611003818a6106d8565b905082810360e084015261101781896106d8565b6001600160401b03881661010085015290506001600160401b03861661012084015282810361014084015261104c81866106aa565b90508281036101608401526110618185610f4f565b9f9e505050505050505050505050505050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b858152606060208201525f6110b5606083018688611074565b8281036040840152838152602080820190600586901b830101865f36829003601e19015b8882101561114657858403601f1901855282358181126110f7575f5ffd5b8a016020810190356001600160401b03811115611112575f5ffd5b803603821315611120575f5ffd5b61112b868284611074565b955050506020830192506020850194506001820191506110d9565b50919b9a5050505050505050505050565b5f82518060208501845e5f92019182525091905056fea2646970667358221220a1a3469e6baa4e985228584ec08d11e1f9432cc1b6dfabc2d34e9fb4dcea463664736f6c634300081c0033",
}

// Contractv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use Contractv3MetaData.ABI instead.
var Contractv3ABI = Contractv3MetaData.ABI

// Contractv3Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Contractv3MetaData.Bin instead.
var Contractv3Bin = Contractv3MetaData.Bin

// DeployContractv3 deploys a new Ethereum contract, binding an instance of Contractv3 to it.
func DeployContractv3(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contractv3, error) {
	parsed, err := Contractv3MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Contractv3Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contractv3{Contractv3Caller: Contractv3Caller{contract: contract}, Contractv3Transactor: Contractv3Transactor{contract: contract}, Contractv3Filterer: Contractv3Filterer{contract: contract}}, nil
}

// Contractv3 is an auto generated Go binding around an Ethereum contract.
type Contractv3 struct {
	Contractv3Caller     // Read-only binding to the contract
	Contractv3Transactor // Write-only binding to the contract
	Contractv3Filterer   // Log filterer for contract events
}

// Contractv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Contractv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Contractv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Contractv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Contractv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Contractv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Contractv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Contractv3Session struct {
	Contract     *Contractv3       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Contractv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Contractv3CallerSession struct {
	Contract *Contractv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Contractv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Contractv3TransactorSession struct {
	Contract     *Contractv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Contractv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Contractv3Raw struct {
	Contract *Contractv3 // Generic contract binding to access the raw methods on
}

// Contractv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Contractv3CallerRaw struct {
	Contract *Contractv3Caller // Generic read-only contract binding to access the raw methods on
}

// Contractv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Contractv3TransactorRaw struct {
	Contract *Contractv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewContractv3 creates a new instance of Contractv3, bound to a specific deployed contract.
func NewContractv3(address common.Address, backend bind.ContractBackend) (*Contractv3, error) {
	contract, err := bindContractv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contractv3{Contractv3Caller: Contractv3Caller{contract: contract}, Contractv3Transactor: Contractv3Transactor{contract: contract}, Contractv3Filterer: Contractv3Filterer{contract: contract}}, nil
}

// NewContractv3Caller creates a new read-only instance of Contractv3, bound to a specific deployed contract.
func NewContractv3Caller(address common.Address, caller bind.ContractCaller) (*Contractv3Caller, error) {
	contract, err := bindContractv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Contractv3Caller{contract: contract}, nil
}

// NewContractv3Transactor creates a new write-only instance of Contractv3, bound to a specific deployed contract.
func NewContractv3Transactor(address common.Address, transactor bind.ContractTransactor) (*Contractv3Transactor, error) {
	contract, err := bindContractv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Contractv3Transactor{contract: contract}, nil
}

// NewContractv3Filterer creates a new log filterer instance of Contractv3, bound to a specific deployed contract.
func NewContractv3Filterer(address common.Address, filterer bind.ContractFilterer) (*Contractv3Filterer, error) {
	contract, err := bindContractv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Contractv3Filterer{contract: contract}, nil
}

// bindContractv3 binds a generic wrapper to an already deployed contract.
func bindContractv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Contractv3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractv3 *Contractv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractv3.Contract.Contractv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractv3 *Contractv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractv3.Contract.Contractv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractv3 *Contractv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractv3.Contract.Contractv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractv3 *Contractv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractv3 *Contractv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractv3 *Contractv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractv3.Contract.contract.Transact(opts, method, params...)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x157cf044.
//
// Solidity: function createNewTask(uint256 numberToBeSquared, string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod) returns(bool)
func (_Contractv3 *Contractv3Transactor) CreateNewTask(opts *bind.TransactOpts, numberToBeSquared *big.Int, name string, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint64, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _Contractv3.contract.Transact(opts, "createNewTask", numberToBeSquared, name, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x157cf044.
//
// Solidity: function createNewTask(uint256 numberToBeSquared, string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod) returns(bool)
func (_Contractv3 *Contractv3Session) CreateNewTask(numberToBeSquared *big.Int, name string, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint64, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _Contractv3.Contract.CreateNewTask(&_Contractv3.TransactOpts, numberToBeSquared, name, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x157cf044.
//
// Solidity: function createNewTask(uint256 numberToBeSquared, string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod) returns(bool)
func (_Contractv3 *Contractv3TransactorSession) CreateNewTask(numberToBeSquared *big.Int, name string, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint64, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _Contractv3.Contract.CreateNewTask(&_Contractv3.TransactOpts, numberToBeSquared, name, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xa99d8347.
//
// Solidity: function raiseAndResolveChallenge((uint256,string,uint64,uint64,uint64,uint64) task, (uint64,uint256) taskResponse, string operatorAddress) returns(bool)
func (_Contractv3 *Contractv3Transactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task AVSTaskTask, taskResponse AVSTaskTaskResponse, operatorAddress string) (*types.Transaction, error) {
	return _Contractv3.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, operatorAddress)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xa99d8347.
//
// Solidity: function raiseAndResolveChallenge((uint256,string,uint64,uint64,uint64,uint64) task, (uint64,uint256) taskResponse, string operatorAddress) returns(bool)
func (_Contractv3 *Contractv3Session) RaiseAndResolveChallenge(task AVSTaskTask, taskResponse AVSTaskTaskResponse, operatorAddress string) (*types.Transaction, error) {
	return _Contractv3.Contract.RaiseAndResolveChallenge(&_Contractv3.TransactOpts, task, taskResponse, operatorAddress)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xa99d8347.
//
// Solidity: function raiseAndResolveChallenge((uint256,string,uint64,uint64,uint64,uint64) task, (uint64,uint256) taskResponse, string operatorAddress) returns(bool)
func (_Contractv3 *Contractv3TransactorSession) RaiseAndResolveChallenge(task AVSTaskTask, taskResponse AVSTaskTaskResponse, operatorAddress string) (*types.Transaction, error) {
	return _Contractv3.Contract.RaiseAndResolveChallenge(&_Contractv3.TransactOpts, task, taskResponse, operatorAddress)
}

// Contractv3TaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the Contractv3 contract.
type Contractv3TaskChallengedSuccessfullyIterator struct {
	Event *Contractv3TaskChallengedSuccessfully // Event containing the contract specifics and raw log

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
func (it *Contractv3TaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Contractv3TaskChallengedSuccessfully)
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
		it.Event = new(Contractv3TaskChallengedSuccessfully)
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
func (it *Contractv3TaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Contractv3TaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Contractv3TaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the Contractv3 contract.
type Contractv3TaskChallengedSuccessfully struct {
	TaskIndex  uint64
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0x586973008e6ed96c5552e4527b5f525f2863837da7b1f3f591e32b0a183a9354.
//
// Solidity: event TaskChallengedSuccessfully(uint64 indexed taskIndex, address indexed challenger)
func (_Contractv3 *Contractv3Filterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint64, challenger []common.Address) (*Contractv3TaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Contractv3.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &Contractv3TaskChallengedSuccessfullyIterator{contract: _Contractv3.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0x586973008e6ed96c5552e4527b5f525f2863837da7b1f3f591e32b0a183a9354.
//
// Solidity: event TaskChallengedSuccessfully(uint64 indexed taskIndex, address indexed challenger)
func (_Contractv3 *Contractv3Filterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *Contractv3TaskChallengedSuccessfully, taskIndex []uint64, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Contractv3.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Contractv3TaskChallengedSuccessfully)
				if err := _Contractv3.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
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
func (_Contractv3 *Contractv3Filterer) ParseTaskChallengedSuccessfully(log types.Log) (*Contractv3TaskChallengedSuccessfully, error) {
	event := new(Contractv3TaskChallengedSuccessfully)
	if err := _Contractv3.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
