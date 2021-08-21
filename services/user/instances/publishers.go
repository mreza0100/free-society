package instances

type Publishers interface {
	DeleteUser(userId uint64) error
	IsFollowingGroup(userId uint64, followings []uint64) (map[uint64]bool, error)
}
