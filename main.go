package main

import (
	"fmt"
	"github.com/alofeoluwafemi/klay-oracle/node"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path"
	"path/filepath"
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

	jobsDir, ok := os.LookupEnv("JOBS_PATH")
	if !ok {
		fmt.Fprintln(os.Stderr, "JOBS_PATH is not set")
		os.Exit(1)
	}

	wd, err = os.Getwd()
	if err != nil {
		log.Fatalf("%v", err)
	}

	jobsPath := filepath.Join(wd, jobsDir)

	node.Boot(jobsPath)
	node.Run()
}
