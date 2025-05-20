package models

import (
	"time"

	"github.com/google/uuid"
)

type PokePost struct {
    ID        uuid.UUID  `gorm:"primaryKey"`
    UserID    uuid.UUID  `gorm:"not null"`
    Caption   string     `gorm:"type:text"`
    ImageURL  string     `gorm:"type:text;not null"`
    CreatedAt time.Time

    User      User       `gorm:"foreignKey:UserID"`
    Likes     []PokePostLike
    Comments  []PokePostComment
}

type PokePostLike struct {
    ID        uuid.UUID `gorm:"primaryKey"`
    UserID    uuid.UUID `gorm:"not null"`
    PokePostID uuid.UUID `gorm:"not null"`
}

type PokePostComment struct {
    ID         uuid.UUID `gorm:"primaryKey"`
    UserID     uuid.UUID `gorm:"not null"`
    PokePostID uuid.UUID `gorm:"not null"`
    Content    string    `gorm:"type:text;not null"`
    CreatedAt  time.Time
}
