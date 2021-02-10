package handlers

import (
	"context"
	pb "microServiceBoilerplate/proto/generated/user"
	domain "microServiceBoilerplate/services/user/domain"
	userNats "microServiceBoilerplate/services/user/nats"

	"github.com/mreza0100/golog"
)

type handlers struct {
	srv        domain.Sevice
	lgr        *golog.Core
	publishers userNats.Publishers

	pb.UnimplementedUserServiceServer
}

type NewHandlersOpts struct {
	Srv        domain.Sevice
	Lgr        *golog.Core
	Publishers userNats.Publishers
}

func NewHandlers(opts NewHandlersOpts) pb.UserServiceServer {
	return &handlers{
		srv:        opts.Srv,
		lgr:        opts.Lgr.With("In handlers: "),
		publishers: opts.Publishers,
	}
}

func (s *handlers) NewUser(ctx context.Context, in *pb.NewUserRequest) (*pb.NewUserResponse, error) {
	id, err := s.srv.NewUser(in)

	return &pb.NewUserResponse{
		Id: id,
	}, err
}

func (s *handlers) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return s.srv.GetUser(in.Id)
}

func (s *handlers) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	errsCh := make(chan error, 2)

	// flow request to services.DeleteUser and publish request to post service to delete user posts
	go func(ch chan error) { ch <- s.srv.DeleteUser(in.Id) }(errsCh)
	go func(ch chan error) { ch <- s.publishers.DeleteUser(in.Id) }(errsCh)

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

func (this *handlers) Validation(ctx context.Context, in *pb.ValidationRequest) (*pb.ValidationResponse, error) {
	userId, err := this.srv.Validation(in.Email, in.Password)

	return &pb.ValidationResponse{
		Id: userId,
	}, err
}
