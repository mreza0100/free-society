package repository

import (
	"github.com/mreza0100/golog"
	gorm "gorm.io/gorm"
)

type likes_read struct {
	lgr *golog.Core
	db  *gorm.DB
}
