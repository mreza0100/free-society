package handlers

import (
	"context"
	pb "freeSociety/proto/generated/security"
)

func (h *handlers) NewUser(_ context.Context, in *pb.NewUserRequest) (*pb.NewUserResponse, error) {
	token, err := h.srv.NewUser(in.UserId, in.Device, in.Password)

	return &pb.NewUserResponse{
		Token: token,
	}, err
}
