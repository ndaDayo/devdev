package entity

import "time"

type Code struct {
	Commit  Commit
	PullReq PullReq
}

type Commit struct {
	Count int
}

type PullReq struct {
	Count       int
	TimeToMerge time.Duration
	Comments    []Comment
}

type Comment struct {
	Count int
}
