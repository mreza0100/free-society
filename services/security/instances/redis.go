package instances

type redis_read interface {
	GetSession(token string) (uint64, error)
}
type redis_write interface {
	NewSession(token string, userId uint64) error
	DeleteSession(token ...string) error
}
type Repo_Redis struct {
	Read  redis_read
	Write redis_write
}
