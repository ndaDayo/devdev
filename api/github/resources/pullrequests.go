package github

import (
	"context"
	"fmt"
	"time"

	entity "github.com/ndaDayo/devdev/domain/entity/activity"
)

type PullRequests []PullRequest

type PullRequest struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ClosedAt  time.Time `json:"closed_at"`
	MergedAt  time.Time `json:"merged_at"`
}

type PullRequestsService service

func (s *PullRequestsService) Get() ([]entity.PullRequest, error) {
	path := fmt.Sprintf("/repos/%v/%v/pulls", p.Owner, p.Repo)
	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	prs := new(PullRequests)
	resp, err := s.client.Do(ctx, req, prs)
	if err != nil {
		return nil, err
	}

	// TODO delete
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var ets []entity.PullRequest
	for _, pr := range *prs {
		e := entity.PullRequest{
			CreatedAt: pr.CreatedAt,
		}
		ets = append(ets, e)
	}

	return ets, nil
}
