package github

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

type CommitsService service

type CommitsParam struct {
	path  Path
	query Query
}

type path struct {
	Owner string
	Repo  string
}

type query struct {
	Sha       string
	Path      string
	Author    string
	Committer string
	Since     string
	Until     string
	PerPage   int
	Page      int
}

type Commits []commit

type commit struct {
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

func (s *CommitsService) Get(ctx context.Context, p CommitsParam) (*Commits, *Response, error) {
	path := fmt.Sprintf("/repos/%v/%v/commits", p.path.Owner, p.path.Repo)
	query := url.Values{}

	query.Add("since", "2023-08-31T00:00:00Z")
	query.Add("until", "2023-09-04T23:59:59Z")

	endpoint := fmt.Sprintf("%s?%s", path, query.Encode())
	req, err := s.client.NewRequest("GET", endpoint)
	if err != nil {
		return nil, nil, err
	}

	commits := new(Commits)
	resp, err := s.client.Do(ctx, req, commits)
	if err != nil {
		return nil, resp, err
	}

	return commits, resp, nil
}