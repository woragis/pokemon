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

    err = db.AutoMigrate(
        &models.BlogPost{},
        &models.ChatMessage{},
        &models.GameGuide{},
        &models.GameGuideTag{},
        &models.Notification{},
        &models.PokePost{},
        &models.PokePostComment{},
        &models.PokePostLike{},
        &models.PokemonSpecies{},
        &models.PokemonGame{},
        &models.Shout{},
        &models.ShoutComment{},
        &models.ShoutLike{},
        &models.Trainer{},
        &models.TrainerPokedexEntry{},
        &models.User{},
        &models.UserFollow{},
        &models.Role{},
        &models.Permission{},
        &models.RolePermission{},
        &models.UserRole{},
        &models.Snap{},
        &models.Comment{},
        &models.Like{},
        &models.ForumCategory{},
        &models.ForumTopic{},
        &models.ForumTopicResponse{},
        &models.ForumTopicLike{},
        &models.ForumTopicComment{},
        &models.ForumTopicView{},
    )
    if err != nil {
        log.Fatalf("Failed to run migrations: %v", err)
    }

    DB = db
    return db
}
