package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
    ID        uuid.UUID `gorm:"primaryKey"`
    UserID    uuid.UUID `gorm:"not null"` // who receives the notification
    Type      string    `gorm:"not null"` // "like", "comment", etc
    Message   string    `gorm:"not null"`
    CreatedAt time.Time
}
