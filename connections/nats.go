package connections

import (
	"freeSociety/configs"
	"log"

	globalConfigs "freeSociety/configs"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
)

func GetDefaultNatsOpts(name string, nConf *globalConfigs.NatsConfigsT) []nats.Option {
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

func GetNatsConnection(lgr *golog.Core, natName string) *nats.Conn {
	nc, err := nats.Connect(
		configs.Nats.Url,
		GetDefaultNatsOpts(natName, configs.Nats)...,
	)
	if err != nil {
		lgr.Fatal("✖✖✖From nats connection: cant connect to nats server. exiting NOW✖✖✖", "\t", err)
	}

	return nc
}
