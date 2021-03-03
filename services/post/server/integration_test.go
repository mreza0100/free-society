package main_test

import (
	"context"
	"errors"
	"fmt"
	"microServiceBoilerplate/configs"
	"microServiceBoilerplate/proto/connections"
	"microServiceBoilerplate/proto/generated/post"
	"microServiceBoilerplate/proto/generated/relation"
	"microServiceBoilerplate/proto/generated/security"
	"microServiceBoilerplate/proto/generated/user"
	"microServiceBoilerplate/utils/test"
	"testing"

	"github.com/mreza0100/golog"
)

var (
	logger = golog.New(golog.InitOpns{
		LogPath:      configs.LogPath,
		Name:         "post service server_test",
		WithTime:     false,
		DebugMode:    true,
		ClearLogFile: false,
	})
	postConn     = connections.PostSrvConn(logger)
	userConn     = connections.UserSrvConn(logger)
	securityConn = connections.SecuritySrvConn(logger)
	relationConn = connections.RelationSrvConn(logger)

	ctx = context.Background()
)

const (
	countPost = 10

	title  = "test title"
	body   = "test body"
	name   = "mamad"
	email  = "mamad@mamad.com"
	gender = "male"
)

var (
	userId  uint64
	postIds []uint64
)

func start(t *testing.T) {
	{
		response, err := userConn.NewUser(ctx, &user.NewUserRequest{
			Name:   name,
			Email:  email,
			Gender: gender,
		})
		test.CheckFail(t, err)
		if response.Id == 0 {
			test.CheckFail(t, errors.New("user id was 0"))
		}
		userId = response.Id
	}
	{
		_, err := securityConn.NewUser(ctx, &security.NewUserRequest{
			UserId:   userId,
			Password: "8888888",
			Device:   "mamads device",
		})
		test.FailIf(t, err != nil, err)
	}
}

func Test_createPost(t *testing.T) {
	start(t)
	postIds = make([]uint64, 0, countPost)

	for i := 0; i < countPost; i++ {
		t.Run(fmt.Sprintf("creat post round: %v", i), func(t *testing.T) {
			t.Parallel()

			response, err := postConn.NewPost(ctx, &post.NewPostRequest{
				Title:  title,
				Body:   body,
				UserId: userId,
			})
			test.CheckFail(t, err)

			postIds = append(postIds, response.Id)
		})
	}

}

func Test_getPostsMultipleTimesAndCheck(t *testing.T) {
	checkData := func(post *post.Post) {
		test.FailIf(t, post.Title != title, errors.New("post.Title != title was true"))
		test.FailIf(t, post.Body != body, errors.New("post.Body!= body was true"))

		test.FailIf(t, post.IsFollowing != false, errors.New("post.IsFollowing was true"))
		test.FailIf(t, post.IsLiked != false, errors.New("post.IsLiked was true"))
		test.FailIf(t, post.Likes != 0, errors.New("post.Likes != 0 was true"))
		test.FailIf(t, post.OwnerId != userId, errors.New("post.OwnerId != userId was true"))

		test.FailIf(t, post.User.Id != post.OwnerId, errors.New("post.User.Id != post.OwnerId was true"))
		test.FailIf(t, post.User.Name != name, errors.New("post.User.Name != name was true"))
		test.FailIf(t, post.User.Email != email, errors.New("post.User.Email != email was true"))
		test.FailIf(t, post.User.Gender != gender, errors.New("post.User.Gender was true"))
	}

	t.Run("get posts and check | with requestor", func(t *testing.T) {
		response, err := postConn.GetPost(ctx, &post.GetPostRequest{
			RequestorId: userId,
			Ids:         postIds,
		})
		test.CheckFail(t, err)

		for _, post := range response.Posts {
			checkData(post)
		}
	})

	t.Run("get posts and check | without requestor", func(t *testing.T) {
		response, err := postConn.GetPost(ctx, &post.GetPostRequest{
			RequestorId: 0,
			Ids:         postIds,
		})
		test.CheckFail(t, err)

		for _, post := range response.Posts {
			checkData(post)
		}
	})

	t.Run("like a post and check like with requestorId and without that", func(t *testing.T) {
		subject := postIds[countPost/2]

		{
			_, err := relationConn.Like(ctx, &relation.LikeRequest{
				LikerId: userId,
				OwnerId: userId,
				PostId:  subject,
			})
			test.FailIf(t, err != nil, err)
		}
		t.Run("check with requestorId", func(t *testing.T) {
			response, err := postConn.GetPost(ctx, &post.GetPostRequest{
				RequestorId: userId,
				Ids:         []uint64{subject},
			})
			post := response.Posts[0]

			test.FailIf(t, err != nil, err)
			test.FailIf(t, len(response.Posts) != 1, errors.New("len was't 1"))
			test.FailIf(t, post.Likes != 1, errors.New("likes was't 1"))
			test.FailIf(t, post.IsLiked != true, errors.New("post did't liked"))
		})

		t.Run("check without requestorId", func(t *testing.T) {
			response, err := postConn.GetPost(ctx, &post.GetPostRequest{
				RequestorId: 0,
				Ids:         []uint64{subject},
			})
			post := response.Posts[0]

			test.FailIf(t, err != nil, err)
			test.FailIf(t, len(response.Posts) != 1, errors.New("len was't 1"))
			test.FailIf(t, post.Likes != 1, errors.New("likes was't 1"))
			test.FailIf(t, post.IsLiked == true, errors.New("post did't liked"))
		})
	})
}

func Test_deletePost(t *testing.T) {
	defer func() {
		_, err := userConn.DeleteUser(ctx, &user.DeleteUserRequest{
			Id: userId,
		})
		test.CheckFail(t, err)
	}()

	for _, postId := range postIds {
		t.Run(fmt.Sprintf("delete post id: %v", postId), func(t *testing.T) {

			go func(id uint64) {
				_, err := postConn.DeletePost(ctx, &post.DeletePostRequest{
					PostId: postId,
					UserId: userId,
				})
				test.CheckFail(t, err)
			}(postId)

		})
	}

}
