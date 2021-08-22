package models

import "time"

type Post struct {
	ID      uint64 `gorm:"primarykey;autoIncrement:true"`
	OwnerId uint64 `gorm:"NOT NULL;"`

	Title string `gorm:"type:text;NOT NULL"`
	Body  string `gorm:"type:text;NOT NULL"`

	PicturesName string `gorm:""`

	CreatedAt time.Time `gorm:"default:NOW();NOT NULL"`
}
