//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/ndaDayo/devdev/infrastructure/api/github"
	activity_uc "github.com/ndaDayo/devdev/usecase/activity"
)

func InitializeResourceFetcher() activity_uc.ResourceFetcher {
	wire.Build(
		github.NewGitHubResourceFetcher,
		wire.Bind(new(activity_uc.ResourceFetcher), new(*github.GitHubResourceFetcher)),
	)
	return nil
}
