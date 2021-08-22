package handlers

import (
	pb "freeSociety/proto/generated/feed"
	"freeSociety/services/feed/instances"

	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Srv        instances.Sevice
	Lgr        *golog.Core
	Publishers instances.Publishers
}

func New(opts *NewOpts) pb.FeedServiceServer {
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
