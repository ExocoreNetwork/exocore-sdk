pragma solidity >=0.8.17;

contract AvsServiceManager {

    address constant DEPOSIT_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000804;
    address constant AVSTASK_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000901;

    bytes4 constant DEPOSIT_FUNCTION_SELECTOR = bytes4(keccak256("depositTo(uint16,bytes,bytes,uint256)"));
    bytes4 constant REGISTERBLSPUB_FUNCTION = bytes4(keccak256("registerBLSPublicKey(string,bytes)"));
    bytes4 constant GETPUB_FUNCTION = bytes4(keccak256("getRegisteredPubkey(address)"));


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

    function registerBLSPublicKey(string operator, bytes memory pubkey) public  returns (bool) {
        (bool success, bytes memory responseOrReason) = AVSTASK_PRECOMPILE_ADDRESS.call(
            abi.encodeWithSelector(
                REGISTERBLSPUB_FUNCTION,
                operator,
                pubkey
            )
        );

        uint256 lastlyUpdatedPrincipleBalance;
        if (success) {
            (, lastlyUpdatedPrincipleBalance) = abi.decode(responseOrReason, (bool, uint256));
        }
        return success;
    }

    function getRegisteredPubkey(address operator) public  returns (bytes32) {
        (bool success, bytes memory responseOrReason) = AVSTASK_PRECOMPILE_ADDRESS.call(
            abi.encodeWithSelector(
                GETPUB_FUNCTION,
                operator
            )
        );

        uint256 lastlyUpdatedPrincipleBalance;
        if (success) {
            (, lastlyUpdatedPrincipleBalance) = abi.decode(responseOrReason, (bool, uint256));
        }
        return bytes32(responseOrReason);
    }

}