package main

import (
	"fmt"
	"github.com/klaytn/klaytn/accounts/keystore"
	"github.com/spf13/cobra"
	"os"
	"path"
)

func newNodeAccountCmd() *cobra.Command {
	var newAccountCmd = &cobra.Command{
		Use:   "node",
		Short: "Wallet for node usage",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	newAccountCmd.AddCommand(newKeyStore())

	return newAccountCmd
}

func newKeyStore() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "account:create",
		Short: "Check or create new keystore for node",
		Run: func(cmd *cobra.Command, args []string) {

			wd, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			setPath, ok := os.LookupEnv("KEYSTORE_PATH")
			if !ok {
				fmt.Fprintln(os.Stderr, "KEYSTORE_PATH is not set")
				os.Exit(1)
			}

			password, ok := os.LookupEnv("KEYSTORE_PASSWORD")
			if !ok {
				fmt.Fprintln(os.Stderr, "KEYSTORE_PASSWORD is not set")
				os.Exit(1)
			}

			keyStorePath := path.Join(wd, setPath)
			ks := keystore.NewKeyStore(keyStorePath,keystore.StandardScryptN, keystore.StandardScryptP)

			account, err := ks.NewAccount(password)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			nodeAddress := account.Address.String()
			faucetUrl := "https://baobab.wallet.klaytn.foundation/faucet"

			fmt.Println("Wallet successfully created.")
			fmt.Printf("Node wallet address %v. \nVisit %v to fund your node before your node can fufill Oracle request\n", nodeAddress, faucetUrl)
		},
	}

	return  cmd
}