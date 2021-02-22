package instances

type read interface {
	GetFollowers(userId uint64) []uint64
	IsFollowingGroup(follower uint64, followings []uint64) (map[uint64]interface{}, error)
}

type write interface {
	SetFollower(follower, following uint64) error
	RemoveFollow(follower, following uint64) error
	DeleteAllRelations(userId uint64) error
}

type Repository struct {
	Read  read
	Write write
}
