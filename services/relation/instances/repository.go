package instances

type (
	followers_read interface {
		GetFollowers(userId uint64) []uint64
		IsFollowing(follower, following uint64) bool
		IsFollowingGroup(follower uint64, followings []uint64) (map[uint64]interface{}, error)
	}

	followers_write interface {
		SetFollower(follower, following uint64) error
		RemoveFollow(follower, following uint64) error
		DeleteAllRelations(userId uint64) error
	}
)

type (
	likes_read interface {
		IsLikedGroup(userId uint64, postIds []uint64) ([]uint64, error)
		CountLikes(postIds []uint64) (CountResult, error)
	}

	likes_write interface {
		Like(likerId, ownerId, postId uint64) error
		UndoLike(likerId, postId uint64) error
	}
)

type CountResult []*struct {
	Count   uint
	Post_id uint64
}

type Repository struct {
	Followers_read  followers_read
	Followers_write followers_write

	Likes_read  likes_read
	Likes_write likes_write
}
