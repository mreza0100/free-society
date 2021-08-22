package handlers

import (
	"context"
	pb "freeSociety/proto/generated/security"
)

func (h *handlers) ChangePassword(_ context.Context, in *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	return &pb.ChangePasswordResponse{}, h.srv.ChangePassword(in.UserId, in.PrevPassword, in.NewPassword)
}
