package pullrequests

import (
	"context"
	"fmt"

	api "github.com/ndaDayo/devdev/api/github"
)

type PullRequestsParam struct {
	Path  Path
	Query Query
}

type Path struct {
	Owner string
	Repo  string
}

type Query struct {
}

func (s *api.PullRequestsService) Get(ctx context.Context, p PullRequestsParam) (*PullRequests, *api.Response, error) {
	path := fmt.Sprintf("/repos/%v/%v/pulls", p.Path.Owner, p.Path.Repo)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return nil, nil, err
	}

	pr := new(PullRequests)
	resp, err := s.client.Do(ctx, req, pr)
	if err != nil {
		return nil, resp, err
	}

	return pr, resp, nil
}
