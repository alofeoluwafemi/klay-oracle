package node

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"math/rand"
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

	// requestId := 0
	adapterId := "efbdab54-4195-11ed-b878-0242ac120002"

	for i := 0; i < len(Jobs); i++ {
		if job := Jobs[i]; job.AdapterId == adapterId {
			jobResponses := job.process()
			randomIndex := rand.Intn(len(jobResponses))
			selected := jobResponses[randomIndex]
			//Get mean and call oracle with response
			fmt.Printf("\nResponse for Oracle %v\n",selected)
		}
	}

	// Make Request to URL Based on Type & Headers
	// Perform Reducer Action
	// Call Oracle with Response
	// Keep Listening for Further event
}

func (job Job) process() []string {
	var results []string

	for i := 0; i < len(job.Feeds); i++ {
		feed := job.Feeds[i]

		client := resty.New()

		for j := 0; j < len(feed.Headers); j++ {
			header := feed.Headers[j]

			for key, value := range header {
				client.SetHeader(key, value)
			}
		}

		// Todo, switch between GET or POST request
		resp, err := client.R().
			EnableTrace().
			Get(feed.Url)

		if err != nil {
			log.Fatalf("Error running job %v, due to error %v", job.Name, err)
		}

		body := resp.String()

		//Run through set of reducers
		for k := 0; k < len(feed.Reducers); k++ {
			reducer := feed.Reducers[k]
			args := reducer.Function

			function := Reducers[args]
			argument := reducer.Args

			response, err := function(body, argument...)

			if err != nil {
				log.Fatalf("Error while running reducer %v on job %v ", reducer, job.Name)
			}

			body = fmt.Sprintf("%v", response)

			log.Printf("Result for %v -> %+v : %v", job.Name, reducer , body)
		}

		results = append(results, body)

	}

	log.Printf("Results for adapter %v are : %+v", job.AdapterId, results )

	return results
}