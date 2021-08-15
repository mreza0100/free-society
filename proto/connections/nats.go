package connections

import (
	"freeSociety/configs"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
)

func GetConnection(lgr *golog.Core, natName string) *nats.Conn {
	var (
		nc  *nats.Conn
		err error
	)

	{
		nc, err = nats.Connect(
			configs.Nats.Url,
			configs.Nats.GetDefaultNatsOpts(natName)...,
		)
	}
	{
		if err != nil {
			lgr.Fatal("✖✖✖From nats connection: cant connect to nats server exiting NOW✖✖✖", "\t", err)
		}
	}
	return nc
}
