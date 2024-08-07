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
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"TaskCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allTaskResponses\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"aggSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"completeTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numberY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"createTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"}],\"name\":\"deregisterAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregisterOperatorFromAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msg_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"fastAggregateVerify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddress\",\"type\":\"address\"}],\"name\":\"getOptInOperators\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"operator\",\"type\":\"string\"}],\"name\":\"getRegisteredPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"getTask\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numberY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCompleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isRewarded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"operatorSignatures\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"}],\"internalType\":\"structAvsServiceContract.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTaskCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"isPublicKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"isPublicKeyRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"minStakeAmount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"avsOwnerAddress\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"assetIds\",\"type\":\"string[]\"},{\"internalType\":\"uint64\",\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minSelfDelegation\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64[]\",\"name\":\"params\",\"type\":\"uint64[]\"}],\"name\":\"registerAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"operator\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubkeyRegistrationMessageHash\",\"type\":\"bytes\"}],\"name\":\"registerBLSPublicKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"registerBLSPublicKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerOperatorToAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"taskContractAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"aggregator\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"avsAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"operatorStatus\",\"type\":\"bytes\"}],\"name\":\"submitProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskIdCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numberY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCompleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isRewarded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"operatorSignatures\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"minStakeAmount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"avsOwnerAddress\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"assetIds\",\"type\":\"string[]\"},{\"internalType\":\"uint64\",\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minSelfDelegation\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64[]\",\"name\":\"params\",\"type\":\"uint64[]\"}],\"name\":\"updateAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161253e38038061253e83398101604081905261002e91610103565b5f80546001600160a01b0319163317905560018055805161005690600590602084019061005d565b50506101c2565b828054828255905f5260205f209081019282156100b0579160200282015b828111156100b057825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061007b565b506100bc9291506100c0565b5090565b5b808211156100bc575f81556001016100c1565b634e487b7160e01b5f52604160045260245ffd5b80516001600160a01b03811681146100fe575f80fd5b919050565b5f6020808385031215610114575f80fd5b82516001600160401b038082111561012a575f80fd5b818501915085601f83011261013d575f80fd5b81518181111561014f5761014f6100d4565b8060051b604051601f19603f83011681018181108582111715610174576101746100d4565b604052918252848201925083810185019188831115610191575f80fd5b938501935b828510156101b6576101a7856100e8565b84529385019392850192610196565b98975050505050505050565b61236f806101cf5f395ff3fe608060405234801561000f575f80fd5b506004361061013d575f3560e01c80636f48e1a2116100b4578063aebe782111610079578063aebe7821146102fb578063c17a340e1461030e578063c208dd9914610316578063c49bb5211461031e578063d94c111b14610331578063de16bf4614610344575f80fd5b80636f48e1a2146102835780637ea5a4cf146102965780638d977672146102a95780638da5cb5b146102d2578063ad8ef91c146102e4575f80fd5b806333a057f51161010557806333a057f5146101e6578063358be137146102095780634bdb04f91461021c57806355b42cbe1461022f57806358b1fa9914610242578063697a1c9114610270575f80fd5b8063025e7c271461014157806319c2e665146101715780631d4c8007146101915780631d65e77e146101b15780632f95670e146101d1575b5f80fd5b61015461014f366004611470565b61034c565b6040516001600160a01b0390911681526020015b60405180910390f35b61018461017f366004611470565b610374565b60405161016891906114b5565b6101a461019f3660046114e2565b61040b565b6040516101689190611553565b6101c46101bf366004611470565b6104c6565b60405161016891906115a8565b6101e46101df366004611709565b61070c565b005b6101f96101f4366004611838565b610ad5565b6040519015158152602001610168565b6101f9610217366004611838565b610b51565b6101f961022a366004611900565b610b82565b61018461023d366004611900565b610beb565b6101f9610250366004611900565b805160208183018101805160048252928201919093012091525460ff1681565b6101f961027e366004611900565b610cb2565b6101f9610291366004611a51565b610cdc565b6101f96102a4366004611709565b610d71565b6102bc6102b7366004611470565b610e43565b6040516101689a99989796959493929190611b82565b5f54610154906001600160a01b031681565b6102ed60015481565b604051908152602001610168565b6101e4610309366004611bea565b610fc6565b6002546102ed565b6101f96111e2565b6101f961032c366004611a51565b61124d565b6101f961033f366004611900565b611291565b6101f9611359565b6005818154811061035b575f80fd5b5f918252602090912001546001600160a01b0316905081565b60036020525f90815260409020805461038c90611c41565b80601f01602080910402602001604051908101604052809291908181526020018280546103b890611c41565b80156104035780601f106103da57610100808354040283529160200191610403565b820191905f5260205f20905b8154815290600101906020018083116103e657829003601f168201915b505050505081565b604051631d4c800760e01b81526001600160a01b03821660048201526060905f90819061090190631d4c8007906024015f604051808303815f875af1158015610456573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261047d9190810190611caf565b9050816104bf5760405162461bcd60e51b815260206004820152600b60248201526a10d85b1b0819985a5b195960aa1b60448201526064015b60405180910390fd5b9392505050565b6104ce61139a565b60026104db600184611d7f565b815481106104eb576104eb611d92565b905f5260205f2090600a0201604051806101600160405290815f8201548152602001600182015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b0316815260200160028201805461054a90611c41565b80601f016020809104026020016040519081016040528092919081815260200182805461057690611c41565b80156105c15780601f10610598576101008083540402835291602001916105c1565b820191905f5260205f20905b8154815290600101906020018083116105a457829003601f168201915b505050918352505060038201546020820152600482015460408201526005820154606082015260068201546080820152600782015460ff808216151560a084015261010090910416151560c082015260088201805460e09092019161062590611c41565b80601f016020809104026020016040519081016040528092919081815260200182805461065190611c41565b801561069c5780601f106106735761010080835404028352916020019161069c565b820191905f5260205f20905b81548152906001019060200180831161067f57829003601f168201915b50505050508152602001600982018054806020026020016040519081016040528092919081815260200182805480156106fc57602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116106de575b5050505050815250509050919050565b5f546001600160a01b031633146107355760405162461bcd60e51b81526004016104b690611da6565b5f8511801561074657506002548511155b6107845760405162461bcd60e51b815260206004820152600f60248201526e3a30b9b5a4b21034b99032b93937b960891b60448201526064016104b6565b5f6002610792600188611d7f565b815481106107a2576107a2611d92565b905f5260205f2090600a0201905085815f0154146107f95760405162461bcd60e51b81526020600482015260146024820152732a30b9b5903237b2b9903737ba1032bc34b9ba1760611b60448201526064016104b6565b600781015460ff161561084e5760405162461bcd60e51b815260206004820152601a60248201527f5461736b20697320616c726561647920636f6d706c657465642e00000000000060448201526064016104b6565b42816006015410156108a25760405162461bcd60e51b815260206004820152601960248201527f5461736b20646561646c696e6520686173207061737365642e0000000000000060448201526064016104b6565b5f5b828110156109405760048484838181106108c0576108c0611d92565b90506020028101906108d29190611df1565b6040516108e0929190611e33565b9081526040519081900360200190205460ff166109385760405162461bcd60e51b8152602060048201526016602482015275383ab135b2bc903737ba103932b3b4b9ba32b932b21760511b60448201526064016104b6565b6001016108a4565b505f6108096001600160a01b03167f7ea5a4cf6ce6f6b705c2660435a26c37251a0c518b51dd1ef25344f05ece43a7836004015484600301546109839190611e42565b60405161099b91908a908a908a908a90602401611e7d565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516109d99190611f2e565b5f604051808303815f865af19150503d805f8114610a12576040519150601f19603f3d011682016040523d82523d5f602084013e610a17565b606091505b5050905080610a685760405162461bcd60e51b815260206004820152601960248201527f66617374416767726567617465566572696679206661696c2e0000000000000060448201526064016104b6565b5f878152600360205260409020610a80868883611f90565b5060078201805460ff191660019081179091558201546040516001600160a01b039091169088907fbb5889c77948badf90e8a5c73d55265e5f5d6e4837a79a78c5669691b897faed905f90a350505050505050565b6040516333a057f560e01b81525f908190610901906333a057f590610b06908a908a908a908a908a90600401612049565b6020604051808303815f875af1158015610b22573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b4691906120b5565b979650505050505050565b60405163358be13760e01b81525f9081906109019063358be13790610b06908a908a908a908a908a90600401612049565b604051634bdb04f960e01b81525f90819061090190634bdb04f990610bab9086906004016114b5565b6020604051808303815f875af1158015610bc7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104bf91906120b5565b604051632ada165f60e11b81526060905f908190610901906355b42cbe90610c179087906004016114b5565b5f604051808303815f875af1158015610c32573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610c5991908101906120d4565b905081610c965760405162461bcd60e51b815260206004820152600b60248201526a10d85b1b0819985a5b195960aa1b60448201526064016104b6565b80806020019051810190610caa91906120d4565b949350505050565b5f600482604051610cc39190611f2e565b9081526040519081900360200190205460ff1692915050565b5f806109016001600160a01b0316636f48e1a28e8e8e8e8e8e8e8e8e8e8e6040518c63ffffffff1660e01b8152600401610d209b9a99989796959493929190612150565b6020604051808303815f875af1158015610d3c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d6091906120b5565b9d9c50505050505050505050505050565b5f806108096001600160a01b03167f7ea5a4cf6ce6f6b705c2660435a26c37251a0c518b51dd1ef25344f05ece43a78888888888604051602401610db9959493929190611e7d565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051610df79190611f2e565b5f604051808303815f865af19150503d805f8114610e30576040519150601f19603f3d011682016040523d82523d5f602084013e610e35565b606091505b509098975050505050505050565b60028181548110610e52575f80fd5b5f9182526020909120600a90910201805460018201546002830180549294506001600160a01b039091169291610e8790611c41565b80601f0160208091040260200160405190810160405280929190818152602001828054610eb390611c41565b8015610efe5780601f10610ed557610100808354040283529160200191610efe565b820191905f5260205f20905b815481529060010190602001808311610ee157829003601f168201915b5050506003840154600485015460058601546006870154600788015460088901805498999598949750929550909360ff8083169461010090930416929091610f4590611c41565b80601f0160208091040260200160405190810160405280929190818152602001828054610f7190611c41565b8015610fbc5780601f10610f9357610100808354040283529160200191610fbc565b820191905f5260205f20905b815481529060010190602001808311610f9f57829003601f168201915b505050505090508a565b5f546001600160a01b03163314610fef5760405162461bcd60e51b81526004016104b690611da6565b610ff761139a565b60018054825233602080840191825260408085018a8152606086018a90526080860189905260a0860188905260c086018790525f60e0870181905261010087018190528251818152938401909252610140860192909252600280549485018155905283517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace600a90940293840190815591517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf840180546001600160a01b0319166001600160a01b039092169190911790555183927f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad001906110f9908261221b565b50606082015160038201556080820151600482015560a0820151600582015560c0820151600682015560e08201516007820180546101008086015115150261ff00199315159390931661ffff19909116179190911790556101208201516008820190611165908261221b565b5061014082015180516111829160098401916020909101906113f9565b50506001805491505f611194836122da565b909155505080516040513391907fdac01f3886a95e591b6ebef63858e5383c5f85b170876b4e18f380599b0a68a5906111d2908a90889088906122f2565b60405180910390a3505050505050565b5f806109016001600160a01b031663c208dd996040518163ffffffff1660e01b81526004016020604051808303815f875af1158015611223573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061124791906120b5565b92915050565b5f806109016001600160a01b031663c49bb5218e8e8e8e8e8e8e8e8e8e8e6040518c63ffffffff1660e01b8152600401610d209b9a99989796959493929190612150565b5f806109016001600160a01b03167fad1a2dbdf16522d18f79bda603c3c31ca300c1eb23e7b72270d1123883b3231233856040516024016112d3929190612316565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516113119190611f2e565b5f604051808303815f865af19150503d805f811461134a576040519150601f19603f3d011682016040523d82523d5f602084013e61134f565b606091505b5090949350505050565b5f806109016001600160a01b031663de16bf466040518163ffffffff1660e01b81526004016020604051808303815f875af1158015611223573d5f803e3d5ffd5b6040518061016001604052805f81526020015f6001600160a01b03168152602001606081526020015f81526020015f81526020015f81526020015f81526020015f151581526020015f1515815260200160608152602001606081525090565b828054828255905f5260205f2090810192821561144c579160200282015b8281111561144c57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611417565b5061145892915061145c565b5090565b5b80821115611458575f815560010161145d565b5f60208284031215611480575f80fd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6104bf6020830184611487565b80356001600160a01b03811681146114dd575f80fd5b919050565b5f602082840312156114f2575f80fd5b6104bf826114c7565b5f8282518085526020808601955060208260051b840101602086015f5b8481101561154657601f19868403018952611534838351611487565b98840198925090830190600101611518565b5090979650505050505050565b602081525f6104bf60208301846114fb565b5f815180845260208085019450602084015f5b8381101561159d5781516001600160a01b031687529582019590820190600101611578565b509495945050505050565b60208152815160208201525f60208301516115ce60408401826001600160a01b03169052565b5060408301516101608060608501526115eb610180850183611487565b915060608501516080850152608085015160a085015260a085015160c085015260c085015160e085015260e085015161010061162a8187018315159052565b860151905061012061163f8682018315159052565b80870151915050601f1961014081878603018188015261165f8584611487565b90880151878203909201848801529350905061167b8382611565565b9695505050505050565b5f8083601f840112611695575f80fd5b5081356001600160401b038111156116ab575f80fd5b6020830191508360208285010111156116c2575f80fd5b9250929050565b5f8083601f8401126116d9575f80fd5b5081356001600160401b038111156116ef575f80fd5b6020830191508360208260051b85010111156116c2575f80fd5b5f805f805f6060868803121561171d575f80fd5b8535945060208601356001600160401b038082111561173a575f80fd5b61174689838a01611685565b9096509450604088013591508082111561175e575f80fd5b5061176b888289016116c9565b969995985093965092949392505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156117b8576117b861177c565b604052919050565b5f6001600160401b038211156117d8576117d861177c565b50601f01601f191660200190565b5f82601f8301126117f5575f80fd5b8135611808611803826117c0565b611790565b81815284602083860101111561181c575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f60a0868803121561184c575f80fd5b85356001600160401b0380821115611862575f80fd5b61186e89838a016117e6565b96506020880135915080821115611883575f80fd5b61188f89838a016117e6565b955060408801359150808211156118a4575f80fd5b6118b089838a016117e6565b945060608801359150808211156118c5575f80fd5b6118d189838a016117e6565b935060808801359150808211156118e6575f80fd5b506118f3888289016117e6565b9150509295509295909350565b5f60208284031215611910575f80fd5b81356001600160401b03811115611925575f80fd5b610caa848285016117e6565b80356001600160401b03811681146114dd575f80fd5b5f6001600160401b0382111561195f5761195f61177c565b5060051b60200190565b5f82601f830112611978575f80fd5b8135602061198861180383611947565b82815260059290921b840181019181810190868411156119a6575f80fd5b8286015b848110156119e35780356001600160401b038111156119c7575f80fd5b6119d58986838b01016117e6565b8452509183019183016119aa565b509695505050505050565b5f82601f8301126119fd575f80fd5b81356020611a0d61180383611947565b8083825260208201915060208460051b870101935086841115611a2e575f80fd5b602086015b848110156119e357611a4481611931565b8352918301918301611a33565b5f805f805f805f805f805f6101608c8e031215611a6c575f80fd5b6001600160401b03808d351115611a81575f80fd5b611a8e8e8e358f016117e6565b9b50611a9c60208e01611931565b9a50611aaa60408e016114c7565b9950611ab860608e016114c7565b9850611ac660808e016114c7565b97508060a08e01351115611ad8575f80fd5b611ae88e60a08f01358f01611969565b96508060c08e01351115611afa575f80fd5b611b0a8e60c08f01358f01611969565b9550611b1860e08e01611931565b9450611b276101008e01611931565b9350806101208e01351115611b3a575f80fd5b611b4b8e6101208f01358f016117e6565b9250806101408e01351115611b5e575f80fd5b50611b708d6101408e01358e016119ee565b90509295989b509295989b9093969950565b8a81526001600160a01b038a166020820152610140604082018190525f90611bac8382018c611487565b90508960608401528860808401528760a08401528660c084015285151560e0840152841515610100840152828103610120840152610d608185611487565b5f805f805f60a08688031215611bfe575f80fd5b85356001600160401b03811115611c13575f80fd5b611c1f888289016117e6565b9860208801359850604088013597606081013597506080013595509350505050565b600181811c90821680611c5557607f821691505b602082108103611c7357634e487b7160e01b5f52602260045260245ffd5b50919050565b5f611c86611803846117c0565b9050828152838383011115611c99575f80fd5b8282602083015e5f602084830101529392505050565b5f6020808385031215611cc0575f80fd5b82516001600160401b0380821115611cd6575f80fd5b818501915085601f830112611ce9575f80fd5b8151611cf761180382611947565b81815260059190911b83018401908481019088831115611d15575f80fd5b8585015b83811015611d5e57805185811115611d2f575f80fd5b8601603f81018b13611d3f575f80fd5b611d508b8983015160408401611c79565b845250918601918601611d19565b5098975050505050505050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561124757611247611d6b565b634e487b7160e01b5f52603260045260245ffd5b6020808252602b908201527f4f6e6c7920636f6e7472616374206f776e65722063616e2063616c6c2074686960408201526a3990333ab731ba34b7b71760a91b606082015260800190565b5f808335601e19843603018112611e06575f80fd5b8301803591506001600160401b03821115611e1f575f80fd5b6020019150368190038213156116c2575f80fd5b818382375f9101908152919050565b8082018082111561124757611247611d6b565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b8581525f6020606081840152611e97606084018789611e55565b8381036040850152848152818101600586901b82018301875f5b88811015611f1d57848303601f190184528135368b9003601e19018112611ed6575f80fd5b8a0186810190356001600160401b03811115611ef0575f80fd5b803603821315611efe575f80fd5b611f09858284611e55565b958801959450505090850190600101611eb1565b50909b9a5050505050505050505050565b5f82518060208501845e5f920191825250919050565b601f821115611f8b57805f5260205f20601f840160051c81016020851015611f695750805b601f840160051c820191505b81811015611f88575f8155600101611f75565b50505b505050565b6001600160401b03831115611fa757611fa761177c565b611fbb83611fb58354611c41565b83611f44565b5f601f841160018114611fec575f8515611fd55750838201355b5f19600387901b1c1916600186901b178355611f88565b5f83815260208120601f198716915b8281101561201b5786850135825560209485019460019092019101611ffb565b5086821015612037575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60a081525f61205b60a0830188611487565b828103602084015261206d8188611487565b905082810360408401526120818187611487565b905082810360608401526120958186611487565b905082810360808401526120a98185611487565b98975050505050505050565b5f602082840312156120c5575f80fd5b815180151581146104bf575f80fd5b5f602082840312156120e4575f80fd5b81516001600160401b038111156120f9575f80fd5b8201601f81018413612109575f80fd5b610caa84825160208401611c79565b5f815180845260208085019450602084015f5b8381101561159d5781516001600160401b03168752958201959082019060010161212b565b5f6101608083526121638184018f611487565b6001600160401b038e1660208501526001600160a01b038d811660408601528c16606085015290506001600160a01b038a16608084015282810360a08401526121ac818a6114fb565b905082810360c08401526121c081896114fb565b6001600160401b03881660e085015290506001600160401b0386166101008401528281036101208401526121f48186611487565b90508281036101408401526122098185612118565b9e9d5050505050505050505050505050565b81516001600160401b038111156122345761223461177c565b612248816122428454611c41565b84611f44565b602080601f83116001811461227b575f84156122645750858301515b5f19600386901b1c1916600185901b1785556122d2565b5f85815260208120601f198616915b828110156122a95788860151825594840194600190910190840161228a565b50858210156122c657878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f600182016122eb576122eb611d6b565b5060010190565b606081525f6123046060830186611487565b60208301949094525060400152919050565b6001600160a01b03831681526040602082018190525f90610caa9083018461148756fea264697066735822122037499049371309827b229bfd23fe5141f831a3d40a9733a4c1b739d1528e927d64736f6c63430008190033",
}

// ContractavsserviceABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractavsserviceMetaData.ABI instead.
var ContractavsserviceABI = ContractavsserviceMetaData.ABI

// ContractavsserviceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractavsserviceMetaData.Bin instead.
var ContractavsserviceBin = ContractavsserviceMetaData.Bin

// DeployContractavsservice deploys a new Ethereum contract, binding an instance of Contractavsservice to it.
func DeployContractavsservice(auth *bind.TransactOpts, backend bind.ContractBackend, _owners []common.Address) (common.Address, *types.Transaction, *Contractavsservice, error) {
	parsed, err := ContractavsserviceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractavsserviceBin), backend, _owners)
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

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Contractavsservice *ContractavsserviceCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Contractavsservice *ContractavsserviceSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Contractavsservice.Contract.Owners(&_Contractavsservice.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Contractavsservice *ContractavsserviceCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Contractavsservice.Contract.Owners(&_Contractavsservice.CallOpts, arg0)
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

// DeregisterAVS is a paid mutator transaction binding the contract method 0x4bdb04f9.
//
// Solidity: function deregisterAVS(string avsName) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) DeregisterAVS(opts *bind.TransactOpts, avsName string) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "deregisterAVS", avsName)
}

// DeregisterAVS is a paid mutator transaction binding the contract method 0x4bdb04f9.
//
// Solidity: function deregisterAVS(string avsName) returns(bool)
func (_Contractavsservice *ContractavsserviceSession) DeregisterAVS(avsName string) (*types.Transaction, error) {
	return _Contractavsservice.Contract.DeregisterAVS(&_Contractavsservice.TransactOpts, avsName)
}

