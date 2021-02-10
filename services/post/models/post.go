package models

import "time"

type Post struct {
	Title     string    `gorm:"type:text;NOT NULL"`
	Body      string    `gorm:"type:text;NOT NULL"`
	OwnerId   uint64    `gorm:"NOT NULL;type:bigserial"`
	ID        uint64    `gorm:"primarykey;autoIncrement:true"`
	CreatedAt time.Time `gorm:"default:NOW();NOT NULL"`
}
