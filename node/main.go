package main

import (
	"github.com/alofeoluwafemi/klay-oracle/adapter"
	"log"
)

func main() {
	err := adapter.LoadJobs()

	if err != nil {
		log.Fatalln(err)
	}
}
