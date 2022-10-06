package klocclient

import (
	"github.com/klaytn/klaytn/client"
	"log"
	"os"
)

const (
	ChainId = 1001
	RPCUrl = "wss://public-node-api.klaytnapi.com/v1/baobab/ws"
)
func Connection() *client.Client {
	conn, err := client.Dial(RPCUrl)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	return conn
}
