package securityNats

import (
	"microServiceBoilerplate/services/security/types"

	"github.com/mreza0100/golog"
)

func NewPublishers(lgr *golog.Core) types.Publishers {
	p := publishers{
		lgr: lgr.With("In publishers: "),
	}

	return &p
}

type publishers struct {
	lgr *golog.Core
}
