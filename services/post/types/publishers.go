package types

type Publishers interface {
	NewPost(userId, postId uint64) error
}
