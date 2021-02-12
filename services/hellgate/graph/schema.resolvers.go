package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"microServiceBoilerplate/proto/generated/post"
	"microServiceBoilerplate/proto/generated/relation"
	pb "microServiceBoilerplate/proto/generated/user"
	"microServiceBoilerplate/services/hellgate/graph/generated"
	models "microServiceBoilerplate/services/hellgate/graph/model"
	"microServiceBoilerplate/services/hellgate/security"
	"microServiceBoilerplate/services/hellgate/validation"
	"microServiceBoilerplate/utils"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.CreateUserInput) (*models.CreateUserRes, error) {
	err := validation.CreateUser(&input)
	if err != nil {
		return nil, err
	}

	response, err := r.userConn.NewUser(ctx, &pb.NewUserRequest{
		Name:     input.Name,
		Email:    input.Email,
		Gender:   input.Gender,
		Password: input.Password,
	})

	if err != nil {
		return nil, utils.GetGRPCMSG(err)
	}

	err = security.SetToken(ctx, response.Id)
	if err != nil {
		return nil, err
	}

	id := int(response.Id)
	return &models.CreateUserRes{
		ID: &id,
	}, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context) (bool, error) {
	userId := security.GetUserId(ctx)

	response, err := r.userConn.DeleteUser(ctx, &pb.DeleteUserRequest{
		Id: userId,
	})

	security.DeleteToken(ctx)

	return response.GetOk(), utils.GetGRPCMSG(err)
}

func (r *mutationResolver) Follow(ctx context.Context, following int) (bool, error) {
	userId := security.GetUserId(ctx)

	_, err := r.RelationConn.Follow(ctx, &relation.FollowRequest{
		Follower:  userId,
		Following: uint64(following),
	})

	return err == nil, utils.GetGRPCMSG(err)
}

func (r *mutationResolver) Unfollow(ctx context.Context, following int) (bool, error) {
	userId := security.GetUserId(ctx)

	_, err := r.RelationConn.Unfollow(ctx, &relation.UnfollowRequest{
		Follower:  userId,
		Following: uint64(following),
	})

	return err == nil, utils.GetGRPCMSG(err)
}

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

	return int(response.Id), utils.GetGRPCMSG(err)
}

func (r *mutationResolver) DeletePost(ctx context.Context, input models.DeletePostInput) (bool, error) {
	userId := security.GetUserId(ctx)

	_, err := r.postConn.DeletePost(ctx, &post.DeletePostRequest{
		PostId: uint64(input.PostID),
		UserId: userId,
	})

	return err != nil, utils.GetGRPCMSG(err)
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (bool, error) {
	response, err := r.userConn.Validation(ctx, &pb.ValidationRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return false, utils.GetGRPCMSG(err)
	}

	err = security.SetToken(ctx, response.GetId())
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	security.DeleteToken(ctx)

	return true, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id int) (*models.User, error) {
	response, err := r.userConn.GetUser(ctx, &pb.GetUserRequest{
		Id: uint64(id),
	})
	if err != nil {
		return nil, utils.GetGRPCMSG(err)
	}

	return &models.User{
		ID:     id,
		Name:   response.Name,
		Email:  response.Email,
		Gender: response.Gender,
	}, nil
}

func (r *queryResolver) GetPost(ctx context.Context, input models.GetPostInput) ([]*models.GetPostRes, error) {
	if len(input.PostIds) > 50 {
		return []*models.GetPostRes{}, errors.New("too many ids")
	}

	ids := make([]uint64, len(input.PostIds))
	{
		for i := 0; i < len(input.PostIds); i++ {
			ids[i] = uint64(input.PostIds[i])
		}
	}
	response, err := r.postConn.GetPost(ctx, &post.GetPostRequest{
		Ids: ids,
	})
	if err != nil {
		return nil, utils.GetGRPCMSG(err)
	}
	result := make([]*models.GetPostRes, len(response.Posts))

	{
		for i := 0; i < len(response.Posts); i++ {
			resVal := response.Posts[i]
			result[i] = &models.GetPostRes{
				Title:   resVal.Title,
				Body:    resVal.Body,
				OwnerID: int(resVal.OwnerId),
				ID:      int(resVal.Id),
			}
		}
	}

	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
