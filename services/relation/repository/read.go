package repository

import (
	"github.com/mreza0100/golog"
	gorm "gorm.io/gorm"
)

type read struct {
	lgr   *golog.Core
	db    *gorm.DB
	write *write
}

func (r *read) GetFollowers(userId uint64) []uint64 {
	const query = `SELECT follower from followers WHERE following=?`
	params := []interface{}{userId}

	tx := r.db.Raw(query, params...)

	data := make([]uint64, 0)
	tx.Scan(&data)

	return data
}

func (r *read) IsFollowing(follower, following uint64) bool {
	const query = `SELECT EXISTS(SELECT 1 FROM followers WHERE following=? AND follower=?)`
	params := []interface{}{following, follower}

	tx := r.db.Raw(query, params...)

	{
		if tx.Error != nil {
			return false
		}
	}

	data := &struct{ Exists bool }{}
	tx.Scan(data)

	return data.Exists
}
