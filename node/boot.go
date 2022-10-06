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
)

func Boot(jobsPath string) {
	err := LoadJobs(jobsPath)

	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Printf("Jobs: %+v\n",Jobs)

	LoadReducers()
}

func Run() {
	conn := klocclient.Connection()

	for i := 0; i < len(Jobs); i++ {

		if job := Jobs[i]; job.Active == true {

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
					log.Printf("Error while listening for Oracle event: %v", err)
				case vLog := <-logs:
					event, err := oracle.ParseNewOracleRequest(vLog)
					if err != nil {
						log.Printf("Error parsing event log: %v\n", err)
					}else{

						log.Printf("Calling Job of Adapter ID %v\n", event.AdapterId)
						log.Printf("With Request ID %v\n", event.RequestId)

						if job := Jobs[i]; job.AdapterId == event.AdapterId {
							jobResponses := job.process()
							randomIndex := rand.Intn(len(jobResponses))
							selected := jobResponses[randomIndex]

							//Get mean and call oracle with response
							fmt.Printf("\nResponse for Oracle %v\n",selected)

							_, trxOpts := klocaccount.LoadAccount()
							data := [32]byte{}
							copy(data[:], []byte(selected))

							trx, err := oracle.FulfillOracleRequest(trxOpts, event.RequestId, data)
							if err != nil {
								log.Printf("Error Fulfilling OracleRequest %v with data %v. reason %v\n", job.Oracle, data, err)
							}else{
								log.Printf("Node called Oracle %v with result %v\n",job.Oracle, selected)
								log.Printf("Transaction Hash %v, Nonce %v", trx.Hash().Hex(), trx.Nonce())
								log.Printf("Transaction Data %v", trx.Data())
							}

						}

					}

				}
			}

		}

	}

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