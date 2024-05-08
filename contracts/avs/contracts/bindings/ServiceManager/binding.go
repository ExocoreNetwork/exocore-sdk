// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractServiceManager

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

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractServiceManagerMetaData contains all meta data concerning the ContractServiceManager contract.
var ContractServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_taskManager\",\"type\":\"address\",\"internalType\":\"contractTaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"taskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620016f6380380620016f683398101604081905262000035916200013e565b6001600160a01b0380851660a052808416608052821660c0528383836200005b62000074565b5050506001600160a01b031660e05250620001a6915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000c55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620001255780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6001600160a01b03811681146200012557600080fd5b600080600080608085870312156200015557600080fd5b8451620001628162000128565b6020860151909450620001758162000128565b6040860151909350620001888162000128565b60608601519092506200019b8162000128565b939692955090935050565b60805160a05160c05160e0516114b4620002426000396000818161017601526106810152600081816103c001528181610512015281816105a901528181610aed01528181610c670152610d060152600081816107470152818161081001526108e40152600081816101eb0152818161027a015281816102fa015281816107bc0152818161088801528181610a280152610bc201526114b46000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80639926ee7d116100715780639926ee7d1461014b578063a364f4da1461015e578063a50a640e14610171578063c4d66de814610198578063e481af9d146101ab578063f2fde38b146101b357600080fd5b806333cfb7b7146100ae57806338c8ee64146100d7578063715018a6146100ec578063750521f5146100f45780638da5cb5b14610107575b600080fd5b6100c16100bc366004610fdf565b6101c6565b6040516100ce9190611003565b60405180910390f35b6100ea6100e5366004610fdf565b610676565b005b6100ea610714565b6100ea610102366004611105565b610728565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03165b6040516001600160a01b0390911681526020016100ce565b6100ea610159366004611156565b6107b1565b6100ea61016c366004610fdf565b61087d565b6101337f000000000000000000000000000000000000000000000000000000000000000081565b6100ea6101a6366004610fdf565b610913565b6100c1610a22565b6100ea6101c1366004610fdf565b610dcf565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa158015610232573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102569190611201565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa1580156102c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102e5919061121a565b90506001600160c01b038116158061037f57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610356573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061037a9190611243565b60ff16155b1561039b57505060408051600081526020810190915292915050565b60006103af826001600160c01b0316610e0a565b90506000805b825181101561047b577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f58483815181106103ff576103ff611266565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa158015610443573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104679190611201565b6104719083611292565b91506001016103b5565b5060008167ffffffffffffffff81111561049757610497611050565b6040519080825280602002602001820160405280156104c0578160200160208202803683370190505b5090506000805b84518110156106695760008582815181106104e4576104e4611266565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610559573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057d9190611201565b905060005b8181101561065e576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa1580156105f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061b91906112a5565b6000015186868151811061063157610631611266565b6001600160a01b03909216602092830291909101909101528461065381611315565b955050600101610582565b5050506001016104c7565b5090979650505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107115760405162461bcd60e51b815260206004820152603560248201527f6f6e6c7945786f636f72655461736b4d616e616765723a206e6f742066726f6d6044820152741032bc37b1b7b932903a30b9b59036b0b730b3b2b960591b60648201526084015b60405180910390fd5b50565b61071c610ecd565b6107266000610f28565b565b610730610ecd565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb3559061077c908490600401611374565b600060405180830381600087803b15801561079657600080fd5b505af11580156107aa573d6000803e3d6000fd5b5050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107f95760405162461bcd60e51b815260040161070890611387565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061084790859085906004016113ff565b600060405180830381600087803b15801561086157600080fd5b505af1158015610875573d6000803e3d6000fd5b505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108c55760405162461bcd60e51b815260040161070890611387565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da9060240161077c565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156109595750825b905060008267ffffffffffffffff1660011480156109765750303b155b905081158015610984575080155b156109a25760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156109cc57845460ff60401b1916600160401b1785555b6109d586610f28565b831561087557845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a84573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa89190611243565b60ff16905080600003610ac957505060408051600081526020810190915290565b6000805b82811015610b7457604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610b3c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b609190611201565b610b6a9083611292565b9150600101610acd565b5060008167ffffffffffffffff811115610b9057610b90611050565b604051908082528060200260200182016040528015610bb9578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c1e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c429190611243565b60ff16811015610dc557604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610cb6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cda9190611201565b905060005b81811015610dbb576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610d54573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d7891906112a5565b60000151858581518110610d8e57610d8e611266565b6001600160a01b039092166020928302919091019091015283610db081611315565b945050600101610cdf565b5050600101610bc0565b5090949350505050565b610dd7610ecd565b6001600160a01b038116610e0157604051631e4fbdf760e01b815260006004820152602401610708565b61071181610f28565b6060600080610e1884610f99565b61ffff1667ffffffffffffffff811115610e3457610e34611050565b6040519080825280601f01601f191660200182016040528015610e5e576020820181803683370190505b5090506000805b825182108015610e76575061010081105b15610dc5576001811b935085841615610ebd578060f81b838381518110610e9f57610e9f611266565b60200101906001600160f81b031916908160001a9053508160010191505b610ec681611315565b9050610e65565b33610eff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146107265760405163118cdaa760e01b8152336004820152602401610708565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b6000805b8215610fc457610fae60018461144a565b9092169180610fbc8161145d565b915050610f9d565b92915050565b6001600160a01b038116811461071157600080fd5b600060208284031215610ff157600080fd5b8135610ffc81610fca565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156110445783516001600160a01b03168352928401929184019160010161101f565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561108957611089611050565b60405290565b600067ffffffffffffffff808411156110aa576110aa611050565b604051601f8501601f19908116603f011681019082821181831017156110d2576110d2611050565b816040528093508581528686860111156110eb57600080fd5b858560208301376000602087830101525050509392505050565b60006020828403121561111757600080fd5b813567ffffffffffffffff81111561112e57600080fd5b8201601f8101841361113f57600080fd5b61114e8482356020840161108f565b949350505050565b6000806040838503121561116957600080fd5b823561117481610fca565b9150602083013567ffffffffffffffff8082111561119157600080fd5b90840190606082870312156111a557600080fd5b6111ad611066565b8235828111156111bc57600080fd5b83019150601f820187136111cf57600080fd5b6111de8783356020850161108f565b815260208301356020820152604083013560408201528093505050509250929050565b60006020828403121561121357600080fd5b5051919050565b60006020828403121561122c57600080fd5b81516001600160c01b0381168114610ffc57600080fd5b60006020828403121561125557600080fd5b815160ff81168114610ffc57600080fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b80820180821115610fc457610fc461127c565b6000604082840312156112b757600080fd5b6040516040810181811067ffffffffffffffff821117156112da576112da611050565b60405282516112e881610fca565b815260208301516bffffffffffffffffffffffff8116811461130957600080fd5b60208201529392505050565b6000600182016113275761132761127c565b5060010190565b6000815180845260005b8181101561135457602081850181015186830182015201611338565b506000602082860101526020601f19601f83011685010191505092915050565b602081526000610ffc602083018461132e565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b038316815260406020820152600082516060604084015261142960a084018261132e565b90506020840151606084015260408401516080840152809150509392505050565b81810381811115610fc457610fc461127c565b600061ffff8083168181036114745761147461127c565b600101939250505056fea2646970667358221220b59c3c603517ce670ce35a8b8f525a296b0f585c4216536d388ef322c49a39c964736f6c63430008170033",
}

// ContractServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractServiceManagerMetaData.ABI instead.
var ContractServiceManagerABI = ContractServiceManagerMetaData.ABI

// ContractServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractServiceManagerMetaData.Bin instead.
var ContractServiceManagerBin = ContractServiceManagerMetaData.Bin

// DeployContractServiceManager deploys a new Ethereum contract, binding an instance of ContractServiceManager to it.
func DeployContractServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _delegationManager common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _taskManager common.Address) (common.Address, *types.Transaction, *ContractServiceManager, error) {
	parsed, err := ContractServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractServiceManagerBin), backend, _delegationManager, _registryCoordinator, _stakeRegistry, _taskManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractServiceManager{ContractServiceManagerCaller: ContractServiceManagerCaller{contract: contract}, ContractServiceManagerTransactor: ContractServiceManagerTransactor{contract: contract}, ContractServiceManagerFilterer: ContractServiceManagerFilterer{contract: contract}}, nil
}

// ContractServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractServiceManager struct {
	ContractServiceManagerCaller     // Read-only binding to the contract
	ContractServiceManagerTransactor // Write-only binding to the contract
	ContractServiceManagerFilterer   // Log filterer for contract events
}

// ContractServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractServiceManagerSession struct {
	Contract     *ContractServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ContractServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractServiceManagerCallerSession struct {
	Contract *ContractServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ContractServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractServiceManagerTransactorSession struct {
	Contract     *ContractServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ContractServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractServiceManagerRaw struct {
	Contract *ContractServiceManager // Generic contract binding to access the raw methods on
}

// ContractServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractServiceManagerCallerRaw struct {
	Contract *ContractServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractServiceManagerTransactorRaw struct {
	Contract *ContractServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractServiceManager creates a new instance of ContractServiceManager, bound to a specific deployed contract.
func NewContractServiceManager(address common.Address, backend bind.ContractBackend) (*ContractServiceManager, error) {
	contract, err := bindContractServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractServiceManager{ContractServiceManagerCaller: ContractServiceManagerCaller{contract: contract}, ContractServiceManagerTransactor: ContractServiceManagerTransactor{contract: contract}, ContractServiceManagerFilterer: ContractServiceManagerFilterer{contract: contract}}, nil
}

// NewContractServiceManagerCaller creates a new read-only instance of ContractServiceManager, bound to a specific deployed contract.
func NewContractServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractServiceManagerCaller, error) {
	contract, err := bindContractServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractServiceManagerCaller{contract: contract}, nil
}

// NewContractServiceManagerTransactor creates a new write-only instance of ContractServiceManager, bound to a specific deployed contract.
func NewContractServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractServiceManagerTransactor, error) {
	contract, err := bindContractServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractServiceManagerTransactor{contract: contract}, nil
}

// NewContractServiceManagerFilterer creates a new log filterer instance of ContractServiceManager, bound to a specific deployed contract.
func NewContractServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractServiceManagerFilterer, error) {
	contract, err := bindContractServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractServiceManagerFilterer{contract: contract}, nil
}

