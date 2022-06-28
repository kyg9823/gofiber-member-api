package model

type Favorite struct {
	// gorm.Model
	Id   string `json:"id" gorm:"primaryKey"`
	Item string `json:"item" gorm:"primaryKey"`
}
