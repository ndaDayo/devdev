package github

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

type CommitsService service

type CommitsParam struct {
	Path  Path
	Query Query
}

type Path struct {
	Owner string
	Repo  string
}

type Query struct {
	Sha       string
	Path      string
	Author    string
	Committer string
	Since     string
	Until     string
	PerPage   int
	Page      int
}

type Commits []CommitDetail

type CommitDetail struct {
	SHA    string    `json:"sha"`
	Commit gitCommit `json:"commit"`
	Author user      `json:"author"`
}

type gitCommit struct {
	Author  author `json:"author"`
	Message string `json:"message"`
}

type author struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type user struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
}

func (s *CommitsService) Get(ctx context.Context, p CommitsParam) (*Commits, error) {
	path := fmt.Sprintf("/repos/%v/%v/commits", p.Path.Owner, p.Path.Repo)
	query := url.Values{}

	query.Add("since", p.Query.Since)
	query.Add("until", p.Query.Until)

	endpoint := fmt.Sprintf("%s?%s", path, query.Encode())
	req, err := s.client.NewRequest("GET", endpoint)
	if err != nil {
		return nil, err
	}

	commits := new(Commits)
	resp, err := s.client.Do(ctx, req, commits)
	if err != nil {
		return nil, err
	}

	slog.Info("success Commits", "count", len(*commits))

	if resp.StatusCode != http.StatusOK {
		slog.Error("failed fetch Commits", "statusCode", resp.StatusCode)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return commits, nil
}
