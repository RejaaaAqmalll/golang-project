package config

import (
	"set-up-Golang/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/latihan?parseTime=true"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(
		&model.User{},
	)

	DB = database

	return database
}
