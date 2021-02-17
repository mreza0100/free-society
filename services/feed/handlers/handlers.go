package handlers

import (
	"context"
	pb "microServiceBoilerplate/proto/generated/feed"
	"microServiceBoilerplate/services/feed/instances"

	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Srv        instances.Sevice
	Lgr        *golog.Core
	Publishers instances.Publishers
}

func New(opts *NewOpts) instances.Handlers {
	return &handlers{
		srv:        opts.Srv,
		lgr:        opts.Lgr,
		publishers: opts.Publishers,
	}
}

type handlers struct {
	srv        instances.Sevice
	lgr        *golog.Core
	publishers instances.Publishers

	pb.UnimplementedFeedServiceServer
}

func (s *handlers) GetFeed(_ context.Context, in *pb.GetFeedRequest) (*pb.GetFeedResponse, error) {
	postIds, err := s.srv.GetFeed(in.UserId, in.Offset, in.Limit)

	return &pb.GetFeedResponse{
		PostIds: postIds,
	}, err
}
