package activity_uc

import (
	"github.com/ndaDayo/devdev/entity"
)

type ActivityFetcher interface {
	FetchActivity(params interface{}) *entity.Activity
}

type Params interface{}

type GithubFetcher struct{}
