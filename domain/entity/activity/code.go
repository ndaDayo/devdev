package entity

import "time"

type Code struct {
	PullRequests []PullRequest
}

type PullRequest struct {
	LeadTime time.Duration
}

func GetLeadTime(createdAt, mergedAt time.Time) time.Duration {
	return mergedAt.Sub(createdAt)
}
