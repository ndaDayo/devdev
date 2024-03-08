package entity

import "time"

type Code struct {
	PullRequests []PullRequest
}

type PullRequest struct {
	CreatedAt time.Time
	MergedAt  time.Time
}

func (pr *PullRequest) LeadTime() time.Duration {
	return pr.MergedAt.Sub(pr.CreatedAt)
}
