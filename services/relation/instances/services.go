package instances

type follow interface {
	Follow(follower, following uint64) error
	Unfollow(following, follower uint64) error

	GetFollowers(userId uint64) []uint64
	DeleteAllRelations(userId uint64) error
	IsFollowing(follower uint64, followings []uint64) (map[uint64]bool, error)
}

type like interface {
	Like(likerId, ownerId uint64, postId string) error
	UndoLike(likerId uint64, postId string) error
	IsLikedGroup(likerId uint64, postIds []string) ([]string, error)
	CountLikes(postIds []string) (CountResult, error)
	DeleteLikes(liker uint64) error
}

type Sevice interface {
	follow
	like
}
