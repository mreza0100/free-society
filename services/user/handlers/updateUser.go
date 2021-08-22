package handlers

import (
	"context"
	pb "freeSociety/proto/generated/user"
)

func (h *handlers) UpdateUser(_ context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return &pb.UpdateUserResponse{}, h.srv.UpdateUser(in.Id, in.Name, in.Gender, in.AvatarFormat, in.Avatar)
}
