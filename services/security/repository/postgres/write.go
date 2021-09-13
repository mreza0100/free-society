package postgres

import (
	"freeSociety/services/security/models"
	dbhelper "freeSociety/utils/dbHelper"

	"github.com/mreza0100/golog"
	gorm "gorm.io/gorm"
)

type write struct {
	lgr *golog.Core
	db  *gorm.DB
}

func (w *write) NewUser(userId uint64, hashPass string) (cc dbhelper.CommandController, err error) {
	cc, err = dbhelper.Transaction(w.db, func(tx *gorm.DB) error {
		const query = `INSERT INTO passwords (user_id, password) VALUES (?, ?)`
		params := []interface{}{userId, hashPass}

		tx = w.db.Exec(query, params...)
		return tx.Error
	})

	return cc, err
}

func (w *write) NewSession(userId uint64, device, token, expireAt string) (sessionId uint64, cc dbhelper.CommandController, err error) {
	cc, err = dbhelper.Transaction(w.db, func(tx *gorm.DB) error {
		const query = `INSERT INTO sessions (user_id, device, token, expire_at) VALUES (?, ?, ?, ?) RETURNING id`
		params := []interface{}{userId, device, token, expireAt}

		tx = w.db.Raw(query, params...)
		if tx.Error != nil {
			return tx.Error
		}

		result := struct{ Id uint64 }{Id: 0}
		tx = tx.Scan(&result)
		sessionId = result.Id

		return tx.Error
	})

	return sessionId, cc, err
}

func (w *write) DeleteSessionByToken(token string) (cc dbhelper.CommandController, err error) {
	cc, err = dbhelper.Transaction(w.db, func(tx *gorm.DB) error {
		const query = `DELETE FROM sessions WHERE token=?`
		params := []interface{}{token}

		tx = w.db.Exec(query, params...)

		return tx.Error
	})

	return cc, err
}

func (w *write) DeleteUserSessions(userId uint64) (sessions []*models.Session, cc dbhelper.CommandController, err error) {
	cc, err = dbhelper.Transaction(w.db, func(tx *gorm.DB) error {
		const query = `DELETE FROM sessions WHERE user_id=? RETURNING *`
		params := []interface{}{userId}

		tx = w.db.Raw(query, params...)
		if tx.Error != nil {
			return tx.Error
		}

		sessions = []*models.Session{}
		tx.Scan(&sessions)

		return tx.Error
	})

	return sessions, cc, err
}

func (w *write) DeletePassword(userId uint64) (cc dbhelper.CommandController, err error) {
	cc, err = dbhelper.Transaction(w.db, func(tx *gorm.DB) error {
		const qeury = `DELETE FROM passwords WHERE user_id=?`
		params := []interface{}{userId}

		tx = w.db.Exec(qeury, params...)
		return tx.Error
	})

	return cc, err
}

func (w *write) DeleteSessionById(sessionId uint64) (session *models.Session, cc dbhelper.CommandController, err error) {
	cc, err = dbhelper.Transaction(w.db, func(tx *gorm.DB) error {
		const qeury = `DELETE FROM sessions WHERE id=? RETURNING *`
		params := []interface{}{sessionId}

		tx = w.db.Raw(qeury, params...)
		if tx.Error != nil {
			return tx.Error
		}

		session := &models.Session{}
		tx.Scan(&session)
		return nil
	})

	return session, cc, err
}

func (w *write) ChangeHashPass(userId uint64, newHashPass string) (cc dbhelper.CommandController, err error) {
	cc, err = dbhelper.Transaction(w.db, func(tx *gorm.DB) error {
		const query = `UPDATE passwords SET password = ? WHERE user_id = ?`
		params := []interface{}{newHashPass, userId}

		tx = w.db.Exec(query, params...)
		return tx.Error
	})

	return cc, err
}
