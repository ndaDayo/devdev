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
	Start time.Time
	End   time.Time
}

type ActivitySource struct {
	github bool
	slack  bool
}

func Get(p ActivityPeriod, s ActivitySource) *entity.Activity {
	return &entity.Activity{}
}
