package github

import (
	"context"
	"fmt"
)

type PullRequestsService service

func (s *PullRequestsService) Get(ctx context.Context, p PullRequestsParam) (*PullRequests, *Response, error) {
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
