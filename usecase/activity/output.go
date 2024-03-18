package usecase

type ActivityOutput struct {
	PullRequests []PullRequest
}

type PullRequest struct {
	LeadTime string
}
