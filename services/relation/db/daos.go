package db

import (
	"errors"

	"github.com/mreza0100/golog"
)

type DAOS struct {
	Lgr *golog.Core
}

func (this *DAOS) SetFollower(follower, following uint64) error {
	const query = `INSERT INTO followers (follower, following) VALUES (?, ?)`
	params := []interface{}{follower, following}

	tx := db.Exec(query, params...)

	{
		if tx.Error != nil {
			this.Lgr.RedLog(tx.Error.Error())
		}

		return tx.Error
	}
}

func (this *DAOS) RemoveFollower(follower, following uint64) error {
	const query = `DELETE FROM followers WHERE follower=? AND following=?`
	params := []interface{}{follower, following}

	tx := db.Exec(query, params...)

	{
		if tx.Error != nil {
			this.Lgr.RedLog(tx.Error.Error())
			return tx.Error
		}
		if tx.RowsAffected == 0 {
			return errors.New("not found any relation")
		}

		return nil
	}
}

func (this *DAOS) GetFollowers(userId uint64) []uint64 {
	const query = `SELECT follower from followers WHERE following=?`
	params := []interface{}{userId}

	tx := db.Raw(query, params...)

	data := make([]uint64, 0)

	tx.Scan(&data)

	return data
}
