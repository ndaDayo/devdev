package usecase

import (
	"errors"
	"fmt"

	entity "github.com/ndaDayo/devdev/domain/entity/activity"
	repository "github.com/ndaDayo/devdev/domain/repository/activity"
)

type ActivityUseCase struct {
	repository repository.Activity
}

func NewActivityUseCase(repo repository.Activity) *ActivityUseCase {
	return &ActivityUseCase{repository: repo}
}

func (u *ActivityUseCase) Run(opts ...func(*Input)) (*entity.Activity, error) {
	options := NewActivityOptionsInput(opts...)
	if options.Source.Code != nil {
		code, err := u.fetchCodeActivity(options.Source.Code)

		if err != nil {
			return nil, fmt.Errorf("failed to fetch GitHub activity: %w", err)
		}

		activity := &entity.Activity{
			CodeActivity: code,
		}

		return activity, nil
	}
	return nil, nil
}

func (u *ActivityUseCase) fetchCodeActivity(params interface{}) (entity.Code, error) {
	cp, ok := params.(*CodeInput)
	if !ok {
		return entity.Code{}, errors.New("invalid params type")
	}

	code, err := u.repository.GetCodeActivity(cp.Owner, cp.Repo, cp.Username)
	if err != nil {
		return entity.Code{}, err
	}

	return code, nil
}
