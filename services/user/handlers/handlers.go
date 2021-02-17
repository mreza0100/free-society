package handlers

import (
	"context"
	pb "microServiceBoilerplate/proto/generated/user"

	"microServiceBoilerplate/services/user/instances"

	"github.com/mreza0100/golog"
)

type NewHandlersOpts struct {
	Lgr        *golog.Core
	Srv        instances.Sevice
	Publishers instances.Publishers
}

func NewHandlers(opts *NewHandlersOpts) instances.Handlers {
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
	id, err := h.srv.NewUser(in)

	return &pb.NewUserResponse{
		Id: id,
	}, err
}

func (h *handlers) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return h.srv.GetUser(in.Id, in.Email)
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
