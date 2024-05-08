// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.17;
contract AvsServiceContract {
    address constant BLS_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000809;
    address constant AVS_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000902;
    address constant TASK_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000901;

    struct Operator {
        bytes publicKey;
        string name;
        uint256 totalRewardsEarned;
        bool isRegistered;
    }

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
    uint256 public slashPercentage;
    mapping(address => Operator) public operators;
    Task[] public tasks;
    mapping(uint256 => bytes) public allTaskResponses;


    event OperatorRegistered(address indexed operatorAddress, string name);
    event TaskCreated(uint256 indexed taskId, address indexed issuer, string description, uint256 reward, uint256 deadline);
    event TaskCompleted(uint256 indexed taskId, address indexed issuer);
    event TaskRewarded(uint256 indexed taskId, address indexed issuer, address[] operators, uint256[] rewards);
    event OperatorSlashed(uint256 indexed taskId, address indexed operatorAddress, uint256 slashAmount);

    constructor(uint256 _slashPercentage) {
        owner = msg.sender;
        taskIdCounter = 1;
        slashPercentage = _slashPercentage;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only contract owner can call this function.");
        _;
    }

    modifier onlyOperator() {
        require(operators[msg.sender].isRegistered, "Only registered operators can call this function.");
        _;
    }

    function registerAVSToExocore(
        string memory avsName,
        string memory avsAddress,
        string memory operatorAddress,
        uint64 action,
        string memory avsOwnerAddress,
        string memory assetID
    ) public  returns (bool) {
        (bool success, ) = AVS_PRECOMPILE_ADDRESS.call(
            abi.encodeWithSelector(
                bytes4(keccak256("AVSAction(string,string,string,uint64,string,string)")),
                avsName,avsAddress,operatorAddress,action,avsOwnerAddress,assetID
            )
        );

        return success;

    }

    function registerOperatorToAVS(bytes memory publicKey, string memory name) public {
        require(!operators[msg.sender].isRegistered, "Operator is already registered.");

        operators[msg.sender] = Operator({
            publicKey: publicKey,
            name: name,
            totalRewardsEarned: 0,
            isRegistered: true
        });

        emit OperatorRegistered(msg.sender, name);
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

        // 初始化映射
        newTask.operators = new address[](0);

        tasks.push(newTask);
        taskIdCounter++;

        emit TaskCreated(newTask.taskId, msg.sender, description, reward, deadline);
    }

    function completeTask(uint256 taskId,bytes calldata aggSig, bytes[] calldata pubkeys) public onlyOwner {
        require(taskId > 0 && taskId <= tasks.length, "taskId is error");
        Task storage task = tasks[taskId - 1];
        require(task.taskId == taskId, "Task does not exist.");

        require(!task.isCompleted, "Task is already completed.");
        require(task.deadline >= block.timestamp, "Task deadline has passed.");

        (bool success,) = BLS_PRECOMPILE_ADDRESS.call(
            abi.encodeWithSelector(
                bytes4(keccak256("fastAggregateVerify(bytes32,bytes,bytes)")),
                bytes32(task.numberX + task.numberY),aggSig,pubkeys
            )
        );
        require(success, "fastAggregateVerify fail.");

        allTaskResponses[taskId] = aggSig;

        task.isCompleted = true;

        emit TaskCompleted(taskId, task.issuer);

    }

    function rewardTask(uint256 taskId, address[] memory operatorAddresses, uint256[] memory rewards) public onlyOwner {
        Task storage task = tasks[taskId - 1];
        require(task.isCompleted, "Task is not completed.");
        require(!task.isRewarded, "Task is already rewarded.");
        require(operatorAddresses.length == rewards.length, "Invalid input length.");

        uint256 totalReward = 0;
        for (uint256 i = 0; i < rewards.length; i++) {
            totalReward += rewards[i];
        }
        require(totalReward <= task.reward, "Total rewards exceed the task reward.");

        task.isRewarded = true;


        emit TaskRewarded(taskId, task.issuer, operatorAddresses, rewards);
    }

    function slashOperator(uint256 taskId, address operatorAddress) public onlyOwner {
        Task storage task = tasks[taskId - 1];
        require(task.isCompleted, "Task is not completed.");
        require(!task.isRewarded, "Task is already rewarded.");
        require(operators[operatorAddress].isRegistered, "Invalid operator address.");

        uint256 slashAmount = calculateSlashAmount(operatorAddress);
        require(slashAmount > 0, "Cannot slash operator for this task.");

        operators[operatorAddress].totalRewardsEarned -= slashAmount;

        emit OperatorSlashed(taskId, operatorAddress, slashAmount);
    }

    function calculateSlashAmount(address operatorAddress) internal view returns (uint256) {
        Operator storage operator = operators[operatorAddress];

        if (operator.isRegistered) {
            uint256 totalCompletedTasks = 0;
            uint256 totalRewardsEarned = operator.totalRewardsEarned;

            if (totalCompletedTasks > 0 && totalRewardsEarned > 0) {
                uint256 averageReward = totalRewardsEarned / totalCompletedTasks;
                uint256 slashAmount = (averageReward * slashPercentage) / 100;

                if (slashAmount > totalRewardsEarned) {
                    slashAmount = totalRewardsEarned;
                }

                return slashAmount;
            }
        }

        return 0;
    }

    function getOperatorCount() public view returns (uint256) {
        return tasks.length;
    }

    function getTaskCount() public view returns (uint256) {
        return tasks.length;
    }

    function getTask(uint256 taskId) public view returns (Task memory) {
        return tasks[taskId - 1];
    }
}