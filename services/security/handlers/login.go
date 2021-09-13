package handlers

import (
	"context"
	pb "freeSociety/proto/generated/security"
)

func (h *handlers) Login(_ context.Context, in *pb.LogInRequest) (*pb.LogInResponse, error) {
	token, err := h.srv.Login(in.UserId, in.Device, in.Password)

	return &pb.LogInResponse{
		Token: token,
	}, err
}
