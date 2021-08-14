package configs

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type subjectsT struct {
	NewPost          string
	DeleteUser       string
	IsUserExist      string
	GetFollowers     string
	GetPosts         string
	GetUsersByIds    string
	IsFollowingGroup string
	IsPostsExists    string
	IsLikedGroup     string
	CountLikes       string
}

type natsConfigsT struct {
	Url            string
	TotalWait      time.Duration
	ReconnectDelay time.Duration
	Subjects       *subjectsT
	Timeout        time.Duration
}

func (nConf *natsConfigsT) GetDefaultNatsOpts(name string) []nats.Option {
	opts := make([]nats.Option, 0, 7)

	disconnectErrHandlerO := nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to: %s, will attempt reconnects for %.0fm", err, nConf.TotalWait.Minutes())
		log.Print("Retrying...")
	})
	reconnectHandlerO := nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	})
	closedHandlerO := nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	})

	opts = append(
		opts,

		nats.Name(name),
		nats.ReconnectWait(nConf.ReconnectDelay),
		nats.MaxReconnects(int(nConf.TotalWait/nConf.ReconnectDelay)),
		disconnectErrHandlerO,
		reconnectHandlerO,
		closedHandlerO,
	)

	return opts
}

var Nats *natsConfigsT

func init() {
	sbjs := &subjectsT{}

	Nats = &natsConfigsT{
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
	}
}
