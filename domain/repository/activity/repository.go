package repository

import (
	"context"

	entity "github.com/ndaDayo/devdev/domain/entity/activity"
)

type Activity interface {
	GetCodeActivity(ctx context.Context, owner, repo, user string) (entity.Code, error)
}
