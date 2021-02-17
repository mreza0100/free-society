package instances

type Sevice interface {
	Follow(follower, following uint64) error
	Unfollow(follower, following uint64) error

	GetFollowers(userId uint64) []uint64
	DeleteAllRelations(userId uint64) error
}
