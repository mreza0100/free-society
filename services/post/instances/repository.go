package instances

import pb "microServiceBoilerplate/proto/generated/post"

type read interface {
	GetPost(postIds []uint64) ([]*pb.Post, error)
}

type write interface {
	NewPost(title, body string, userId uint64) (uint64, error)
	DeletePost(postId, userId uint64) error
	DeleteUserPosts(userId uint64) error
}

type Repository struct {
	Read  read
	Write write
}
