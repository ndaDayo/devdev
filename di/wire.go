//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/ndaDayo/devdev/adapter/api/github"
	presenter "github.com/ndaDayo/devdev/adapter/presenter/activity"
	repository "github.com/ndaDayo/devdev/domain/repository/activity"
	usecase "github.com/ndaDayo/devdev/usecase/activity"
)

var activitySet = wire.NewSet(
	github.NewCodeActivityFetcher,
	wire.Bind(new(repository.Activity), new(*github.CodeActivityFetcher)),
	presenter.NewActivityPresenter,
	usecase.NewActivityUseCase,
)

func InitializeActivityUseCase() *usecase.ActivityUseCase {
	wire.Build(activitySet)
	return nil
}
