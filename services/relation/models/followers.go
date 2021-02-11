package models

import "time"

type Followers struct {
	Follower  uint64 `gorm:"primarykey"`
	Following uint64 `gorm:"primarykey"`

	CreatedAt time.Time `gorm:"NOT NULL;default:NOW()"`
}
