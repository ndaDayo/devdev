// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/ndaDayo/devdev/infrastructure/api/github"
	"github.com/ndaDayo/devdev/usecase/activity"
)

// Injectors from wire.go:

func InitializeResourceFetcher() activity_uc.ResourceFetcher {
	gitHubResourceFetcher := github.NewGitHubResourceFetcher()
	return gitHubResourceFetcher
}
