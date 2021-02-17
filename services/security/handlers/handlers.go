package handlers

import (
	"context"
	pb "microServiceBoilerplate/proto/generated/security"
	"microServiceBoilerplate/services/security/instances"

	"github.com/mreza0100/golog"
)

type NewHandlersOpts struct {
	Srv        instances.Sevice
	Lgr        *golog.Core
	Publishers instances.Publishers
}

func NewHandlers(opts NewHandlersOpts) instances.Handlers {
	return &handlers{
		srv:        opts.Srv,
		lgr:        opts.Lgr.With("In handlers->"),
		publishers: opts.Publishers,
	}
}

type handlers struct {
	srv        instances.Sevice
	lgr        *golog.Core
	publishers instances.Publishers

	pb.UnimplementedSecurityServiceServer
}

func (h *handlers) NewUser(_ context.Context, in *pb.NewUserRequest) (*pb.NewUserResponse, error) {
	var (
		token  string
		err    error
		result *pb.NewUserResponse
	)

	{
		token, err = h.srv.NewUser(in.UserId, in.Device, in.Password)
	}

	{
		result = &pb.NewUserResponse{
			Token: token,
		}
	}

	return result, err
}

func (h *handlers) Login(_ context.Context, in *pb.LogInRequest) (*pb.LogInResponse, error) {
	var (
		token  string
		err    error
		result *pb.LogInResponse
	)

	{
		token, err = h.srv.Login(in.UserId, "", in.Password)
	}

	{
		result = &pb.LogInResponse{
			Token:  token,
			Device: "",
		}
	}
	return result, err
}
func (h *handlers) Logout(_ context.Context, in *pb.LogoutInRequest) (*pb.LogoutInResponse, error) {
	return &pb.LogoutInResponse{}, h.srv.Logout(in.Token)
}
func (h *handlers) GetUserId(_ context.Context, in *pb.GetUserIdRequest) (*pb.GetUserIdResponse, error) {
	var (
		userId uint64
		err    error
		result *pb.GetUserIdResponse
	)

	{
		userId, err = h.srv.GetUserId(in.Token)
	}

	{
		result = &pb.GetUserIdResponse{
			UserId: userId,
		}
	}

	return result, err
}
