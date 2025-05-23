package models

import (
	"time"

	"github.com/google/uuid"
)

// ForumCategory represents predefined categories like Competitive, General, etc.
type ForumCategory struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name  string    `gorm:"uniqueIndex;not null" json:"name"`
	Color string    `gorm:"not null" json:"color"`
	Description string    `json:"description"`
}

// ForumTopic represents a discussion topic in the forum.
type ForumTopic struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	Title string `gorm:"type:varchar(255);not null" json:"title"`

	AuthorID uuid.UUID `gorm:"type:uuid;not null;index" json:"-"`
	Author   User      `gorm:"foreignKey:AuthorID" json:"author"`

	CategoryID uuid.UUID     `gorm:"type:uuid;not null;index" json:"-"`
	Category   ForumCategory `gorm:"foreignKey:CategoryID" json:"category"`

	Pinned bool `gorm:"default:false" json:"pinned"`

	RepliesCount int64 `gorm:"default:0" json:"replies"`
	LikesCount   int64 `gorm:"default:0" json:"likes"`
	ViewsCount   int64 `gorm:"default:0" json:"views"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ForumTopicResponse struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Author       string    `json:"author"`
	AuthorAvatar string    `json:"authorAvatar"`
	Date         string    `json:"date"`
	Replies      int64     `json:"replies"`
	Likes        int64     `json:"likes"`
	Views        int64     `json:"views"`
	Category     string    `json:"category"`
	Pinned       bool      `json:"pinned"`
}
