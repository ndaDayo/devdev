package activity_uc

import (
	"time"

	entity "github.com/ndaDayo/devdev/entity/activity"
)

type ActivityFetcher interface {
	FetchActivity(params interface{}) *entity.Activity
}

type Params interface{}

type CodeActivityFetcher struct{}

type CodeFetcher interface {
	GetCodeActivity(owner, repo string) (*CodeActivity, error)
}

type CodeActivity struct {
	TimeToMerge time.Duration
}
