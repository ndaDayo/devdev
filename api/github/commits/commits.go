package commits

import (
	"context"
	"fmt"
	"net/url"
)

type CommitsService Service

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

func (s *CommitsService) Get(ctx context.Context, p CommitsParam) (*Commits, *Response, error) {
	path := fmt.Sprintf("/repos/%v/%v/commits", p.Path.Owner, p.Path.Repo)
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
