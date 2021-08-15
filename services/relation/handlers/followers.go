package handlers

import (
	"context"
	"errors"
	pb "freeSociety/proto/generated/relation"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *handlers) Follow(_ context.Context, in *pb.FollowRequest) (*empty.Empty, error) {
	{
		isExist := h.publishers.IsUserExist(in.Following)
		if !isExist {
			return &emptypb.Empty{}, errors.New("user not exist")
		}
	}

	return &emptypb.Empty{}, h.srv.Follow(in.Follower, in.Following)
}

func (h *handlers) Unfollow(_ context.Context, in *pb.UnfollowRequest) (*empty.Empty, error) {
	err := h.srv.Unfollow(in.Follower, in.Following)

	return &emptypb.Empty{}, err
}

func (h *handlers) GetFollowers(userId uint64) []uint64 {
	return h.srv.GetFollowers(userId)
}
