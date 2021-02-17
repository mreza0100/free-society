package instances

import (
	pb "microServiceBoilerplate/proto/generated/feed"
)

type Handlers interface {
	pb.FeedServiceServer
}
