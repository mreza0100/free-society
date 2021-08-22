package handlers

import (
	"context"
	pb "freeSociety/proto/generated/user"
)

func (h *handlers) DeleteUser(_ context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	errsCh := make(chan error, 2)

	// flow request to services.DeleteUser and publish request to post service to delete user posts
	go func(ch chan error) { ch <- h.srv.DeleteUser(in.Id) }(errsCh)
	go func(ch chan error) { ch <- h.publishers.DeleteUser(in.Id) }(errsCh)

	// start from 1 not 0
	for i := 1; i < cap(errsCh); i++ {
		if err := <-errsCh; err != nil {
			return &pb.DeleteUserResponse{
				Ok: false,
			}, err
		}
	}

	return &pb.DeleteUserResponse{
		Ok: true,
	}, nil
}
