package instances

type read interface {
	GetFollowers(userId uint64) []uint64
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
