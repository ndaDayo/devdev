package activity_uc

import (
	"time"

	entity "github.com/ndaDayo/devdev/entity/activity"
)

type ActivityFetcher interface {
	FetchActivity(params interface{}) *entity.Activity
}

type Params interface{}

type CodeFetcher interface {
	GetCodeActivity(c CodeParams) (*CodeActivity, error)
}

type PullRequestsParams struct {
	Owner    string
	Repo     string
	Username string
}

type ResourceGetter interface {
	GetResource(resource interface{}) (interface{}, error)
}

type CodeActivity struct {
	TimeToMerge time.Duration
}
