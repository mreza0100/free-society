package instances

import (
	pb "freeSociety/proto/generated/post"
)

type Handlers interface {
	pb.PostServiceServer
}
