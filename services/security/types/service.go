package types

type Sevice interface {
	NewUser(userId uint64, password string) (string, error)
	Login(userId uint64, password string) (string, error)
}
