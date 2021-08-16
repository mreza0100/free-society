package connections

import (
	"time"

	"google.golang.org/grpc"
)

func getGRPCDefaultOptions(timeout time.Duration, extra ...grpc.DialOption) (opts []grpc.DialOption) {
	opts = append(opts, extra...)

	opts = append(opts, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(timeout))

	return
}
