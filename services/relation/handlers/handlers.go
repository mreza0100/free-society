package handlers

import (
	"context"
	"errors"
	pb "microServiceBoilerplate/proto/generated/relation"
	"microServiceBoilerplate/services/relation/types"

	"github.com/mreza0100/golog"
)

type NewHandlersOpts struct {
	Srv        types.Sevice
	Lgr        *golog.Core
	Publishers types.Publishers
}

func NewHandlers(opts NewHandlersOpts) types.Handlers {
	return &handlers{
		srv:        opts.Srv,
		lgr:        opts.Lgr.With("In handlers: "),
		Publishers: opts.Publishers,
	}
}

type handlers struct {
	srv        types.Sevice
	lgr        *golog.Core
	Publishers types.Publishers

	pb.UnimplementedRelationServiceServer
}

func (h *handlers) Follow(_ context.Context, in *pb.FollowRequest) (*pb.FollowResponse, error) {
	isExist := h.Publishers.IsUserExist(in.Following)

	if !isExist {
		return &pb.FollowResponse{}, errors.New("user not exist")
	}

	err := h.srv.Follow(in.Follower, in.Following)

	return &pb.FollowResponse{}, err
}

func (h *handlers) Unfollow(_ context.Context, in *pb.UnfollowRequest) (*pb.UnfollowResponse, error) {
	err := h.srv.Unfollow(in.Follower, in.Following)

	return &pb.UnfollowResponse{}, err
}

func (this *handlers) GetFollowers(userId uint64) []uint64 {
	return this.srv.GetFollowers(userId)
}
