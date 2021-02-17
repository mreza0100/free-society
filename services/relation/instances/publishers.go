package instances

type Publishers interface {
	IsUserExist(userId uint64) bool
}
