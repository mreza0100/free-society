package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"microServiceBoilerplate/proto/generated/feed"
	"microServiceBoilerplate/proto/generated/post"
	"microServiceBoilerplate/proto/generated/relation"
	models "microServiceBoilerplate/services/hellgate/graph/model"
	"microServiceBoilerplate/services/hellgate/security"
	"microServiceBoilerplate/services/hellgate/validation"
	"microServiceBoilerplate/utils"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input models.CreatePostInput) (int, error) {
	userId := security.GetUserId(ctx)

	err := validation.CreatePost(&input)
	if err != nil {
		return 0, err
	}

	response, err := r.postConn.NewPost(ctx, &post.NewPostRequest{
		Title:  input.Title,
		Body:   input.Body,
		UserId: userId,
	})
	if err != nil {
		return 0, utils.GetGRPCMSG(err)
	}

	return int(response.Id), nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, input models.DeletePostInput) (bool, error) {
	userId := security.GetUserId(ctx)

	_, err := r.postConn.DeletePost(ctx, &post.DeletePostRequest{
		PostId: uint64(input.PostID),
		UserId: userId,
	})

	return err != nil, utils.GetGRPCMSG(err)
}

func (r *mutationResolver) Like(ctx context.Context, postID int, ownerID int) (bool, error) {
	userId := security.GetUserId(ctx)

	_, err := r.relationConn.Like(ctx, &relation.LikeRequest{
		LikerId: userId,
		PostId:  uint64(postID),
		OwnerId: uint64(ownerID),
	})

	return err == nil, nil
}

func (r *mutationResolver) UndoLike(ctx context.Context, postID int) (bool, error) {
	userId := security.GetUserId(ctx)

	_, err := r.relationConn.UndoLike(ctx, &relation.UndoLikeRequest{
		LikerId: userId,
		PostId:  uint64(postID),
	})

	return err == nil, nil
}

func (r *queryResolver) GetPost(ctx context.Context, input []int) ([]*models.Post, error) {
	var (
		ids         []uint64
		requestorId uint64
		rawPosts    []*post.Post
		result      []*models.Post
	)

	{
		if len(input) > 50 {
			return []*models.Post{}, errors.New("too many ids")
		}
		requestorId, _ = security.GetOptionalId(ctx)
	}

	{
		ids = make([]uint64, len(input))
		{
			for i := 0; i < len(input); i++ {
				ids[i] = uint64(input[i])
			}
		}
	}
	{
		response, err := r.postConn.GetPost(ctx, &post.GetPostRequest{
			Ids:         ids,
			RequestorId: requestorId,
		})
		if err != nil {
			return nil, utils.GetGRPCMSG(err)
		}
		rawPosts = response.Posts
		r.Lgr.InfoLog(rawPosts)
	}
	{
		result = make([]*models.Post, len(rawPosts))

		for idx, p := range rawPosts {
			result[idx] = &models.Post{
				Title:       p.Title,
				Body:        p.Body,
				ID:          int(p.Id),
				OwnerID:     int(p.OwnerId),
				IsFollowing: p.IsFollowing,
				User: &models.User{
					ID:     int(p.User.Id),
					Name:   p.User.Name,
					Email:  p.User.Email,
					Gender: p.User.Gender,
				},
			}
		}
	}

	return result, nil
}

func (r *queryResolver) GetFeed(ctx context.Context, offset int, limit int) ([]*models.Post, error) {
	var (
		userId   uint64
		postIds  []uint64
		rawPosts []*post.Post
		posts    []*models.Post
	)

	{
		if limit > 50 {
			return nil, errors.New("limit must be less then 50")
		}
		userId = security.GetUserId(ctx)
	}
	{
		response, err := r.feedConn.GetFeed(ctx, &feed.GetFeedRequest{
			UserId: userId,
			Offset: uint64(offset),
			Limit:  uint64(limit),
		})
		if err != nil {
			return nil, err
		}
		postIds = response.PostIds
	}
	{
		response, err := r.postConn.GetPost(ctx, &post.GetPostRequest{
			Ids:         postIds,
			RequestorId: userId,
		})
		if err != nil {
			return nil, err
		}
		rawPosts = response.Posts
	}

	{
		posts = make([]*models.Post, len(rawPosts))
		for idx, p := range rawPosts {
			posts[idx] = &models.Post{
				Title:       p.Title,
				Body:        p.Body,
				ID:          int(p.Id),
				OwnerID:     int(p.OwnerId),
				IsFollowing: p.IsFollowing,
				User: &models.User{
					ID:     int(p.User.Id),
					Name:   p.User.Name,
					Email:  p.User.Email,
					Gender: p.User.Gender,
				},
			}
		}
	}

	return posts, nil
}
