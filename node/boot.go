package node

import (
	"fmt"
	"log"
)

func Boot(jobsPath string) {
	err := LoadJobs(jobsPath)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(Jobs)

	LoadReducers()
}

func Run() {
//Loop through Jobs
//Listen to Oracle for a Certain Event on Address
//When Event fires

oracle := "0xDC77A7A497C9a9A7C086E3d57bFb753fF2cFa414"

for i := 0; i < len(Jobs); i++{
	if job := Jobs[i]; job.Oracle == oracle {
		job.process()
	}
}


//Make Request to URL Based on Type & Headers
//Perform Reducer Action
//Call Oracle with Response
//Keep Listening for Further event
}

func (job Job) process() {
	log.Printf(job.Name)
}