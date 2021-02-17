package repository

import (
	"github.com/mreza0100/golog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	gorm "gorm.io/gorm"
)

type write struct {
	lgr  *golog.Core
	db   *gorm.DB
	read *read
}

func (w *write) SetFollower(follower, following uint64) error {
	{
		if w.read.IsFollowing(follower, following) {
			return status.Error(codes.AlreadyExists, "already following")
		}
	}
	{
		const query = `INSERT INTO followers (follower, following) VALUES (?, ?)`
		params := []interface{}{follower, following}

		tx := w.db.Exec(query, params...)

		{
			if tx.Error != nil {
				w.lgr.Debug.RedLog(tx.Error.Error())
			}
		}

		return tx.Error
	}
}

func (w *write) RemoveFollow(follower, following uint64) error {
	const query = `DELETE FROM followers WHERE follower=? AND following=?`
	params := []interface{}{follower, following}

	tx := w.db.Exec(query, params...)

	{
		if tx.Error != nil {
			w.lgr.RedLog(tx.Error.Error())
			return tx.Error
		}
		if tx.RowsAffected == 0 {
			return status.Error(codes.NotFound, "not found any relation")
		}

		return nil
	}
}

func (w *write) DeleteAllRelations(userId uint64) error {
	const qeury = `DELETE FROM followers WHERE follower=? OR following=?`
	params := []interface{}{userId, userId}

	tx := w.db.Exec(qeury, params...)

	return tx.Error
}
