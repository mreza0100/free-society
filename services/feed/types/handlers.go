package types

import (
	pb "microServiceBoilerplate/proto/generated/feed"
)

type Handlers interface {
	pb.FeedServiceServer
}
