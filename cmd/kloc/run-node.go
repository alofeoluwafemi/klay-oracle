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
	var cmd = &cobra.Command{
		Use:   "run:watch",
		Short: "Start watching to requests Oracle fulfill",
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