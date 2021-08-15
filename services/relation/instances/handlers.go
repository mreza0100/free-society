package instances

import (
	pb "freeSociety/proto/generated/relation"
)

type Handlers interface {
	pb.RelationServiceServer
}
