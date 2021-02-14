package db

import (
	"fmt"

	"github.com/mreza0100/golog"
)

type DAOS struct {
	Lgr *golog.Core
}

func (this *DAOS) parseId(userId uint64) string {
	return fmt.Sprintf("feed_user_id_%v", userId)
}

func (this *DAOS) GetFeed(userId, offset, limit uint64) ([]uint64, error) {
	vals := db.LRange(this.parseId(userId), int64(offset), int64(limit))
	if vals.Err() != nil {
		this.Lgr.Debug.RedLog("error in GetFeed: ", vals.Err())
		return nil, vals.Err()
	}

	ids := make([]uint64, 0)
	vals.ScanSlice(&ids)

	return ids, nil
}

func (this *DAOS) SetPostOnFeeds(userId, postId uint64, followers []uint64) error {

	for _, f := range followers {
		db.LPush(this.parseId(f), postId)
	}

	return nil
}
