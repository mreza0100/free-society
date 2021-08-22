package handlers

import (
	"context"
	pb "freeSociety/proto/generated/relation"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *handlers) Unfollow(_ context.Context, in *pb.UnfollowRequest) (*empty.Empty, error) {
	err := h.srv.Unfollow(in.Follower, in.Following)

	return &emptypb.Empty{}, err
}