// DeregisterAVS is a paid mutator transaction binding the contract method 0x4bdb04f9.
//
// Solidity: function deregisterAVS(string avsName) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) DeregisterAVS(avsName string) (*types.Transaction, error) {
	return _Contractavsservice.Contract.DeregisterAVS(&_Contractavsservice.TransactOpts, avsName)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xde16bf46.
//
// Solidity: function deregisterOperatorFromAVS() returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "deregisterOperatorFromAVS")
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xde16bf46.
//
// Solidity: function deregisterOperatorFromAVS() returns(bool)
func (_Contractavsservice *ContractavsserviceSession) DeregisterOperatorFromAVS() (*types.Transaction, error) {
	return _Contractavsservice.Contract.DeregisterOperatorFromAVS(&_Contractavsservice.TransactOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xde16bf46.
//
// Solidity: function deregisterOperatorFromAVS() returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) DeregisterOperatorFromAVS() (*types.Transaction, error) {
	return _Contractavsservice.Contract.DeregisterOperatorFromAVS(&_Contractavsservice.TransactOpts)
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

// GetOptInOperators is a paid mutator transaction binding the contract method 0x1d4c8007.
//
// Solidity: function getOptInOperators(address avsAddress) returns(string[])
func (_Contractavsservice *ContractavsserviceTransactor) GetOptInOperators(opts *bind.TransactOpts, avsAddress common.Address) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "getOptInOperators", avsAddress)
}

