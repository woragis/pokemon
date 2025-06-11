package forum

import (
	"errors"
	"pokemon/internal/domains/user"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Topic struct {
	ID           uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:generate_uuid_v4()"`
	UserID       uuid.UUID      `json:"user_id" gorm:"type:uuid;not null;index"`
	Title        string         `json:"title" gorm:"not null"`
	Content      string         `json:"content" gorm:"not null"`
	Pinned       bool           `json:"pinned" gorm:"default:false"`

	ViewCount    int64          `json:"view_count" gorm:"default:0"`
	CommentCount int64          `json:"comment_count" gorm:"default:0"`
	LikeCount    int64          `json:"like_count" gorm:"default:0"`
	DislikeCount int64          `json:"dislike_count" gorm:"default:0"`

	User         user.User      `json:"user" gorm:"foreignKey:UserID"`

	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
    DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type Response struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Author       string    `json:"author"`
	AuthorAvatar string    `json:"author_avatar"`
	Date         string    `json:"date"`
	Replies      int64     `json:"replies"`
	Likes        int64     `json:"likes"`
	Views        int64     `json:"views"`
	Category     string    `json:"category"`
	Pinned       bool      `json:"pinned"`
}

func (t *Topic) Validate() error {
	// Normalize input
	t.Title = strings.TrimSpace(t.Title)
	t.Content = strings.TrimSpace(t.Content)

	// Title validations
	if t.Title == "" {
		return errors.New("title is required")
	}
	if len(t.Title) > 200 {
		return errors.New("title cannot be longer than 200 characters")
	}

	// Content validations
	if t.Content == "" {
		return errors.New("content is required")
	}
	if len(t.Content) > 5000 {
		return errors.New("content cannot be longer than 5000 characters")
	}

	// UserID must be valid
	if t.UserID == uuid.Nil {
		return errors.New("user_id is required and must be valid")
	}

	return nil
}
