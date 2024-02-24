package activity_uc

import (
	"time"

	"github.com/ndaDayo/devdev/entity"
)

type ActivityOptions struct {
	period ActivityPeriod
}

type ActivityPeriod struct {
	Start time.Time
	End   time.Time
}

func Get(p ActivityPeriod) *entity.Activity {

	return &entity.Activity{}
}
