package main

import (
	"fmt"
	"github.com/alofeoluwafemi/klay-oracle/node"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

func runNode() *cobra.Command {
	var newAccountCmd = &cobra.Command{
		Use:   "node",
		Short: "Run node",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	newAccountCmd.AddCommand(newKeyStore())

	return newAccountCmd
}

func watch() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "oracle:watch",
		Short: "Start watching for requests to fulfill",
		Run: func(cmd *cobra.Command, args []string) {
			jobsDir, ok := os.LookupEnv("JOBS_PATH")
			if !ok {
				fmt.Fprintln(os.Stderr, "JOBS_PATH is not set")
				os.Exit(1)
			}

			wd, err := os.Getwd()
			if err != nil {
				log.Fatalf("%v", err)
			}

			jobsPath := filepath.Join(wd, jobsDir)

			node.Boot(jobsPath)
			node.Run()
		},
	}

	return cmd
}