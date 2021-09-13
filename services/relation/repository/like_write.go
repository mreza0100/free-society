package repository

import (
	"errors"
	dbhelper "freeSociety/utils/dbHelper"

	"github.com/mreza0100/golog"
	gorm "gorm.io/gorm"
)

type likes_write struct {
	lgr *golog.Core
	db  *gorm.DB
}

func (w *likes_write) Like(likerId, ownerId uint64, postId string) (cc dbhelper.CommandController, err error) {
	cc, err = dbhelper.Transaction(w.db, func(tx *gorm.DB) error {
		const query = `INSERT INTO likes (liker_id, owner_id, post_id) VALUES (?, ?, ?)`
		params := []interface{}{likerId, ownerId, postId}

		tx = w.db.Exec(query, params...)
		return tx.Error
	})

	return cc, err
}

func (w *likes_write) UndoLike(likerId uint64, postId string) (cc dbhelper.CommandController, err error) {
	cc, err = dbhelper.Transaction(w.db, func(tx *gorm.DB) error {
		const query = `DELETE FROM likes WHERE liker_id=? AND post_id=?`
		params := []interface{}{likerId, postId}

		tx = w.db.Exec(query, params...)
		if tx.Error != nil {
			return tx.Error
		}

		if tx.RowsAffected != 1 {
			return errors.New("not found")
		}
		return tx.Error
	})

	return cc, err
}

func (w *likes_write) PurgeUserLikes(liker uint64) (cc dbhelper.CommandController, err error) {
	cc, err = dbhelper.Transaction(w.db, func(tx *gorm.DB) error {
		const query = `DELETE FROM likes WHERE liker_id=?`
		params := []interface{}{liker}

		tx = w.db.Exec(query, params...)

		return tx.Error
	})

	return cc, err
}
