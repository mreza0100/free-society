package securityNats

import (
	"log"
	"microServiceBoilerplate/configs"

	"github.com/nats-io/nats.go"
)

var nc *nats.Conn

const natName = "Security Service"

func init() {
	nConnection, err := nats.Connect(
		configs.Nats.Url,
		configs.Nats.GetDefaultNatsOpts(natName)...,
	)
	if err != nil {
		log.Fatal("✖✖✖From nats connection: cant connect to nats server exiting NOW✖✖✖", "\t", err)
	}

	nc = nConnection
}
