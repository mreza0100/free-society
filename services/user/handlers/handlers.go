package handlers

import (
	pb "freeSociety/proto/generated/user"

	"freeSociety/services/user/instances"

	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Lgr        *golog.Core
	Srv        instances.Sevice
	Publishers instances.Publishers
}

func New(opts *NewOpts) pb.UserServiceServer {
	return &handlers{
		lgr:        opts.Lgr.With("In handlers->"),
		srv:        opts.Srv,
		publishers: opts.Publishers,
	}
}

type handlers struct {
	lgr        *golog.Core
	srv        instances.Sevice
	publishers instances.Publishers

	pb.UnimplementedUserServiceServer
}
