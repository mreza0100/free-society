package handlers

import (
	"context"
	pb "freeSociety/proto/generated/security"
)

func (h *handlers) Logout(_ context.Context, in *pb.LogoutInRequest) (*pb.LogoutInResponse, error) {
	return &pb.LogoutInResponse{}, h.srv.Logout(in.Token)
}
