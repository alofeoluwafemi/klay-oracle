package main

import (
	"fmt"
	"github.com/alofeoluwafemi/klay-oracle/node/klocaccount"
	"github.com/joho/godotenv"
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

	created, keyPath := klocaccount.IsCreated()

	fmt.Println(created)
	fmt.Println(keyPath)
}
