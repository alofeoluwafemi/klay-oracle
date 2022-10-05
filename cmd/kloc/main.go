package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"os"
	"path"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("%v",err)
		os.Exit(1)
	}

	err = godotenv.Load(path.Join(wd,"cmd",".env"))
	if err != nil {
		fmt.Printf("Error loading .env file")
		os.Exit(1)
	}

	var klocCmd = &cobra.Command{
		Use:   "kloc",
		Short: "Oracle Protocol for Klaytn Network",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	klocCmd.AddCommand(versionCmd)
	klocCmd.AddCommand(NodeAccountCmd())
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