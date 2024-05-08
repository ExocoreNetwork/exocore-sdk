// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractavsservice

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

// AvsServiceContractTask is an auto generated low-level Go binding around an user-defined struct.
type AvsServiceContractTask struct {
	TaskId             *big.Int
	Issuer             common.Address
	Description        string
	NumberX            *big.Int
	NumberY            *big.Int
	Reward             *big.Int
	Deadline           *big.Int
	IsCompleted        bool
	IsRewarded         bool
	OperatorSignatures []byte
	Operators          []common.Address
}

// ContractavsserviceMetaData contains all meta data concerning the Contractavsservice contract.
var ContractavsserviceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashPercentage\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"OperatorSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"TaskCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"rewards\",\"type\":\"uint256[]\"}],\"name\":\"TaskRewarded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allTaskResponses\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"aggSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"completeTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numberY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"createTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"getTask\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numberY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCompleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isRewarded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"operatorSignatures\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"}],\"internalType\":\"structAvsServiceContract.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTaskCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardsEarned\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"avsAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"operatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"action\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"avsOwnerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"assetID\",\"type\":\"string\"}],\"name\":\"registerAVSToExocore\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"registerOperatorToAVS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"operatorAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rewards\",\"type\":\"uint256[]\"}],\"name\":\"rewardTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"name\":\"slashOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskIdCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numberY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCompleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isRewarded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"operatorSignatures\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50604051612148380380612148833981016040819052602b916047565b5f80546001600160a01b0319163317905560018055600255605d565b5f602082840312156056575f80fd5b5051919050565b6120de8061006a5f395ff3fe608060405234801561000f575f80fd5b50600436106100f0575f3560e01c806386cfceea11610093578063ad8ef91c11610063578063ad8ef91c14610236578063aebe78211461023f578063bd58199a14610252578063c17a340e14610140575f80fd5b806386cfceea146101ad5780638d977672146101c05780638da5cb5b146101e9578063a24d3c1114610213575f80fd5b80631d65e77e116100ce5780631d65e77e146101525780632f95670e146101725780633a3e890714610187578063648b14c51461019a575f80fd5b806313e7c9d8146100f457806319c2e665146101205780631a4a7c6114610140575b5f80fd5b6101076101023660046114ef565b61025b565b604051610117949392919061153d565b60405180910390f35b61013361012e36600461157a565b610392565b6040516101179190611591565b6004545b604051908152602001610117565b61016561016036600461157a565b610429565b60405161011791906115e6565b6101856101803660046116b9565b61066f565b005b610185610195366004611781565b6109a2565b6101856101a836600461187d565b610be5565b6101856101bb3660046119b2565b610e2d565b6101d36101ce36600461157a565b610f44565b6040516101179a99989796959493929190611a27565b5f546101fb906001600160a01b031681565b6040516001600160a01b039091168152602001610117565b610226610221366004611ab6565b6110c7565b6040519015158152602001610117565b61014460015481565b61018561024d366004611b8d565b61119c565b61014460025481565b60036020525f908152604090208054819061027590611be4565b80601f01602080910402602001604051908101604052809291908181526020018280546102a190611be4565b80156102ec5780601f106102c3576101008083540402835291602001916102ec565b820191905f5260205f20905b8154815290600101906020018083116102cf57829003601f168201915b50505050509080600101805461030190611be4565b80601f016020809104026020016040519081016040528092919081815260200182805461032d90611be4565b80156103785780601f1061034f57610100808354040283529160200191610378565b820191905f5260205f20905b81548152906001019060200180831161035b57829003601f168201915b50505050600283015460039093015491929160ff16905084565b60056020525f9081526040902080546103aa90611be4565b80601f01602080910402602001604051908101604052809291908181526020018280546103d690611be4565b80156104215780601f106103f857610100808354040283529160200191610421565b820191905f5260205f20905b81548152906001019060200180831161040457829003601f168201915b505050505081565b6104316113fe565b600461043e600184611c30565b8154811061044e5761044e611c49565b905f5260205f2090600a0201604051806101600160405290815f8201548152602001600182015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020016002820180546104ad90611be4565b80601f01602080910402602001604051908101604052809291908181526020018280546104d990611be4565b80156105245780601f106104fb57610100808354040283529160200191610524565b820191905f5260205f20905b81548152906001019060200180831161050757829003601f168201915b505050918352505060038201546020820152600482015460408201526005820154606082015260068201546080820152600782015460ff808216151560a084015261010090910416151560c082015260088201805460e09092019161058890611be4565b80601f01602080910402602001604051908101604052809291908181526020018280546105b490611be4565b80156105ff5780601f106105d6576101008083540402835291602001916105ff565b820191905f5260205f20905b8154815290600101906020018083116105e257829003601f168201915b505050505081526020016009820180548060200260200160405190810160405280929190818152602001828054801561065f57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610641575b5050505050815250509050919050565b5f546001600160a01b031633146106a15760405162461bcd60e51b815260040161069890611c5d565b60405180910390fd5b5f851180156106b257506004548511155b6106f05760405162461bcd60e51b815260206004820152600f60248201526e3a30b9b5a4b21034b99032b93937b960891b6044820152606401610698565b5f60046106fe600188611c30565b8154811061070e5761070e611c49565b905f5260205f2090600a0201905085815f0154146107655760405162461bcd60e51b81526020600482015260146024820152732a30b9b5903237b2b9903737ba1032bc34b9ba1760611b6044820152606401610698565b600781015460ff16156107ba5760405162461bcd60e51b815260206004820152601a60248201527f5461736b20697320616c726561647920636f6d706c657465642e0000000000006044820152606401610698565b428160060154101561080e5760405162461bcd60e51b815260206004820152601960248201527f5461736b20646561646c696e6520686173207061737365642e000000000000006044820152606401610698565b5f6108096001600160a01b03167f45104d62855afd6bee93c27ad3ce757c879542d68d0d2f5d7cf869489b4b7d45836004015484600301546108509190611ca8565b60405161086891908a908a908a908a90602401611ce3565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516108a69190611d94565b5f604051808303815f865af19150503d805f81146108df576040519150601f19603f3d011682016040523d82523d5f602084013e6108e4565b606091505b50509050806109355760405162461bcd60e51b815260206004820152601960248201527f66617374416767726567617465566572696679206661696c2e000000000000006044820152606401610698565b5f87815260056020526040902061094d868883611df6565b5060078201805460ff191660019081179091558201546040516001600160a01b039091169088907fbb5889c77948badf90e8a5c73d55265e5f5d6e4837a79a78c5669691b897faed905f90a350505050505050565b5f546001600160a01b031633146109cb5760405162461bcd60e51b815260040161069890611c5d565b5f60046109d9600185611c30565b815481106109e9576109e9611c49565b5f9182526020909120600a90910201600781015490915060ff16610a485760405162461bcd60e51b81526020600482015260166024820152752a30b9b59034b9903737ba1031b7b6b83632ba32b21760511b6044820152606401610698565b6007810154610100900460ff1615610a9e5760405162461bcd60e51b81526020600482015260196024820152782a30b9b59034b99030b63932b0b23c903932bbb0b93232b21760391b6044820152606401610698565b6001600160a01b0382165f908152600360208190526040909120015460ff16610b095760405162461bcd60e51b815260206004820152601960248201527f496e76616c6964206f70657261746f7220616464726573732e000000000000006044820152606401610698565b5f610b13836113b8565b90505f8111610b705760405162461bcd60e51b8152602060048201526024808201527f43616e6e6f7420736c617368206f70657261746f7220666f722074686973207460448201526330b9b59760e11b6064820152608401610698565b6001600160a01b0383165f9081526003602052604081206002018054839290610b9a908490611c30565b90915550506040518181526001600160a01b0384169085907fa2e524bd0f71903485fbb3d6d50cb305f61005ceea2047c3ac92aa7e0d1043069060200160405180910390a350505050565b5f546001600160a01b03163314610c0e5760405162461bcd60e51b815260040161069890611c5d565b5f6004610c1c600186611c30565b81548110610c2c57610c2c611c49565b5f9182526020909120600a90910201600781015490915060ff16610c8b5760405162461bcd60e51b81526020600482015260166024820152752a30b9b59034b9903737ba1031b7b6b83632ba32b21760511b6044820152606401610698565b6007810154610100900460ff1615610ce15760405162461bcd60e51b81526020600482015260196024820152782a30b9b59034b99030b63932b0b23c903932bbb0b93232b21760391b6044820152606401610698565b8151835114610d2a5760405162461bcd60e51b815260206004820152601560248201527424b73b30b634b21034b7383aba103632b733ba341760591b6044820152606401610698565b5f805b8351811015610d6557838181518110610d4857610d48611c49565b602002602001015182610d5b9190611ca8565b9150600101610d2d565b508160050154811115610dc85760405162461bcd60e51b815260206004820152602560248201527f546f74616c20726577617264732065786365656420746865207461736b2072656044820152643bb0b9321760d91b6064820152608401610698565b60078201805461ff00191661010017905560018201546040516001600160a01b039091169086907f907649ae5ab1beeecf203ab18eeb4acc8270d8db46fbf7e7becd29b5d2d8d45790610e1e9088908890611eaf565b60405180910390a35050505050565b335f908152600360208190526040909120015460ff1615610e905760405162461bcd60e51b815260206004820152601f60248201527f4f70657261746f7220697320616c726561647920726567697374657265642e006044820152606401610698565b6040805160808101825283815260208082018490525f8284018190526001606084015233815260039091529190912081518190610ecd9082611f31565b5060208201516001820190610ee29082611f31565b5060408281015160028301556060909201516003909101805460ff19169115159190911790555133907f11a85ea4a40584362c3d9c17685709a2e02b466ac78d5eb00b6aff73d90f580590610f38908490611591565b60405180910390a25050565b60048181548110610f53575f80fd5b5f9182526020909120600a90910201805460018201546002830180549294506001600160a01b039091169291610f8890611be4565b80601f0160208091040260200160405190810160405280929190818152602001828054610fb490611be4565b8015610fff5780601f10610fd657610100808354040283529160200191610fff565b820191905f5260205f20905b815481529060010190602001808311610fe257829003601f168201915b5050506003840154600485015460058601546006870154600788015460088901805498999598949750929550909360ff808316946101009093041692909161104690611be4565b80601f016020809104026020016040519081016040528092919081815260200182805461107290611be4565b80156110bd5780601f10611094576101008083540402835291602001916110bd565b820191905f5260205f20905b8154815290600101906020018083116110a057829003601f168201915b505050505090508a565b5f806109026001600160a01b03167fb6c1c1ca0705e625b8d2de3b8a502fa0f8743a5761c7e20f734657df647a422689898989898960405160240161111196959493929190611ff0565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b031990941693909317909252905161114f9190611d94565b5f604051808303815f865af19150503d805f8114611188576040519150601f19603f3d011682016040523d82523d5f602084013e61118d565b606091505b50909998505050505050505050565b5f546001600160a01b031633146111c55760405162461bcd60e51b815260040161069890611c5d565b6111cd6113fe565b60018054825233602080840191825260408085018a8152606086018a90526080860189905260a0860188905260c086018790525f60e0870181905261010087018190528251818152938401909252610140860192909252600480549485018155905283517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b600a90940293840190815591517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19c840180546001600160a01b0319166001600160a01b039092169190911790555183927f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19d01906112cf9082611f31565b50606082015160038201556080820151600482015560a0820151600582015560c0820151600682015560e08201516007820180546101008086015115150261ff00199315159390931661ffff1990911617919091179055610120820151600882019061133b9082611f31565b50610140820151805161135891600984019160209091019061145d565b50506001805491505f61136a8361206c565b909155505080516040513391907fdac01f3886a95e591b6ebef63858e5383c5f85b170876b4e18f380599b0a68a5906113a8908a9088908890612084565b60405180910390a3505050505050565b6001600160a01b0381165f90815260036020819052604082209081015460ff16156113f65760028101545f906113f3565b9695505050505050565b50505b505f92915050565b6040518061016001604052805f81526020015f6001600160a01b03168152602001606081526020015f81526020015f81526020015f81526020015f81526020015f151581526020015f1515815260200160608152602001606081525090565b828054828255905f5260205f209081019282156114b0579160200282015b828111156114b057825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061147b565b506114bc9291506114c0565b5090565b5b808211156114bc575f81556001016114c1565b80356001600160a01b03811681146114ea575f80fd5b919050565b5f602082840312156114ff575f80fd5b611508826114d4565b9392505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b608081525f61154f608083018761150f565b8281036020840152611561818761150f565b6040840195909552505090151560609091015292915050565b5f6020828403121561158a575f80fd5b5035919050565b602081525f611508602083018461150f565b5f815180845260208085019450602084015f5b838110156115db5781516001600160a01b0316875295820195908201906001016115b6565b509495945050505050565b60208152815160208201525f602083015161160c60408401826001600160a01b03169052565b50604083015161016080606085015261162961018085018361150f565b915060608501516080850152608085015160a085015260a085015160c085015260c085015160e085015260e08501516101006116688187018315159052565b860151905061012061167d8682018315159052565b80870151915050601f1961014081878603018188015261169d858461150f565b9088015187820390920184880152935090506113e983826115a3565b5f805f805f606086880312156116cd575f80fd5b8535945060208601356001600160401b03808211156116ea575f80fd5b818801915088601f8301126116fd575f80fd5b81358181111561170b575f80fd5b89602082850101111561171c575f80fd5b602083019650809550506040880135915080821115611739575f80fd5b818801915088601f83011261174c575f80fd5b81358181111561175a575f80fd5b8960208260051b850101111561176e575f80fd5b9699959850939650602001949392505050565b5f8060408385031215611792575f80fd5b823591506117a2602084016114d4565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156117e7576117e76117ab565b604052919050565b5f6001600160401b03821115611807576118076117ab565b5060051b60200190565b5f82601f830112611820575f80fd5b81356020611835611830836117ef565b6117bf565b8083825260208201915060208460051b870101935086841115611856575f80fd5b602086015b84811015611872578035835291830191830161185b565b509695505050505050565b5f805f6060848603121561188f575f80fd5b833592506020808501356001600160401b03808211156118ad575f80fd5b818701915087601f8301126118c0575f80fd5b81356118ce611830826117ef565b81815260059190911b8301840190848101908a8311156118ec575f80fd5b938501935b8285101561191157611902856114d4565b825293850193908501906118f1565b965050506040870135925080831115611928575f80fd5b505061193686828701611811565b9150509250925092565b5f6001600160401b03831115611958576119586117ab565b61196b601f8401601f19166020016117bf565b905082815283838301111561197e575f80fd5b828260208301375f602084830101529392505050565b5f82601f8301126119a3575f80fd5b61150883833560208501611940565b5f80604083850312156119c3575f80fd5b82356001600160401b03808211156119d9575f80fd5b818501915085601f8301126119ec575f80fd5b6119fb86833560208501611940565b93506020850135915080821115611a10575f80fd5b50611a1d85828601611994565b9150509250929050565b8a81526001600160a01b038a166020820152610140604082018190525f90611a518382018c61150f565b90508960608401528860808401528760a08401528660c084015285151560e0840152841515610100840152828103610120840152611a8f818561150f565b9d9c50505050505050505050505050565b80356001600160401b03811681146114ea575f80fd5b5f805f805f8060c08789031215611acb575f80fd5b86356001600160401b0380821115611ae1575f80fd5b611aed8a838b01611994565b97506020890135915080821115611b02575f80fd5b611b0e8a838b01611994565b96506040890135915080821115611b23575f80fd5b611b2f8a838b01611994565b9550611b3d60608a01611aa0565b94506080890135915080821115611b52575f80fd5b611b5e8a838b01611994565b935060a0890135915080821115611b73575f80fd5b50611b8089828a01611994565b9150509295509295509295565b5f805f805f60a08688031215611ba1575f80fd5b85356001600160401b03811115611bb6575f80fd5b611bc288828901611994565b9860208801359850604088013597606081013597506080013595509350505050565b600181811c90821680611bf857607f821691505b602082108103611c1657634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b81810381811115611c4357611c43611c1c565b92915050565b634e487b7160e01b5f52603260045260245ffd5b6020808252602b908201527f4f6e6c7920636f6e7472616374206f776e65722063616e2063616c6c2074686960408201526a3990333ab731ba34b7b71760a91b606082015260800190565b80820180821115611c4357611c43611c1c565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b8581525f6020606081840152611cfd606084018789611cbb565b8381036040850152848152818101600586901b82018301875f5b88811015611d8357848303601f190184528135368b9003601e19018112611d3c575f80fd5b8a0186810190356001600160401b03811115611d56575f80fd5b803603821315611d64575f80fd5b611d6f858284611cbb565b958801959450505090850190600101611d17565b50909b9a5050505050505050505050565b5f82518060208501845e5f920191825250919050565b601f821115611df157805f5260205f20601f840160051c81016020851015611dcf5750805b601f840160051c820191505b81811015611dee575f8155600101611ddb565b50505b505050565b6001600160401b03831115611e0d57611e0d6117ab565b611e2183611e1b8354611be4565b83611daa565b5f601f841160018114611e52575f8515611e3b5750838201355b5f19600387901b1c1916600186901b178355611dee565b5f83815260208120601f198716915b82811015611e815786850135825560209485019460019092019101611e61565b5086821015611e9d575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b604080825283519082018190525f906020906060840190828701845b82811015611ef05781516001600160a01b031684529284019290840190600101611ecb565b505050838103828501528451808252858301918301905f5b81811015611f2457835183529284019291840191600101611f08565b5090979650505050505050565b81516001600160401b03811115611f4a57611f4a6117ab565b611f5e81611f588454611be4565b84611daa565b602080601f831160018114611f91575f8415611f7a5750858301515b5f19600386901b1c1916600185901b178555611fe8565b5f85815260208120601f198616915b82811015611fbf57888601518255948401946001909101908401611fa0565b5085821015611fdc57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b60c081525f61200260c083018961150f565b8281036020840152612014818961150f565b90508281036040840152612028818861150f565b90506001600160401b0386166060840152828103608084015261204b818661150f565b905082810360a084015261205f818561150f565b9998505050505050505050565b5f6001820161207d5761207d611c1c565b5060010190565b606081525f612096606083018661150f565b6020830194909452506040015291905056fea26469706673582212201a1d70b18c573a261436a170dcb7e89f1c1281906d8879f54beb8a7add22017c64736f6c63430008190033",
}

// ContractavsserviceABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractavsserviceMetaData.ABI instead.
var ContractavsserviceABI = ContractavsserviceMetaData.ABI

// ContractavsserviceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractavsserviceMetaData.Bin instead.
var ContractavsserviceBin = ContractavsserviceMetaData.Bin

// DeployContractavsservice deploys a new Ethereum contract, binding an instance of Contractavsservice to it.
func DeployContractavsservice(auth *bind.TransactOpts, backend bind.ContractBackend, _slashPercentage *big.Int) (common.Address, *types.Transaction, *Contractavsservice, error) {
	parsed, err := ContractavsserviceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractavsserviceBin), backend, _slashPercentage)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contractavsservice{ContractavsserviceCaller: ContractavsserviceCaller{contract: contract}, ContractavsserviceTransactor: ContractavsserviceTransactor{contract: contract}, ContractavsserviceFilterer: ContractavsserviceFilterer{contract: contract}}, nil
}

// Contractavsservice is an auto generated Go binding around an Ethereum contract.
type Contractavsservice struct {
	ContractavsserviceCaller     // Read-only binding to the contract
	ContractavsserviceTransactor // Write-only binding to the contract
	ContractavsserviceFilterer   // Log filterer for contract events
}

// ContractavsserviceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractavsserviceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractavsserviceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractavsserviceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractavsserviceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractavsserviceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractavsserviceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractavsserviceSession struct {
	Contract     *Contractavsservice // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractavsserviceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractavsserviceCallerSession struct {
	Contract *ContractavsserviceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContractavsserviceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractavsserviceTransactorSession struct {
	Contract     *ContractavsserviceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContractavsserviceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractavsserviceRaw struct {
	Contract *Contractavsservice // Generic contract binding to access the raw methods on
}

// ContractavsserviceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractavsserviceCallerRaw struct {
	Contract *ContractavsserviceCaller // Generic read-only contract binding to access the raw methods on
}

// ContractavsserviceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractavsserviceTransactorRaw struct {
	Contract *ContractavsserviceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractavsservice creates a new instance of Contractavsservice, bound to a specific deployed contract.
func NewContractavsservice(address common.Address, backend bind.ContractBackend) (*Contractavsservice, error) {
	contract, err := bindContractavsservice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contractavsservice{ContractavsserviceCaller: ContractavsserviceCaller{contract: contract}, ContractavsserviceTransactor: ContractavsserviceTransactor{contract: contract}, ContractavsserviceFilterer: ContractavsserviceFilterer{contract: contract}}, nil
}

// NewContractavsserviceCaller creates a new read-only instance of Contractavsservice, bound to a specific deployed contract.
func NewContractavsserviceCaller(address common.Address, caller bind.ContractCaller) (*ContractavsserviceCaller, error) {
	contract, err := bindContractavsservice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractavsserviceCaller{contract: contract}, nil
}

// NewContractavsserviceTransactor creates a new write-only instance of Contractavsservice, bound to a specific deployed contract.
func NewContractavsserviceTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractavsserviceTransactor, error) {
	contract, err := bindContractavsservice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractavsserviceTransactor{contract: contract}, nil
}

// NewContractavsserviceFilterer creates a new log filterer instance of Contractavsservice, bound to a specific deployed contract.
func NewContractavsserviceFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractavsserviceFilterer, error) {
	contract, err := bindContractavsservice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractavsserviceFilterer{contract: contract}, nil
}

