# Adapter

The Adapter defines the job specification for getting a particular offchain data for an Oracle request. Adn adapter is made up of different parts which are explained below.

```json
// Adapter for KLAY Price in USD

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

* **Status**: Options are either boolean true or false. False means the Adpater will not be loaded by the node, while True as you can guess mean it will be loaded and the oracle address of the adapter will be watched by the node for any requests and fulfill them.&#x20;
* **Name:** Name of your adapter. It can be anything, but for general understanding naming your adapter in a verbose manner is advisable.
* **Job Type**: For version v0.0.5 Beta, it's currently not being used. The two options available are either DATA\__FEED, RANDOM_\_NUMBER.
* **Adapter ID:** This is a unique **** 32-length string that must be unique through the Job folder. You make use of [https://www.uuidgenerator.net/](https://www.uuidgenerator.net/) and remove all dashes in between the UUID generated for you on that site to make use of.
* **Oracle**: This is the address of the Oracle smart contract deployed for this particular Job. It is one Oracle per job. Find the Oracle smart contract here  [https://remix.ethereum.org/#version=soljson-v0.8.7+commit.e28d00a7.js\&optimize=false\&runs=200\&gist=13d672362bd84375b51c90b2a7ee5b7b](https://remix.ethereum.org/#version=soljson-v0.8.7+commit.e28d00a7.js\&optimize=false\&runs=200\&gist=13d672362bd84375b51c90b2a7ee5b7b)****
* **Feeds:** This is probably the most important part of the adapter. This section defines the URLs to get the off-chain data and the rules of how the response should be handled by defining Reducers for each URL section. There are different types of reducers, currently 5 of them :&#x20;
  * **PARSE**: For transversing through JSON content
  * **MUL**: For multiplying a response value, usually needed to change the decimal to a value that can be passed to a smart contract since we can't pass decimals.
  * **ADD:** For adding a response value.
  * **DIV:** For dividing a response value.
  * **SUB:** For subtracting a response value.
