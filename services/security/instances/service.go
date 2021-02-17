package instances

type Sevice interface {
	NewUser(userId uint64, device, password string) (string, error)
	Login(userId uint64, device, password string) (string, error)
	Logout(token string) error
	GetUserId(token string) (uint64, error)
	PurgeUser(userId uint64) error
}
