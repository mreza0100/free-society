package instances

type Sevice interface {
	GetFeed(userId, offset, limit uint64) ([]string, error)
	SetPost(userId uint64, postId string) error
	DeleteFeed(userId uint64) error
	Reshare(userId uint64, postId string) error
}
