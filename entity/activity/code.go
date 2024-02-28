package entity

import "time"

type Code struct {
	PullReq PullReq
}

type PullReq struct {
	Count       int
	TimeToMerge time.Duration
	Comments    []Comment
}

type Comment struct {
	Count int
}
