package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"freeSociety/proto/generated/relation"
	"freeSociety/services/hellgate/security"
	"freeSociety/utils"
)

func (r *mutationResolver) Follow(ctx context.Context, following int) (bool, error) {
	userId := security.GetUserId(ctx)

	_, err := r.relationConn.Follow(ctx, &relation.FollowRequest{
		Follower:  userId,
		Following: uint64(following),
	})

	return err == nil, utils.GetGRPCMSG(err)
}

func (r *mutationResolver) Unfollow(ctx context.Context, following int) (bool, error) {
	userId := security.GetUserId(ctx)

	_, err := r.relationConn.Unfollow(ctx, &relation.UnfollowRequest{
		Follower:  userId,
		Following: uint64(following),
	})

	return err == nil, utils.GetGRPCMSG(err)
}
