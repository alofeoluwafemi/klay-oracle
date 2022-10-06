// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "./Oracle.sol";

contract RandomNumberConsumer {

    Oracle public oracleAddress;

    uint public rand;

    function requestRandomNumber(
        Oracle _oracleAddress,
        string memory _adapterId
    ) external returns(bool) {

        oracleAddress = _oracleAddress;

        oracleAddress.newOracleRequest(this.setRandomNumber.selector, _adapterId, address(this));

        return true;
    }

    function setRandomNumber(
        uint _rand
    ) external {

        require(msg.sender == address(oracleAddress),"RN Consumer: Permission Denied");

        rand = _rand;
    }

}