package instances

import (
	pb "freeSociety/proto/generated/post"
)

type Sevice interface {
	NewPost(title, body string, userId uint64, pictures []*pb.Picture) (uint64, error)
	GetPost(requestorId uint64, postIds []uint64) ([]*pb.Post, error)
	DeletePost(postId, userId uint64) error
	DeleteUserPosts(userId uint64) error
	IsPostsExists(postIds []uint64) ([]uint64, error)
}
