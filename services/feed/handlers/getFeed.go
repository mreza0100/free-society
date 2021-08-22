package handlers

import (
	"context"
	pb "freeSociety/proto/generated/feed"
)

func (s *handlers) GetFeed(_ context.Context, in *pb.GetFeedRequest) (*pb.GetFeedResponse, error) {
	postIds, err := s.srv.GetFeed(in.UserId, in.Offset, in.Limit)

	return &pb.GetFeedResponse{
		PostIds: postIds,
	}, err
}
