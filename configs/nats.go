package configs

import (
	"time"

	"github.com/nats-io/nats.go"
)

type subjectsT struct {
	NewPost             string
	DeleteUser          string
	IsUserExist         string
	GetFollowers        string
	GetPosts            string
	GetUsersByIds       string
	IsFollowingGroup    string
	IsPostsExists       string
	IsLikedGroup        string
	CountLikes          string
	SetLikeNotification string
}

type NatsConfigsT struct {
	Url            string
	TotalWait      time.Duration
	ReconnectDelay time.Duration
	Subjects       *subjectsT
	Timeout        time.Duration
}

var NatsConfigs *NatsConfigsT

func init() {
	sbjs := new(subjectsT)

	NatsConfigs = &NatsConfigsT{
		Url:            nats.DefaultURL,
		TotalWait:      2 * time.Minute,
		ReconnectDelay: time.Second,
		Subjects:       sbjs,
		Timeout:        time.Second,
	}

	{
		sbjs.DeleteUser = "user.delete"
		sbjs.NewPost = "post.new"
		sbjs.IsUserExist = "user.is_exist"
		sbjs.GetFollowers = "relation.get_followers"
		sbjs.GetPosts = "post.get"
		sbjs.GetUsersByIds = "user.get_users"
		sbjs.IsFollowingGroup = "relation.is_following_group"
		sbjs.IsPostsExists = "post.is_exists"
		sbjs.IsLikedGroup = "post.is_liked_group"
		sbjs.CountLikes = "post.count_likes"
		sbjs.SetLikeNotification = "notification.set_like"
	}
}
