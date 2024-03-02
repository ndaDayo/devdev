package api

import "context"

type PullRequestsParam struct {
	Path  Path
	Query Query
}

type Path struct {
	Owner string
	Repo  string
}

type Query struct {
}

type PullRequests interface {
	Get(ctx context.Context, p PullRequestsParam) (*PullRequests, *Response, error)
}
