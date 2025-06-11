package blog

import (
	"pokemon/internal/domains/user"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title     string         `gorm:"not null" json:"title"`
	Content   string         `gorm:"type:text;not null" json:"content"`

	AuthorID  uuid.UUID      `gorm:"type:uuid;not null" json:"author_id"`                         // FK field
	Author    user.User      `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"author"` // preloadable
	Pinned    bool           `gorm:"default:false" json:"pinned"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

/****************
 * API RESPONSE *
 ****************/
type Response struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Author       string    `json:"author"`
	AuthorAvatar string    `json:"author_avatar"`
	Date         string    `json:"date"`       // formatted date string (e.g., RFC3339 or "Jan 2, 2006")
	Replies      int64     `json:"replies"`    // total comments or replies
	Likes        int64     `json:"likes"`      // total likes
	Views        int64     `json:"views"`      // total views
	Category     string    `json:"category"`   // category name or slug
	Pinned       bool      `json:"pinned"`     // pinned to top?
}
