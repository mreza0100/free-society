package repository

import (
	"github.com/mreza0100/golog"
	gorm "gorm.io/gorm"
)

type followers_read struct {
	lgr *golog.Core
	db  *gorm.DB
}

func (r *followers_read) GetFollowers(userId uint64) []uint64 {
	const query = `SELECT follower from followers WHERE following=?`
	params := []interface{}{userId}

	tx := r.db.Raw(query, params...)

	data := make([]uint64, 0)
	tx.Scan(&data)

	return data
}

func (r *followers_read) IsFollowing(follower, following uint64) bool {
	const query = `SELECT EXISTS(SELECT 1 FROM followers WHERE following=? AND follower=?)`
	params := []interface{}{following, follower}

	tx := r.db.Raw(query, params...)

	if tx.Error != nil {
		return false
	}

	data := struct{ Exists bool }{}
	tx.Scan(&data)

	return data.Exists
}

func (r *followers_read) IsFollowingGroup(follower uint64, followings []uint64) (map[uint64]interface{}, error) {
	const query = `SELECT following FROM followers WHERE follower=? AND following IN(?)`
	params := []interface{}{follower, followings}

	{
		var (
			rawResult []uint64
			result    map[uint64]interface{}
			tx        *gorm.DB
		)

		{
			tx = r.db.Raw(query, params...)
			if tx.Error != nil {
				return nil, tx.Error
			}
		}
		{
			rawResult = make([]uint64, 0)
			tx.Scan(&rawResult)
		}
		{
			result = make(map[uint64]interface{}, len(rawResult))
			for _, i := range rawResult {
				result[i] = struct{}{}
			}
		}

		return result, tx.Error
	}
}
