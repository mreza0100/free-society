package instances

import (
	pb "freeSociety/proto/generated/post"
)

type Sevice interface {
	NewPost(title, body string, userId uint64, pictures []*pb.Picture) (string, error)
	GetPost(requestorId uint64, postIds []string) ([]*pb.Post, error)
	DeletePost(postId string, userId uint64) error
	DeleteUserPosts(userId uint64) error
	IsPostsExists(postIds []string) ([]string, error)
}