// bindContractavsservice binds a generic wrapper to an already deployed contract.
func bindContractavsservice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractavsserviceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractavsservice *ContractavsserviceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractavsservice.Contract.ContractavsserviceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractavsservice *ContractavsserviceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractavsservice.Contract.ContractavsserviceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractavsservice *ContractavsserviceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractavsservice.Contract.ContractavsserviceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractavsservice *ContractavsserviceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractavsservice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractavsservice *ContractavsserviceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractavsservice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractavsservice *ContractavsserviceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractavsservice.Contract.contract.Transact(opts, method, params...)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x19c2e665.
//
// Solidity: function allTaskResponses(uint256 ) view returns(bytes)
func (_Contractavsservice *ContractavsserviceCaller) AllTaskResponses(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x19c2e665.
//
// Solidity: function allTaskResponses(uint256 ) view returns(bytes)
func (_Contractavsservice *ContractavsserviceSession) AllTaskResponses(arg0 *big.Int) ([]byte, error) {
	return _Contractavsservice.Contract.AllTaskResponses(&_Contractavsservice.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x19c2e665.
//
// Solidity: function allTaskResponses(uint256 ) view returns(bytes)
func (_Contractavsservice *ContractavsserviceCallerSession) AllTaskResponses(arg0 *big.Int) ([]byte, error) {
	return _Contractavsservice.Contract.AllTaskResponses(&_Contractavsservice.CallOpts, arg0)
}

// GetOperatorCount is a free data retrieval call binding the contract method 0x1a4a7c61.
//
// Solidity: function getOperatorCount() view returns(uint256)
func (_Contractavsservice *ContractavsserviceCaller) GetOperatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "getOperatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorCount is a free data retrieval call binding the contract method 0x1a4a7c61.
//
// Solidity: function getOperatorCount() view returns(uint256)
func (_Contractavsservice *ContractavsserviceSession) GetOperatorCount() (*big.Int, error) {
	return _Contractavsservice.Contract.GetOperatorCount(&_Contractavsservice.CallOpts)
}

// GetOperatorCount is a free data retrieval call binding the contract method 0x1a4a7c61.
//
// Solidity: function getOperatorCount() view returns(uint256)
func (_Contractavsservice *ContractavsserviceCallerSession) GetOperatorCount() (*big.Int, error) {
	return _Contractavsservice.Contract.GetOperatorCount(&_Contractavsservice.CallOpts)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 taskId) view returns((uint256,address,string,uint256,uint256,uint256,uint256,bool,bool,bytes,address[]))
func (_Contractavsservice *ContractavsserviceCaller) GetTask(opts *bind.CallOpts, taskId *big.Int) (AvsServiceContractTask, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "getTask", taskId)

	if err != nil {
		return *new(AvsServiceContractTask), err
	}

	out0 := *abi.ConvertType(out[0], new(AvsServiceContractTask)).(*AvsServiceContractTask)

	return out0, err

}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 taskId) view returns((uint256,address,string,uint256,uint256,uint256,uint256,bool,bool,bytes,address[]))
func (_Contractavsservice *ContractavsserviceSession) GetTask(taskId *big.Int) (AvsServiceContractTask, error) {
	return _Contractavsservice.Contract.GetTask(&_Contractavsservice.CallOpts, taskId)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 taskId) view returns((uint256,address,string,uint256,uint256,uint256,uint256,bool,bool,bytes,address[]))
func (_Contractavsservice *ContractavsserviceCallerSession) GetTask(taskId *big.Int) (AvsServiceContractTask, error) {
	return _Contractavsservice.Contract.GetTask(&_Contractavsservice.CallOpts, taskId)
}

// GetTaskCount is a free data retrieval call binding the contract method 0xc17a340e.
//
// Solidity: function getTaskCount() view returns(uint256)
func (_Contractavsservice *ContractavsserviceCaller) GetTaskCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "getTaskCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTaskCount is a free data retrieval call binding the contract method 0xc17a340e.
//
// Solidity: function getTaskCount() view returns(uint256)
func (_Contractavsservice *ContractavsserviceSession) GetTaskCount() (*big.Int, error) {
	return _Contractavsservice.Contract.GetTaskCount(&_Contractavsservice.CallOpts)
}

// GetTaskCount is a free data retrieval call binding the contract method 0xc17a340e.
//
// Solidity: function getTaskCount() view returns(uint256)
func (_Contractavsservice *ContractavsserviceCallerSession) GetTaskCount() (*big.Int, error) {
	return _Contractavsservice.Contract.GetTaskCount(&_Contractavsservice.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bytes publicKey, string name, uint256 totalRewardsEarned, bool isRegistered)
func (_Contractavsservice *ContractavsserviceCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (struct {
	PublicKey          []byte
	Name               string
	TotalRewardsEarned *big.Int
	IsRegistered       bool
}, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "operators", arg0)

	outstruct := new(struct {
		PublicKey          []byte
		Name               string
		TotalRewardsEarned *big.Int
		IsRegistered       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PublicKey = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.TotalRewardsEarned = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsRegistered = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bytes publicKey, string name, uint256 totalRewardsEarned, bool isRegistered)
func (_Contractavsservice *ContractavsserviceSession) Operators(arg0 common.Address) (struct {
	PublicKey          []byte
	Name               string
	TotalRewardsEarned *big.Int
	IsRegistered       bool
}, error) {
	return _Contractavsservice.Contract.Operators(&_Contractavsservice.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bytes publicKey, string name, uint256 totalRewardsEarned, bool isRegistered)
func (_Contractavsservice *ContractavsserviceCallerSession) Operators(arg0 common.Address) (struct {
	PublicKey          []byte
	Name               string
	TotalRewardsEarned *big.Int
	IsRegistered       bool
}, error) {
	return _Contractavsservice.Contract.Operators(&_Contractavsservice.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contractavsservice *ContractavsserviceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contractavsservice *ContractavsserviceSession) Owner() (common.Address, error) {
	return _Contractavsservice.Contract.Owner(&_Contractavsservice.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contractavsservice *ContractavsserviceCallerSession) Owner() (common.Address, error) {
	return _Contractavsservice.Contract.Owner(&_Contractavsservice.CallOpts)
}

// SlashPercentage is a free data retrieval call binding the contract method 0xbd58199a.
//
// Solidity: function slashPercentage() view returns(uint256)
func (_Contractavsservice *ContractavsserviceCaller) SlashPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "slashPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashPercentage is a free data retrieval call binding the contract method 0xbd58199a.
//
// Solidity: function slashPercentage() view returns(uint256)
func (_Contractavsservice *ContractavsserviceSession) SlashPercentage() (*big.Int, error) {
	return _Contractavsservice.Contract.SlashPercentage(&_Contractavsservice.CallOpts)
}

// SlashPercentage is a free data retrieval call binding the contract method 0xbd58199a.
//
// Solidity: function slashPercentage() view returns(uint256)
func (_Contractavsservice *ContractavsserviceCallerSession) SlashPercentage() (*big.Int, error) {
	return _Contractavsservice.Contract.SlashPercentage(&_Contractavsservice.CallOpts)
}

// TaskIdCounter is a free data retrieval call binding the contract method 0xad8ef91c.
//
// Solidity: function taskIdCounter() view returns(uint256)
func (_Contractavsservice *ContractavsserviceCaller) TaskIdCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "taskIdCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskIdCounter is a free data retrieval call binding the contract method 0xad8ef91c.
//
// Solidity: function taskIdCounter() view returns(uint256)
func (_Contractavsservice *ContractavsserviceSession) TaskIdCounter() (*big.Int, error) {
	return _Contractavsservice.Contract.TaskIdCounter(&_Contractavsservice.CallOpts)
}

// TaskIdCounter is a free data retrieval call binding the contract method 0xad8ef91c.
//
// Solidity: function taskIdCounter() view returns(uint256)
func (_Contractavsservice *ContractavsserviceCallerSession) TaskIdCounter() (*big.Int, error) {
	return _Contractavsservice.Contract.TaskIdCounter(&_Contractavsservice.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 taskId, address issuer, string description, uint256 numberX, uint256 numberY, uint256 reward, uint256 deadline, bool isCompleted, bool isRewarded, bytes operatorSignatures)
func (_Contractavsservice *ContractavsserviceCaller) Tasks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TaskId             *big.Int
	Issuer             common.Address
	Description        string
	NumberX            *big.Int
	NumberY            *big.Int
	Reward             *big.Int
	Deadline           *big.Int
	IsCompleted        bool
	IsRewarded         bool
	OperatorSignatures []byte
}, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		TaskId             *big.Int
		Issuer             common.Address
		Description        string
		NumberX            *big.Int
		NumberY            *big.Int
		Reward             *big.Int
		Deadline           *big.Int
		IsCompleted        bool
		IsRewarded         bool
		OperatorSignatures []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TaskId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Issuer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Description = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.NumberX = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.NumberY = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.IsCompleted = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.IsRewarded = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.OperatorSignatures = *abi.ConvertType(out[9], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 taskId, address issuer, string description, uint256 numberX, uint256 numberY, uint256 reward, uint256 deadline, bool isCompleted, bool isRewarded, bytes operatorSignatures)
func (_Contractavsservice *ContractavsserviceSession) Tasks(arg0 *big.Int) (struct {
	TaskId             *big.Int
	Issuer             common.Address
	Description        string
	NumberX            *big.Int
	NumberY            *big.Int
	Reward             *big.Int
	Deadline           *big.Int
	IsCompleted        bool
	IsRewarded         bool
	OperatorSignatures []byte
}, error) {
	return _Contractavsservice.Contract.Tasks(&_Contractavsservice.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 taskId, address issuer, string description, uint256 numberX, uint256 numberY, uint256 reward, uint256 deadline, bool isCompleted, bool isRewarded, bytes operatorSignatures)
func (_Contractavsservice *ContractavsserviceCallerSession) Tasks(arg0 *big.Int) (struct {
	TaskId             *big.Int
	Issuer             common.Address
	Description        string
	NumberX            *big.Int
	NumberY            *big.Int
	Reward             *big.Int
	Deadline           *big.Int
	IsCompleted        bool
	IsRewarded         bool
	OperatorSignatures []byte
}, error) {
	return _Contractavsservice.Contract.Tasks(&_Contractavsservice.CallOpts, arg0)
}

// CompleteTask is a paid mutator transaction binding the contract method 0x2f95670e.
//
// Solidity: function completeTask(uint256 taskId, bytes aggSig, bytes[] pubkeys) returns()
func (_Contractavsservice *ContractavsserviceTransactor) CompleteTask(opts *bind.TransactOpts, taskId *big.Int, aggSig []byte, pubkeys [][]byte) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "completeTask", taskId, aggSig, pubkeys)
}

// CompleteTask is a paid mutator transaction binding the contract method 0x2f95670e.
//
// Solidity: function completeTask(uint256 taskId, bytes aggSig, bytes[] pubkeys) returns()
func (_Contractavsservice *ContractavsserviceSession) CompleteTask(taskId *big.Int, aggSig []byte, pubkeys [][]byte) (*types.Transaction, error) {
	return _Contractavsservice.Contract.CompleteTask(&_Contractavsservice.TransactOpts, taskId, aggSig, pubkeys)
}

// CompleteTask is a paid mutator transaction binding the contract method 0x2f95670e.
//
// Solidity: function completeTask(uint256 taskId, bytes aggSig, bytes[] pubkeys) returns()
func (_Contractavsservice *ContractavsserviceTransactorSession) CompleteTask(taskId *big.Int, aggSig []byte, pubkeys [][]byte) (*types.Transaction, error) {
	return _Contractavsservice.Contract.CompleteTask(&_Contractavsservice.TransactOpts, taskId, aggSig, pubkeys)
}

// CreateTask is a paid mutator transaction binding the contract method 0xaebe7821.
//
// Solidity: function createTask(string description, uint256 numberX, uint256 numberY, uint256 reward, uint256 deadline) returns()
func (_Contractavsservice *ContractavsserviceTransactor) CreateTask(opts *bind.TransactOpts, description string, numberX *big.Int, numberY *big.Int, reward *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "createTask", description, numberX, numberY, reward, deadline)
}

// CreateTask is a paid mutator transaction binding the contract method 0xaebe7821.
//
// Solidity: function createTask(string description, uint256 numberX, uint256 numberY, uint256 reward, uint256 deadline) returns()
func (_Contractavsservice *ContractavsserviceSession) CreateTask(description string, numberX *big.Int, numberY *big.Int, reward *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Contractavsservice.Contract.CreateTask(&_Contractavsservice.TransactOpts, description, numberX, numberY, reward, deadline)
}

// CreateTask is a paid mutator transaction binding the contract method 0xaebe7821.
//
// Solidity: function createTask(string description, uint256 numberX, uint256 numberY, uint256 reward, uint256 deadline) returns()
func (_Contractavsservice *ContractavsserviceTransactorSession) CreateTask(description string, numberX *big.Int, numberY *big.Int, reward *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Contractavsservice.Contract.CreateTask(&_Contractavsservice.TransactOpts, description, numberX, numberY, reward, deadline)
}

// RegisterAVSToExocore is a paid mutator transaction binding the contract method 0xa24d3c11.
//
// Solidity: function registerAVSToExocore(string avsName, string avsAddress, string operatorAddress, uint64 action, string avsOwnerAddress, string assetID) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) RegisterAVSToExocore(opts *bind.TransactOpts, avsName string, avsAddress string, operatorAddress string, action uint64, avsOwnerAddress string, assetID string) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "registerAVSToExocore", avsName, avsAddress, operatorAddress, action, avsOwnerAddress, assetID)
}

// RegisterAVSToExocore is a paid mutator transaction binding the contract method 0xa24d3c11.
//
// Solidity: function registerAVSToExocore(string avsName, string avsAddress, string operatorAddress, uint64 action, string avsOwnerAddress, string assetID) returns(bool)
func (_Contractavsservice *ContractavsserviceSession) RegisterAVSToExocore(avsName string, avsAddress string, operatorAddress string, action uint64, avsOwnerAddress string, assetID string) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterAVSToExocore(&_Contractavsservice.TransactOpts, avsName, avsAddress, operatorAddress, action, avsOwnerAddress, assetID)
}

// RegisterAVSToExocore is a paid mutator transaction binding the contract method 0xa24d3c11.
//
// Solidity: function registerAVSToExocore(string avsName, string avsAddress, string operatorAddress, uint64 action, string avsOwnerAddress, string assetID) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) RegisterAVSToExocore(avsName string, avsAddress string, operatorAddress string, action uint64, avsOwnerAddress string, assetID string) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterAVSToExocore(&_Contractavsservice.TransactOpts, avsName, avsAddress, operatorAddress, action, avsOwnerAddress, assetID)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x86cfceea.
//
// Solidity: function registerOperatorToAVS(bytes publicKey, string name) returns()
func (_Contractavsservice *ContractavsserviceTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, publicKey []byte, name string) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "registerOperatorToAVS", publicKey, name)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x86cfceea.
//
// Solidity: function registerOperatorToAVS(bytes publicKey, string name) returns()
func (_Contractavsservice *ContractavsserviceSession) RegisterOperatorToAVS(publicKey []byte, name string) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterOperatorToAVS(&_Contractavsservice.TransactOpts, publicKey, name)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x86cfceea.
//
// Solidity: function registerOperatorToAVS(bytes publicKey, string name) returns()
func (_Contractavsservice *ContractavsserviceTransactorSession) RegisterOperatorToAVS(publicKey []byte, name string) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterOperatorToAVS(&_Contractavsservice.TransactOpts, publicKey, name)
}

// RewardTask is a paid mutator transaction binding the contract method 0x648b14c5.
//
// Solidity: function rewardTask(uint256 taskId, address[] operatorAddresses, uint256[] rewards) returns()
func (_Contractavsservice *ContractavsserviceTransactor) RewardTask(opts *bind.TransactOpts, taskId *big.Int, operatorAddresses []common.Address, rewards []*big.Int) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "rewardTask", taskId, operatorAddresses, rewards)
}

