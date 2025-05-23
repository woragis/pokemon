package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Snap struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID        uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	MediaURL      string    `json:"media_url" gorm:"not null"`
	MediaType     string    `json:"media_type" gorm:"type:text;not null"`
	Caption       string    `json:"caption"`
	Tags          pq.StringArray  `json:"tags" gorm:"type:text[]"`
	Pokemon       pq.StringArray  `json:"pokemon" gorm:"type:text[]"`
	LikesCount    int       `json:"likes_count" gorm:"default:0"`
	CommentsCount int       `json:"comments_count" gorm:"default:0"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Comment struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SnapID    uuid.UUID `json:"snap_id" gorm:"type:uuid;not null;index"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Content   string    `json:"content" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
}

type Like struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SnapID    uuid.UUID `json:"snap_id" gorm:"type:uuid;not null;index"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at"`
}
