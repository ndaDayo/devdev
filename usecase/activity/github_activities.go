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
		Owner: gp.Username,
		Repo:  gp.Repo,
	}

	commits, err := api.GetResource(prm)
	if err != nil {
		fmt.Println("err", err)
	}

	cm, ok := commits.(*api.Commits)
	if !ok {
		return nil, errors.New("")
	}

	for _, c := range *cm {
		fmt.Println(c.Commit.Message)
	}

	return &entity.Activity{}, nil
}
