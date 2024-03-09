package usecase

import "time"

type Output struct {
	CodeOutput CodeOutput
}

type CodeOutput struct {
	CodeActivity CodeActivity
}

type CodeActivity struct {
	LeadTime time.Duration
}
