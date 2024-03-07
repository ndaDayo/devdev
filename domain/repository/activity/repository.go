package repository

import entity "github.com/ndaDayo/devdev/entity/activity"

type Activity interface {
	GetCodeActivity(owner, repo, user string) (*entity.Code, error)
}
