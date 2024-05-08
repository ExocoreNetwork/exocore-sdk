// SPDX-License-Identifier: BUSL-1.1
pragma solidity >0.8.9;

import "@openzeppelin/contracts/interfaces/IERC1271.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

/**
 * @title Library of utilities for making EIP1271-compliant signature checks.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
library EIP1271SignatureUtils {
    // bytes4(keccak256("isValidSignature(bytes32,bytes)")
    bytes4 internal constant EIP1271_MAGICVALUE = 0x1626ba7e;

    /**
     * @notice Checks @param signature is a valid signature of @param digestHash from @param signer.
     * If the `signer` contains no code -- i.e. it is not (yet, at least) a contract address, then checks using standard ECDSA logic
     * Otherwise, passes on the signature to the signer to verify the signature and checks that it returns the `EIP1271_MAGICVALUE`.
     */
    function checkSignature_EIP1271(
        address signer,
        bytes32 digestHash,
        bytes memory signature
    ) internal pure {
        /**
         * check validity of signature:
         * `signature` must be a valid ECDSA signature from `signer`, indicating their intention for this action
         */
        require(
            ECDSA.recover(digestHash, signature) == signer,
            "EIP1271SignatureUtils.checkSignature_EIP1271: signature not from signer"
        );
    }
}
