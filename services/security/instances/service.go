package instances

import "microServiceBoilerplate/services/security/models"

type Sevice interface {
	NewUser(userId uint64, device, password string) (string, error)
	Login(userId uint64, device, password string) (string, error)
	Logout(token string) error
	GetUserId(token string) (uint64, error)
	PurgeUser(userId uint64) error
	GetSessions(userId uint64) ([]*models.Session, error)
	DeleteSession(sessionId uint64) error
}
