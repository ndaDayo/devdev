package activity_uc

import (
	"time"

	"github.com/ndaDayo/devdev/entity"
)

type ActivityOptions struct {
	period ActivityPeriod
	source ActivitySource
}

type ActivityPeriod struct {
	start time.Time
	end   time.Time
}

type ActivitySource struct {
	github bool
	slack  bool
}

func Get(op ActivityOptions) *entity.Activity {
	return &entity.Activity{}
}
