package presenter

import (
	entity "github.com/ndaDayo/devdev/domain/entity/activity"
	usecase "github.com/ndaDayo/devdev/usecase/activity"
)

type activityPresenter struct{}

func NewActivityPresenter() usecase.ActivityPresenter {
	return activityPresenter{}
}

func (ac activityPresenter) Output(e entity.Activity) usecase.ActivityOutput {
	pulls := e.CodeActivity.PullRequests
	var p = make([]usecase.PullRequest, 0)

	for _, pull := range pulls {
		p = append(p, usecase.PullRequest{
			LeadTime: pull.GetLeadTime(),
		})
	}

	o := usecase.ActivityOutput{
		PullRequests: p,
	}

	return o
}
