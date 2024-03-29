package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"freeSociety/configs"
	"freeSociety/proto/generated/feed"
	"freeSociety/proto/generated/post"
	"freeSociety/proto/generated/relation"
	models "freeSociety/services/hellgate/graph/model"
	"freeSociety/services/hellgate/security"
	"freeSociety/services/hellgate/validation"
	"freeSociety/utils"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input models.CreatePostInput) (string, error) {
	var (
		userId      = security.GetUserId(ctx)
		rawPictures = make([]*graphql.Upload, 0, 4)
		pictures    = make([]*post.Picture, 0, 4)
	)

	{
		if input.Image1 != nil {
			rawPictures = append(rawPictures, input.Image1)
		}
		if input.Image2 != nil {
			rawPictures = append(rawPictures, input.Image2)
		}
		if input.Image3 != nil {
			rawPictures = append(rawPictures, input.Image3)
		}
		if input.Image4 != nil {
			rawPictures = append(rawPictures, input.Image4)
		}
	}
	{ // check if the user is allowed to upload this pictures
		for _, image := range rawPictures {
			if image.Size > configs.Picture_size_limit {
				return "", errors.New("image size is too large")
			}
		}

		for _, image := range rawPictures {
			// check image content type to be exacly an image
			if image.ContentType != "image/jpeg" && image.ContentType != "image/png" {
				return "", errors.New("image type is not a image")
			}
		}
	}
	{
		if err := validation.CreatePost(&input); err != nil {
			return "", err
		}
	}
	{
		for _, image := range rawPictures {
			pictureContent, err := io.ReadAll(image.File)
			if err != nil {
				return "", err
			}
			pictures = append(pictures, &post.Picture{
				Name:    image.Filename,
				Content: pictureContent,
			})
		}
	}

	response, err := r.postConn.NewPost(ctx, &post.NewPostRequest{
		Title:    input.Title,
		Body:     input.Body,
		UserId:   userId,
		Pictures: pictures,
	})

	return response.Id, utils.GetGRPCMSG(err)
}

func (r *mutationResolver) DeletePost(ctx context.Context, postID string) (bool, error) {
	userId := security.GetUserId(ctx)

	_, err := r.postConn.DeletePost(ctx, &post.DeletePostRequest{
		PostId: postID,
		UserId: userId,
	})

	return err == nil, utils.GetGRPCMSG(err)
}

func (r *mutationResolver) Like(ctx context.Context, postID string, ownerID int) (bool, error) {
	userId := security.GetUserId(ctx)

	_, err := r.relationConn.Like(ctx, &relation.LikeRequest{
		LikerId: userId,
		PostId:  postID,
		OwnerId: uint64(ownerID),
	})

	return err == nil, nil
}

func (r *mutationResolver) UndoLike(ctx context.Context, postID string) (bool, error) {
	userId := security.GetUserId(ctx)

	_, err := r.relationConn.UndoLike(ctx, &relation.UndoLikeRequest{
		LikerId: userId,
		PostId:  postID,
	})

	return err == nil, nil
}

func (r *mutationResolver) ResharePost(ctx context.Context, postID string) (bool, error) {
	userId := security.GetUserId(ctx)

	_, err := r.feedConn.Reshare(ctx, &feed.ReshareRequest{
		UserId: userId,
		PostId: postID,
	})

	return err == nil, err
}

func (r *queryResolver) GetPost(ctx context.Context, postIds []string) ([]*models.Post, error) {
	var (
		uIds        []string
		requestorId uint64
		rawPosts    []*post.Post
		posts       []*models.Post
	)

	{
		if len(postIds) > 50 {
			return nil, errors.New("too many postIds")
		}
		requestorId, _ = security.GetOptionalId(ctx)
	}

	{
		uIds = make([]string, len(postIds))
		{
			for i := 0; i < len(postIds); i++ {
				uIds[i] = postIds[i]
			}
		}
	}
	{
		response, err := r.postConn.GetPost(ctx, &post.GetPostRequest{
			Ids:         uIds,
			RequestorId: requestorId,
		})
		if err != nil {
			return nil, utils.GetGRPCMSG(err)
		}
		rawPosts = response.Posts
	}
	{
		posts = make([]*models.Post, len(rawPosts))
		for idx, p := range rawPosts {
			posts[idx] = &models.Post{
				Title: p.Title,
				Body:  p.Body,
				ID:    p.Id,

				OwnerID:     int(p.OwnerId),
				Likes:       int(p.Likes),
				IsLiked:     p.IsLiked,
				PictureUrls: p.PictureUrls,

				User: &models.User{
					IsFollowing: p.IsFollowing,

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

func (r *queryResolver) GetFeed(ctx context.Context, offset int, limit int) ([]*models.Post, error) {
	var (
		userId   uint64
		postIds  []string
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
				Title: p.Title,
				Body:  p.Body,
				ID:    p.Id,

				OwnerID:     int(p.OwnerId),
				Likes:       int(p.Likes),
				IsLiked:     p.IsLiked,
				PictureUrls: p.PictureUrls,

				User: &models.User{
					IsFollowing: p.IsFollowing,

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
