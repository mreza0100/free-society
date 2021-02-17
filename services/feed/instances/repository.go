package instances

type read interface {
	GetFeed(userId, offset, limit uint64) ([]uint64, error)
}
type write interface {
	SetPostOnFeeds(userId, postId uint64, followers []uint64) error
	DeleteFeed(userId uint64) error
}

type Repository struct {
	Read  read
	Write write
}
