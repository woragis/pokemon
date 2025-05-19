package models

import (
	"time"
)

type Shout struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Content   string    `gorm:"type:varchar(280);not null"` // 280 chars like Twitter
	CreatedAt time.Time
	User      User `gorm:"foreignKey:UserID"`
	Likes     []ShoutLike
	Comments  []ShoutComment
}

type ShoutLike struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint `gorm:"not null"`
	ShoutID uint `gorm:"not null"`
}

type ShoutComment struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	ShoutID   uint      `gorm:"not null"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time
}
