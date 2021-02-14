package types

import (
	pb "microServiceBoilerplate/proto/generated/post"
)

type Handlers interface {
	pb.PostServiceServer
}
