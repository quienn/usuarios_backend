package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	dsn = "root@tcp(127.0.0.1:3306)/android?charset=utf8mb4&parseTime=True&loc=Local"
)

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Failed to migrate database!")
	}

	DB = db
}
