package entities

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	UUID string `gorm:"not null"`
	Name string
}
