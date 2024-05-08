// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./libraries/BytesLib.sol";
import "./TaskManager.sol";
import "./ServiceManagerBase.sol";

contract ServiceManager is ServiceManagerBase {
    TaskManager public immutable taskManager;

    /// @notice when applied to a function, ensures that the function is only callable by the `registryCoordinator`.
    modifier onlyTaskManager() {
        require(
            msg.sender == address(taskManager),
            "onlyExocoreTaskManager: not from exocore task manager"
        );
        _;
    }

    constructor(
        IDelegationManager _delegationManager,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry,
        TaskManager _taskManager
    )
        ServiceManagerBase(
            _delegationManager,
            _registryCoordinator,
            _stakeRegistry
        )
    {
        taskManager = _taskManager;
    }

    /// @notice Called in the event of challenge resolution, in order to forward a call to the Slasher, which 'freezes' the `operator`.
    /// @dev The Slasher contract is under active development and its interface expected to change.
    ///      We recommend writing slashing logic without integrating with the Slasher at this point in time.
    function freezeOperator(address operatorAddr) external onlyTaskManager {
        // slasher.freezeOperator(operatorAddr);
    }
}
