package repository

import (
	"errors"

	"github.com/mreza0100/golog"
	gorm "gorm.io/gorm"
)

type likes_write struct {
	lgr *golog.Core
	db  *gorm.DB
}

func (w *likes_write) Like(likerId, ownerId, postId uint64) error {
	const query = `INSERT INTO likes (liker_id, owner_id, post_id) VALUES (?, ?, ?)`
	params := []interface{}{likerId, ownerId, postId}

	tx := w.db.Exec(query, params...)

	return tx.Error
}

func (w *likes_write) UndoLike(likerId, postId uint64) error {
	const query = `DELETE FROM likes WHERE liker_id=? AND post_id=?`
	params := []interface{}{likerId, postId}

	tx := w.db.Exec(query, params...)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected != 1 {
		return errors.New("not found")
	}

	return nil
}
