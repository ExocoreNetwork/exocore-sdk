// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.17;

contract Storage {

    uint256 number;
    address constant DEPOSIT_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000804;
    bytes4 constant DEPOSIT_FUNCTION_SELECTOR = bytes4(keccak256("depositTo(uint16,bytes,bytes,uint256)"));


    function requestDeposit(uint16 srcChainId, bytes memory token, bytes memory depositor,uint256 amount) public  returns (bytes32) {
        (bool success, bytes memory responseOrReason) = DEPOSIT_PRECOMPILE_ADDRESS.call(
            abi.encodeWithSelector(
                DEPOSIT_FUNCTION_SELECTOR,
                srcChainId,
                token,
                depositor,
                amount
            )
        );

        uint256 lastlyUpdatedPrincipleBalance;
        if (success) {
            (, lastlyUpdatedPrincipleBalance) = abi.decode(responseOrReason, (bool, uint256));
        }
        return bytes32(responseOrReason);
    }

    /**
 * @notice Forwards a call to EigenLayer's DelegationManager contract to confirm operator registration with the AVS
     * @param operator The address of the operator to register.
     * @param operatorSignature The signature, salt, and expiry of the operator's signature.
     */
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external;

    /**
     * @notice Forwards a call to EigenLayer's DelegationManager contract to confirm operator deregistration from the AVS
     * @param operator The address of the operator to deregister.
     */
    function deregisterOperatorFromAVS(address operator) external;

    function checkSignatures(
        bytes32 msgHash,
        bytes calldata quorumNumbers,
        uint32 referenceBlockNumber,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    )
    external
    view
    returns (
        QuorumStakeTotals memory,
        bytes32
    );
}
