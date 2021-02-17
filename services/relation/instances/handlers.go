package instances

import (
	pb "microServiceBoilerplate/proto/generated/relation"
)

type Handlers interface {
	pb.RelationServiceServer
}
