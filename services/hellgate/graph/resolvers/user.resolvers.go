package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"freeSociety/configs"
	securityPb "freeSociety/proto/generated/security"
	"freeSociety/proto/generated/user"
	models "freeSociety/services/hellgate/graph/model"
	"freeSociety/services/hellgate/security"
	"freeSociety/services/hellgate/validation"
	"freeSociety/utils"
	"freeSociety/utils/files"
	"io"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (int, error) {
	var (
		avatarFormat = ""
		avatarBytes  = []byte{}
		userId       uint64
	)
	{
		err := validation.CreateUser(&input)
		if err != nil {
			return 0, err
		}
	}
	{
		if input.Avatar != nil {
			if configs.Avatar_max_size < input.Avatar.Size {
				return 0, fmt.Errorf("Avatar size is bigger than %v MB", configs.Avatar_max_size/1024/1024)
			}
			// check type of input.avatar to be a image
			if input.Avatar.ContentType != "image/jpeg" && input.Avatar.ContentType != "image/png" {
				return 0, errors.New("image type is not a image")
			}

			pictureContent, err := io.ReadAll(input.Avatar.File)
			if err != nil {
				return 0, err
			}

			avatarBytes = pictureContent
			avatarFormat = files.GetFileFormat(input.Avatar.Filename)
		}
	}

	{
		userRes, err := r.userConn.NewUser(ctx, &user.NewUserRequest{
			Name:         input.Name,
			Email:        input.Email,
			Gender:       input.Gender,
			Avatar:       avatarBytes,
			AvatarFormat: avatarFormat,
		})
		if err != nil {
			return 0, utils.GetGRPCMSG(err)
		}
		userId = userRes.Id
	}
	{
		securityRes, err := r.SecurityConn.NewUser(ctx, &securityPb.NewUserRequest{
			UserId:   userId,
			Password: input.Password,
			Device:   security.GetUserAgent(ctx),
		})
		if err != nil {
			return 0, err
		}
		security.SetToken(ctx, securityRes.GetToken())
	}

	return int(userId), nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context) (bool, error) {
	userId := security.GetUserId(ctx)
	fmt.Println(userId)

	response, err := r.userConn.DeleteUser(ctx, &user.DeleteUserRequest{
		Id: userId,
	})

	if err == nil {
		security.DeleteToken(ctx)
	}

	return response.GetOk(), utils.GetGRPCMSG(err)
}

func (r *mutationResolver) EditUser(ctx context.Context, userData models.UpdateUserInput) (bool, error) {
	userId := security.GetUserId(ctx)

	var (
		avatarFormat string
		avatarBytes  = []byte{}
	)

	if userData.Avatar != nil {
		var err error

		if userData.Avatar.ContentType != "image/jpeg" && userData.Avatar.ContentType != "image/png" {
			return false, errors.New("image type is not a image")
		}
		avatarFormat = files.GetFileFormat(userData.Avatar.Filename)
		avatarBytes, err = io.ReadAll(userData.Avatar.File)
		if err != nil {
			return false, err
		}
	}

	_, err := r.userConn.UpdateUser(ctx, &user.UpdateUserRequest{
		Id:     userId,
		Name:   userData.Name,
		Gender: userData.Gender,

		AvatarFormat: avatarFormat,
		Avatar:       avatarBytes,
	})

	return err == nil, err
}

func (r *queryResolver) GetUser(ctx context.Context, id int) (*models.User, error) {
	userId := security.GetUserId(ctx)

	response, err := r.userConn.GetUser(ctx, &user.GetUserRequest{
		Id:          uint64(id),
		RequestorId: userId,
	})
	if err != nil {
		return nil, utils.GetGRPCMSG(err)
	}

	return &models.User{
		ID:          id,
		Name:        response.Name,
		Email:       response.Email,
		Gender:      response.Gender,
		IsFollowing: response.IsFollowing,
		Avatar:      response.AvatarPath,
	}, nil
}
