package get

import (
	"time"
)

type CodeActivityFetcher struct {
	ResourceFetcher ResourceFetcher
}

type ResourceFetcher interface {
	GetResource(resource interface{}) (interface{}, error)
}

type CodeActivity struct {
	TimeToMerge time.Duration
}
