package instances

type Publishers interface {
	IsUserExist(userId uint64) bool
	IsPostsExists(postIds ...uint64) ([]uint64, error)
	LikeNotify(userId, likerId, postId uint64) (uint64, error)
}
