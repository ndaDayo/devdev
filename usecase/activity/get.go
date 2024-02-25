package activity_uc

import (
	"time"

	"github.com/ndaDayo/devdev/entity"
)

type ActivityOptions struct {
	ActivitySource
}

type ActivityPeriod struct {
	start time.Time
	end   time.Time
}

type ActivitySource struct {
	Github       bool
	Slack        bool
	GithubParams *GithubParams
}

func Get(op ActivityOptions) *entity.Activity {
	return &entity.Activity{}
}
