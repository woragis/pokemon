package team

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/****************
 * INTERACTIONS *
 ****************/

// --- Team Like ---
type TeamLike struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_team_like" json:"user_id"`
	TeamID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_team_like" json:"team_id"`
}

// --- Team View ---
type TeamView struct {
	ID       uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID   *uuid.UUID `gorm:"type:uuid;index" json:"user_id,omitempty"` // optional
	TeamID   uuid.UUID  `gorm:"type:uuid;not null;index" json:"team_id"`
	ViewedAt time.Time  `gorm:"autoCreateTime" json:"viewed_at"`
}

// --- Team Save ---
type TeamSave struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_team_save" json:"user_id"`
	TeamID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_team_save" json:"team_id"`
	Team   Team      `gorm:"foreignKey:TeamID" json:"team"`
}

// --- Team Comment with nesting support ---
type TeamComment struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	TeamID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"team_id"`
	ParentID  *uuid.UUID     `gorm:"type:uuid;index" json:"parent_id,omitempty"` // nullable, for nested comments
	Content   string         `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
