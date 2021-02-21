package configs

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type SubjectsT struct {
	NewPost             string
	DeleteUser          string
	IsUserExist_REQUEST string
	GetFollowers        string
	GetPosts            string
	GetUsersByIds       string
}

type natsConfigsT struct {
	Url            string
	TotalWait      time.Duration
	ReconnectDelay time.Duration
	Subjects       *SubjectsT
	Timeout        time.Duration
}

func (this *natsConfigsT) GetDefaultNatsOpts(name string) []nats.Option {
	opts := make([]nats.Option, 0, 7)

	disconnectErrHandlerO := nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to: %s, will attempt reconnects for %.0fm", err, this.TotalWait.Minutes())
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
		nats.ReconnectWait(this.ReconnectDelay),
		nats.MaxReconnects(int(this.TotalWait/this.ReconnectDelay)),
		disconnectErrHandlerO,
		reconnectHandlerO,
		closedHandlerO,
	)

	return opts
}

var Nats *natsConfigsT

func init() {
	sbjs := &SubjectsT{}

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
		sbjs.IsUserExist_REQUEST = "user.is_exist"
		sbjs.GetFollowers = "relation.get_followers"
		sbjs.GetPosts = "posts.get"
		sbjs.GetUsersByIds = "user.get_users"
	}
}
