package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogPost struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title     string         `gorm:"not null" json:"title"`
	Content   string         `gorm:"type:text;not null" json:"content"`

	AuthorID  uuid.UUID      `gorm:"type:uuid;not null" json:"author_id"`                         // FK field
	Author    User           `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"author"` // preloadable

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
