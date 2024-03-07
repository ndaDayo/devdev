package repository

import entity "github.com/ndaDayo/devdev/domain/entity/activity"

type Activity interface {
	GetCodeActivity(owner, repo, user string) (entity.Code, error)
}
