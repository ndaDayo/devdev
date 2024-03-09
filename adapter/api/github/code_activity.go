package github

import (
	github "github.com/ndaDayo/devdev/adapter/api/github/resources"
	entity "github.com/ndaDayo/devdev/domain/entity/activity"
)

type CodeActivityFetcher struct{}

func NewCodeActivityFetcher() *CodeActivityFetcher {
	return &CodeActivityFetcher{}
}

func (c *CodeActivityFetcher) GetCodeActivity(owner, repo, user string) (entity.Code, error) {
	client := github.NewClient(github.WithToken())
	pr, err := pullRequest(client, owner, repo)
	if err != nil {
		return entity.Code{}, err
	}

	code := entity.Code{
		PullRequests: pr,
	}

	return code, nil
}

func pullRequest(c *github.Client, owner, repo string) ([]entity.PullRequest, error) {
	prs, err := c.PullRequests.Get(owner, repo)
	if err != nil {
		return nil, err
	}

	var entities []entity.PullRequest
	for _, pr := range prs {
		lt := entity.NewLeadTime(pr.CreatedAt, pr.MergedAt)
		e := entity.NewPullRequest(lt)

		entities = append(entities, e)
	}

	return entities, nil
}
