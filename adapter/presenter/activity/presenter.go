package presenter

import (
	entity "github.com/ndaDayo/devdev/domain/entity/activity"
	usecase "github.com/ndaDayo/devdev/usecase/activity"
)

type ActivityPresenter struct{}

func NewActivityPresenter() usecase.ActivityPresenter {
	return ActivityPresenter{}
}

func (p ActivityPresenter) Output(entity.Activity) usecase.ActivityOutput {
	return usecase.ActivityOutput{}
}
