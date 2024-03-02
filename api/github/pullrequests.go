package github

import (
	"context"
	"fmt"

	entity "github.com/ndaDayo/devdev/entity/activity"
	activity_uc "github.com/ndaDayo/devdev/usecase/activity"
)

type PullRequestsService service

func (s *PullRequestsService) FetchPullRequests(ctx context.Context, p activity_uc.PullRequestsParams) (*entity.PullReq, *Response, error) {
	path := fmt.Sprintf("/repos/%v/%v/pulls", p.Owner, p.Repo)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return nil, nil, err
	}

	pr := new(entity.PullReq)
	resp, err := s.client.Do(ctx, req, pr)
	if err != nil {
		return nil, resp, err
	}

	return pr, resp, nil
}
