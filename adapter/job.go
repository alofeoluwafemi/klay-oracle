package adapter

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type AdapterType interface {
	getTypeName() string
}

const (
	TypeFeed = "DATA_FEED"
	TypeRDN = "RANDOM_NUMBER"
)

type Reducer struct {
	Function string `json:"function"`
	Args []interface{} `json:"args"`
}

type Feed struct {
	Url string `json:"url"`
	RequestType string `json:"request_type"`
	Headers []map[string]string `json:"headers"`
	Reducers []Reducer `json:"reducers"`
}

type Job struct {
	Active bool `json:"active"`
	Name string `json:"name"`
	JobType string `json:"job_type"`
	AdapterId string `json:"adapter_id"`
	Oracle string `json:"oracle"`
	Feeds []Feed `json:"feeds"`
}

var (
	Jobs []*Job
	Fsys fs.FS
)


func LoadJobs() error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	jobsPath := filepath.Join(wd, "adapter","jobs")

	Fsys = os.DirFS(jobsPath)

	err = fs.WalkDir(Fsys, ".", walkJobs)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

func walkJobs(path string, d fs.DirEntry, err error) error  {
	configFileName := filepath.Base(path)

	if configFileName != "." {
		contentContent, err := fs.ReadFile(Fsys, configFileName)
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		job := new(Job)
		err = json.Unmarshal(contentContent, job)
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		//Only load active jobs should be loaded
		if job.Active == true {
			Jobs = append(Jobs, job)

			fmt.Printf("%+v\n",job.Name)
		}


	}

	return nil

}