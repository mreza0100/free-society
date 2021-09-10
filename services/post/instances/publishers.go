package instances

import (
	pb "freeSociety/proto/generated/post"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Publishers interface {
	NewPost(userId uint64, postId string) error
	GetUsers(userIds []uint64) (map[uint64]*pb.User, error)
	IsFollowingGroup(userId uint64, followings []uint64) (map[uint64]bool, error)
	GetCounts(postIds []string) (map[string]uint64, error)
	IsLikedGroup(liker uint64, postIds []string) (map[string]*emptypb.Empty, error)
}
