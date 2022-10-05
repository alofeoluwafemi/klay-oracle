package klocclient

import (
	"github.com/klaytn/klaytn/client"
	"log"
	"os"
)

func Connection() *client.Client {
	conn, err := client.Dial("wss://public-node-api.klaytnapi.com/v1/baobab/ws")
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	return conn
}
