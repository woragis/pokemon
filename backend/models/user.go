package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Username string    `gorm:"unique;not null" json:"username"`
	Name     string    `json:"name"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Password string    `gorm:"not null" json:"-"`
	Role     string    `gorm:"default:user" json:"role"`

	// Trainer  Trainer   `gorm:"constraint:OnDelete:CASCADE"` // establishes relation
}

type Trainer struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID   uuid.UUID `gorm:"type:uuid;not null;unique"`
	Nickname string    `gorm:"not null"`
	Avatar   string

	User     User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
