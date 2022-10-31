# Node

**OracleInterface Contract:** https://remix.ethereum.org/#version=soljson-v0.8.7+commit.e28d00a7.js\&optimize=false\&runs=200\&gist=5ee9cf2d97d51ab47fddd185b6923522

**Oracle Contract:** https://remix.ethereum.org/#version=soljson-v0.8.7+commit.e28d00a7.js\&optimize=false\&runs=200\&gist=13d672362bd84375b51c90b2a7ee5b7b

**PriceConsumer Contract:** https://remix.ethereum.org/#version=soljson-v0.8.7+commit.e28d00a7.js\&optimize=false\&runs=200\&gist=b76e7e86cb82f5f6ad54fcc2a1545f44\&evmVersion=null

Deploy the Oracle contract, and pass your node address. You can get your node address by running `kloc node account:info` an adapter id of `adapter/jobs/klay_usd.adapter.json` to the Oracle constructor.

Copy the Oracle address and replace it in the `adapter/jobs/klay_usd.adapter.json` file. Deploy the Price consumer contract also and afterward run `kloc node run:watch` to start the node to start fulfilling Oracle requests.

Call the function `requestKlayPrice` and pass the node address and adapter id, if you check the running terminal you will see logs of the Oracle request processing.

Call the `price` getter function to see the price response.
