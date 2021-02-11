package db

import (
	"github.com/mreza0100/golog"
)

type DAOS struct {
	Lgr *golog.Core
}

func (this *DAOS) SetFollower(following, follower uint64) error {
	const query = `INSERT INTO followers (following, follower) VALUES (?, ?)`
	params := []interface{}{following, follower}

	tx := DB.Exec(query, params...)

	if tx.Error != nil {
		this.Lgr.RedLog(tx.Error.Error())
	}

	return tx.Error
}
