package models

import (
	"time"

	"github.com/google/uuid"
)

type Shout struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID       uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	User         User           `gorm:"foreignKey:UserID" json:"user"`
	Content      string         `gorm:"type:varchar(280);not null" json:"content"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`

	ReshoutOfID  *uuid.UUID     `gorm:"type:uuid;index" json:"reshout_of_id,omitempty"`
	ReshoutOf    *Shout         `gorm:"foreignKey:ReshoutOfID" json:"reshout_of,omitempty"`
	QuoteContent *string        `gorm:"type:varchar(280)" json:"quote_content,omitempty"`
	Likes        []ShoutLike    `gorm:"constraint:OnDelete:CASCADE" json:"likes"`
	Comments     []ShoutComment `gorm:"constraint:OnDelete:CASCADE" json:"comments"`

	IsFlagged    bool           `gorm:"default:false" json:"is_flagged"`
}

type ShoutLike struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID  uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_shout_like" json:"user_id"`
	ShoutID uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_shout_like" json:"shout_id"`
}

type ShoutComment struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	ShoutID   uuid.UUID `gorm:"type:uuid;not null;index" json:"shout_id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time `json:"created_at"`

	IsFlagged bool `gorm:"default:false" json:"is_flagged"`
}
