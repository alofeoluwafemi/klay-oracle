// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "./Oracle.sol";

contract Consumer {

    Oracle public oracleAddress;

    uint public price;

    function requestEthereumPrice(
        Oracle _oracleAddress,
        string memory _adapterId
    ) external returns(bool) {

        oracleAddress = _oracleAddress;

        oracleAddress.newOracleRequest(this.setKLAYUSD.selector, _adapterId, address(this));

        return true;
    }

    function setKLAYUSD(
        uint _price
    ) external {

        require(msg.sender == address(oracleAddress),"Consumer: Permission Denied");

        price = _price;
    }

}