package usecase

import (
	"context"
	"errors"
	"fmt"

	entity "github.com/ndaDayo/devdev/domain/entity/activity"
	repository "github.com/ndaDayo/devdev/domain/repository/activity"
)

type (
	ActivityUseCase struct {
		repository repository.Activity
		presenter  ActivityPresenter
	}

	ActivityPresenter interface {
		Output(entity.Activity) ActivityOutput
	}

	ActivityOutput struct {
		OutPut Output
	}
)

func NewActivityUseCase(
	repo repository.Activity,
	presenter ActivityPresenter,
) *ActivityUseCase {
	return &ActivityUseCase{
		repository: repo,
		presenter:  presenter,
	}
}

func (u *ActivityUseCase) Run(opts ...func(*Input)) (ActivityOutput, error) {
	options := NewActivityOptionsInput(opts...)
	activity := entity.Activity{}

	if options.Source.Code != nil {
		code, err := u.fetchCodeActivity(options.Source.Code)

		if err != nil {
			return ActivityOutput{}, fmt.Errorf("failed to fetch GitHub activity: %w", err)
		}

		activity = entity.Activity{
			CodeActivity: code,
		}

		return u.presenter.Output(activity), nil
	}

	return u.presenter.Output(activity), nil
}

func (u *ActivityUseCase) fetchCodeActivity(params interface{}) (entity.Code, error) {
	cp, ok := params.(*CodeInput)
	if !ok {
		return entity.Code{}, errors.New("invalid params type")
	}

	ctx := context.Background()
	c := repository.Criteria{
		Owner: cp.Owner,
		Repo:  cp.Repo,
		User:  cp.Username,
	}
	code, err := u.repository.GetCodeActivity(ctx, c)

	if err != nil {
		return entity.Code{}, err
	}

	return code, nil
}
