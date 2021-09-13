package instances

import (
	"freeSociety/services/security/models"
	dbhelper "freeSociety/utils/dbHelper"
)

type postgres_read interface {
	GetHashPass(userId uint64) (string, error)
	GetUserIdByToken(token string) (uint64, error)
	GetUserToken(userId uint64) []string
	GetSessions(userId uint64) ([]*models.Session, error)
}
type postgres_write interface {
	NewUser(userId uint64, hashPass string) (cc dbhelper.CommandController, err error)
	NewSession(userId uint64, device, token, expireAt string) (sessionId uint64, cc dbhelper.CommandController, err error)
	DeleteSessionByToken(token string) (cc dbhelper.CommandController, err error)
	DeleteUserSessions(userId uint64) (sessions []*models.Session, cc dbhelper.CommandController, err error)
	DeletePassword(userId uint64) (cc dbhelper.CommandController, err error)
	DeleteSessionById(sessionId uint64) (session *models.Session, cc dbhelper.CommandController, err error)
	ChangeHashPass(userId uint64, newHashPass string) (cc dbhelper.CommandController, err error)
}

type Repo_Postgres struct {
	Read  postgres_read
	Write postgres_write
}
