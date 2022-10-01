package adapter

type AdapterType interface {
	getTypeName() string
}

type TypeFeed string

type TypeRDN string


func (f TypeFeed) getTypeName() string{
	return "DATA_FEED"
}

func (f TypeRDN) getTypeName() string{
	return "RANDOM_NUMBER"
}

type Job struct {
	Active bool `json:"active"`
	AdapterName string `json:"adapter_name"`
	JobType AdapterType `json:"job_type"`
	AdapterId string `json:"adapter_id"`
	Oracle string `json:"oracle"`
}

var jobs []Job