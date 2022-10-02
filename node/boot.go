package node

import (
	"github.com/alofeoluwafemi/klay-oracle/adapter"
	"log"
)

func Boot() {
	err := adapter.LoadJobs()

	if err != nil {
		log.Fatalln(err)
	}

	LoadReducers()
}

func Run() {
//Loop through Jobs
//Listen to Oracle for a Certain Event on Address
//When Event fires
//Loop through Feeds
//Make Request to URL Based on Type & Headers
//Perform Reducer Action
//Call Oracle with Response
//Keep Listening for Further event
}