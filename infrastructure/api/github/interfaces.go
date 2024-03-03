package github

type PullRequestsParam struct {
	Path  Path
	Query Query
}

type Path struct {
	Owner string
	Repo  string
}

type Query struct {
}
