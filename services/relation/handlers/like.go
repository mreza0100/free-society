package handlers

import (
	"context"
	pb "freeSociety/proto/generated/relation"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *handlers) Like(_ context.Context, in *pb.LikeRequest) (*empty.Empty, error) {
	{
		isExists, err := h.publishers.IsPostsExists(in.PostId)

		if err != nil {
			return &emptypb.Empty{}, err
		}
		if len(isExists) != 1 {
			return &emptypb.Empty{}, status.Error(codes.NotFound, "post not found")
		}

		err = h.srv.Like(in.LikerId, in.OwnerId, in.PostId)
		if err != nil {
			return &emptypb.Empty{}, err
		}
	}

	{
		_, err := h.publishers.LikeNotify(in.OwnerId, in.LikerId, in.PostId)
		if err != nil {
			return nil, err
		}
	}

	return &emptypb.Empty{}, nil
}
