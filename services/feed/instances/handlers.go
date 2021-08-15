package instances

import (
	pb "freeSociety/proto/generated/feed"
)

type Handlers interface {
	pb.FeedServiceServer
}
