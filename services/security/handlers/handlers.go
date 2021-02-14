package handlers

import (
	"context"
	pb "microServiceBoilerplate/proto/generated/security"
	"microServiceBoilerplate/services/security/types"

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
		publishers: opts.Publishers,
	}
}

type handlers struct {
	srv        types.Sevice
	lgr        *golog.Core
	publishers types.Publishers

	pb.UnimplementedSecurityServiceServer
}

func (h *handlers) NewUser(ctx context.Context, in *pb.NewUserRequest) (*pb.NewUserResponse, error) {
	return &pb.NewUserResponse{}, nil
}

func (h *handlers) Login(context.Context, *pb.LogInRequest) (*pb.LogInResponse, error) {
	return &pb.LogInResponse{}, nil
}
