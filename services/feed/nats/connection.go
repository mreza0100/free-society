package feedNats

import (
	"log"
	"microServiceBoilerplate/configs"
	"os"

	"github.com/nats-io/nats.go"
)

var nc *nats.Conn

const natName = "Feed Service"

func init() {
	nConnection, err := nats.Connect(
		configs.Nats.Url,
		configs.Nats.GetDefaultNatsOpts(natName)...,
	)
	if err != nil {
		log.Fatal("✖✖✖From nats connection: cant connect to nats server exiting NOW✖✖✖")
		log.Fatal(err)
		os.Exit(1)
	}

	nc = nConnection
}
