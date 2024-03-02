package activity_uc

import (
	"context"
	"time"

	entity "github.com/ndaDayo/devdev/entity/activity"
)

type ActivityFetcher interface {
	FetchActivity(params interface{}) *entity.Activity
}

type Params interface{}

type CodeActivityFetcher struct{}

type CodeFetcher interface {
	GetCodeActivity(c CodeParams) (*CodeActivity, error)
}

type PullRequestsParams struct {
	Owner    string
	Repo     string
	Username string
}

type PullRequestsFetcher interface {
	FetchPullRequests(ctx context.Context, params PullRequestsParams) (*entity.PullReq, error)
}

type CodeActivity struct {
	TimeToMerge time.Duration
}
