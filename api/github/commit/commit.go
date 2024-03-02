package commit

import (
	"context"
	"fmt"
)

type CommitService Service

type CommitParam struct {
	Owner string
	Repo  string
	Ref   string
}

func (s *CommitService) Get(ctx context.Context, p CommitParam) (*Commit, *Response, error) {
	path := fmt.Sprintf("/repos/%v/%v/commits/%v", p.Owner, p.Repo, p.Ref)
	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return nil, nil, err
	}

	commit := new(Commit)
	resp, err := s.client.Do(ctx, req, commit)
	if err != nil {
		return nil, resp, err
	}

	return commit, resp, nil
}
