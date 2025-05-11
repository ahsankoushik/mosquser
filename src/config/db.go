package config

import (
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func ConnectDB() *gorm.DB {
	once.Do(func() {
		var err error
		DB, err = gorm.Open(sqlite.Open("data.db"))
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}

		// Configure connection pool
		// sqlDB, _ := db.DB()
		// sqlDB.SetMaxOpenConns(10)           // Maximum number of open connections
		// sqlDB.SetMaxIdleConns(5)            // Maximum number of idle connections
		// sqlDB.SetConnMaxLifetime(time.Hour) // Maximum lifetime of a connection
	})
	return DB
}
