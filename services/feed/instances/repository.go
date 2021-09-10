package instances

type read interface {
	GetFeed(userId, offset, limit uint64) ([]string, error)
}
type write interface {
	SetPostOnFeeds(userId uint64, postId string, followers []uint64) error
	DeleteFeed(userId uint64) error
}

type Repository struct {
	Read  read
	Write write
}
