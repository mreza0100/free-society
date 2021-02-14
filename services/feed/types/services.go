package types

type Sevice interface {
	GetFeed(userId, offset, limit uint64) ([]uint64, error)
	SetPost(userId, postId uint64) error
}
