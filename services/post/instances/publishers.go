package instances

import (
	pb "microServiceBoilerplate/proto/generated/post"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Publishers interface {
	NewPost(userId, postId uint64) error
	GetUsers(userIds []uint64) (map[uint64]*pb.User, error)
	IsFollowingGroup(userId uint64, followings []uint64) (map[uint64]bool, error)
	GetCounts(postIds []uint64) (map[uint64]uint64, error)
	IsLikedGroup(liker uint64, postIds []uint64) (map[uint64]*emptypb.Empty, error)
}
