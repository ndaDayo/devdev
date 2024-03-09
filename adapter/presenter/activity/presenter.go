package presenter

import (
	entity "github.com/ndaDayo/devdev/domain/entity/activity"
	usecase "github.com/ndaDayo/devdev/usecase/activity"
)

type activityPresenter struct{}

func NewActivityPresenter() usecase.ActivityPresenter {
	return activityPresenter{}
}

func (p activityPresenter) Output(entity.Activity) usecase.ActivityOutput {
	return usecase.ActivityOutput{}
}
