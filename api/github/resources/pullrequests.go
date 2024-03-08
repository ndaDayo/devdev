package github

import (
	"context"
	"fmt"
	"net/http"
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

func (s *PullRequestsService) Get(owner, repo string) ([]entity.PullRequest, error) {
	path := fmt.Sprintf("/repos/%v/%v/pulls?state=all", owner, repo)
	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	prs := new(PullRequests)
	resp, err := s.client.Do(ctx, req, prs)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pullrequests: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var entities []entity.PullRequest
	for _, pr := range *prs {
		lt := entity.NewLeadTime(pr.CreatedAt, pr.MergedAt)
		e := entity.NewPullRequest(lt)

		entities = append(entities, e)
	}

	return entities, nil
}
