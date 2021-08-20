package main_test

import (
	"context"
	"errors"
	"fmt"
	"freeSociety/connections"
	"freeSociety/proto/generated/post"
	"freeSociety/proto/generated/relation"
	"freeSociety/utils/test"
	"sync"
	"testing"

	"github.com/mreza0100/golog"
)

var (
	logger = golog.New(golog.InitOpns{
		LogPath:   "",
		Name:      "post test",
		WithTime:  false,
		DebugMode: true,
	})
	postConn     = connections.PostSrvConn(logger)
	relationConn = connections.RelationSrvConn(logger)

	ctx = context.Background()
)

const (
	countPost = 10

	title  = "test title"
	body   = "test body"
	name   = "mamad"
	email  = "mamad123@mamad.com"
	gender = "male"
)

var (
	userId  uint64
	postIds []uint64
)

func TestMain(m *testing.M) {
	var deleteUser func()
	userId, deleteUser = test.NewUser(test.NewUserOpns{
		Lgr:    logger,
		Name:   name,
		Email:  email,
		Gender: gender,
	})
	defer deleteUser()

	m.Run()
}

func Test_createPost(t *testing.T) {
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
	var wg sync.WaitGroup

	for _, postId := range postIds {
		wg.Add(1)
		t.Run(fmt.Sprintf("delete post id: %v", postId), func(t *testing.T) {

			go func(id uint64) {
				defer wg.Done()

				_, err := postConn.DeletePost(ctx, &post.DeletePostRequest{
					PostId: postId,
					UserId: userId,
				})
				test.CheckFail(t, err)
			}(postId)

		})
	}
}
