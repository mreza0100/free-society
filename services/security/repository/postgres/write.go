package postgres

import (
	"github.com/mreza0100/golog"
	gorm "gorm.io/gorm"
)

type write struct {
	lgr *golog.Core
	db  *gorm.DB
}

func (w *write) NewUser(userId uint64, hashPass string) error {
	const query = `INSERT INTO passwords (user_id, password) VALUES (?, ?)`
	params := []interface{}{userId, hashPass}

	tx := w.db.Exec(query, params...)

	return tx.Error
}

func (w *write) NewSession(userId uint64, device, token string) (sessionId uint64, err error) {
	const query = `INSERT INTO sessions (user_id, device, token) VALUES (?, ?, ?) RETURNING id`
	params := []interface{}{userId, device, token}

	tx := w.db.Raw(query, params...)
	if tx.Error != nil {
		return sessionId, tx.Error
	}

	data := struct{ Id uint64 }{Id: 0}
	tx = tx.Scan(&data)
	if tx.Error != nil {
		w.lgr.BugHunter(tx.Error.Error())
		return sessionId, tx.Error
	}

	return data.Id, nil
}

// ! deletes
func (w *write) DeleteSessionByToken(token string) error {
	const query = `DELETE FROM sessions WHERE token=?`
	params := []interface{}{token}

	tx := w.db.Exec(query, params...)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (w *write) DeleteUserSessions(userId uint64) error {
	const query = `DELETE FROM sessions WHERE user_id=?`
	params := []interface{}{userId}

	tx := w.db.Exec(query, params...)

	return tx.Error
}

func (w *write) DeletePassword(userId uint64) error {
	const qeury = `DELETE FROM passwords WHERE user_id=?`
	params := []interface{}{userId}

	tx := w.db.Exec(qeury, params...)

	return tx.Error
}
