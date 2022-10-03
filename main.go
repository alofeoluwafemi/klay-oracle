package main

import (
	"github.com/alofeoluwafemi/klay-oracle/node"
	"log"
	"os"
	"path/filepath"
)

func main() {
	jobsDir := os.Getenv("JOBS_PATH")

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("%v", err)
	}

	jobsPath := filepath.Join(wd, jobsDir)

	node.Boot(jobsPath)
	node.Run()
}
