# Kloc Command

The **kloc** command is used to interact with the KlayOracle node.

To install it, build the cmd folder of the repository by running&#x20;

```shell
go install ./cmd/kloc/...
```

Confirm it works by checking version

```shell
kloc version
```

Create a Klaytn wallet for your Node to fulfill Oracle requests with.

```
kloc node account:create
```

```shell
// Output for a new Node Wallet

2022/10/05 14:14:24 Using existing account...
2022/10/05 14:14:27 Wallet successfully created.
2022/10/05 14:14:27 Node wallet address 0xcAD0444951A5faB031b62fcF717eCb5a1e8d7Dc7. 
Visit https://baobab.wallet.klaytn.foundation/faucet to fund your node before your node can fulfill Oracle request

// Output For Existing Node Wallet

2022/10/05 14:14:24 Generating new account...
2022/10/05 14:14:27 Wallet successfully created.
2022/10/05 14:14:27 Node wallet address 0xcAD0444951A5faB031b62fcF717eCb5a1e8d7Dc7. 
Visit https://baobab.wallet.klaytn.foundation/faucet to fund your node before your node can fulfill Oracle request
```

Get your Node wallet info

```shell
 kloc node account:info
```

```shell
// Output

2022/10/05 15:29:16 Account: 0xcAD0444951A5faB031b62fcF717eCb5a1e8d7Dc7
2022/10/05 15:29:20 Balance: 150.00
2022/10/05 15:29:20 View on explorer https://baobab.scope.klaytn.com/account/0xcAD0444951A5faB031b62fcF717eCb5a1e8d7Dc7?tabId=txList
```

{% hint style="info" %}
Fund your Node wallet on testnet. Visit https://baobab.wallet.klaytn.foundation/faucet
{% endhint %}
