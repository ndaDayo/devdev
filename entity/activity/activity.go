package entity

type Activity struct {
	GithubActivity GithubActivity
}

func NewActivity(g GithubActivity) *Activity {
	a := &Activity{
		GithubActivity: g,
	}

	return a
}
