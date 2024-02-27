package activity_uc

import (
	"errors"
	"fmt"

	"github.com/ndaDayo/devdev/api"
	entity "github.com/ndaDayo/devdev/entity/activity"
)

type GithubParams struct {
	Username string
	Repo     string
}

func (g GithubFetcher) FetchActivity(params interface{}) (*entity.Activity, error) {
	gp, ok := params.(*GithubParams)
	if !ok {
		return nil, errors.New("invalid parameters for GitHub fetcher")
	}

	prm := api.CommitsParam{
		Path: api.Path{
			Owner: gp.Username,
			Repo:  gp.Repo,
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

	totalLen := 0
	for _, c := range *cm {
		prm := api.CommitParam{
			Owner: gp.Username,
			Repo:  gp.Repo,
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

		totalLen += cm.Stats.Total
	}

	activity := &entity.Activity{
		Github: entity.Github{
			TotalLen: totalLen,
		},
	}

	return activity, nil
}
