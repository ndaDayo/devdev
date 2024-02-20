package api

import (
	"context"
	"fmt"
	"time"
)

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

type CommitsParams struct {
	Owner string
	Repo  string
}

func NewCommitsParams(owner, repo string) *CommitsParams {
	return &CommitsParams{
		Owner: owner,
		Repo:  repo,
	}
}

type CommitsService service

func (s *CommitsService) Get(ctx context.Context, owner, repo string) (*Commits, *Response, error) {
	u := fmt.Sprintf("/repos/%v/%v/commits", owner, repo)
	req, err := s.client.NewRequest("GET", u)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, resp, err
	}

}
