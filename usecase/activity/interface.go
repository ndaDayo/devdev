package activity_uc

import (
	"time"

	entity "github.com/ndaDayo/devdev/entity/activity"
)

type ActivityFetcher interface {
	FetchActivity(params interface{}) *entity.Activity
}

type Params interface{}

type CodeActivityFetcher struct {
	PullRequestsFetcher PullRequestsFetcher
}

type CodeFetcher interface {
	GetCodeActivity(c CodeParams) (*CodeActivity, error)
}

type PullRequestsParams struct {
	Owner    string
	Repo     string
	Username string
}

type PullRequestsFetcher interface {
	FetchPullRequests(params PullRequestsParams) ([]entity.PullRequest, error)
}

type CodeActivity struct {
	TimeToMerge time.Duration
}
