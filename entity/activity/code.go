package entity

import "time"

type Code struct {
	PullReq PullReq
}

type PullReq struct {
	Count       int
	TimeToMerge time.Duration
}
