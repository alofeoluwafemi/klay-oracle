package node

import (
	"context"
	"fmt"
	"github.com/alofeoluwafemi/klay-oracle/node/klocaccount"
	"github.com/alofeoluwafemi/klay-oracle/node/klocclient"
	"github.com/alofeoluwafemi/klay-oracle/node/klocoracle"
	"github.com/go-resty/resty/v2"
	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"log"
	"math/rand"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func Boot(jobsPath string) {
	err := LoadJobs(jobsPath)

	if err != nil {
		log.Fatalln(err)
	}

	LoadReducers()
}

func Run() {
	defer wg.Done()

	wg.Add(len(Jobs))	//It will never be done anyways

	for i := 0; i < len(Jobs); i++ {

		if job := Jobs[i]; job.Active == true {

			go watch(job)
		}
	}

	wg.Wait()
	log.Printf("Watching all active jobs...")
}

func watch(job *Job) {
	conn := klocclient.Connection()

	log.Printf("Watching active Job %v\n", job.Name)
	log.Printf("Listening to Oracle %v\n", job.Oracle)

	oracle, err := klocoracle.NewOracle(common.HexToAddress(job.Oracle), conn)

	if err != nil {
		log.Printf("Error with job %v, cannot get Oracle due to %v", job.Name, err)
	}

	oracleAddress := common.HexToAddress(job.Oracle)
	query := klaytn.FilterQuery{
		Addresses: []common.Address{oracleAddress},
	}

	logs := make(chan types.Log)
	subscription, err := conn.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Printf("Oracle Subscription error: %v",err)
		log.Printf("Skipping Oracle: %v",err)
	}

	for {
		select {
		case err := <-subscription.Err():
			log.Printf("Error while listening for Oracle %v event due to: %v",job.Oracle, err)
			watch(job)
		case vLog := <-logs:
			event, err := oracle.ParseNewOracleRequest(vLog)
			if err != nil {
				log.Printf("Error parsing event log: %v\n", err)
			}else{
				log.Printf("Calling Job of Adapter ID %v\n", event.AdapterId)
				log.Printf("With Request ID %v\n", event.RequestId)

				if job.AdapterId == event.AdapterId {
					jobResponses := job.process()
					randomIndex := rand.Intn(len(jobResponses))
					selected := jobResponses[randomIndex]

					//Get mean and call oracle with response
					fmt.Printf("\nResponse for Oracle %v\n",selected)

					_, trxOpts := klocaccount.LoadAccount()

					var hexSelected string

					intSelected, err := strconv.ParseInt(selected,10, 64)
					if err != nil {
						log.Printf("Error parsing voted result %v due to %v\n", selected,  err)

						hexSelected = strconv.FormatInt(0, 16)
					}else{
						hexSelected = strconv.FormatInt(intSelected, 16)
					}

					data := common.HexToHash(hexSelected)

					trx, err := oracle.FulfillOracleRequest(trxOpts, event.RequestId, data)
					if err != nil {
						log.Printf("Error Fulfilling OracleRequest %v with data %v. reason %v\n", job.Oracle, data, err.Error())
					}else{
						log.Printf("Node called Oracle %v with result %v and data %v\n",job.Oracle, selected, data)
						log.Printf("Transaction Hash %v, Nonce %v\n", trx.Hash().Hex(), trx.Nonce())
						log.Printf("Transaction Data %v\n", trx.Data())
					}

				}

			}

		}
	}
}

func bytes32FromString(s string) [32]byte {
	var b32 [32]byte
	copy(b32[:], s)
	return b32
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