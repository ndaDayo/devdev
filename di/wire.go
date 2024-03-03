//go:build wireinject
// +build wireinject

package di

import (
	activity_uc "github.com/ndaDayo/devdev/usecase/activity"
)

func Initialize() *activity_uc.CodeActivityFetcher {
	return nil
}
