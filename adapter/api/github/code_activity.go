package github

import (
	"context"

	github "github.com/ndaDayo/devdev/adapter/api/github/resources"
	entity "github.com/ndaDayo/devdev/domain/entity/activity"
)

type CodeActivityFetcher struct{}

func NewCodeActivityFetcher() *CodeActivityFetcher {
	return &CodeActivityFetcher{}
}

func (c *CodeActivityFetcher) GetCodeActivity(ctx context.Context, owner, repo, user string) (entity.Code, error) {
	client := github.NewClient(github.WithToken())
	pr, err := pullRequest(ctx, client, owner, repo)
	if err != nil {
		return entity.Code{}, err
	}

	code := entity.Code{
		PullRequests: pr,
	}

	return code, nil
}

func pullRequest(ctx context.Context, c *github.Client, owner, repo string) ([]entity.PullRequest, error) {
	prs, err := c.PullRequests.Get(ctx, owner, repo)
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
