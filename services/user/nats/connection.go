package userNats

import (
	"log"
	"microServiceBoilerplate/configs"

	"github.com/nats-io/nats.go"
)

var nc *nats.Conn

const natName = "User Service"

func init() {
	nConnection, err := nats.Connect(
		configs.NatsConfigs.Url,
		configs.NatsConfigs.GetDefaultNatsOpts(natName)...,
	)
	if err != nil {
		log.Fatal("✖✖✖From nats connection: cant connect to nats server exiting NOW✖✖✖", "\t", err)
	}

	nc = nConnection
}
