package activity_uc

import (
	"errors"

	entity "github.com/ndaDayo/devdev/entity/activity"
)

type CodeParams struct {
	Owner    string
	Repo     string
	Username string
}

func (c CodeActivityFetcher) FetchActivity(params interface{}) (*entity.Activity, error) {
	cp, ok := params.(*CodeParams)
	if !ok {
		return nil, errors.New("invalid params type")
	}

	p := PullRequestsParams{
		Owner:    cp.Owner,
		Repo:     cp.Repo,
		Username: cp.Username,
	}

	pr, err := c.PullRequestsFetcher.FetchPullRequests(p)
	if err != nil {
		return nil, err
	}

	ac := &entity.Activity{
		CodeActivity: entity.Code{
			PullRequests: pr,
		},
	}

	return ac, nil
}
