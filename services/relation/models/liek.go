package models

import "time"

type Like struct {
	LikerId uint64 `gorm:"primarykey"`
	OwnerId uint64 `gorm:"primarykey"`
	PostId  string `gorm:"primarykey"`

	CreatedAt time.Time `gorm:"NOT NULL;default:NOW()"`
}
