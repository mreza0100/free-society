package handlers

import (
	"context"
	pb "microServiceBoilerplate/proto/generated/relation"

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
	}

	return &emptypb.Empty{}, h.srv.Like(in.LikerId, in.OwnerId, in.PostId)
}

func (h *handlers) UndoLike(_ context.Context, in *pb.UndoLikeRequest) (*empty.Empty, error) {
	return &emptypb.Empty{}, h.srv.UndoLike(in.LikerId, in.PostId)
}
