package pullrequests

import (
	"context"
	"fmt"

	api "github.com/ndaDayo/devdev/api/github"
)

type PullRequestsService struct {
	Client *api.Client
}

func NewPullRequestsService(c *api.Client) *PullRequestsService {
	return &PullRequestsService{Client: c}
}

func (s *PullRequestsService) Get(ctx context.Context, p api.PullRequestsParam) (*PullRequests, *api.Response, error) {
	path := fmt.Sprintf("/repos/%v/%v/pulls", p.Path.Owner, p.Path.Repo)

	req, err := s.Client.NewRequest("GET", path)
	if err != nil {
		return nil, nil, err
	}

	pr := new(PullRequests)
	resp, err := s.Client.Do(ctx, req, pr)
	if err != nil {
		return nil, resp, err
	}

	return pr, resp, nil
}
