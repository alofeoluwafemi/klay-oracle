// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "./OracleInterface.sol";

contract Oracle is OracleInterface {

    address public nodeAddress;

    string public adapterId;

    struct Request {
        uint    requestId;
        address nodeAddress;
        string  adapterId;
        bytes4  callbackFunctionId;
        address callBackContract;
        bytes32   data;
    }

    Request[] public requests;

    event NewOracleRequest(string adapterId, uint indexed requestId);

    constructor(address _nodeAddress,string memory _adapterId) {
        nodeAddress = _nodeAddress;
        adapterId = _adapterId;
    }

    function newOracleRequest(
        bytes4 callbackFunctionId,
        string memory adapterID,
        address callBackContract
    ) external override returns(bool) {
        uint requestId = requests.length;

        Request memory request = Request(
            requestId,
            nodeAddress,
            adapterID,
            callbackFunctionId,
            callBackContract,
            bytes32("")
        );

        requests.push(request);

        emit NewOracleRequest(adapterID, requestId);

        return true;
    }

    function fulfillOracleRequest(
        uint requestId,
        bytes32 data
    ) external override returns(bool) {

        require(msg.sender == nodeAddress,"Oracle: Permission needed for node");

        Request memory request = requests[requestId];

        (bool success, ) = request.callBackContract.call(abi.encodeWithSelector(request.callbackFunctionId, data));

        return success;
    }

}