package repository

import (
	"context"

	entity "github.com/ndaDayo/devdev/domain/entity/activity"
)

type Criteria struct {
	Owner string
	Repo  string
	User  string
	Since string
	Until string
}

type Activity interface {
	GetCodeActivity(ctx context.Context, criteria Criteria) (entity.Code, error)
}