// GetOptInOperators is a paid mutator transaction binding the contract method 0x1d4c8007.
//
// Solidity: function getOptInOperators(address avsAddress) returns(string[])
func (_Contractavsservice *ContractavsserviceSession) GetOptInOperators(avsAddress common.Address) (*types.Transaction, error) {
	return _Contractavsservice.Contract.GetOptInOperators(&_Contractavsservice.TransactOpts, avsAddress)
}

// GetOptInOperators is a paid mutator transaction binding the contract method 0x1d4c8007.
//
// Solidity: function getOptInOperators(address avsAddress) returns(string[])
func (_Contractavsservice *ContractavsserviceTransactorSession) GetOptInOperators(avsAddress common.Address) (*types.Transaction, error) {
	return _Contractavsservice.Contract.GetOptInOperators(&_Contractavsservice.TransactOpts, avsAddress)
}

// GetRegisteredPubkey is a paid mutator transaction binding the contract method 0x55b42cbe.
//
// Solidity: function getRegisteredPubkey(string operator) returns(bytes)
func (_Contractavsservice *ContractavsserviceTransactor) GetRegisteredPubkey(opts *bind.TransactOpts, operator string) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "getRegisteredPubkey", operator)
}

// GetRegisteredPubkey is a paid mutator transaction binding the contract method 0x55b42cbe.
//
// Solidity: function getRegisteredPubkey(string operator) returns(bytes)
func (_Contractavsservice *ContractavsserviceSession) GetRegisteredPubkey(operator string) (*types.Transaction, error) {
	return _Contractavsservice.Contract.GetRegisteredPubkey(&_Contractavsservice.TransactOpts, operator)
}

