package handlers

import (
	"context"
	pb "freeSociety/proto/generated/relation"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *handlers) UndoLike(_ context.Context, in *pb.UndoLikeRequest) (*empty.Empty, error) {
	return &emptypb.Empty{}, h.srv.UndoLike(in.LikerId, in.PostId)
}
