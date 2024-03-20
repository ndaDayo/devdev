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

	p, err := filter(*prs, param.Since, param.Until)
	if err != nil {
		return nil, fmt.Errorf("failed to filter pullrequests: %w", err)
	}

	slog.Info("success fetch PullRequest", "count", len(p))

	return p, nil
}

func filter(prs PullRequests, since, until string) ([]PullRequest, error) {
	var filteredPRs []PullRequest

	sinceTime, err := time.Parse(time.RFC3339, since)
	if err != nil {
		return nil, fmt.Errorf("invalid 'since' format: %w", err)
	}

	untilTime, err := time.Parse(time.RFC3339, until)
	if err != nil {
		return nil, fmt.Errorf("invalid 'until' format: %w", err)
	}

	for _, pr := range prs {
		if pr.CreatedAt.After(sinceTime) && pr.CreatedAt.Before(untilTime) {
			filteredPRs = append(filteredPRs, pr)
		}
	}

	return filteredPRs, nil
}
