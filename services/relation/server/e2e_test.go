package main_test

import (
	"context"
	"errors"
	"freeSociety/connections"
	"freeSociety/proto/generated/feed"
	"freeSociety/proto/generated/post"
	"freeSociety/proto/generated/relation"
	"freeSociety/services/relation/domain"
	"freeSociety/utils/test"
	"testing"
	"time"

	"github.com/mreza0100/golog"
)

const (
	title = "test title"
	body  = "test body"
)

const (
	user1Name   = "mamad1"
	user1Email  = "mamad1@mamad.com"
	user1Gender = "male"

	user2Name   = "mamad2"
	user2Email  = "mamad2@mamad.com"
	user2Gender = "male"
)

var (
	logger = golog.New(golog.InitOpns{
		LogPath:   "",
		Name:      "relation test",
		WithTime:  false,
		DebugMode: true,
	})

	services = domain.New(&domain.NewOpts{
		Lgr: logger,
	})

	postConn     = connections.PostSrvConn(logger)
	relationConn = connections.RelationSrvConn(logger)
	feedConn     = connections.FeedSrvConn(logger)

	ctx = context.Background()
)

var (
	user1 uint64
	user2 uint64

	user1Post uint64
	user2Post uint64
)

func TestMain(m *testing.M) {
	{
		var deleteUser func()
		user1, deleteUser = test.NewUser(test.NewUserOpns{
			Lgr:    logger,
			Name:   user1Name,
			Email:  user1Email,
			Gender: user1Gender,
		})
		defer deleteUser()
	}
	{
		var deleteUser func()
		user2, deleteUser = test.NewUser(test.NewUserOpns{
			Lgr:    logger,
			Name:   user2Name,
			Email:  user2Email,
			Gender: user2Gender,
		})
		defer deleteUser()
	}
	m.Run()
}

func Test_follow(t *testing.T) {
	{
		{
			_, err := relationConn.Follow(ctx, &relation.FollowRequest{
				Following: user2,
				Follower:  user1,
			})
			test.FailIf(t, err != nil, err)
		}

		{
			followers := services.GetFollowers(user2)
			test.FailIf(t, len(followers) != 1, errors.New("len(followers) != 1"))
			test.FailIf(t, followers[0] != user1, errors.New("followers[0] != followerId"))
		}
	}
	{
		{
			_, err := relationConn.Follow(ctx, &relation.FollowRequest{
				Following: user1,
				Follower:  user2,
			})
			test.FailIf(t, err != nil, err)
		}

		{
			followers := services.GetFollowers(user1)
			test.FailIf(t, len(followers) != 1, errors.New("len(followers) != 1"))
			test.FailIf(t, followers[0] != user2, errors.New("followers[0] != followerId"))
		}
	}
}

func Test_unFollow(t *testing.T) {
	// follow it again
	// we have works to do...
	defer Test_follow(t)
	{
		_, err := relationConn.Unfollow(ctx, &relation.UnfollowRequest{
			Following: user1,
			Follower:  user2,
		})
		test.FailIf(t, err != nil, err)
	}
	{
		_, err := relationConn.Unfollow(ctx, &relation.UnfollowRequest{
			Following: user2,
			Follower:  user1,
		})
		test.FailIf(t, err != nil, err)
	}
}

func Test_createPostAndCheckFeed(t *testing.T) {
	{
		{
			response, err := postConn.NewPost(ctx, &post.NewPostRequest{
				Title:  title,
				Body:   body,
				UserId: user1,
			})
			test.FailIf(t, err != nil, err)
			user1Post = response.Id
		}
		{
			response, err := postConn.NewPost(ctx, &post.NewPostRequest{
				Title:  title,
				Body:   body,
				UserId: user2,
			})
			test.FailIf(t, err != nil, err)
			user2Post = response.Id
		}
	}
	{
		// postConn.NewPost will publish a newPost event on nats network
		// in this while feed subscribers are listining to newPost event
		// its take a little time for feed to hear the event and write it to redis...
		time.Sleep(time.Second)
	}
	{
		{
			response, err := feedConn.GetFeed(ctx, &feed.GetFeedRequest{
				UserId: user1,
				Offset: 0,
				Limit:  50,
			})
			test.FailIf(t, err != nil, err)
			test.FailIf(t, len(response.PostIds) != 1, errors.New("len was't 1"))
			test.FailIf(t, response.PostIds[0] != user2Post, errors.New("postId from feed != generated postId"))
		}
		{
			response, err := feedConn.GetFeed(ctx, &feed.GetFeedRequest{
				UserId: user2,
				Offset: 0,
				Limit:  50,
			})
			test.FailIf(t, err != nil, err)
			test.FailIf(t, len(response.PostIds) != 1, errors.New("len was't 1"))
			test.FailIf(t, response.PostIds[0] != user1Post, errors.New("postId from feed != generated postId"))
		}
	}
}

func Test_like(t *testing.T) {
	{
		_, err := relationConn.Like(ctx, &relation.LikeRequest{
			LikerId: user1,
			OwnerId: user2,
			PostId:  user2Post,
		})
		test.FailIf(t, err != nil, err)
	}

	{
		_, err := relationConn.Like(ctx, &relation.LikeRequest{
			LikerId: user2,
			OwnerId: user1,
			PostId:  user1Post,
		})
		test.FailIf(t, err != nil, err)
	}
}

func Test_undoLike(t *testing.T) {
	// $there is no reason for this one =)))
	defer Test_like(t)
	{
		_, err := relationConn.UndoLike(ctx, &relation.UndoLikeRequest{
			LikerId: user1,
			PostId:  user2Post,
		})
		test.FailIf(t, err != nil, err)
	}

	{
		_, err := relationConn.UndoLike(ctx, &relation.UndoLikeRequest{
			LikerId: user2,
			PostId:  user1Post,
		})
		test.FailIf(t, err != nil, err)
	}
}
