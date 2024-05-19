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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashPercentage\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"OperatorSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"TaskCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"rewards\",\"type\":\"uint256[]\"}],\"name\":\"TaskRewarded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allTaskResponses\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"aggSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"completeTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numberY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"createTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msg_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"fastAggregateVerify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"getTask\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numberY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCompleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isRewarded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"operatorSignatures\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"}],\"internalType\":\"structAvsServiceContract.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTaskCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"isPublicKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"isPublicKeyRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardsEarned\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"avsAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"operatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"action\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"avsOwnerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"assetID\",\"type\":\"string\"}],\"name\":\"registerAVSToExocore\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"registerOperatorToAVS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"operatorAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rewards\",\"type\":\"uint256[]\"}],\"name\":\"rewardTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"name\":\"slashOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskIdCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numberY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCompleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isRewarded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"operatorSignatures\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50604051612426380380612426833981016040819052602b916047565b5f80546001600160a01b0319163317905560018055600255605d565b5f602082840312156056575f80fd5b5051919050565b6123bc8061006a5f395ff3fe608060405234801561000f575f80fd5b5060043610610111575f3560e01c80637ea5a4cf1161009e578063a24d3c111161006e578063a24d3c1114610298578063ad8ef91c146102ab578063aebe7821146102b4578063bd58199a146102c7578063c17a340e14610161575f80fd5b80637ea5a4cf1461021f57806386cfceea146102325780638d977672146102455780638da5cb5b1461026e575f80fd5b80632f95670e116100e45780632f95670e146101935780633a3e8907146101a857806358b1fa99146101bb578063648b14c5146101f9578063697a1c911461020c575f80fd5b806313e7c9d81461011557806319c2e665146101415780631a4a7c61146101615780631d65e77e14610173575b5f80fd5b610128610123366004611731565b6102d0565b604051610138949392919061177f565b60405180910390f35b61015461014f3660046117bc565b610407565b60405161013891906117d3565b6004545b604051908152602001610138565b6101866101813660046117bc565b61049e565b6040516101389190611828565b6101a66101a136600461197f565b6106e4565b005b6101a66101b63660046119f2565b610ab6565b6101e96101c9366004611acb565b805160208183018101805160068252928201919093012091525460ff1681565b6040519015158152602001610138565b6101a6610207366004611b92565b610cf9565b6101e961021a366004611acb565b610f41565b6101e961022d36600461197f565b610f6b565b6101a6610240366004611c55565b61103d565b6102586102533660046117bc565b611186565b6040516101389a99989796959493929190611cb4565b5f54610280906001600160a01b031681565b6040516001600160a01b039091168152602001610138565b6101e96102a6366004611d43565b611309565b61016560015481565b6101a66102c2366004611e1a565b6113de565b61016560025481565b60036020525f90815260409020805481906102ea90611e71565b80601f016020809104026020016040519081016040528092919081815260200182805461031690611e71565b80156103615780601f1061033857610100808354040283529160200191610361565b820191905f5260205f20905b81548152906001019060200180831161034457829003601f168201915b50505050509080600101805461037690611e71565b80601f01602080910402602001604051908101604052809291908181526020018280546103a290611e71565b80156103ed5780601f106103c4576101008083540402835291602001916103ed565b820191905f5260205f20905b8154815290600101906020018083116103d057829003601f168201915b50505050600283015460039093015491929160ff16905084565b60056020525f90815260409020805461041f90611e71565b80601f016020809104026020016040519081016040528092919081815260200182805461044b90611e71565b80156104965780601f1061046d57610100808354040283529160200191610496565b820191905f5260205f20905b81548152906001019060200180831161047957829003601f168201915b505050505081565b6104a6611640565b60046104b3600184611ebd565b815481106104c3576104c3611ed6565b905f5260205f2090600a0201604051806101600160405290815f8201548152602001600182015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b0316815260200160028201805461052290611e71565b80601f016020809104026020016040519081016040528092919081815260200182805461054e90611e71565b80156105995780601f1061057057610100808354040283529160200191610599565b820191905f5260205f20905b81548152906001019060200180831161057c57829003601f168201915b505050918352505060038201546020820152600482015460408201526005820154606082015260068201546080820152600782015460ff808216151560a084015261010090910416151560c082015260088201805460e0909201916105fd90611e71565b80601f016020809104026020016040519081016040528092919081815260200182805461062990611e71565b80156106745780601f1061064b57610100808354040283529160200191610674565b820191905f5260205f20905b81548152906001019060200180831161065757829003601f168201915b50505050508152602001600982018054806020026020016040519081016040528092919081815260200182805480156106d457602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116106b6575b5050505050815250509050919050565b5f546001600160a01b031633146107165760405162461bcd60e51b815260040161070d90611eea565b60405180910390fd5b5f8511801561072757506004548511155b6107655760405162461bcd60e51b815260206004820152600f60248201526e3a30b9b5a4b21034b99032b93937b960891b604482015260640161070d565b5f6004610773600188611ebd565b8154811061078357610783611ed6565b905f5260205f2090600a0201905085815f0154146107da5760405162461bcd60e51b81526020600482015260146024820152732a30b9b5903237b2b9903737ba1032bc34b9ba1760611b604482015260640161070d565b600781015460ff161561082f5760405162461bcd60e51b815260206004820152601a60248201527f5461736b20697320616c726561647920636f6d706c657465642e000000000000604482015260640161070d565b42816006015410156108835760405162461bcd60e51b815260206004820152601960248201527f5461736b20646561646c696e6520686173207061737365642e00000000000000604482015260640161070d565b5f5b828110156109215760068484838181106108a1576108a1611ed6565b90506020028101906108b39190611f35565b6040516108c1929190611f77565b9081526040519081900360200190205460ff166109195760405162461bcd60e51b8152602060048201526016602482015275383ab135b2bc903737ba103932b3b4b9ba32b932b21760511b604482015260640161070d565b600101610885565b505f6108096001600160a01b03167f7ea5a4cf6ce6f6b705c2660435a26c37251a0c518b51dd1ef25344f05ece43a7836004015484600301546109649190611f86565b60405161097c91908a908a908a908a90602401611fc1565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516109ba9190612072565b5f604051808303815f865af19150503d805f81146109f3576040519150601f19603f3d011682016040523d82523d5f602084013e6109f8565b606091505b5050905080610a495760405162461bcd60e51b815260206004820152601960248201527f66617374416767726567617465566572696679206661696c2e00000000000000604482015260640161070d565b5f878152600560205260409020610a618688836120d4565b5060078201805460ff191660019081179091558201546040516001600160a01b039091169088907fbb5889c77948badf90e8a5c73d55265e5f5d6e4837a79a78c5669691b897faed905f90a350505050505050565b5f546001600160a01b03163314610adf5760405162461bcd60e51b815260040161070d90611eea565b5f6004610aed600185611ebd565b81548110610afd57610afd611ed6565b5f9182526020909120600a90910201600781015490915060ff16610b5c5760405162461bcd60e51b81526020600482015260166024820152752a30b9b59034b9903737ba1031b7b6b83632ba32b21760511b604482015260640161070d565b6007810154610100900460ff1615610bb25760405162461bcd60e51b81526020600482015260196024820152782a30b9b59034b99030b63932b0b23c903932bbb0b93232b21760391b604482015260640161070d565b6001600160a01b0382165f908152600360208190526040909120015460ff16610c1d5760405162461bcd60e51b815260206004820152601960248201527f496e76616c6964206f70657261746f7220616464726573732e00000000000000604482015260640161070d565b5f610c27836115fa565b90505f8111610c845760405162461bcd60e51b8152602060048201526024808201527f43616e6e6f7420736c617368206f70657261746f7220666f722074686973207460448201526330b9b59760e11b606482015260840161070d565b6001600160a01b0383165f9081526003602052604081206002018054839290610cae908490611ebd565b90915550506040518181526001600160a01b0384169085907fa2e524bd0f71903485fbb3d6d50cb305f61005ceea2047c3ac92aa7e0d1043069060200160405180910390a350505050565b5f546001600160a01b03163314610d225760405162461bcd60e51b815260040161070d90611eea565b5f6004610d30600186611ebd565b81548110610d4057610d40611ed6565b5f9182526020909120600a90910201600781015490915060ff16610d9f5760405162461bcd60e51b81526020600482015260166024820152752a30b9b59034b9903737ba1031b7b6b83632ba32b21760511b604482015260640161070d565b6007810154610100900460ff1615610df55760405162461bcd60e51b81526020600482015260196024820152782a30b9b59034b99030b63932b0b23c903932bbb0b93232b21760391b604482015260640161070d565b8151835114610e3e5760405162461bcd60e51b815260206004820152601560248201527424b73b30b634b21034b7383aba103632b733ba341760591b604482015260640161070d565b5f805b8351811015610e7957838181518110610e5c57610e5c611ed6565b602002602001015182610e6f9190611f86565b9150600101610e41565b508160050154811115610edc5760405162461bcd60e51b815260206004820152602560248201527f546f74616c20726577617264732065786365656420746865207461736b2072656044820152643bb0b9321760d91b606482015260840161070d565b60078201805461ff00191661010017905560018201546040516001600160a01b039091169086907f907649ae5ab1beeecf203ab18eeb4acc8270d8db46fbf7e7becd29b5d2d8d45790610f32908890889061218d565b60405180910390a35050505050565b5f600682604051610f529190612072565b9081526040519081900360200190205460ff1692915050565b5f806108096001600160a01b03167f7ea5a4cf6ce6f6b705c2660435a26c37251a0c518b51dd1ef25344f05ece43a78888888888604051602401610fb3959493929190611fc1565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051610ff19190612072565b5f604051808303815f865af19150503d805f811461102a576040519150601f19603f3d011682016040523d82523d5f602084013e61102f565b606091505b509098975050505050505050565b335f908152600360208190526040909120015460ff16156110a05760405162461bcd60e51b815260206004820152601f60248201527f4f70657261746f7220697320616c726561647920726567697374657265642e00604482015260640161070d565b6040805160808101825283815260208082018490525f82840181905260016060840152338152600390915291909120815181906110dd908261220f565b50602082015160018201906110f2908261220f565b5060408281015160028301556060909201516003909101805460ff19169115159190911790555160019060069061112a908590612072565b908152604051908190036020018120805492151560ff199093169290921790915533907f11a85ea4a40584362c3d9c17685709a2e02b466ac78d5eb00b6aff73d90f58059061117a9084906117d3565b60405180910390a25050565b60048181548110611195575f80fd5b5f9182526020909120600a90910201805460018201546002830180549294506001600160a01b0390911692916111ca90611e71565b80601f01602080910402602001604051908101604052809291908181526020018280546111f690611e71565b80156112415780601f1061121857610100808354040283529160200191611241565b820191905f5260205f20905b81548152906001019060200180831161122457829003601f168201915b5050506003840154600485015460058601546006870154600788015460088901805498999598949750929550909360ff808316946101009093041692909161128890611e71565b80601f01602080910402602001604051908101604052809291908181526020018280546112b490611e71565b80156112ff5780601f106112d6576101008083540402835291602001916112ff565b820191905f5260205f20905b8154815290600101906020018083116112e257829003601f168201915b505050505090508a565b5f806109026001600160a01b03167fb6c1c1ca0705e625b8d2de3b8a502fa0f8743a5761c7e20f734657df647a4226898989898989604051602401611353969594939291906122ce565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516113919190612072565b5f604051808303815f865af19150503d805f81146113ca576040519150601f19603f3d011682016040523d82523d5f602084013e6113cf565b606091505b50909998505050505050505050565b5f546001600160a01b031633146114075760405162461bcd60e51b815260040161070d90611eea565b61140f611640565b60018054825233602080840191825260408085018a8152606086018a90526080860189905260a0860188905260c086018790525f60e0870181905261010087018190528251818152938401909252610140860192909252600480549485018155905283517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b600a90940293840190815591517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19c840180546001600160a01b0319166001600160a01b039092169190911790555183927f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19d0190611511908261220f565b50606082015160038201556080820151600482015560a0820151600582015560c0820151600682015560e08201516007820180546101008086015115150261ff00199315159390931661ffff1990911617919091179055610120820151600882019061157d908261220f565b50610140820151805161159a91600984019160209091019061169f565b50506001805491505f6115ac8361234a565b909155505080516040513391907fdac01f3886a95e591b6ebef63858e5383c5f85b170876b4e18f380599b0a68a5906115ea908a9088908890612362565b60405180910390a3505050505050565b6001600160a01b0381165f90815260036020819052604082209081015460ff16156116385760028101545f90611635565b9695505050505050565b50505b505f92915050565b6040518061016001604052805f81526020015f6001600160a01b03168152602001606081526020015f81526020015f81526020015f81526020015f81526020015f151581526020015f1515815260200160608152602001606081525090565b828054828255905f5260205f209081019282156116f2579160200282015b828111156116f257825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906116bd565b506116fe929150611702565b5090565b5b808211156116fe575f8155600101611703565b80356001600160a01b038116811461172c575f80fd5b919050565b5f60208284031215611741575f80fd5b61174a82611716565b9392505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b608081525f6117916080830187611751565b82810360208401526117a38187611751565b6040840195909552505090151560609091015292915050565b5f602082840312156117cc575f80fd5b5035919050565b602081525f61174a6020830184611751565b5f815180845260208085019450602084015f5b8381101561181d5781516001600160a01b0316875295820195908201906001016117f8565b509495945050505050565b60208152815160208201525f602083015161184e60408401826001600160a01b03169052565b50604083015161016080606085015261186b610180850183611751565b915060608501516080850152608085015160a085015260a085015160c085015260c085015160e085015260e08501516101006118aa8187018315159052565b86015190506101206118bf8682018315159052565b80870151915050601f196101408187860301818801526118df8584611751565b90880151878203909201848801529350905061162b83826117e5565b5f8083601f84011261190b575f80fd5b5081356001600160401b03811115611921575f80fd5b602083019150836020828501011115611938575f80fd5b9250929050565b5f8083601f84011261194f575f80fd5b5081356001600160401b03811115611965575f80fd5b6020830191508360208260051b8501011115611938575f80fd5b5f805f805f60608688031215611993575f80fd5b8535945060208601356001600160401b03808211156119b0575f80fd5b6119bc89838a016118fb565b909650945060408801359150808211156119d4575f80fd5b506119e18882890161193f565b969995985093965092949392505050565b5f8060408385031215611a03575f80fd5b82359150611a1360208401611716565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715611a5857611a58611a1c565b604052919050565b5f82601f830112611a6f575f80fd5b81356001600160401b03811115611a8857611a88611a1c565b611a9b601f8201601f1916602001611a30565b818152846020838601011115611aaf575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60208284031215611adb575f80fd5b81356001600160401b03811115611af0575f80fd5b611afc84828501611a60565b949350505050565b5f6001600160401b03821115611b1c57611b1c611a1c565b5060051b60200190565b5f82601f830112611b35575f80fd5b81356020611b4a611b4583611b04565b611a30565b8083825260208201915060208460051b870101935086841115611b6b575f80fd5b602086015b84811015611b875780358352918301918301611b70565b509695505050505050565b5f805f60608486031215611ba4575f80fd5b833592506020808501356001600160401b0380821115611bc2575f80fd5b818701915087601f830112611bd5575f80fd5b8135611be3611b4582611b04565b81815260059190911b8301840190848101908a831115611c01575f80fd5b938501935b82851015611c2657611c1785611716565b82529385019390850190611c06565b965050506040870135925080831115611c3d575f80fd5b5050611c4b86828701611b26565b9150509250925092565b5f8060408385031215611c66575f80fd5b82356001600160401b0380821115611c7c575f80fd5b611c8886838701611a60565b93506020850135915080821115611c9d575f80fd5b50611caa85828601611a60565b9150509250929050565b8a81526001600160a01b038a166020820152610140604082018190525f90611cde8382018c611751565b90508960608401528860808401528760a08401528660c084015285151560e0840152841515610100840152828103610120840152611d1c8185611751565b9d9c50505050505050505050505050565b80356001600160401b038116811461172c575f80fd5b5f805f805f8060c08789031215611d58575f80fd5b86356001600160401b0380821115611d6e575f80fd5b611d7a8a838b01611a60565b97506020890135915080821115611d8f575f80fd5b611d9b8a838b01611a60565b96506040890135915080821115611db0575f80fd5b611dbc8a838b01611a60565b9550611dca60608a01611d2d565b94506080890135915080821115611ddf575f80fd5b611deb8a838b01611a60565b935060a0890135915080821115611e00575f80fd5b50611e0d89828a01611a60565b9150509295509295509295565b5f805f805f60a08688031215611e2e575f80fd5b85356001600160401b03811115611e43575f80fd5b611e4f88828901611a60565b9860208801359850604088013597606081013597506080013595509350505050565b600181811c90821680611e8557607f821691505b602082108103611ea357634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b81810381811115611ed057611ed0611ea9565b92915050565b634e487b7160e01b5f52603260045260245ffd5b6020808252602b908201527f4f6e6c7920636f6e7472616374206f776e65722063616e2063616c6c2074686960408201526a3990333ab731ba34b7b71760a91b606082015260800190565b5f808335601e19843603018112611f4a575f80fd5b8301803591506001600160401b03821115611f63575f80fd5b602001915036819003821315611938575f80fd5b818382375f9101908152919050565b80820180821115611ed057611ed0611ea9565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b8581525f6020606081840152611fdb606084018789611f99565b8381036040850152848152818101600586901b82018301875f5b8881101561206157848303601f190184528135368b9003601e1901811261201a575f80fd5b8a0186810190356001600160401b03811115612034575f80fd5b803603821315612042575f80fd5b61204d858284611f99565b958801959450505090850190600101611ff5565b50909b9a5050505050505050505050565b5f82518060208501845e5f920191825250919050565b601f8211156120cf57805f5260205f20601f840160051c810160208510156120ad5750805b601f840160051c820191505b818110156120cc575f81556001016120b9565b50505b505050565b6001600160401b038311156120eb576120eb611a1c565b6120ff836120f98354611e71565b83612088565b5f601f841160018114612130575f85156121195750838201355b5f19600387901b1c1916600186901b1783556120cc565b5f83815260208120601f198716915b8281101561215f578685013582556020948501946001909201910161213f565b508682101561217b575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b604080825283519082018190525f906020906060840190828701845b828110156121ce5781516001600160a01b0316845292840192908401906001016121a9565b505050838103828501528451808252858301918301905f5b81811015612202578351835292840192918401916001016121e6565b5090979650505050505050565b81516001600160401b0381111561222857612228611a1c565b61223c816122368454611e71565b84612088565b602080601f83116001811461226f575f84156122585750858301515b5f19600386901b1c1916600185901b1785556122c6565b5f85815260208120601f198616915b8281101561229d5788860151825594840194600190910190840161227e565b50858210156122ba57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b60c081525f6122e060c0830189611751565b82810360208401526122f28189611751565b905082810360408401526123068188611751565b90506001600160401b038616606084015282810360808401526123298186611751565b905082810360a084015261233d8185611751565b9998505050505050505050565b5f6001820161235b5761235b611ea9565b5060010190565b606081525f6123746060830186611751565b6020830194909452506040015291905056fea2646970667358221220e03a135f6a65f01376cbb6d13c03a888290a883800bae08a570898d8ad1fc68f64736f6c63430008190033",
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

