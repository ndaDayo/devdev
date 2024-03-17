package entity

import "time"

type Code struct {
	PullRequests []PullRequest
	Commits      []Commit
}

type PullRequest struct {
	leadTime LeadTime
}

type Commit struct {
	Author string
}

func NewPullRequest(l LeadTime) PullRequest {
	return PullRequest{leadTime: l}
}

type LeadTime struct {
	time time.Duration
}

func NewLeadTime(createdAt, mergedAt time.Time) LeadTime {
	t := mergedAt.Sub(createdAt)
	return LeadTime{time: t}
}
