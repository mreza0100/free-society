package handlers

import (
	"context"
	pb "freeSociety/proto/generated/security"
	"freeSociety/services/security/instances"
	"freeSociety/services/security/models"

	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Srv        instances.Sevice
	Lgr        *golog.Core
	Publishers instances.Publishers
}

func New(opts *NewOpts) pb.SecurityServiceServer {
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
		token, err = h.srv.Login(in.UserId, in.Device, in.Password)
	}

	{
		result = &pb.LogInResponse{
			Token: token,
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

func (h *handlers) GetSessions(_ context.Context, in *pb.GetSessionsRequest) (*pb.GetSessionsResponse, error) {
	var (
		result    []*models.Session
		converted []*pb.Session
		err       error
	)

	{
		result, err = h.srv.GetSessions(in.UserId)
		if err != nil {
			return &pb.GetSessionsResponse{}, err
		}
	}

	{
		converted = make([]*pb.Session, len(result))
		for idx, i := range result {
			converted[idx] = &pb.Session{
				SessionId: i.ID,
				Device:    i.Device,
				CreatedAt: uint64(i.CreatedAt.Unix()),
			}
		}
	}

	return &pb.GetSessionsResponse{
		Sessions: converted,
	}, err
}

func (h *handlers) DeleteSession(_ context.Context, in *pb.DeleteSessionRequest) (*pb.DeleteSessionResponse, error) {
	return &pb.DeleteSessionResponse{}, h.srv.DeleteSession(in.SessionId)
}

func (h *handlers) ChangePassword(_ context.Context, in *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	return &pb.ChangePasswordResponse{}, h.srv.ChangePassword(in.UserId, in.PrevPassword, in.NewPassword)
}
