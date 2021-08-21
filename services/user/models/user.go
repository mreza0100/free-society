package models

import "time"

type User struct {
	ID         uint64    `gorm:"primarykey;autoIncrement:true"`
	CreatedAt  time.Time `gorm:"NOT NULL;default:NOW()"`
	Name       string    `gorm:"type:text;NOT NULL"`
	Gender     string    `gorm:"type:text;NOT NULL"`
	Email      string    `gorm:"index:unique;NOT NULL"`
	AvatarPath string    `gorm:"NOT NULL"`
}
