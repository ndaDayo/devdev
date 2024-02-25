package activity_uc

import (
	"time"

	entity "github.com/ndaDayo/devdev/entity/activity"
)

type activityOptions func(*activitySource)

type activityPeriod struct {
	start time.Time
	end   time.Time
}

type activitySource struct {
	Github       bool
	Slack        bool
	githubParams *githubParams
	slackParams  *slackParams
}

type slackParams struct {
	Username string
}

func WithGithub(prm *githubParams) activityOptions {
	return func(as *activitySource) {
		as.Github = true
		as.githubParams = prm
	}
}

func WithSlack(prm *slackParams) activityOptions {
	return func(as *activitySource) {
		as.Slack = true
		as.slackParams = prm
	}
}

func NewActivityOptions(opts ...activityOptions) *activitySource {
	as := &activitySource{}

	for _, op := range opts {
		op(as)
	}

	return as
}

func Get(op activityOptions) *entity.Activity {
	return &entity.Activity{}
}
