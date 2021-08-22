package repository

import (
	"errors"
	"freeSociety/configs"
	"strings"

	"github.com/mreza0100/golog"
	"gorm.io/gorm"
)

type write struct {
	lgr *golog.Core
	db  *gorm.DB
}

func (w *write) NewPost(title, body string, userId uint64, picturesName []string) (uint64, error) {
	const query = `INSERT INTO posts (title, body, owner_id, pictures_name) VALUES (?, ?, ?, ? ) RETURNING id`
	params := []interface{}{title, body, userId}

	if len(picturesName) > 0 {
		params = append(params, strings.Join(picturesName, ","))
	} else {
		params = append(params, nil)
	}

	for i := 0; i < configs.Max_picture_per_post; i++ {
		if len(picturesName) > i {
			params = append(params, picturesName[i])
		} else {
			params = append(params, nil)
		}
	}

	tx := w.db.Raw(query, params...)
	if tx.Error != nil {
		return 0, tx.Error
	}

	result := struct{ Id uint64 }{}
	tx.Scan(&result)

	return result.Id, nil
}

func (w *write) DeletePost(postId, userId uint64) (picturesName string, err error) {
	const query = `DELETE FROM posts WHERE id=? AND owner_id=? RETURNING pictures_name`
	params := []interface{}{postId, userId}

	tx := w.db.Raw(query, params...)
	if tx.Error != nil {
		return "", tx.Error
	}

	result := struct{ PicturesName string }{}
	tx.Scan(&result)

	if tx.RowsAffected != 1 {
		return "", errors.New("Cant find post")
	}

	return result.PicturesName, tx.Error
}

func (w *write) DeleteUserPosts(userId uint64) ([]struct{ PicturesName string }, error) {
	const query = `DELETE FROM posts WHERE owner_id= ? RETURNING pictures_name`
	params := []interface{}{userId}

	tx := w.db.Raw(query, params...)
	if tx.Error != nil {
		return []struct{ PicturesName string }{}, tx.Error
	}
	result := make([]struct{ PicturesName string }, 0)
	tx.Scan(&result)

	return result, tx.Error
}
