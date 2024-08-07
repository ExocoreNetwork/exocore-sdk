pragma solidity ^0.8.0;

contract AVSTask {
    enum TaskStatus {Created, Accepted, Completed, Verified}

    struct Task {
        uint256 taskId;
        address creator;
        string requirements;
        string acceptanceCriteria;
        TaskStatus status;
        address operator;
        string result;
    }

    Task[] public tasks;

    event TaskCreated(uint256 indexed taskId, address indexed creator);
    event TaskAccepted(uint256 indexed taskId, address indexed operator);
    event TaskCompleted(uint256 indexed taskId, address indexed operator);
    event TaskVerified(uint256 indexed taskId, address indexed verifier);

    modifier onlyCreator(uint256 _taskId) {
        require(tasks[_taskId].creator == msg.sender, "Only task creator can call this function.");
        _;
    }

    modifier onlyOperator(uint256 _taskId) {
        require(tasks[_taskId].operator == msg.sender, "Only task operator can call this function.");
        _;
    }

    modifier onlyVerifier(uint256 _taskId) {
        require(tasks[_taskId].status == TaskStatus.Completed, "Task must be completed for verification.");
        require(tasks[_taskId].creator != msg.sender, "Task creator cannot be the verifier.");
        _;
    }

    function createTask(string memory _requirements, string memory _acceptanceCriteria) external {
        uint256 taskId = tasks.length;
        Task memory newTask = Task(taskId, msg.sender, _requirements, _acceptanceCriteria, TaskStatus.Created, address(0), "");
        tasks.push(newTask);
        emit TaskCreated(taskId, msg.sender);
    }

    function acceptTask(uint256 _taskId) external {
        require(tasks[_taskId].status == TaskStatus.Created, "Task must be in the created status.");
        tasks[_taskId].status = TaskStatus.Accepted;
        tasks[_taskId].operator = msg.sender;
        emit TaskAccepted(_taskId, msg.sender);
    }

    function completeTask(uint256 _taskId, string memory _result) external onlyOperator(_taskId) {
        require(tasks[_taskId].status == TaskStatus.Accepted, "Task must be in the accepted status.");
        tasks[_taskId].status = TaskStatus.Completed;
        tasks[_taskId].result = _result;
        emit TaskCompleted(_taskId, msg.sender);
    }

    function verifyTask(uint256 _taskId) external onlyVerifier(_taskId) {
        tasks[_taskId].status = TaskStatus.Verified;
        emit TaskVerified(_taskId, msg.sender);
    }
}