// GetRegisteredPubkey is a paid mutator transaction binding the contract method 0x55b42cbe.
//
// Solidity: function getRegisteredPubkey(string operator) returns(bytes)
func (_Contractavsservice *ContractavsserviceTransactorSession) GetRegisteredPubkey(operator string) (*types.Transaction, error) {
	return _Contractavsservice.Contract.GetRegisteredPubkey(&_Contractavsservice.TransactOpts, operator)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0xc49bb521.
//
// Solidity: function registerAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) RegisterAVS(opts *bind.TransactOpts, avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "registerAVS", avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0xc49bb521.
//
// Solidity: function registerAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_Contractavsservice *ContractavsserviceSession) RegisterAVS(avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterAVS(&_Contractavsservice.TransactOpts, avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0xc49bb521.
//
// Solidity: function registerAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) RegisterAVS(avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterAVS(&_Contractavsservice.TransactOpts, avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0x358be137.
//
// Solidity: function registerBLSPublicKey(string operator, string name, bytes pubKey, bytes pubkeyRegistrationSignature, bytes pubkeyRegistrationMessageHash) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) RegisterBLSPublicKey(opts *bind.TransactOpts, operator string, name string, pubKey []byte, pubkeyRegistrationSignature []byte, pubkeyRegistrationMessageHash []byte) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "registerBLSPublicKey", operator, name, pubKey, pubkeyRegistrationSignature, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0x358be137.
//
// Solidity: function registerBLSPublicKey(string operator, string name, bytes pubKey, bytes pubkeyRegistrationSignature, bytes pubkeyRegistrationMessageHash) returns(bool)
func (_Contractavsservice *ContractavsserviceSession) RegisterBLSPublicKey(operator string, name string, pubKey []byte, pubkeyRegistrationSignature []byte, pubkeyRegistrationMessageHash []byte) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterBLSPublicKey(&_Contractavsservice.TransactOpts, operator, name, pubKey, pubkeyRegistrationSignature, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0x358be137.
//
// Solidity: function registerBLSPublicKey(string operator, string name, bytes pubKey, bytes pubkeyRegistrationSignature, bytes pubkeyRegistrationMessageHash) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) RegisterBLSPublicKey(operator string, name string, pubKey []byte, pubkeyRegistrationSignature []byte, pubkeyRegistrationMessageHash []byte) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterBLSPublicKey(&_Contractavsservice.TransactOpts, operator, name, pubKey, pubkeyRegistrationSignature, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey0 is a paid mutator transaction binding the contract method 0xd94c111b.
//
// Solidity: function registerBLSPublicKey(bytes publicKey) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) RegisterBLSPublicKey0(opts *bind.TransactOpts, publicKey []byte) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "registerBLSPublicKey0", publicKey)
}

// RegisterBLSPublicKey0 is a paid mutator transaction binding the contract method 0xd94c111b.
//
// Solidity: function registerBLSPublicKey(bytes publicKey) returns(bool)
func (_Contractavsservice *ContractavsserviceSession) RegisterBLSPublicKey0(publicKey []byte) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterBLSPublicKey0(&_Contractavsservice.TransactOpts, publicKey)
}

