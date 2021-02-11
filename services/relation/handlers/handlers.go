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
	isExist := h.Publishers.IsUserExist(in.Follower)

	h.lgr.Log("isExist: ", isExist)
	h.lgr.Log("in.Follower: ", in.Follower)

	if !isExist {
		return &pb.FollowResponse{}, errors.New("user not exist")
	}

	err := h.srv.Follow(in.Following, in.Follower)

	return &pb.FollowResponse{}, err
}
