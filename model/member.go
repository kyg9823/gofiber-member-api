package model

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
