package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Snap represents a media post.
type Snap struct {
	ID            uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID        uuid.UUID      `json:"user_id" gorm:"type:uuid;not null"`
	MediaURL      string         `json:"media_url" gorm:"not null"`
	MediaType     string         `json:"media_type" gorm:"type:text;not null"`
	Caption       string         `json:"caption"`
	Tags          pq.StringArray `json:"tags" gorm:"type:text[]"`
	Pokemon       pq.StringArray `json:"pokemon" gorm:"type:text[]"`
	LikesCount    int            `json:"likes_count" gorm:"default:0"`
	CommentsCount int            `json:"comments_count" gorm:"default:0"`
	ReportCount   int            `json:"report_count" gorm:"default:0"`
	Status        string         `json:"status" gorm:"type:text;default:'active'"` // active, flagged, removed
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

// SnapComment represents a comment on a Snap.
type SnapComment struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SnapID      uuid.UUID `json:"snap_id" gorm:"type:uuid;not null;index"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Content     string    `json:"content" gorm:"not null"`
	ReportCount int       `json:"report_count" gorm:"default:0"`
	Status      string    `json:"status" gorm:"type:text;default:'active'"` // active, flagged, removed
	CreatedAt   time.Time `json:"created_at"`
}

// SnapLike represents a like on a Snap.
type SnapLike struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SnapID    uuid.UUID `json:"snap_id" gorm:"type:uuid;not null;index"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at"`
}

// SnapReport represents a user report on a Snap.
type SnapReport struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SnapID     uuid.UUID `json:"snap_id" gorm:"type:uuid;not null;index;uniqueIndex:idx_snap_report_unique"`
	ReporterID uuid.UUID `json:"reporter_id" gorm:"type:uuid;not null;index;uniqueIndex:idx_snap_report_unique"`
	Reason     string    `json:"reason" gorm:"type:text;not null"` // e.g. spam, abuse
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
}

// SnapCommentReport represents a user report on a SnapComment.
type SnapCommentReport struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SnapCommentID uuid.UUID `json:"snap_comment_id" gorm:"type:uuid;not null;index;uniqueIndex:idx_snap_comment_report_unique"`
	ReporterID    uuid.UUID `json:"reporter_id" gorm:"type:uuid;not null;index;uniqueIndex:idx_snap_comment_report_unique"`
	Reason        string    `json:"reason" gorm:"type:text;not null"`
	Comment       string    `json:"comment"`
	CreatedAt     time.Time `json:"created_at"`
}
