package handlers

import (
	"context"
	pb "freeSociety/proto/generated/security"
)

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
