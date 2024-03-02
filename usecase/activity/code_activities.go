package activity_uc

import (
	"errors"
	"fmt"

	entity "github.com/ndaDayo/devdev/entity/activity"
)

type CodeParams struct {
	Owner    string
	Repo     string
	Username string
}

func (c CodeActivityFetcher) FetchActivity(params interface{}) (*entity.Activity, error) {
	cp, ok := params.(CodeParams)
	if !ok {
		return nil, errors.New("invalid params type")
	}

	cm, err := ca.getCommits()
	if err != nil {
		fmt.Println("Error fetching commits:", err)
		return nil, err
	}

	for _, c := range *cm {
		_, err := ca.getCommit(c.SHA)
		if err != nil {
			return nil, errors.New("")
		}

	}

	activity := entity.NewActivity()

	return activity, nil
}
