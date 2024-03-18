package entity

import "time"

type Activity struct {
	CodeActivity Code
}

type Code struct {
	PullRequests []PullRequest
	Commits      []Commit
}

type PullRequest struct {
	leadTime leadTime
}

type Commit struct {
	Author string
}

func NewPullRequest(l leadTime) PullRequest {
	return PullRequest{leadTime: l}
}

type leadTime struct {
	time time.Duration
}

func NewLeadTime(createdAt, mergedAt time.Time) leadTime {
	t := mergedAt.Sub(createdAt)
	return leadTime{time: t}
}

func (p PullRequest) GetLeadTime() string {
	return p.leadTime.time.String()
}
