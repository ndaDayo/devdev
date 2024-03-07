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
	code  *CodeInput
	slack *SlackInput
}

type Params interface{}

type CodeInput struct {
	Owner    string
	Repo     string
	Username string
}

func WithGithub(prm *CodeInput) func(*Input) {
	return func(opts *Input) {
		opts.Source.code = prm
	}
}

type SlackInput struct {
	Username string
}

func WithSlack(prm *SlackInput) func(*Input) {
	return func(opts *Input) {
		opts.Source.slack = prm
	}
}

func NewActivityOptionsInput(opts ...func(*Input)) *Input {
	options := &Input{}

	for _, opt := range opts {
		opt(options)
	}
	return options
}
