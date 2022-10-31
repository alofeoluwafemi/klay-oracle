# About

### What is KlayOracle

KlayOracle adopts the design used by leading and existing oracles to bring offchain data to Klaytn Smart Contracts. Its straight forward to setup to access offchain data such as prices, sport data, weather data, random numbers or any external API usecase suites your project purpose.

### Why Use KlayOracle

**Reliable Data Feed:** Get reliable price, weather, and external API data for your Defi and GameFi apps. Additionally, obtain uniquely generated numbers for your smart contract use cases.

**Customizable Adapter:** Without relying on external data feeds, which may be fraudulent, you can quickly set up adapters to aggregate API data for any off-chain data you need to allow.

**Self Hosted Node:** KlayOracle currently relies on you configuring your own adapters and node. to manage the node or nodes and have faith that you are receiving expected off-chain data that is genuine.

### Products

- Random Numbers
- Data Feed

### Architechture
![Watch the video](https://github.com/alofeoluwafemi/klay-oracle/blob/development/docs/.gitbook/assets/define.png)


### How it Works 
https://user-images.githubusercontent.com/7295729/198984762-910136cb-5ed0-4af8-a748-8f3bdcdc0ea9.mov

### Key differences between KlayOracle and other Oracles like Chainlink.

Firstly, we would like to point out that Chainlink is currently only active on the Klaytn Baobab Testnet and has not yet been made available on Mainnet, despite being in development since August 2021. They presently have two price feeds, primarily KLAY/USD and WEMIX/USD, and as you can expect, we also have both on KlayOracle and the Baobab Testnet. By building primarily on Klaytn at the beginning, we can more quickly launch features critically needed by developers on the Klaytn blockchain.

Secondly, our guiding principle is simplicity… Right now, anyone with a little technical know-how can set up a KlayOracle node and adapter more easily than they would Chainlink nodes and jobs. While this is expected with the simple data feeds we have now, KlayOracle’s simplicity is something we aim to maintain even with the addition of more complex data feeds.

We believe that by lowering the barrier to setting up nodes, oracles and adapters, we can onboard many more participants on the network. An increased number of nodes will help maintain secure and decentralized data on the Klaytn network, as more participants supply data, instead of just a few participants.

We envision ourselves as a less complex, less expensive alternative to Chainlink that also achieves genuine decentralization by enabling users to host more nodes (think x10 - x100 nodes than chainlink runs).



## Testnet Baobab Price Oracles

All contracts have been verified on **Klaytn Baobab Testnet**.

### Node 
[0xcAD0444951A5faB031b62fcF717eCb5a1e8d7Dc7](https://baobab.scope.klaytn.com/account/0xcAD0444951A5faB031b62fcF717eCb5a1e8d7Dc7?tabId=txList)

### KLAY/USD
Oracle: [0xCC4377b912c4517Fe895817c6a7c6937D92A70B3](https://baobab.scope.klaytn.com/account/0xCC4377b912c4517Fe895817c6a7c6937D92A70B3)  
Adapter ID: efbdab54419511edb8780242ac120002  
Price Consumer Sample Contract: [0x8e892CE230eEe07598eed68EA326e308A3d2687D](https://baobab.scope.klaytn.com/account/0x8e892CE230eEe07598eed68EA326e308A3d2687D)

### WEMIX/USD
Oracle: [0xb56c6b973688A8fbA3D6bB536b7CdFFC7b46252A](https://baobab.scope.klaytn.com/account/0xb56c6b973688A8fbA3D6bB536b7CdFFC7b46252A)  
Adapter ID: be498ee0521011edbdc30242ac120002  
Price Consumer Sample Contract: [0x6aB624819D2082801bA2d44a64DbE1201b608e4f](https://baobab.scope.klaytn.com/account/0x6aB624819D2082801bA2d44a64DbE1201b608e4f)
