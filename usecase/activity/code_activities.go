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

type CodeActivity struct {
	Owner string
	Repo  string
}

func NewCodeActivity(owner, repo string) *CodeActivity {
	return &CodeActivity{Owner: owner, Repo: repo}
}

func (c CodeActivityFetcher) FetchActivity(params interface{}) (*entity.Activity, error) {
	cp, ok := params.(CodeParams)
	if !ok {
		return nil, errors.New("invalid params type")
	}

	ca := NewCodeActivity(cp.Username, cp.Repo)
	cm, err := ca.getCommits()
	if err != nil {
		fmt.Println("Error fetching commits:", err)
		return nil, err
	}

	for _, c := range *cm {
	}

	activity := entity.NewActivity()

	return activity, nil
}

func (ca *CodeActivity) getCommits() (*api.Commits, error) {
	prm := api.CommitsParam{
		Path: api.Path{
			Owner: ca.Owner,
			Repo:  ca.Repo,
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

func (ca *CodeActivity) getCommit(ref string) (*api.Commit, error) {
	prm := api.CommitParam{
		Owner: ca.Owner,
		Repo:  ca.Repo,
		Ref:   ref,
	}

	commit, err := api.GetResource(prm)
	if err != nil {
		fmt.Println("err", err)
	}

	cm, ok := commit.(*api.Commit)
	if !ok {
		return nil, errors.New("")
	}

	return cm, nil
}
