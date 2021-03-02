package handlers

import (
	pb "microServiceBoilerplate/proto/generated/relation"
	"microServiceBoilerplate/services/relation/instances"

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
		lgr:        opts.Lgr.With("In handlers->"),
		publishers: opts.Publishers,
	}
}

type handlers struct {
	srv        instances.Sevice
	lgr        *golog.Core
	publishers instances.Publishers

	pb.UnimplementedRelationServiceServer
}
