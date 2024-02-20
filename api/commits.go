package api

import (
	"time"
)

type Commits []commit

type commit struct {
	SHA    string    `json:"sha"`
	Commit gitCommit `json:"commit"`
	Author user      `json:"author"`
}

type gitCommit struct {
	Author  author `json:"author"`
	Message string `json:"message"`
}

type author struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type user struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
}

type CommitsParams struct {
	Owner string
	Repo  string
}

func NewCommitsParams(owner, repo string) *CommitsParams {
	return &CommitsParams{
		Owner: owner,
		Repo:  repo,
	}
}
