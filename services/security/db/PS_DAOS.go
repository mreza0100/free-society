package db

import (
	"github.com/mreza0100/golog"
)

type PS_DAOS struct {
	Lgr *golog.Core
}

func (this *PS_DAOS) NewUser(userId uint64, hashPass string) error {
	const query = `INSERT INTO passwords (user_id, password) VALUES (?, ?)`
	params := []interface{}{userId, hashPass}

	tx := psDB.Exec(query, params...)

	return tx.Error
}

func (this *PS_DAOS) NewSession(userId uint64, device, token string) (sessionId uint64, err error) {
	const query = `INSERT INTO sessions (user_id, device, token) VALUES (?, ?, ?) RETURNING id`
	params := []interface{}{userId, device, token}

	tx := psDB.Raw(query, params...)
	if tx.Error != nil {
		return sessionId, tx.Error
	}

	data := struct{ Id uint64 }{Id: 0}
	tx = tx.Scan(&data)
	if tx.Error != nil {
		this.Lgr.BugHunter(tx.Error.Error())
		return sessionId, tx.Error
	}

	return data.Id, nil
}

func (this *PS_DAOS) GetHashPass(userId uint64) (string, error) {
	const query = `SELECT password FROM passwords WHERE user_id=?`
	params := []interface{}{userId}

	tx := psDB.Raw(query, params...)
	if tx.Error != nil {
		return "", tx.Error
	}

	data := struct{ Password string }{}

	tx.Scan(&data)

	return data.Password, nil
}

func (this *PS_DAOS) GetUserIdByToken(token string) (uint64, error) {
	const query = `SELECT user_id FROM sessions WHERE token=?`
	params := []interface{}{token}

	tx := psDB.Raw(query, params...)
	if tx.Error != nil {
		return 0, tx.Error
	}

	data := struct{ UserId uint64 }{}

	tx.Scan(&data)

	return data.UserId, nil
}

func (this *PS_DAOS) DeleteSessionByToken(token string) error {
	const query = `DELETE FROM sessions WHERE token=?`
	params := []interface{}{token}

	tx := psDB.Exec(query, params...)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
