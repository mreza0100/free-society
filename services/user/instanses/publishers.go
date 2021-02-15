package instanses

type Publishers interface {
	DeleteUser(userId uint64) error
}
