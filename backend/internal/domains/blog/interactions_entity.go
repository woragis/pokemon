package blog

import (
	"pokemon/internal/domains/user"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/****************
 * INTERACTIONS *
 ****************/

// --- Post Like ---
type PostLike struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_post_like" json:"user_id"`
	PostID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_post_like" json:"post_id"`
}

// --- Post View ---
type PostView struct {
	ID       uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID   *uuid.UUID `gorm:"type:uuid;index" json:"user_id,omitempty"` // optional
	PostID   uuid.UUID  `gorm:"type:uuid;not null;index" json:"post_id"`
	ViewedAt time.Time  `gorm:"autoCreateTime" json:"viewed_at"`
}

// --- Post Save ---
type PostSave struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_post_save" json:"user_id"`
	PostID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_post_save" json:"post_id"`
	Post   Post      `gorm:"foreignKey:PostID" json:"post"`
}

// --- Post Comment with nesting support ---
type PostComment struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	User      user.User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	PostID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"post_id"`
	Post      Post           `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"post,omitempty"`
	ParentID  *uuid.UUID     `gorm:"type:uuid;index" json:"parent_id,omitempty"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
