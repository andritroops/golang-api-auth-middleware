package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=123456 dbname=golang port=5431 sslmode=disable TimeZone=Asia/jakarta"

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err)
	}

	DB = db

}
