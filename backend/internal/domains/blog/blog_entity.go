package blog

import (
	"errors"
	"pokemon/internal/domains/user"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title     string         `gorm:"not null" json:"title"`
	Content   string         `gorm:"type:text;not null" json:"content"`

	AuthorID  uuid.UUID      `gorm:"type:uuid;not null" json:"author_id"`                         // FK field
	Author    user.User      `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"author"` // preloadable
	Pinned    bool           `gorm:"default:false" json:"pinned"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

/****************
 * API RESPONSE *
 ****************/
type Response struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Author       string    `json:"author"`
	AuthorAvatar string    `json:"author_avatar"`
	Date         string    `json:"date"`       // formatted date string (e.g., RFC3339 or "Jan 2, 2006")
	Replies      int64     `json:"replies"`    // total comments or replies
	Likes        int64     `json:"likes"`      // total likes
	Views        int64     `json:"views"`      // total views
	Category     string    `json:"category"`   // category name or slug
	Pinned       bool      `json:"pinned"`     // pinned to top?
}

/***************
 * VALIDATIONS *
 ***************/

func (p *Post) Validate() error {
	// Trim spaces to avoid " " being accepted
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)

	if p.Title == "" {
		return errors.New("title is required")
	}
	if len(p.Title) > 200 {
		return errors.New("title cannot be longer than 200 characters")
	}

	if p.Content == "" {
		return errors.New("content is required")
	}

	// Check if AuthorID is zero (all zeros UUID means not set)
	if p.AuthorID == uuid.Nil {
		return errors.New("author_id is required and must be valid")
	}

	return nil
}
