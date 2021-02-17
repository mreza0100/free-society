package repository

import (
	"fmt"

	"github.com/mreza0100/golog"
)

type helpers struct {
	lgr *golog.Core
}

func (this *helpers) parseId(userId uint64) string {
	return fmt.Sprintf("feed_user_id_%v", userId)
}
