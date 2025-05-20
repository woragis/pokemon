package models

import (
	"time"

	"github.com/google/uuid"
)

type Shout struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID       uuid.UUID      `gorm:"type:uuid;not null;index"`
	User         User           `gorm:"foreignKey:UserID"`

	Content      string         `gorm:"type:varchar(280);not null"`

	CreatedAt    time.Time
	UpdatedAt    time.Time      // Good practice for edit history

	// Reshout / Quote-shout
	ReshoutOfID  *uuid.UUID     `gorm:"type:uuid;index"`
	ReshoutOf    *Shout         `gorm:"foreignKey:ReshoutOfID"`

	QuoteContent *string        `gorm:"type:varchar(280)"`

	// Relationships
	Likes        []ShoutLike    `gorm:"constraint:OnDelete:CASCADE"`
	Comments     []ShoutComment `gorm:"constraint:OnDelete:CASCADE"`

	// AIReplyContent *string     // Future feature
}

type ShoutLike struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	ShoutID   uuid.UUID `gorm:"type:uuid;not null;index"`

	// Prevent duplicate likes per user per shout
	// Add a unique index
	// gorm:"uniqueIndex:idx_user_shout_like"
}

type ShoutComment struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	ShoutID   uuid.UUID `gorm:"type:uuid;not null;index"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time
}
