package usecase

import "time"

type Input struct {
	Source activitySource
	Period activityPeriod
}

type activityPeriod struct {
	start time.Time
	end   time.Time
}

type activitySource struct {
	Code  *CodeInput
	Slack *slackInput
}

type Params interface{}

type CodeInput struct {
	Owner    string
	Repo     string
	Username string
	Since    string
	Until    string
}

func WithGithub(prm *CodeInput) func(*Input) {
	return func(opts *Input) {
		opts.Source.Code = prm
	}
}

type slackInput struct {
	Username string
}

func WithSlack(prm *slackInput) func(*Input) {
	return func(opts *Input) {
		opts.Source.Slack = prm
	}
}

func NewActivityOptionsInput(opts ...func(*Input)) *Input {
	options := &Input{}
	for _, opt := range opts {
		opt(options)
	}
	return options
}