// bindContractServiceManager binds a generic wrapper to an already deployed contract.
func bindContractServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractServiceManager *ContractServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractServiceManager.Contract.ContractServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractServiceManager *ContractServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractServiceManager.Contract.ContractServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractServiceManager *ContractServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractServiceManager.Contract.ContractServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractServiceManager *ContractServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractServiceManager *ContractServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractServiceManager *ContractServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractServiceManager.Contract.contract.Transact(opts, method, params...)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractServiceManager *ContractServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractServiceManager *ContractServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractServiceManager *ContractServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractServiceManager *ContractServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractServiceManager *ContractServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractServiceManager.Contract.GetRestakeableStrategies(&_ContractServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractServiceManager *ContractServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractServiceManager.Contract.GetRestakeableStrategies(&_ContractServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractServiceManager *ContractServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractServiceManager *ContractServiceManagerSession) Owner() (common.Address, error) {
	return _ContractServiceManager.Contract.Owner(&_ContractServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractServiceManager *ContractServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractServiceManager.Contract.Owner(&_ContractServiceManager.CallOpts)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_ContractServiceManager *ContractServiceManagerCaller) TaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractServiceManager.contract.Call(opts, &out, "taskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_ContractServiceManager *ContractServiceManagerSession) TaskManager() (common.Address, error) {
	return _ContractServiceManager.Contract.TaskManager(&_ContractServiceManager.CallOpts)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_ContractServiceManager *ContractServiceManagerCallerSession) TaskManager() (common.Address, error) {
	return _ContractServiceManager.Contract.TaskManager(&_ContractServiceManager.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractServiceManager *ContractServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractServiceManager *ContractServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractServiceManager *ContractServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractServiceManager.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractServiceManager *ContractServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractServiceManager.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractServiceManager *ContractServiceManagerSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractServiceManager.Contract.FreezeOperator(&_ContractServiceManager.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractServiceManager *ContractServiceManagerTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractServiceManager.Contract.FreezeOperator(&_ContractServiceManager.TransactOpts, operatorAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractServiceManager *ContractServiceManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _ContractServiceManager.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractServiceManager *ContractServiceManagerSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ContractServiceManager.Contract.Initialize(&_ContractServiceManager.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractServiceManager *ContractServiceManagerTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ContractServiceManager.Contract.Initialize(&_ContractServiceManager.TransactOpts, initialOwner)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractServiceManager *ContractServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractServiceManager *ContractServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractServiceManager.Contract.RegisterOperatorToAVS(&_ContractServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractServiceManager *ContractServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractServiceManager.Contract.RegisterOperatorToAVS(&_ContractServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractServiceManager *ContractServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractServiceManager *ContractServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractServiceManager.Contract.RenounceOwnership(&_ContractServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractServiceManager *ContractServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractServiceManager.Contract.RenounceOwnership(&_ContractServiceManager.TransactOpts)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ContractServiceManager *ContractServiceManagerTransactor) SetMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractServiceManager.contract.Transact(opts, "setMetadataURI", _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ContractServiceManager *ContractServiceManagerSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractServiceManager.Contract.SetMetadataURI(&_ContractServiceManager.TransactOpts, _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ContractServiceManager *ContractServiceManagerTransactorSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractServiceManager.Contract.SetMetadataURI(&_ContractServiceManager.TransactOpts, _metadataURI)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractServiceManager *ContractServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractServiceManager *ContractServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractServiceManager.Contract.TransferOwnership(&_ContractServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractServiceManager *ContractServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractServiceManager.Contract.TransferOwnership(&_ContractServiceManager.TransactOpts, newOwner)
}

// ContractServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractServiceManager contract.
type ContractServiceManagerInitializedIterator struct {
	Event *ContractServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractServiceManagerInitialized)
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
		it.Event = new(ContractServiceManagerInitialized)
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
func (it *ContractServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractServiceManagerInitialized represents a Initialized event raised by the ContractServiceManager contract.
type ContractServiceManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractServiceManager *ContractServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractServiceManagerInitializedIterator{contract: _ContractServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractServiceManager *ContractServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractServiceManagerInitialized)
				if err := _ContractServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractServiceManager *ContractServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractServiceManagerInitialized, error) {
	event := new(ContractServiceManagerInitialized)
	if err := _ContractServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractServiceManager contract.
type ContractServiceManagerOwnershipTransferredIterator struct {
	Event *ContractServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractServiceManagerOwnershipTransferred)
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
func (it *ContractServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractServiceManager contract.
type ContractServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractServiceManager *ContractServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractServiceManagerOwnershipTransferredIterator{contract: _ContractServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractServiceManager *ContractServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractServiceManagerOwnershipTransferred)
				if err := _ContractServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractServiceManager *ContractServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractServiceManagerOwnershipTransferred, error) {
	event := new(ContractServiceManagerOwnershipTransferred)
	if err := _ContractServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
