package handlers

import (
	"context"
	pb "freeSociety/proto/generated/user"

	"freeSociety/services/user/instances"

	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Lgr        *golog.Core
	Srv        instances.Sevice
	Publishers instances.Publishers
}

func New(opts *NewOpts) instances.Handlers {
	return &handlers{
		lgr:        opts.Lgr.With("In handlers->"),
		srv:        opts.Srv,
		publishers: opts.Publishers,
	}
}

type handlers struct {
	lgr        *golog.Core
	srv        instances.Sevice
	publishers instances.Publishers

	pb.UnimplementedUserServiceServer
}

func (h *handlers) NewUser(ctx context.Context, in *pb.NewUserRequest) (*pb.NewUserResponse, error) {
	id, err := h.srv.NewUser(in.Name, in.Email, in.Gender, in.AvatarFormat, in.Avatar)

	return &pb.NewUserResponse{
		Id: id,
	}, err
}

func (h *handlers) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return h.srv.GetUser(in.RequestorId, in.Id, in.Email)
}

func (h *handlers) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	errsCh := make(chan error, 2)

	// flow request to services.DeleteUser and publish request to post service to delete user posts
	go func(ch chan error) { ch <- h.srv.DeleteUser(in.Id) }(errsCh)
	go func(ch chan error) { ch <- h.publishers.DeleteUser(in.Id) }(errsCh)

	// start from 1 not 0
	for i := 1; i < cap(errsCh); i++ {
		if err := <-errsCh; err != nil {
			return &pb.DeleteUserResponse{
				Ok: false,
			}, err
		}
	}

	return &pb.DeleteUserResponse{
		Ok: true,
	}, nil
}
