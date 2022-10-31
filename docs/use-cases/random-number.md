# Random Number

## Adapter

```json
{
  "active": true,
  "name": "RAND",
  "job_type": "RANDOM_NUMBER",
  "adapter_id": "550a752a435111edb8780242ac120002",
  "oracle": "0x3472059945ee170660a9A97892a3cf77857Eba3A",
  "feeds": [
    {
      "url": "https://random-number2.p.rapidapi.com?min=9999&max=999999",
      "headers": [
        {"Content-Type" :  "application/json"},
        {"X-RapidAPI-Key" :  "7b51d2e2d0msh13d3226d725be02p1eea9cjsn45f324819d23"},
        {"X-RapidAPI-Host" :  "random-number2.p.rapidapi.com"}
      ],
      "request_type": "GET",
      "reducers": [
        {
          "function": "PARSE",
          "args": ["data","random_number"]
        }
      ]
    }
  ]
}
```

## Oracle

{% embed url="https://remix.ethereum.org/#version=soljson-v0.8.7+commit.e28d00a7.js&optimize=false&runs=200&gist=13d672362bd84375b51c90b2a7ee5b7b" %}

## Random Number Consumer

```solidity
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
```
