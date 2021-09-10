package models

import "time"

type Notification struct {
	ID        uint64    `gorm:"primarykey;autoIncrement:true"`
	CreatedAt time.Time `gorm:"default:NOW();NOT NULL"`

	UserId uint64 `gorm:"NOT NULL"`

	IsLike  bool   `gorm:"NOT NULL;default:FALSE"`
	LikerId uint64 `gorm:""`
	PostId  string `gorm:""`

	Seen bool `gorm:"NOT NULL;default:FALSE"`
}
