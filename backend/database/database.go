package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"pokemon/config"
	"pokemon/models"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
    config.LoadEnv()
    dsn := config.GetDatabaseURL()

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    err = db.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatalf("Failed to run migrations: %v", err)
    }

    DB = db
    return db
}
