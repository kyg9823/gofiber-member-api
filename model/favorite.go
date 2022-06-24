package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	Id   int32  `json:"id" gorm:"primaryKey"`
	Item string `json:"item" gorm:"primaryKey"`
}
