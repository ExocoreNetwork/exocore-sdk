pragma solidity >=0.8.17;

contract TestProxy {

    address constant BLS_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000809;
    address constant AVS_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000902;
    address constant TASK_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000901;



    bytes4 constant BLS_FUNCTION_SELECTOR = bytes4(keccak256("generatePrivateKey()"));
    mapping(address => bytes) public results;

    event callgeneratePrivateKey(bool indexed success, bytes indexed latestAssetState);


    function generatePrivateKey() public  returns (bytes32) {
        (bool success, bytes memory responseOrReason) = BLS_PRECOMPILE_ADDRESS.call(
            abi.encodeWithSelector(
                BLS_FUNCTION_SELECTOR
            )
        );

        results[0x0000000000000000000000000000000000000111] = responseOrReason;
        emit callgeneratePrivateKey(success, responseOrReason);


    }
    function publicKey(bytes memory pri) public  returns (bytes32) {
        (bool success, bytes memory responseOrReason) = BLS_PRECOMPILE_ADDRESS.call(
            abi.encodeWithSelector(
                bytes4(keccak256("publicKey(bytes)")),
                pri
            )
        );

        results[0x0000000000000000000000000000000000000112] = responseOrReason;

    }


    function AVSAction(
        string memory avsName,
        string memory avsAddress,
        string memory operatorAddress,
        uint64 action,
        string memory avsOwnerAddress,
        string memory assetID) public  returns (bytes32) {

        (bool success, bytes memory responseOrReason) = AVS_PRECOMPILE_ADDRESS.call(
            abi.encodeWithSelector(
                bytes4(keccak256("AVSAction(string,string,string,uint64,string,string)")),
                avsName,avsAddress,operatorAddress,action,avsOwnerAddress,assetID
            )
        );

        results[0x0000000000000000000000000000000000000113] = responseOrReason;

    }


    function registerAVSTask(
        string memory TaskContractAddress,
        string memory Name,
        string memory MetaInfo) public  returns (bytes32) {

        (bool success, bytes memory responseOrReason) = TASK_PRECOMPILE_ADDRESS.call(
            abi.encodeWithSelector(
                bytes4(keccak256("registerAVSTask(string,string,string)")),
                TaskContractAddress,Name,MetaInfo
            )
        );

        results[0x0000000000000000000000000000000000000114] = responseOrReason;

    }

}