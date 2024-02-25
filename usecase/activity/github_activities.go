package activity_uc

import (
	"context"
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

	client := api.NewClient(api.WithToken())
	ctx := context.Background()

	prm := api.CommitsParam{
		Owner: gp.Username,
		Repo:  gp.Repo,
	}

	commits, _, err := client.Commits.Get(ctx, prm)
	if err != nil {
		fmt.Println("err", err)
	}

	for _, c := range *commits {
		fmt.Println(c.Commit.Message)
	}

	return &entity.Activity{}, nil
}
