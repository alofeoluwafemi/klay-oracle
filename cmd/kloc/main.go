package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = godotenv.Load(path.Join(wd,"cmd",".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var klocCmd = &cobra.Command{
		Use:   "kloc",
		Short: "Oracle Protocol for Klaytn Network",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	klocCmd.AddCommand(versionCmd)
	klocCmd.AddCommand(newNodeAccountCmd())
	//klocCmd.AddCommand(nodeAccountSummary)
	//klocCmd.AddCommand(nodeBalanceCmd())
	//klocCmd.AddCommand(listJobsCmd())

	err = klocCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}