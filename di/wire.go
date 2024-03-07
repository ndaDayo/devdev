//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/ndaDayo/devdev/api/github"
	repository "github.com/ndaDayo/devdev/domain/repository/activity"
	usecase "github.com/ndaDayo/devdev/usecase/activity"
)

func ProvideActivity(repo *github.CodeActivityFetcher) repository.Activity {
	return repo
}

var activitySet = wire.NewSet(
	github.NewCodeActivityFetcher,
	wire.Bind(new(repository.Activity), new(*github.CodeActivityFetcher)),
	usecase.NewActivityUseCase,
)

func InitializeActivityUseCase() *usecase.ActivityUseCase {
	wire.Build(activitySet)
	return nil
}
