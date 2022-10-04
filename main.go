package main

import (
	"fmt"
	"github.com/alofeoluwafemi/klay-oracle/node"
	"log"
	"os"
	"path/filepath"
)

func main() {
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
}
