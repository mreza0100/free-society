package types

type Sevice interface {
	Follow(follower, following uint64) error
	Unfollow(follower, following uint64) error
}
