package github

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type PullRequests []PullRequest

type PullRequest struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ClosedAt  time.Time `json:"closed_at"`
	MergedAt  time.Time `json:"merged_at"`
}

type PullRequestsService service

type PullsParam struct {
	Owner   string
	Repo    string
	State   string
	PerPage string
	Since   string
	Until   string
}

func (s *PullRequestsService) Get(ctx context.Context, param PullsParam) ([]PullRequest, error) {
	req, err := s.client.NewRequest("GET", s.client.Payload(param))
	if err != nil {
		return nil, fmt.Errorf("failed to construct NewRequest: %w", err)
	}

	prs := new(PullRequests)
	resp, err := s.client.Do(ctx, req, prs)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pullrequests: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var p []PullRequest
	for _, pr := range *prs {
		p = append(p, pr)
	}

	slog.Info("success fetch PullRequest", "count", len(p))

	return p, nil
}
