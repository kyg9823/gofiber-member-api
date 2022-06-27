package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}
