package models

import (
	"time"

	"github.com/google/uuid"
)

type GameGuide struct {
    ID          uuid.UUID `gorm:"primaryKey"`
    Title       string    `gorm:"type:varchar(255);not null"`
    Slug        string    `gorm:"uniqueIndex;not null"`
    Summary     string    `gorm:"type:text"`
    Content     string    `gorm:"type:text;not null"`
    AuthorID    uuid.UUID `gorm:"not null"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    Tags        []GameGuideTag `gorm:"many2many:game_guide_tags_relation"`
}

type GameGuideTag struct {
    ID    uuid.UUID `gorm:"primaryKey"`
    Name  string    `gorm:"uniqueIndex;not null"`
}
