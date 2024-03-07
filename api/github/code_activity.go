package github

import entity "github.com/ndaDayo/devdev/domain/entity/activity"

type CodeActivityFetcher struct{}

func NewCodeActivityFetcher() *CodeActivityFetcher {
	return &CodeActivityFetcher{}
}

func (*CodeActivityFetcher) GetCodeActivity(owner, repo, user string) (entity.Code, error) {
	return entity.Code{}, nil
}
