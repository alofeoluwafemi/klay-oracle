package main

import (
	"github.com/alofeoluwafemi/klay-oracle/adapter"
	"github.com/alofeoluwafemi/klay-oracle/node"
	"log"
)

func main() {
	err := adapter.LoadJobs()

	if err != nil {
		log.Fatalln(err)
	}

	node.LoadReducers()
}
