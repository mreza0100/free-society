package instances

import natsPb "freeSociety/proto/generated/nats"

type Publishers interface {
	GetFollowers(userId uint64) ([]uint64, error)
	GetPosts(postIds []uint64) ([]*natsPb.Post, error)
}