// RewardTask is a paid mutator transaction binding the contract method 0x648b14c5.
//
// Solidity: function rewardTask(uint256 taskId, address[] operatorAddresses, uint256[] rewards) returns()
func (_Contractavsservice *ContractavsserviceSession) RewardTask(taskId *big.Int, operatorAddresses []common.Address, rewards []*big.Int) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RewardTask(&_Contractavsservice.TransactOpts, taskId, operatorAddresses, rewards)
}

// RewardTask is a paid mutator transaction binding the contract method 0x648b14c5.
//
// Solidity: function rewardTask(uint256 taskId, address[] operatorAddresses, uint256[] rewards) returns()
func (_Contractavsservice *ContractavsserviceTransactorSession) RewardTask(taskId *big.Int, operatorAddresses []common.Address, rewards []*big.Int) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RewardTask(&_Contractavsservice.TransactOpts, taskId, operatorAddresses, rewards)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x3a3e8907.
//
// Solidity: function slashOperator(uint256 taskId, address operatorAddress) returns()
func (_Contractavsservice *ContractavsserviceTransactor) SlashOperator(opts *bind.TransactOpts, taskId *big.Int, operatorAddress common.Address) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "slashOperator", taskId, operatorAddress)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x3a3e8907.
//
// Solidity: function slashOperator(uint256 taskId, address operatorAddress) returns()
func (_Contractavsservice *ContractavsserviceSession) SlashOperator(taskId *big.Int, operatorAddress common.Address) (*types.Transaction, error) {
	return _Contractavsservice.Contract.SlashOperator(&_Contractavsservice.TransactOpts, taskId, operatorAddress)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x3a3e8907.
//
// Solidity: function slashOperator(uint256 taskId, address operatorAddress) returns()
func (_Contractavsservice *ContractavsserviceTransactorSession) SlashOperator(taskId *big.Int, operatorAddress common.Address) (*types.Transaction, error) {
	return _Contractavsservice.Contract.SlashOperator(&_Contractavsservice.TransactOpts, taskId, operatorAddress)
}

// ContractavsserviceOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the Contractavsservice contract.
type ContractavsserviceOperatorRegisteredIterator struct {
	Event *ContractavsserviceOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ContractavsserviceOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractavsserviceOperatorRegistered)
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
		it.Event = new(ContractavsserviceOperatorRegistered)
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
func (it *ContractavsserviceOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractavsserviceOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractavsserviceOperatorRegistered represents a OperatorRegistered event raised by the Contractavsservice contract.
type ContractavsserviceOperatorRegistered struct {
	OperatorAddress common.Address
	Name            string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x11a85ea4a40584362c3d9c17685709a2e02b466ac78d5eb00b6aff73d90f5805.
//
// Solidity: event OperatorRegistered(address indexed operatorAddress, string name)
func (_Contractavsservice *ContractavsserviceFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operatorAddress []common.Address) (*ContractavsserviceOperatorRegisteredIterator, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contractavsservice.contract.FilterLogs(opts, "OperatorRegistered", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractavsserviceOperatorRegisteredIterator{contract: _Contractavsservice.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x11a85ea4a40584362c3d9c17685709a2e02b466ac78d5eb00b6aff73d90f5805.
//
// Solidity: event OperatorRegistered(address indexed operatorAddress, string name)
func (_Contractavsservice *ContractavsserviceFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractavsserviceOperatorRegistered, operatorAddress []common.Address) (event.Subscription, error) {

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contractavsservice.contract.WatchLogs(opts, "OperatorRegistered", operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractavsserviceOperatorRegistered)
				if err := _Contractavsservice.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x11a85ea4a40584362c3d9c17685709a2e02b466ac78d5eb00b6aff73d90f5805.
//
// Solidity: event OperatorRegistered(address indexed operatorAddress, string name)
func (_Contractavsservice *ContractavsserviceFilterer) ParseOperatorRegistered(log types.Log) (*ContractavsserviceOperatorRegistered, error) {
	event := new(ContractavsserviceOperatorRegistered)
	if err := _Contractavsservice.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractavsserviceOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the Contractavsservice contract.
type ContractavsserviceOperatorSlashedIterator struct {
	Event *ContractavsserviceOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *ContractavsserviceOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractavsserviceOperatorSlashed)
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
		it.Event = new(ContractavsserviceOperatorSlashed)
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
func (it *ContractavsserviceOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractavsserviceOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractavsserviceOperatorSlashed represents a OperatorSlashed event raised by the Contractavsservice contract.
type ContractavsserviceOperatorSlashed struct {
	TaskId          *big.Int
	OperatorAddress common.Address
	SlashAmount     *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0xa2e524bd0f71903485fbb3d6d50cb305f61005ceea2047c3ac92aa7e0d104306.
//
// Solidity: event OperatorSlashed(uint256 indexed taskId, address indexed operatorAddress, uint256 slashAmount)
func (_Contractavsservice *ContractavsserviceFilterer) FilterOperatorSlashed(opts *bind.FilterOpts, taskId []*big.Int, operatorAddress []common.Address) (*ContractavsserviceOperatorSlashedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contractavsservice.contract.FilterLogs(opts, "OperatorSlashed", taskIdRule, operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractavsserviceOperatorSlashedIterator{contract: _Contractavsservice.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0xa2e524bd0f71903485fbb3d6d50cb305f61005ceea2047c3ac92aa7e0d104306.
//
// Solidity: event OperatorSlashed(uint256 indexed taskId, address indexed operatorAddress, uint256 slashAmount)
func (_Contractavsservice *ContractavsserviceFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *ContractavsserviceOperatorSlashed, taskId []*big.Int, operatorAddress []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Contractavsservice.contract.WatchLogs(opts, "OperatorSlashed", taskIdRule, operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractavsserviceOperatorSlashed)
				if err := _Contractavsservice.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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

// ParseOperatorSlashed is a log parse operation binding the contract event 0xa2e524bd0f71903485fbb3d6d50cb305f61005ceea2047c3ac92aa7e0d104306.
//
// Solidity: event OperatorSlashed(uint256 indexed taskId, address indexed operatorAddress, uint256 slashAmount)
func (_Contractavsservice *ContractavsserviceFilterer) ParseOperatorSlashed(log types.Log) (*ContractavsserviceOperatorSlashed, error) {
	event := new(ContractavsserviceOperatorSlashed)
	if err := _Contractavsservice.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractavsserviceTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the Contractavsservice contract.
type ContractavsserviceTaskCompletedIterator struct {
	Event *ContractavsserviceTaskCompleted // Event containing the contract specifics and raw log

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
func (it *ContractavsserviceTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractavsserviceTaskCompleted)
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
		it.Event = new(ContractavsserviceTaskCompleted)
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
func (it *ContractavsserviceTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractavsserviceTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractavsserviceTaskCompleted represents a TaskCompleted event raised by the Contractavsservice contract.
type ContractavsserviceTaskCompleted struct {
	TaskId *big.Int
	Issuer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0xbb5889c77948badf90e8a5c73d55265e5f5d6e4837a79a78c5669691b897faed.
//
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed issuer)
func (_Contractavsservice *ContractavsserviceFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskId []*big.Int, issuer []common.Address) (*ContractavsserviceTaskCompletedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Contractavsservice.contract.FilterLogs(opts, "TaskCompleted", taskIdRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &ContractavsserviceTaskCompletedIterator{contract: _Contractavsservice.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0xbb5889c77948badf90e8a5c73d55265e5f5d6e4837a79a78c5669691b897faed.
//
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed issuer)
func (_Contractavsservice *ContractavsserviceFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractavsserviceTaskCompleted, taskId []*big.Int, issuer []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Contractavsservice.contract.WatchLogs(opts, "TaskCompleted", taskIdRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractavsserviceTaskCompleted)
				if err := _Contractavsservice.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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

// ParseTaskCompleted is a log parse operation binding the contract event 0xbb5889c77948badf90e8a5c73d55265e5f5d6e4837a79a78c5669691b897faed.
//
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed issuer)
func (_Contractavsservice *ContractavsserviceFilterer) ParseTaskCompleted(log types.Log) (*ContractavsserviceTaskCompleted, error) {
	event := new(ContractavsserviceTaskCompleted)
	if err := _Contractavsservice.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractavsserviceTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the Contractavsservice contract.
type ContractavsserviceTaskCreatedIterator struct {
	Event *ContractavsserviceTaskCreated // Event containing the contract specifics and raw log

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
func (it *ContractavsserviceTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractavsserviceTaskCreated)
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
		it.Event = new(ContractavsserviceTaskCreated)
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
func (it *ContractavsserviceTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractavsserviceTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractavsserviceTaskCreated represents a TaskCreated event raised by the Contractavsservice contract.
type ContractavsserviceTaskCreated struct {
	TaskId      *big.Int
	Issuer      common.Address
	Description string
	Reward      *big.Int
	Deadline    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0xdac01f3886a95e591b6ebef63858e5383c5f85b170876b4e18f380599b0a68a5.
//
// Solidity: event TaskCreated(uint256 indexed taskId, address indexed issuer, string description, uint256 reward, uint256 deadline)
func (_Contractavsservice *ContractavsserviceFilterer) FilterTaskCreated(opts *bind.FilterOpts, taskId []*big.Int, issuer []common.Address) (*ContractavsserviceTaskCreatedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Contractavsservice.contract.FilterLogs(opts, "TaskCreated", taskIdRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &ContractavsserviceTaskCreatedIterator{contract: _Contractavsservice.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0xdac01f3886a95e591b6ebef63858e5383c5f85b170876b4e18f380599b0a68a5.
//
// Solidity: event TaskCreated(uint256 indexed taskId, address indexed issuer, string description, uint256 reward, uint256 deadline)
func (_Contractavsservice *ContractavsserviceFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractavsserviceTaskCreated, taskId []*big.Int, issuer []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Contractavsservice.contract.WatchLogs(opts, "TaskCreated", taskIdRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractavsserviceTaskCreated)
				if err := _Contractavsservice.contract.UnpackLog(event, "TaskCreated", log); err != nil {
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

// ParseTaskCreated is a log parse operation binding the contract event 0xdac01f3886a95e591b6ebef63858e5383c5f85b170876b4e18f380599b0a68a5.
//
// Solidity: event TaskCreated(uint256 indexed taskId, address indexed issuer, string description, uint256 reward, uint256 deadline)
func (_Contractavsservice *ContractavsserviceFilterer) ParseTaskCreated(log types.Log) (*ContractavsserviceTaskCreated, error) {
	event := new(ContractavsserviceTaskCreated)
	if err := _Contractavsservice.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractavsserviceTaskRewardedIterator is returned from FilterTaskRewarded and is used to iterate over the raw logs and unpacked data for TaskRewarded events raised by the Contractavsservice contract.
type ContractavsserviceTaskRewardedIterator struct {
	Event *ContractavsserviceTaskRewarded // Event containing the contract specifics and raw log

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
func (it *ContractavsserviceTaskRewardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractavsserviceTaskRewarded)
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
		it.Event = new(ContractavsserviceTaskRewarded)
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
func (it *ContractavsserviceTaskRewardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractavsserviceTaskRewardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractavsserviceTaskRewarded represents a TaskRewarded event raised by the Contractavsservice contract.
type ContractavsserviceTaskRewarded struct {
	TaskId    *big.Int
	Issuer    common.Address
	Operators []common.Address
	Rewards   []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskRewarded is a free log retrieval operation binding the contract event 0x907649ae5ab1beeecf203ab18eeb4acc8270d8db46fbf7e7becd29b5d2d8d457.
//
// Solidity: event TaskRewarded(uint256 indexed taskId, address indexed issuer, address[] operators, uint256[] rewards)
func (_Contractavsservice *ContractavsserviceFilterer) FilterTaskRewarded(opts *bind.FilterOpts, taskId []*big.Int, issuer []common.Address) (*ContractavsserviceTaskRewardedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Contractavsservice.contract.FilterLogs(opts, "TaskRewarded", taskIdRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &ContractavsserviceTaskRewardedIterator{contract: _Contractavsservice.contract, event: "TaskRewarded", logs: logs, sub: sub}, nil
}

// WatchTaskRewarded is a free log subscription operation binding the contract event 0x907649ae5ab1beeecf203ab18eeb4acc8270d8db46fbf7e7becd29b5d2d8d457.
//
// Solidity: event TaskRewarded(uint256 indexed taskId, address indexed issuer, address[] operators, uint256[] rewards)
func (_Contractavsservice *ContractavsserviceFilterer) WatchTaskRewarded(opts *bind.WatchOpts, sink chan<- *ContractavsserviceTaskRewarded, taskId []*big.Int, issuer []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Contractavsservice.contract.WatchLogs(opts, "TaskRewarded", taskIdRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractavsserviceTaskRewarded)
				if err := _Contractavsservice.contract.UnpackLog(event, "TaskRewarded", log); err != nil {
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

// ParseTaskRewarded is a log parse operation binding the contract event 0x907649ae5ab1beeecf203ab18eeb4acc8270d8db46fbf7e7becd29b5d2d8d457.
//
// Solidity: event TaskRewarded(uint256 indexed taskId, address indexed issuer, address[] operators, uint256[] rewards)
func (_Contractavsservice *ContractavsserviceFilterer) ParseTaskRewarded(log types.Log) (*ContractavsserviceTaskRewarded, error) {
	event := new(ContractavsserviceTaskRewarded)
	if err := _Contractavsservice.contract.UnpackLog(event, "TaskRewarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
