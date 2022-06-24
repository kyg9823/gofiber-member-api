package database

import (
	"github.com/kyg9823/gofiber-member-api/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("model.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&model.Member{})

	db.Create(&model.Member{Id: 1, Name: "Kim Youngkook", Email: "kyg9823@gmail.com"})

	DBConn = db
}
