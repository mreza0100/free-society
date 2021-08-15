package instances

import (
	pb "freeSociety/proto/generated/security"
)

type Handlers interface {
	pb.SecurityServiceServer
}
