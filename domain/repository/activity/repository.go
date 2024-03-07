package repository

import entity "github.com/ndaDayo/devdev/entity/activity"

type Activity interface {
	Get(owner, repo, user string) (*entity.Activity, error)
}
