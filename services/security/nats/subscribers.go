package securityNats

import (
	"microServiceBoilerplate/services/security/types"

	"github.com/mreza0100/golog"
)

func InitialNatsSubs(srv types.Sevice, lgr *golog.Core) {
	_ = subscribers{
		srv: srv,
		lgr: lgr.With("In subscribers => "),
	}
	lgr.GreenLog("✅ subscribers has been attached to nats")
}

type subscribers struct {
	srv types.Sevice
	lgr *golog.Core
}