// RegisterBLSPublicKey0 is a paid mutator transaction binding the contract method 0xd94c111b.
//
// Solidity: function registerBLSPublicKey(bytes publicKey) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) RegisterBLSPublicKey0(publicKey []byte) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterBLSPublicKey0(&_Contractavsservice.TransactOpts, publicKey)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0xc208dd99.
//
// Solidity: function registerOperatorToAVS() returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "registerOperatorToAVS")
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0xc208dd99.
//
// Solidity: function registerOperatorToAVS() returns(bool)
func (_Contractavsservice *ContractavsserviceSession) RegisterOperatorToAVS() (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterOperatorToAVS(&_Contractavsservice.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0xc208dd99.
//
// Solidity: function registerOperatorToAVS() returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) RegisterOperatorToAVS() (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterOperatorToAVS(&_Contractavsservice.TransactOpts)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x33a057f5.
//
// Solidity: function submitProof(string taskId, string taskContractAddress, string aggregator, string avsAddress, bytes operatorStatus) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) SubmitProof(opts *bind.TransactOpts, taskId string, taskContractAddress string, aggregator string, avsAddress string, operatorStatus []byte) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "submitProof", taskId, taskContractAddress, aggregator, avsAddress, operatorStatus)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x33a057f5.
//
// Solidity: function submitProof(string taskId, string taskContractAddress, string aggregator, string avsAddress, bytes operatorStatus) returns(bool)
func (_Contractavsservice *ContractavsserviceSession) SubmitProof(taskId string, taskContractAddress string, aggregator string, avsAddress string, operatorStatus []byte) (*types.Transaction, error) {
	return _Contractavsservice.Contract.SubmitProof(&_Contractavsservice.TransactOpts, taskId, taskContractAddress, aggregator, avsAddress, operatorStatus)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x33a057f5.
//
// Solidity: function submitProof(string taskId, string taskContractAddress, string aggregator, string avsAddress, bytes operatorStatus) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) SubmitProof(taskId string, taskContractAddress string, aggregator string, avsAddress string, operatorStatus []byte) (*types.Transaction, error) {
	return _Contractavsservice.Contract.SubmitProof(&_Contractavsservice.TransactOpts, taskId, taskContractAddress, aggregator, avsAddress, operatorStatus)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x6f48e1a2.
//
// Solidity: function updateAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) UpdateAVS(opts *bind.TransactOpts, avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "updateAVS", avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x6f48e1a2.
//
// Solidity: function updateAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_Contractavsservice *ContractavsserviceSession) UpdateAVS(avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _Contractavsservice.Contract.UpdateAVS(&_Contractavsservice.TransactOpts, avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x6f48e1a2.
//
// Solidity: function updateAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) UpdateAVS(avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _Contractavsservice.Contract.UpdateAVS(&_Contractavsservice.TransactOpts, avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
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
