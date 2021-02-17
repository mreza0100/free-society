package postgres

import (
	"github.com/mreza0100/golog"
	"gorm.io/gorm"
)

type read struct {
	lgr *golog.Core
	db  *gorm.DB
}

func (r *read) GetHashPass(userId uint64) (string, error) {
	const query = `SELECT password FROM passwords WHERE user_id=?`
	params := []interface{}{userId}

	tx := r.db.Raw(query, params...)
	if tx.Error != nil {
		return "", tx.Error
	}

	data := struct{ Password string }{}

	tx.Scan(&data)

	return data.Password, nil
}

func (r *read) GetUserIdByToken(token string) (uint64, error) {
	const query = `SELECT user_id FROM sessions WHERE token=?`
	params := []interface{}{token}

	tx := r.db.Raw(query, params...)
	if tx.Error != nil {
		return 0, tx.Error
	}

	data := struct{ UserId uint64 }{}
	tx.Scan(&data)

	return data.UserId, nil
}

func (w *read) GetUserToken(userId uint64) []string {
	const query = `SELECT token FROM sessions WHERE user_id=?`
	params := []interface{}{userId}

	tx := w.db.Raw(query, params...)

	data := make([]string, 0)
	tx.Scan(&data)

	return data
}
