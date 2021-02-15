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

func (h *handlers) NewUser(_ context.Context, in *pb.NewUserRequest) (*pb.NewUserResponse, error) {
	result, err := h.srv.NewUser(in.UserId, in.Device, in.Password)

	return &pb.NewUserResponse{
		Token: result,
	}, err
}

func (h *handlers) Login(_ context.Context, in *pb.LogInRequest) (*pb.LogInResponse, error) {
	token, err := h.srv.Login(in.UserId, "", in.Password)

	return &pb.LogInResponse{
		Token: token,
	}, err
}
func (h *handlers) Logout(_ context.Context, in *pb.LogoutInRequest) (*pb.LogoutInResponse, error) {
	return nil, h.srv.Logout(in.Token)
}
func (h *handlers) GetUserId(_ context.Context, in *pb.GetUserIdRequest) (*pb.GetUserIdResponse, error) {
	userId, err := h.srv.GetUserId(in.Token)

	return &pb.GetUserIdResponse{
		UserId: userId,
	}, err
}
