package entity

import "time"

type Code struct {
	PullRequests []PullRequest
}

type PullRequest struct {
	CreatedAt time.Time
}
