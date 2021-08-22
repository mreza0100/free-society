package handlers

import (
	"context"
	pb "freeSociety/proto/generated/user"
)

func (h *handlers) NewUser(_ context.Context, in *pb.NewUserRequest) (*pb.NewUserResponse, error) {
	id, err := h.srv.NewUser(in.Name, in.Email, in.Gender, in.AvatarFormat, in.Avatar)

	return &pb.NewUserResponse{
		Id: id,
	}, err
}
