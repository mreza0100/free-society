package instances

import (
	pb "microServiceBoilerplate/proto/generated/security"
)

type Handlers interface {
	pb.SecurityServiceServer
}