// IsPublicKey is a free data retrieval call binding the contract method 0x697a1c91.
//
// Solidity: function isPublicKey(bytes publicKey) view returns(bool)
func (_Contractavsservice *ContractavsserviceCaller) IsPublicKey(opts *bind.CallOpts, publicKey []byte) (bool, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "isPublicKey", publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPublicKey is a free data retrieval call binding the contract method 0x697a1c91.
//
// Solidity: function isPublicKey(bytes publicKey) view returns(bool)
func (_Contractavsservice *ContractavsserviceSession) IsPublicKey(publicKey []byte) (bool, error) {
	return _Contractavsservice.Contract.IsPublicKey(&_Contractavsservice.CallOpts, publicKey)
}

// IsPublicKey is a free data retrieval call binding the contract method 0x697a1c91.
//
// Solidity: function isPublicKey(bytes publicKey) view returns(bool)
func (_Contractavsservice *ContractavsserviceCallerSession) IsPublicKey(publicKey []byte) (bool, error) {
	return _Contractavsservice.Contract.IsPublicKey(&_Contractavsservice.CallOpts, publicKey)
}

// IsPublicKeyRegistered is a free data retrieval call binding the contract method 0x58b1fa99.
//
// Solidity: function isPublicKeyRegistered(bytes ) view returns(bool)
func (_Contractavsservice *ContractavsserviceCaller) IsPublicKeyRegistered(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "isPublicKeyRegistered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPublicKeyRegistered is a free data retrieval call binding the contract method 0x58b1fa99.
//
// Solidity: function isPublicKeyRegistered(bytes ) view returns(bool)
func (_Contractavsservice *ContractavsserviceSession) IsPublicKeyRegistered(arg0 []byte) (bool, error) {
	return _Contractavsservice.Contract.IsPublicKeyRegistered(&_Contractavsservice.CallOpts, arg0)
}

// IsPublicKeyRegistered is a free data retrieval call binding the contract method 0x58b1fa99.
//
// Solidity: function isPublicKeyRegistered(bytes ) view returns(bool)
func (_Contractavsservice *ContractavsserviceCallerSession) IsPublicKeyRegistered(arg0 []byte) (bool, error) {
	return _Contractavsservice.Contract.IsPublicKeyRegistered(&_Contractavsservice.CallOpts, arg0)
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

// FastAggregateVerify is a paid mutator transaction binding the contract method 0x7ea5a4cf.
//
// Solidity: function fastAggregateVerify(bytes32 msg_, bytes signature, bytes[] pubkeys) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) FastAggregateVerify(opts *bind.TransactOpts, msg_ [32]byte, signature []byte, pubkeys [][]byte) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "fastAggregateVerify", msg_, signature, pubkeys)
}

// FastAggregateVerify is a paid mutator transaction binding the contract method 0x7ea5a4cf.
//
// Solidity: function fastAggregateVerify(bytes32 msg_, bytes signature, bytes[] pubkeys) returns(bool)
func (_Contractavsservice *ContractavsserviceSession) FastAggregateVerify(msg_ [32]byte, signature []byte, pubkeys [][]byte) (*types.Transaction, error) {
	return _Contractavsservice.Contract.FastAggregateVerify(&_Contractavsservice.TransactOpts, msg_, signature, pubkeys)
}

// FastAggregateVerify is a paid mutator transaction binding the contract method 0x7ea5a4cf.
//
// Solidity: function fastAggregateVerify(bytes32 msg_, bytes signature, bytes[] pubkeys) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) FastAggregateVerify(msg_ [32]byte, signature []byte, pubkeys [][]byte) (*types.Transaction, error) {
	return _Contractavsservice.Contract.FastAggregateVerify(&_Contractavsservice.TransactOpts, msg_, signature, pubkeys)
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
