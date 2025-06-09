package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(dsn string) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    
    return db, nil
}

func MigrateModels(db *gorm.DB) {
    err := db.AutoMigrate(
    )
    if err != nil {
        log.Fatalf("Failed to run migrations: %v", err)
    }
}
