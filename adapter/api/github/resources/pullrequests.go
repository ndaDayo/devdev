package github

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	repository "github.com/ndaDayo/devdev/domain/repository/activity"
)

type PullRequests []PullRequest

type PullRequest struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ClosedAt  time.Time `json:"closed_at"`
	MergedAt  time.Time `json:"merged_at"`
}

type PullRequestsService service

func (s *PullRequestsService) Get(ctx context.Context, criteria repository.Criteria) ([]PullRequest, error) {
	path := fmt.Sprintf("/repos/%v/%v/pulls?state=all", criteria.Owner, criteria.Repo)

	query := url.Values{}
	query.Add("since", criteria.Since)
	query.Add("until", criteria.Until)

	endpoint := fmt.Sprintf("%s?%s", path, query.Encode())
	req, err := s.client.NewRequest("GET", endpoint)

	if err != nil {
		return nil, fmt.Errorf("failed to construct NewRequest: %w", err)
	}

	prs := new(PullRequests)
	resp, err := s.client.Do(ctx, req, prs)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pullrequests: %w", err)
	}

	slog.Info("success fetch PullRequest", "count", len(*prs))

	if resp.StatusCode != http.StatusOK {
		slog.Error("failed fetch PullRequests", "statusCode", resp.StatusCode)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var p []PullRequest
	for _, pr := range *prs {
		p = append(p, pr)
	}

	return p, nil
}
