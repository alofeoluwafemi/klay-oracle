package main

import (
	"github.com/alofeoluwafemi/klay-oracle/node/klocaccount"
	"github.com/spf13/cobra"
	"log"
)

func accountInfo() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "account:info",
		Short: "Display wallet info for node",
		Run: func(cmd *cobra.Command, args []string) {

			account, _ := klocaccount.LoadAccount()

			log.Printf("Account: %v\n", account.Address.String())

			balance := klocaccount.Balance()

			log.Printf("Balance: %v\n", balance)

			log.Printf("View on explorer %v/%v?tabId=txList","https://baobab.scope.klaytn.com/account",account.Address.String())
		},
	}

	return  cmd
}