package handlers

import (
	"context"
	pb "freeSociety/proto/generated/security"
)

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
