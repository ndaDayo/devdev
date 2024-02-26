package activity_uc

import (
	"time"

	entity "github.com/ndaDayo/devdev/entity/activity"
)

type activityOptions struct {
	Source activitySource
	Period activityPeriod
}

type activityPeriod struct {
	start time.Time
	end   time.Time
}

type activitySource struct {
	githubParams *GithubParams
	slackParams  *SlackParams
}

type SlackParams struct {
	Username string
}

func WithGithub(prm *GithubParams) func(*activityOptions) {
	return func(opts *activityOptions) {
		opts.Source.githubParams = prm
	}
}

func WithSlack(prm *SlackParams) func(*activityOptions) {
	return func(opts *activityOptions) {
		opts.Source.slackParams = prm
	}
}

func NewActivityOptions() *activityOptions {
	return &activityOptions{
		Source: activitySource{},
		Period: activityPeriod{},
	}
}

func Get(op activityOptions) *entity.Activity {
	return &entity.Activity{}
}
