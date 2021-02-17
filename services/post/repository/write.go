package repository

import (
	"errors"

	"github.com/mreza0100/golog"
	"gorm.io/gorm"
)

type write struct {
	lgr *golog.Core
	db  *gorm.DB
}

func (w *write) NewPost(title, body string, userId uint64) (uint64, error) {
	const query = `INSERT INTO posts (title, body, owner_id) VALUES (?, ?, ?) RETURNING id`
	params := []interface{}{title, body, userId}

	tx := w.db.Raw(query, params...)
	if tx.Error != nil {
		return 0, tx.Error
	}

	result := struct{ Id uint64 }{}
	tx.Scan(&result)

	return result.Id, nil
}

func (w *write) DeletePost(postId, userId uint64) error {
	const query = `DELETE FROM posts WHERE id=? AND owner_id=?`
	params := []interface{}{postId, userId}

	tx := w.db.Exec(query, params...)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected != 1 {
		return errors.New("Cant find post")
	}

	return nil
}

func (w *write) DeleteUserPosts(userId uint64) error {
	const query = `DELETE FROM posts WHERE owner_id=?`
	params := []interface{}{userId}

	tx := w.db.Exec(query, params...)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		w.lgr.Debug.CyanLog("tx.RowsAffected: ", tx.RowsAffected)
		w.lgr.Debug.InfoLog("user dont have any posts")
		return errors.New("user dont have any posts")
	}

	return nil
}
