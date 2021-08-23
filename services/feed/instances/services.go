package instances

type Sevice interface {
	GetFeed(userId, offset, limit uint64) ([]uint64, error)
	SetPost(userId, postId uint64) error
	DeleteFeed(userId uint64) error
	Reshare(userId, postId uint64) error
}
