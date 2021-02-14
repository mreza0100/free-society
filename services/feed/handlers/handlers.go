package handlers

import (
	"context"
	pb "microServiceBoilerplate/proto/generated/feed"
	"microServiceBoilerplate/services/feed/types"

	"github.com/mreza0100/golog"
)

type HandlersOpts struct {
	Srv        types.Sevice
	Lgr        *golog.Core
	Publishers types.Publishers
}

func NewHandlers(opts *HandlersOpts) types.Handlers {
	return &handlers{
		srv:        opts.Srv,
		lgr:        opts.Lgr,
		publishers: opts.Publishers,
	}
}

type handlers struct {
	srv        types.Sevice
	lgr        *golog.Core
	publishers types.Publishers

	pb.UnimplementedFeedServiceServer
}

func (s *handlers) GetFeed(_ context.Context, in *pb.GetFeedRequest) (*pb.GetFeedResponse, error) {
	postIds, err := s.srv.GetFeed(in.UserId, in.Offset, in.Limit)

	return &pb.GetFeedResponse{
		PostIds: postIds,
	}, err
}
