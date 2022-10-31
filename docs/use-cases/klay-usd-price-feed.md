# Klay/USD Price Feed

## Adapter

```json
{
  "active": true,
  "name": "KLAY_USD",
  "job_type": "DATA_FEED",
  "adapter_id": "efbdab54419511edb8780242ac120002",
  "oracle": "0xCC4377b912c4517Fe895817c6a7c6937D92A70B3",
  "feeds": [
    {
      "url": "https://min-api.cryptocompare.com/data/pricemultifull?fsyms=KLAY&tsyms=USD",
      "headers": [
        {"Content-Type" :  "application/json"}
      ],
      "request_type": "GET",
      "reducers": [
        {
          "function": "PARSE",
          "args": ["RAW","KLAY","USD","PRICE"]
        },
        {
          "function": "MUL",
          "args": [1000000000]
        }
      ]
    },
    {
      "url": "https://rest.coinapi.io/v1/exchangerate/KLAY/USD",
      "headers": [
        {"X-CoinAPI-Key": "209D7E0E-9330-4706-B693-6E2ED10279A5"}
      ],
      "reducers": [
        {
          "function": "PARSE",
          "args": ["rate"]
        },
        {
          "function": "MUL",
          "args": [1000000000]
        }
      ]
    }
  ]
}
```

## Oracle

{% embed url="https://remix.ethereum.org/#version=soljson-v0.8.7+commit.e28d00a7.js&optimize=false&runs=200&gist=13d672362bd84375b51c90b2a7ee5b7b" %}

## Price Consumer

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "./Oracle.sol";

contract PriceConsumer {

    Oracle public oracleAddress;

    uint public price;

    function requestKlayPrice(
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
```



##
