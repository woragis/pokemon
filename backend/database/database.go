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

        // Snaps
        &models.Snap{}, // Post
        &models.SnapComment{}, // Comments
        &models.SnapLike{}, // Likes
        &models.SnapReport{}, // Post Report
        &models.SnapCommentReport{}, // Comment Report

        // Forum
        // Still Missing: Topic Comment Like, Topic Comment Comment, Topic Report Support, Topic Comment Report Support
        &models.ForumCategory{}, // Category
        &models.ForumTopic{}, // Topic
        &models.ForumTopicLike{}, // Topic Like
        &models.ForumTopicComment{}, // Topic Comment
        &models.ForumTopicView{}, // Topic View
        &models.ForumTopicResponse{}, // API Response
    )
    if err != nil {
        log.Fatalf("Failed to run migrations: %v", err)
    }

    DB = db
    return db
}
