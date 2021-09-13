package handlers

import (
	"context"
	pb "freeSociety/proto/generated/user"
)

func (h *handlers) DeleteUser(_ context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := h.srv.DeleteUser(in.Id)

	return &pb.DeleteUserResponse{
		Ok: err == nil,
	}, err
}
