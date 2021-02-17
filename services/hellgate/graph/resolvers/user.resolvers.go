package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	securityPb "microServiceBoilerplate/proto/generated/security"
	"microServiceBoilerplate/proto/generated/user"
	models "microServiceBoilerplate/services/hellgate/graph/model"
	"microServiceBoilerplate/services/hellgate/security"
	"microServiceBoilerplate/services/hellgate/validation"
	"microServiceBoilerplate/utils"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.CreateUserInput) (int, error) {
	err := validation.CreateUser(&input)
	if err != nil {
		return 0, err
	}

	userRes, err := r.userConn.NewUser(ctx, &user.NewUserRequest{
		Name:   input.Name,
		Email:  input.Email,
		Gender: input.Gender,
	})
	if err != nil {
		return 0, utils.GetGRPCMSG(err)
	}
	{
		securityRes, err := r.SecurityConn.NewUser(ctx, &securityPb.NewUserRequest{
			UserId:   userRes.Id,
			Password: input.Password,
			Device:   "",
		})
		if err != nil {
			return 0, err
		}
		security.SetToken(ctx, securityRes.GetToken())
	}

	return int(userRes.Id), nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context) (bool, error) {
	userId := security.GetUserId(ctx)

	response, err := r.userConn.DeleteUser(ctx, &user.DeleteUserRequest{
		Id: userId,
	})

	security.DeleteToken(ctx)

	return response.GetOk(), utils.GetGRPCMSG(err)
}

func (r *queryResolver) GetUser(ctx context.Context, id int) (*models.User, error) {
	response, err := r.userConn.GetUser(ctx, &user.GetUserRequest{
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
