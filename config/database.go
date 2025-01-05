package config

import (
	"gin-samples/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to SQLite in-memory database: %v", err)
	}

	err = db.AutoMigrate(&entity.Greeting{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
