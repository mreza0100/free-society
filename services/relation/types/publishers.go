package types

type Publishers interface {
	IsUserExist(userId uint64) bool
}
