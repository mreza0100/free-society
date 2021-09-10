package instances

type Publishers interface {
	IsUserExist(userId uint64) bool
	IsPostsExists(postIds ...string) ([]string, error)
	LikeNotify(userId, likerId uint64, postId string) (uint64, error)
}
