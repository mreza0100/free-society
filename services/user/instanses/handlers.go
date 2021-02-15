package instanses

import (
	pb "microServiceBoilerplate/proto/generated/user"
)

type Handlers interface {
	pb.UserServiceServer
}
