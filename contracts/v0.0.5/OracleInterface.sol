// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

interface OracleInterface {

    function fulfillOracleRequest(
        uint requestId,
        bytes32 data
    ) external returns (bool);

    function newOracleRequest(
        bytes4 callbackFunctionId,
        string memory adapterID,
        address callBackContract
    ) external returns (bool);
}
