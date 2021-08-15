package instances

import (
	"freeSociety/services/security/models"
)

type postgres_read interface {
	GetHashPass(userId uint64) (string, error)
	GetUserIdByToken(token string) (uint64, error)
	GetUserToken(userId uint64) []string
	GetSessions(userId uint64) ([]*models.Session, error)
}
type postgres_write interface {
	NewUser(userId uint64, hashPass string) error
	NewSession(userId uint64, device, token, expireAt string) (sessionId uint64, err error)
	DeleteSessionByToken(token string) error
	DeleteUserSessions(userId uint64) ([]*models.Session, error)
	DeletePassword(userId uint64) error
	DeleteSessionById(sessionId uint64) (*models.Session, error)
	ChangeHashPass(userId uint64, hashPass string) error
}

type Repo_Postgres struct {
	Read  postgres_read
	Write postgres_write
}
