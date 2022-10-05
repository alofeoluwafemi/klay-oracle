

![What is KlayOracle](https://s3.amazonaws.com/alofe.oluwafemi/Klay+Oracle+Hero++Banner.jpg)

## Learn More

[Visit KlayOracle Website](https://klayoracle.com)

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
| KlayOracle v0.0.5 Node   | In Progress        | https://bit.ly/3SMqPGY      |
| Developer Documentation   | In Progress        | https://bit.ly/3SMqPGY      |
| KlayOracle Baobab Deploy   | Pending        | -      |
| Video Tutorial   | Pending        | -      |

### Post Hackathon Plans
- Build and release KlayOracle Node V1.0 for testnet
- Build a token reward system for Node runner
- Build a token reward system for Node runner
- Get beta tester to build using our oracle on the test net
- Smart contract and code audit
- Live launch Q1 2023


### Usecase
- **NFT**: To generate randomly unique numbers for NFT ID.
- **Gaming**: Fetch sport data from Oracle, for a betting smart contract.
- **Defi**: Get Price pair for your staking or exchange smart contract.

### Architechture
![KlayOracle Architechture](https://s3.amazonaws.com/alofe.oluwafemi/KlayOracle+Architechture.jpg)

### Components
- **Adapter**
- **Oracle Contract**
- **KlayOracle Node**


# Setup

Create this environment file in `cmd/.env.

```dotenv
KEYSTORE_PASSWORD=1234567890Abcd
KEYSTORE_PATH=node/account
```

Build `kloc command`

```console
go install ./cmd/kloc/...
```

Confirm it works by checking version

```console
kloc version
```

Create a Klaytn wallet for your Node to fulfill Oracle request with

```console
kloc node account:create
```

**Output:**

For New Account

```console
2022/10/05 14:14:24 Using existing account...
2022/10/05 14:14:27 Wallet successfully created.
2022/10/05 14:14:27 Node wallet address 0xcAD0444951A5faB031b62fcF717eCb5a1e8d7Dc7. 
Visit https://baobab.wallet.klaytn.foundation/faucet to fund your node before your node can fufill Oracle request
```

For Existing Account

```console
2022/10/05 14:14:24 Generating new account...
2022/10/05 14:14:27 Wallet successfully created.
2022/10/05 14:14:27 Node wallet address 0xcAD0444951A5faB031b62fcF717eCb5a1e8d7Dc7. 
Visit https://baobab.wallet.klaytn.foundation/faucet to fund your node before your node can fufill Oracle request

```

Get your Node account info

```console
 kloc node account:info
```

**NB:** Fund your Node wallet on testnet. Visit https://baobab.wallet.klaytn.foundation/faucet

## Set Adapter Path for Node

```markdown
JOBS_PATH=adapters/jobs
```

Add this to your IDE env like Golang or your Terminal albeit temporary.