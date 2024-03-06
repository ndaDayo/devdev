package get

import (
	"time"
)

type Params interface{}

type PullRequestsParams struct {
	Owner    string
	Repo     string
	Username string
}

type CodeActivityFetcher struct {
	ResourceFetcher ResourceFetcher
}

type ResourceFetcher interface {
	GetResource(resource interface{}) (interface{}, error)
}

type CodeActivity struct {
	TimeToMerge time.Duration
}
