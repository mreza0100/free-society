package db

import (
	"errors"
	pb "microServiceBoilerplate/proto/generated/post"

	"github.com/mreza0100/golog"
)

type DAOS struct {
	Lgr *golog.Core
}

func (this *DAOS) NewPost(title, body string, userId uint64) (uint64, error) {
	const query = `INSERT INTO posts (title, body, owner_id) VALUES (?, ?, ?) RETURNING id`
	params := []interface{}{title, body, userId}

	tx := DB.Raw(query, params...)
	if tx.Error != nil {
		return 0, tx.Error
	}

	result := struct{ Id uint64 }{}
	tx.Scan(&result)

	return result.Id, nil
}

func (this *DAOS) GetPost(postIds []uint64) ([]*pb.GetPostResponseInner, error) {
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

	tx := DB.Raw(query, params...)

	if tx.Error != nil {
		return []*pb.GetPostResponseInner{}, tx.Error
	}

	result := make([]*pb.GetPostResponseInner, 0, len(postIds))
	tx.Scan(&result)

	return result, nil
}

func (this *DAOS) DeletePost(postId, userId uint64) error {
	const query = `DELETE FROM posts WHERE id=? AND owner_id=?`
	params := []interface{}{postId, userId}

	tx := DB.Exec(query, params...)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected != 1 {
		return errors.New("Cant find post")
	}

	return nil
}

func (this *DAOS) DeleteUserPosts(userId uint64) error {
	const query = `DELETE FROM posts WHERE owner_id=?`
	params := []interface{}{userId}

	tx := DB.Exec(query, params...)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		this.Lgr.Log("tx.RowsAffected: ", tx.RowsAffected)
		this.Lgr.Log("Cant find posts")
		return errors.New("Cant find posts")
	}

	return nil
}
