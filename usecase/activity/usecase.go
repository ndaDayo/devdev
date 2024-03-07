package usecase

import (
	"errors"
	"fmt"
	"time"

	entity "github.com/ndaDayo/devdev/domain/entity/activity"
	repository "github.com/ndaDayo/devdev/domain/repository/activity"
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

type Params interface{}

type PullRequestsParams struct {
	Owner    string
	Repo     string
	Username string
}

type CodeParams struct {
	Owner    string
	Repo     string
	Username string
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

type ActivityUseCase struct {
	repository repository.Activity
}

func NewActivityUseCase(repo repository.Activity) *ActivityUseCase {
	return &ActivityUseCase{repository: repo}
}

func (u *ActivityUseCase) Get(opts ...func(*ActivityOptions)) (*entity.Activity, error) {
	options := NewActivityOptions(opts...)
	if options.Source.codeParams != nil {
		activity, err := u.FetchActivity(options.Source.codeParams)

		if err != nil {
			return nil, fmt.Errorf("failed to fetch GitHub activity: %w", err)
		}

		return activity, nil
	}
	return nil, nil
}

func (u *ActivityUseCase) FetchActivity(params interface{}) (*entity.Activity, error) {
	cp, ok := params.(*CodeParams)
	if !ok {
		return nil, errors.New("invalid params type")
	}

	code, err := u.repository.GetCodeActivity(cp.Owner, cp.Repo, cp.Username)
	if err != nil {
		return nil, err
	}

	ac := &entity.Activity{
		CodeActivity: code,
	}

	return ac, nil
}
