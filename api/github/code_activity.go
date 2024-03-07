package github

import (
	repository "github.com/ndaDayo/devdev/domain/repository/activity"
	entity "github.com/ndaDayo/devdev/entity/activity"
)

type CodeActivityFetcher struct {
	Repository repository.Activity
}

func NewCodeActivityFetcher(repository repository.Activity) *CodeActivityFetcher {
	return &CodeActivityFetcher{Repository: repository}
}

func (*CodeActivityFetcher) GetCodeActivity(owner, repo, user string) (*entity.Code, error) {
	return nil, nil
}
