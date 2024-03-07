package github

import (
	github "github.com/ndaDayo/devdev/api/github/resources"
	entity "github.com/ndaDayo/devdev/domain/entity/activity"
)

type CodeActivityFetcher struct{}

func NewCodeActivityFetcher() *CodeActivityFetcher {
	return &CodeActivityFetcher{}
}

func (c *CodeActivityFetcher) GetCodeActivity(owner, repo, user string) (entity.Code, error) {
	client := github.NewClient(github.WithToken())

	pr, err := client.PullRequests.Get(owner, repo)
	if err != nil {
		return entity.Code{}, err
	}

	code := entity.Code{
		PullRequests: pr,
	}

	return code, nil
}
