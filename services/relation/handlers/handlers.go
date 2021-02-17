package handlers

import (
	"context"
	"errors"
	pb "microServiceBoilerplate/proto/generated/relation"
	"microServiceBoilerplate/services/relation/instances"

	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Srv        instances.Sevice
	Lgr        *golog.Core
	Publishers instances.Publishers
}

func New(opts *NewOpts) instances.Handlers {
	return &handlers{
		srv:        opts.Srv,
		lgr:        opts.Lgr.With("In handlers->"),
		Publishers: opts.Publishers,
	}
}

type handlers struct {
	srv        instances.Sevice
	lgr        *golog.Core
	Publishers instances.Publishers

	pb.UnimplementedRelationServiceServer
}

func (h *handlers) Follow(_ context.Context, in *pb.FollowRequest) (*pb.FollowResponse, error) {
	{
		isExist := h.Publishers.IsUserExist(in.Following)
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
