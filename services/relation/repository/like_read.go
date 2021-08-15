package repository

import (
	"freeSociety/services/relation/instances"

	"github.com/mreza0100/golog"
	gorm "gorm.io/gorm"
)

type likes_read struct {
	lgr *golog.Core
	db  *gorm.DB
}

func (r *likes_read) IsLikedGroup(userId uint64, postIds []uint64) ([]uint64, error) {
	const query = `SELECT post_id FROM likes WHERE liker_id=? AND post_id IN(?)`
	params := []interface{}{userId, postIds}

	tx := r.db.Raw(query, params...)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := make([]uint64, 0)
	tx.Scan(&result)

	return result, nil
}

func (r *likes_read) CountLikes(postIds []uint64) (instances.CountResult, error) {
	const query = `SELECT COUNT(*), post_id FROM likes WHERE post_id IN(?) GROUP BY post_id`
	params := []interface{}{postIds}

	tx := r.db.Raw(query, params...)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := make(instances.CountResult, 0)
	tx.Scan(&result)

	return result, nil
}
