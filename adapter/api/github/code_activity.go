package github

import (
	"context"

	github "github.com/ndaDayo/devdev/adapter/api/github/resources"
	entity "github.com/ndaDayo/devdev/domain/entity/activity"
	repository "github.com/ndaDayo/devdev/domain/repository/activity"
)

type CodeActivityFetcher struct{}

func NewCodeActivityFetcher() *CodeActivityFetcher {
	return &CodeActivityFetcher{}
}

func (c *CodeActivityFetcher) GetCodeActivity(ctx context.Context, criteria repository.Criteria) (entity.Code, error) {
	client := github.NewClient(github.WithToken())
	pr, err := pullRequest(ctx, client, criteria)
	if err != nil {
		return entity.Code{}, err
	}

	commits, err := commits(ctx, client, criteria)

	code := entity.Code{
		PullRequests: pr,
	}

	return code, nil
}

func pullRequest(ctx context.Context, c *github.Client, criteria repository.Criteria) ([]entity.PullRequest, error) {
	prs, err := c.PullRequests.Get(ctx, criteria)
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

func commits(ctx context.Context, c *github.Client, criteria repository.Criteria) ([]entity.Commit, error) {
	p := github.CommitsParam{
		Path: github.Path{
			Owner: criteria.Owner,
			Repo:  criteria.Repo,
		},
		Query: github.Query{
			Since: criteria.Since,
			Until: criteria.Until,
		},
	}
	cmts, err := c.Commits.Get(ctx, p)
	if err != nil {
		return nil, err
	}

	var entities []entity.Commit
	for _, cmt := range *cmts {
		commit := entity.Commit{
			Author: cmt.Commit.Author.Name,
		}

		entities = append(entities, commit)
	}

	return entities, nil
}
