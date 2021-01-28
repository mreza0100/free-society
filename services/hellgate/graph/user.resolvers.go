package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	pb "microServiceBoilerplate/proto/generated/user"
	"microServiceBoilerplate/services/hellgate/graph/generated"
	models "microServiceBoilerplate/services/hellgate/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.CreateUserInput) (*models.CreateUserRes, error) {
	response, err := r.userConn.NewUser(ctx, &pb.NewUserRequest{
		Name:     input.Name,
		Email:    input.Email,
		Gender:   input.Gender,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}

	id := (int)(response.Id)

	return &models.CreateUserRes{
		ID: &id,
	}, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (bool, error) {
	response, err := r.userConn.DeleteUserById(ctx, &pb.DeleteUserByIdRequest{
		Id: uint64(id),
	})

	return response.GetOk(), err
}

func (r *queryResolver) GetUser(ctx context.Context, id int) (*models.User, error) {
	response, err := r.userConn.GetUserById(ctx, &pb.GetUserByIdRequest{
		Id: uint64(id),
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:     id,
		Name:   response.Name,
		Email:  response.Email,
		Gender: response.Gender,
	}, nil
}

func (r *queryResolver) GetUsers(ctx context.Context) (*models.GetUserRes, error) {
	response, err := r.userConn.GetUsers(ctx, &pb.GetUsersRequest{})
	if err != nil {
		return nil, err
	}

	rawUsers := response.Users
	users := make([]*models.User, len(rawUsers))

	for idx, rawUser := range rawUsers {
		var user models.User = models.User{
			ID:     int(rawUser.Id),
			Name:   rawUser.Name,
			Email:  rawUser.Email,
			Gender: rawUser.Gender,
		}
		users[idx] = &user
	}

	return &models.GetUserRes{
		Users: users,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
