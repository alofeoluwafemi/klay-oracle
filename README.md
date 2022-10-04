

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

## Set Adapter Path for Node

```markdown
JOBS_PATH=adapter/jobs go run main.go 
```
