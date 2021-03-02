package handlers

import (
	"context"
	"errors"
	pb "microServiceBoilerplate/proto/generated/relation"
)

func (h *handlers) Follow(_ context.Context, in *pb.FollowRequest) (*pb.FollowResponse, error) {
	{
		isExist := h.publishers.IsUserExist(in.Following)
		if !isExist {
			return &pb.FollowResponse{}, errors.New("user not exist")
		}
	}

	return &pb.FollowResponse{}, h.srv.Follow(in.Follower, in.Following)
}

func (h *handlers) Unfollow(_ context.Context, in *pb.UnfollowRequest) (*pb.UnfollowResponse, error) {
	err := h.srv.Unfollow(in.Follower, in.Following)

	return &pb.UnfollowResponse{}, err
}

func (this *handlers) GetFollowers(userId uint64) []uint64 {
	return this.srv.GetFollowers(userId)
}
