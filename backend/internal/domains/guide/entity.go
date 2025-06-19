package gamguide

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type GameGuide struct {
	ID            uuid.UUID      `gorm:"primaryKey" json:"id"`
	Title         string         `gorm:"type:varchar(255);not null" json:"title"`
	Slug          string         `gorm:"uniqueIndex;not null" json:"slug"`
	Summary       string         `gorm:"type:text" json:"summary"`
	Content       string         `gorm:"type:text;not null" json:"content"`
	CoverImageURL string         `gorm:"type:text" json:"cover_image_url"`
	AuthorID      uuid.UUID      `gorm:"not null" json:"author_id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	Tags          []GameGuideTag `gorm:"many2many:game_guide_tags_relation" json:"tags"`
}

type GameGuideTag struct {
	ID   uuid.UUID `gorm:"primaryKey" json:"id"`
	Name string    `gorm:"uniqueIndex;not null" json:"name"`
}

type GameGuideView struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	GuideID   uuid.UUID `gorm:"type:uuid;not null;index" json:"guide_id"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type GameGuideLike struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	GuideID   uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_guide_like" json:"guide_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_guide_like" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

/***************
 * VALIDATIONS *
 ***************/

func (g *GameGuide) Validate() error {
	if g.Title == "" {
		return errors.New("title is required")
	}
	if g.Slug == "" {
		return errors.New("slug is required")
	}
	if g.Content == "" {
		return errors.New("content is required")
	}
	if g.AuthorID == uuid.Nil {
		return errors.New("author_id is required")
	}
	return nil
}
