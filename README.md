## Learn More

**Website**: https://klayoracle.com

**Documentation**: https://klayoracle.gitbook.io/klayoracle-documentation/

### What is KlayOracle

KlayOracle adopts the design used by leading and existing oracles to bring offchain data to Klaytn Smart Contracts. Its straight forward to setup to access offchain data such as prices, sport data, weather data, random numbers or any external API usecase suites your project purpose.

### Why Use KlayOracle

**Reliable Data Feed:** Get reliable price, weather, and external API data for your Defi and GameFi apps. Additionally, obtain uniquely generated numbers for your smart contract use cases.

**Customizable Adapter:** Without relying on external data feeds, which may be fraudulent, you can quickly set up adapters to aggregate API data for any off-chain data you need to allow.

**Self Hosted Node:** KlayOracle currently relies on you configuring your own adapters and node. to manage the node or nodes and have faith that you are receiving expected off-chain data that is genuine.

### Products

- Random Numbers
- Data Feed

### Hackathon Milestone

| Goal      | Status | Comment     |
| :---        |    :----:   |          ---: |
| Figma Design for Website      | Done      | https://bit.ly/3LWehun    |
| KlayOracle Website   | Done        | https://klayoracle.com      |
| KlayOracle v0.0.5 Adapter   | Done        | https://bit.ly/3C8fLNy      |
| KlayOracle v0.0.5 Oracle Contract   | Done        | https://bit.ly/3SMqPGY      |
| KlayOracle v0.0.5 Node   | Done        | https://bit.ly/3SMqPGY      |
| Developer Documentation   | Done        | https://bit.ly/3SMqPGY      |
| KlayOracle Baobab Deploy   | Done        | -      |
| Video Tutorial   | Done        | https://www.youtube.com/watch?v=nC-LlVKuNL0      |

### Post Hackathon Plans
- Build and release KlayOracle Node V1.0 for testnet
- Build a token reward system for Node runner
- Build a token reward system for Node runner
- Get beta tester to build using our oracle on the test net
- Smart contract and code audit
- Live launch Q1 2023


### Use cases
- **NFT**: To generate randomly unique numbers for NFT ID.
- **Gaming**: Fetch sport data from Oracle, for a betting smart contract.
- **Defi**: Get Price pair for your staking or exchange smart contract.

### Architechture
![KlayOracle Architechture](https://www.klayoracle.com/images/define.png)

### Components
- **Adapter**
- **Oracle Contract**
- **KlayOracle Node**
- **Kloc Command**


# Setup

### Create this environment file in `cmd/.env.

```dotenv
KEYSTORE_PASSWORD=1234567890Abcd
KEYSTORE_PATH=node/klocaccount
JOBS_PATH=adapter/jobs
```

`KEYSTORE_PASSWORD` and `KEYSTORE_PATH` is the password to unlock your Node wallet and path to store your node wallet keystore file.

### Build kloc Command

```console
go install ./cmd/kloc/...
```

Confirm it works by checking version

```console
kloc version
```

### Create a Klaytn wallet for your Node to fulfill Oracle request with

```console
kloc node account:create
```

**Output:**

For New Node Wallet

```console
2022/10/05 14:14:24 Using existing account...
2022/10/05 14:14:27 Wallet successfully created.
2022/10/05 14:14:27 Node wallet address 0xcAD0444951A5faB031b62fcF717eCb5a1e8d7Dc7. 
Visit https://baobab.wallet.klaytn.foundation/faucet to fund your node before your node can fufill Oracle request
```

For Existing Node Wallet

```console
2022/10/05 14:14:24 Generating new account...
2022/10/05 14:14:27 Wallet successfully created.
2022/10/05 14:14:27 Node wallet address 0xcAD0444951A5faB031b62fcF717eCb5a1e8d7Dc7. 
Visit https://baobab.wallet.klaytn.foundation/faucet to fund your node before your node can fufill Oracle request

```

Get your Node wallet info

```console
 kloc node account:info
```

Output

```console
2022/10/05 15:29:16 Account: 0xcAD0444951A5faB031b62fcF717eCb5a1e8d7Dc7
2022/10/05 15:29:20 Balance: 150.00
2022/10/05 15:29:20 View on explorer https://baobab.scope.klaytn.com/account/0xcAD0444951A5faB031b62fcF717eCb5a1e8d7Dc7?tabId=txList
```

**NB:** Fund your Node wallet on testnet. Visit https://baobab.wallet.klaytn.foundation/faucet


### Deploy Oracle Contract

**OracleInterface Contract:** https://remix.ethereum.org/#version=soljson-v0.8.7+commit.e28d00a7.js&optimize=false&runs=200&gist=5ee9cf2d97d51ab47fddd185b6923522

**Oracle Contract:** https://remix.ethereum.org/#version=soljson-v0.8.7+commit.e28d00a7.js&optimize=false&runs=200&gist=13d672362bd84375b51c90b2a7ee5b7b

**PriceConsumer Contract:** https://remix.ethereum.org/#version=soljson-v0.8.7+commit.e28d00a7.js&optimize=false&runs=200&gist=b76e7e86cb82f5f6ad54fcc2a1545f44&evmVersion=null

Deploy the Oracle contract, pass your node address. You can get your node address by running `kloc node account:info` and adapter id of `adapter/jobs/klay_usd.adapter.json` to the Oracle constructor.

Copy the Oracle address and replace it in the `adapter/jobs/klay_usd.adapter.json` file. Deploy the Price consumer contract also and after wards run `kloc node run:watch` to start the node to start fulfilling Oracle requests.

Call the function `requestKlayPrice` and pass the node address and adapter id, if you check the running terminal you will see logs of the Oracle request procssing.

Call the `price` getter function to see the price response.


![Deploy Klaytn Oracle](https://github.com/alofeoluwafemi/klay-oracle/blob/development/docs/img/Screenshot-57.png)

![Deploy PriceConsumer Klaytn](https://github.com/alofeoluwafemi/klay-oracle/blob/development/docs/img/Screenshot-47.png)
