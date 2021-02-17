package repository

import (
	pb "microServiceBoilerplate/proto/generated/post"

	"github.com/mreza0100/golog"
	"gorm.io/gorm"
)

type read struct {
	lgr *golog.Core
	db  *gorm.DB
}

func (r *read) GetPost(postIds []uint64) ([]*pb.Post, error) {
	query := `SELECT * FROM posts WHERE `
	params := make([]interface{}, 0, len(postIds))

	{
		for i := 0; i < len(postIds); i++ {
			query += "id=? OR "
			params = append(params, postIds[i])
		}
		// remove last "OR "
		query = query[:len(query)-3]
	}

	tx := r.db.Raw(query, params...)

	if tx.Error != nil {
		return []*pb.Post{}, tx.Error
	}

	result := make([]*pb.Post, 0, len(postIds))
	tx.Scan(&result)

	return result, nil
}
