package models

import "time"

type Password struct {
	Password  string    `gorm:"type:text;NOT NULL"`
	UserId    uint64    `gorm:"primarykey;NOT NULL"`
	CreatedAt time.Time `gorm:"default:NOW();NOT NULL" json:"-"`
}
