package get

import (
	"fmt"
	"time"

	entity "github.com/ndaDayo/devdev/entity/activity"
)

type ActivityOptions struct {
	Source activitySource
	Period activityPeriod
}

type activityPeriod struct {
	start time.Time
	end   time.Time
}

type activitySource struct {
	codeParams  *CodeParams
	slackParams *SlackParams
}

type SlackParams struct {
	Username string
}

func WithGithub(prm *CodeParams) func(*ActivityOptions) {
	return func(opts *ActivityOptions) {
		opts.Source.codeParams = prm
	}
}

func WithSlack(prm *SlackParams) func(*ActivityOptions) {
	return func(opts *ActivityOptions) {
		opts.Source.slackParams = prm
	}
}

func NewActivityOptions(opts ...func(*ActivityOptions)) *ActivityOptions {
	options := &ActivityOptions{}

	for _, opt := range opts {
		opt(options)
	}
	return options
}

func Get(opts ...func(*ActivityOptions)) (*entity.Activity, error) {
	options := NewActivityOptions(opts...)
	if options.Source.codeParams != nil {
		gf := CodeActivityFetcher{}
		activity, err := gf.FetchActivity(options.Source.codeParams)

		if err != nil {
			return nil, fmt.Errorf("failed to fetch GitHub activity: %w", err)
		}

		return activity, nil
	}
	return nil, nil
}
