package instances

type postgres_read interface {
	GetHashPass(userId uint64) (string, error)
	GetUserIdByToken(token string) (uint64, error)
	GetUserToken(userId uint64) []string
}
type postgres_write interface {
	NewUser(userId uint64, hashPass string) error
	NewSession(userId uint64, device, token string) (sessionId uint64, err error)
	DeleteSessionByToken(token string) error
	DeleteUserSessions(userId uint64) error
	DeletePassword(userId uint64) error
}

type Repo_Postgres struct {
	Read  postgres_read
	Write postgres_write
}
