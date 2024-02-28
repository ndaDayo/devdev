package activity_uc

import entity "github.com/ndaDayo/devdev/entity/activity"

type ActivityFetcher interface {
	FetchActivity(params interface{}) *entity.Activity
}

type Params interface{}

type CodeActivityFetcher struct{}
