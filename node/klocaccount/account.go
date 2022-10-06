package klocaccount

import (
	"context"
	"fmt"
	"github.com/alofeoluwafemi/klay-oracle/node/klocclient"
	"github.com/klaytn/klaytn/accounts"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/accounts/keystore"
	"io/fs"
	"log"
	"math/big"
	"os"
	"path"
)

func Variables() (keyPath string,keyPass string, keyFullPath string){
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	keyPath, ok := os.LookupEnv("KEYSTORE_PATH")
	if !ok {
		fmt.Fprintln(os.Stderr, "KEYSTORE_PATH is not set")
		os.Exit(1)
	}

	keyPass, ok = os.LookupEnv("KEYSTORE_PASSWORD")
	if !ok {
		fmt.Fprintln(os.Stderr, "KEYSTORE_PASSWORD is not set")
		os.Exit(1)
	}

	keyFullPath = path.Join(wd, keyPath)


	return keyPath, keyPass, keyFullPath
}

func IsCreated() (bool, string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	setPath, ok := os.LookupEnv("KEYSTORE_PATH")
	if !ok {
		fmt.Fprintln(os.Stderr, "KEYSTORE_PATH is not set")
		os.Exit(1)
	}

	keyStorePath := path.Join(wd, setPath)

	fsys := os.DirFS(keyStorePath)

	match, err := fs.Glob(fsys, "UTC--*-*-*-*-*.*--*")

	if len(match) > 0 {
		return true, path.Join(keyStorePath,match[0])
	}

	return false, ""
}

func LoadAccount() (accounts.Account, *bind.TransactOpts) {
	_, keystoreFile := IsCreated()
	_, password, keyStorePath := Variables()

	ks := keystore.NewKeyStore(keyStorePath, keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := os.ReadFile(keystoreFile)
	if err != nil {
		log.Fatal(err)
	}

	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = ks.Unlock(account, password)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	client := klocclient.Connection()
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("Cannot suggest Gas Price %v", err)
		os.Exit(1)
	}

	//gasLimit, err := client.EstimateGas(context.Background(), klaytn.CallMsg{GasPrice: gasPrice})
	//if err != nil {
	//	fmt.Printf("Cannot suggest Gas Limit %v", err)
	//	os.Exit(1)
	//}

	nonce, err := client.NonceAt(context.Background(), account.Address, nil)
	if err != nil {
		fmt.Printf("Cannot suggest Nonce %v", err)
		os.Exit(1)
	}

	trxOpts := bind.NewKeyedTransactorWithKeystore(account.Address, ks, big.NewInt(klocclient.ChainId))
	trxOpts.GasPrice = gasPrice
	trxOpts.GasLimit = 200000
	trxOpts.Nonce = new(big.Int).SetUint64(nonce)

	if err := os.RemoveAll(keystoreFile); err != nil {
		log.Fatal(err)
	}

	return account, trxOpts
}

func Balance() string {
	conn := klocclient.Connection()
	account, _ := LoadAccount()

	balance, err := conn.BalanceAt(context.Background(), account.Address ,nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	balanceInWei, _ := new(big.Float).SetInt(balance).Float64()
	mantissa, _ := new(big.Float).SetUint64(1000000000000000000).Float64()

	balanceInEther := balanceInWei / mantissa

	return fmt.Sprintf("%.2f", big.NewFloat(balanceInEther))
}