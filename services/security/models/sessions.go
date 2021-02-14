package models

import "time"

type Session struct {
	ID        uint64    `gorm:"primarykey;autoIncrement:true"`
	CreatedAt time.Time `gorm:"NOT NULL;default:NOW()"`

	Token string `gorm:"NOT NULL;type:text"`

	Device string `gorm:"default:'';type:text"`
	UserId uint64 `gorm:"primarykey;NOT NULL"`
}
