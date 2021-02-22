package instances

import (
	pb "microServiceBoilerplate/proto/generated/post"
)

type Publishers interface {
	NewPost(userId, postId uint64) error
	GetUsers(userIds []uint64) (map[uint64]*pb.User, error)
	IsFollowingGroup(userId uint64, followings []uint64) map[uint64]bool
}
