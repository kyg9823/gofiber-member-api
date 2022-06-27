package database

import (
	"errors"
	"os"

	"github.com/kyg9823/gofiber-member-api/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func ConnectDB() {
	_, err := os.Stat("model.db")
	IsDBExist := !errors.Is(err, os.ErrNotExist)

	db, err := gorm.Open(sqlite.Open("model.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	DBConn = db

	if !IsDBExist {
		db.AutoMigrate(&model.Member{})
		db.AutoMigrate(&model.Favorite{})
		db.AutoMigrate(&model.User{})

		db.Create(&model.Member{Id: 1, Name: "Kim Youngkook", Email: "kyg9823@gmail.com", Favorites: []model.Favorite{
			{
				Id:   1,
				Item: "Food",
			},
			{
				Id:   1,
				Item: "Go",
			},
		},
		})
		// $2a$14$r0yhSEbqBt7lVwtcWMw2X.YMwMxSvgwllvGnn2oAuRJrGzvNmVlYO
		db.Create(&model.User{
			Username: "admin",
			Name:     "Administrator",
			Password: "$2a$14$gllaLu.cdKQRgguEzi1v1.w19MkTFeJo1VbIJ2ABIrs7Qf6uapU1G",
		})
	}
}
