package activity_uc

import (
	"errors"
	"fmt"

	"github.com/ndaDayo/devdev/api"
	entity "github.com/ndaDayo/devdev/entity/activity"
)

type CodeParams struct {
	Username string
	Repo     string
	query    api.Query
}

func (c CodeActivityFetcher) FetchActivity(params interface{}) (*entity.Activity, error) {
	cp, ok := params.(CodeParams)
	if !ok {
		return nil, errors.New("invalid params type")
	}

	cm, err := getCommits(cp)
	if err != nil {
		fmt.Println("Error fetching commits:", err)
		return nil, err
	}

	for _, c := range *cm {
		prm := api.CommitParam{
			Owner: cp.Username,
			Repo:  cp.Repo,
			Ref:   c.SHA,
		}

		commit, err := api.GetResource(prm)
		if err != nil {
			fmt.Println("err", err)
		}

		cm, ok := commit.(*api.Commit)
		if !ok {
			return nil, errors.New("")
		}
	}

	activity := entity.NewActivity()

	return activity, nil
}

func getCommits(cp CodeParams) (*api.Commits, error) {
	prm := api.CommitsParam{
		Path: api.Path{
			Owner: cp.Username,
			Repo:  cp.Repo,
		},
		Query: api.Query{
			Since: "",
			Until: "",
		},
	}

	commits, err := api.GetResource(prm)
	if err != nil {
		fmt.Println("err", err)
	}

	cm, ok := commits.(*api.Commits)
	if !ok {
		return nil, errors.New("")
	}

	return cm, nil

}
