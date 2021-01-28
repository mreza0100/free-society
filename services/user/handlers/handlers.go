package handlers

import (
	"context"
	pb "microServiceBoilerplate/proto/generated/user"
	domain "microServiceBoilerplate/services/user/domain"
)

type handlers struct {
	srv domain.Sevice
	pb.UnimplementedUserServiceServer
}

func NewHandlers(srv domain.Sevice) pb.UserServiceServer {
	return &handlers{
		srv: srv,
	}
}

func (s *handlers) NewUser(ctx context.Context, in *pb.NewUserRequest) (*pb.NewUserResponse, error) {
	id, err := s.srv.NewUser(in)

	return &pb.NewUserResponse{
		Id: id,
	}, err
}

func (s *handlers) GetUserById(ctx context.Context, in *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	return s.srv.GetUserById(in.Id)
}
func (s *handlers) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users, err := s.srv.GetUsers()
	return &pb.GetUsersResponse{
		Users: users,
	}, err
}
func (s *handlers) DeleteUserById(ctx context.Context, in *pb.DeleteUserByIdRequest) (*pb.DeleteUserByIdResponse, error) {
	err := s.srv.DeleteUserById(in.Id)

	if err != nil {
		return &pb.DeleteUserByIdResponse{
			Ok: false,
		}, err
	}

	return &pb.DeleteUserByIdResponse{
		Ok: true,
	}, nil
}
