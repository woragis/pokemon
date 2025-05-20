package models

import (
	"time"

	"github.com/google/uuid"
)

type ChatMessage struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	SenderID  uuid.UUID `gorm:"not null"`
	ReceiverID uuid.UUID `gorm:"not null"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time
}
