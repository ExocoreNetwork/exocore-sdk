pragma solidity >=0.8.17;
import "./iavs.sol" as avs;
contract AvsServiceContract {
    address constant BLS_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000809;
    address constant AVS_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000901;

    struct Task {
        uint256 taskId;
        address issuer;
        string description;
        uint256 numberX;
        uint256 numberY;
        uint256 reward;
        uint256 deadline;
        bool isCompleted;
        bool isRewarded;
        bytes operatorSignatures;
        address[] operators;
    }

    address public owner;
    uint256 public taskIdCounter;
    Task[] public tasks;
    mapping(uint256 => bytes) public allTaskResponses;
    mapping(bytes => bool) public isPublicKeyRegistered;


    event OperatorRegistered(address indexed operatorAddress, string name);
    event TaskCreated(uint256 indexed taskId, address indexed issuer, string description, uint256 reward, uint256 deadline);
    event TaskCompleted(uint256 indexed taskId, address indexed issuer);

    constructor() {
        owner = msg.sender;
        taskIdCounter = 1;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only contract owner can call this function.");
        _;
    }



    function registerAVS(
        string memory avsName,
        uint64 minStakeAmount,
        address taskAddr,
        address slashAddr,
        address rewardAddr,
        string[] memory avsOwnerAddress,
        string[] memory assetIds,
        uint64 avsUnbondingPeriod,
        uint64 minSelfDelegation,
        string memory epochIdentifier,
        uint64[] memory params
    ) public returns (bool) {
        bool success =  avs.AVSMANAGER_CONTRACT.registerAVS(
            avsName,
            minStakeAmount,
            taskAddr,
            slashAddr,
            rewardAddr,
            avsOwnerAddress,
            assetIds,
            avsUnbondingPeriod,
            minSelfDelegation,
            epochIdentifier,
            params
        );
        return success;
    }


    function updateAVS(
        string memory avsName,
        uint64 minStakeAmount,
        address taskAddr,
        address slashAddr,
        address rewardAddr,
        string[] memory avsOwnerAddress,
        string[] memory assetIds,
        uint64 avsUnbondingPeriod,
        uint64 minSelfDelegation,
        string memory epochIdentifier,
        uint64[] memory params
    ) public returns (bool) {
        bool success =  avs.AVSMANAGER_CONTRACT.updateAVS(
            avsName,
            minStakeAmount,
            taskAddr,
            slashAddr,
            rewardAddr,
            avsOwnerAddress,
            assetIds,
            avsUnbondingPeriod,
            minSelfDelegation,
            epochIdentifier,
            params

        );
        return success;
    }


    function deregisterAVS(string memory avsName) public returns (bool) {
        bool success = avs.AVSMANAGER_CONTRACT.deregisterAVS(
            avsName
        );
        return success;
    }

    function registerOperatorToAVS() public returns (bool) {

        bool success = avs.AVSMANAGER_CONTRACT.registerOperatorToAVS(

        );
        return success;
    }

    function deregisterOperatorFromAVS() public returns (bool) {

        bool success = avs.AVSMANAGER_CONTRACT.deregisterOperatorFromAVS(

        );
        return success;
    }


    function submitProof(
        string memory taskId,
        string memory taskContractAddress,
        string memory aggregator,
        string memory avsAddress,
        bytes memory operatorStatus
    ) public returns (bool) {
        bool success = avs.AVSMANAGER_CONTRACT.submitProof(
            taskId,
            taskContractAddress,
            aggregator,
            avsAddress,
            operatorStatus
        );
        return success;
    }

    function registerBLSPublicKey(
        string memory operator,
        string memory name,
        bytes memory pubKey,
        bytes memory pubkeyRegistrationSignature,
        bytes memory pubkeyRegistrationMessageHash
    ) public returns (bool) {

        bool success = avs.AVSMANAGER_CONTRACT.registerBLSPublicKey(
            operator,
            name,
            pubKey,
            pubkeyRegistrationSignature,
            pubkeyRegistrationMessageHash
        );
        return success;
    }

    function getRegisteredPubkey(string memory operator) public returns (bytes memory) {


        bool success; bytes memory data = avs.AVSMANAGER_CONTRACT.getRegisteredPubkey(
        operator
    );

        require(success, "Call failed");
        return abi.decode(data, (bytes));
    }

    function getOptInOperators(address avsAddress) public returns (string[] memory) {
        bool success; string[] memory data = avs.AVSMANAGER_CONTRACT.getOptInOperators(
        avsAddress
    );


        require(success, "Call failed");
        return data;
    }

    function fastAggregateVerify(
        bytes32 msg_,
        bytes calldata signature,
        bytes[] calldata pubkeys
    ) public  returns (bool) {
        (bool success, ) = BLS_PRECOMPILE_ADDRESS.call(
            abi.encodeWithSelector(
                bytes4(keccak256("fastAggregateVerify(bytes32,bytes,bytes[])")),
                msg_,signature,pubkeys
            )
        );

        return success;

    }


    function registerBLSPublicKey(
        bytes memory publicKey
    ) public  returns (bool) {
        (bool success, ) = AVS_PRECOMPILE_ADDRESS.call(
            abi.encodeWithSelector(
                bytes4(keccak256("registerBLSPublicKey(string,bytes)")),
                msg.sender,publicKey
            )
        );
        return success;
    }

    function createTask(string memory description, uint256 numberX,uint256 numberY,uint256 reward, uint256 deadline) public onlyOwner {
        Task memory newTask;
        newTask.taskId = taskIdCounter;
        newTask.issuer = msg.sender;
        newTask.description = description;
        newTask.numberX = numberX;
        newTask.numberY = numberY;
        newTask.reward = reward;
        newTask.deadline = deadline;
        newTask.isCompleted = false;
        newTask.isRewarded = false;

        // initialize the mapping
        newTask.operators = new address[](0);

        tasks.push(newTask);
        taskIdCounter++;

        emit TaskCreated(newTask.taskId, msg.sender, description, reward, deadline);
    }
    function isPublicKey(bytes memory publicKey) public view returns (bool) {
        return isPublicKeyRegistered[publicKey];
    }

    function completeTask(uint256 taskId,bytes calldata aggSig, bytes[] calldata pubkeys) public onlyOwner {
        require(taskId > 0 && taskId <= tasks.length, "taskId is error");
        Task storage task = tasks[taskId - 1];
        require(task.taskId == taskId, "Task does not exist.");

        require(!task.isCompleted, "Task is already completed.");
        require(task.deadline >= block.timestamp, "Task deadline has passed.");
        for (uint256 i = 0; i < pubkeys.length; i++) {
            require(isPublicKeyRegistered[pubkeys[i]], "pubkey not registered.");
        }

        (bool success,) = BLS_PRECOMPILE_ADDRESS.call(
            abi.encodeWithSelector(
                bytes4(keccak256("fastAggregateVerify(bytes32,bytes,bytes[])")),
                bytes32(task.numberX + task.numberY),aggSig,pubkeys
            )
        );
        require(success, "fastAggregateVerify fail.");

        allTaskResponses[taskId] = aggSig;

        task.isCompleted = true;

        emit TaskCompleted(taskId, task.issuer);

    }




    function getTaskCount() public view returns (uint256) {
        return tasks.length;
    }

    function getTask(uint256 taskId) public view returns (Task memory) {
        return tasks[taskId - 1];
    }
}