package klocclient

import (
	"github.com/klaytn/klaytn/client"
	"log"
	"os"
)

const (
	ChainId = 1001
	RPCUrl = "wss://api.baobab.klaytn.net:8652"
)
func Connection() *client.Client {
	conn, err := client.Dial(RPCUrl)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	return conn
}
