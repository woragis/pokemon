package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username string    `gorm:"unique;not null"`
	Name     string
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Role     string    `gorm:"default:user"`

	// Trainer  Trainer   `gorm:"constraint:OnDelete:CASCADE"` // establishes relation
}

type Trainer struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID   uuid.UUID `gorm:"type:uuid;not null;unique"`
	Nickname string    `gorm:"not null"`
	Avatar   string

	User     User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
