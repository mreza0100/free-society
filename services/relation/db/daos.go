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

	tx := DB.Exec(query, params...)

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

	tx := DB.Exec(query, params...)

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
