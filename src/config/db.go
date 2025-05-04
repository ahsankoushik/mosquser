package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data.db"))
	if err != nil {
		log.Fatal("Unable to connect to database")
	}
	return db
}

var DB *gorm.DB = ConnectDB()
