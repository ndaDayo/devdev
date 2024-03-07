//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	usecase "github.com/ndaDayo/devdev/usecase/activity"
)

func InitializeActivityUseCase() *usecase.ActivityUseCase {
	wire.Build(activitySet)
	return nil
}
