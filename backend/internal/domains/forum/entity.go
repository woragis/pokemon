package forum

import (
	"pokemon/internal/domains/user"
	"time"

	"github.com/google/uuid"
)

/*
Future features
🧠 Additional Tips

    Add DB triggers or application logic to update RepliesCount, LikesCount, and ViewsCount for performance.

    Add soft delete (gorm.DeletedAt) if you want users to remove comments or likes.

    Add ParentID to ForumTopicComment for nested discussions.

Let me know if you want threaded replies, reactions (like ❤️ or 😂), or a notification system.
*/

// ForumCategory represents predefined categories like Competitive, General, etc.
type Category struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name  string    `gorm:"uniqueIndex;not null" json:"name"`
	Color string    `gorm:"not null" json:"color"`
	Description string    `json:"description"`
}

// ForumTopic represents a discussion topic in the forum.
type Topic struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	Title string `gorm:"type:varchar(255);not null" json:"title"`
	Content string `gorm:"type:varchar(255);not null" json:"content"`

	AuthorID uuid.UUID `gorm:"type:uuid;not null;index" json:"-"`
	Author   user.User `gorm:"foreignKey:AuthorID" json:"author"`

	CategoryID uuid.UUID     `gorm:"type:uuid;not null;index" json:"-"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`

	Pinned bool `gorm:"default:false" json:"pinned"`

	RepliesCount int64 `gorm:"default:0" json:"replies"`
	LikesCount   int64 `gorm:"default:0" json:"likes"`
	ViewsCount   int64 `gorm:"default:0" json:"views"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

/****************
 * API RESPONSE *
 ****************/
type Response struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Author       string    `json:"author"`
	AuthorAvatar string    `json:"author_avatar"`
	Date         string    `json:"date"`
	Replies      int64     `json:"replies"`
	Likes        int64     `json:"likes"`
	Views        int64     `json:"views"`
	Category     string    `json:"category"`
	Pinned       bool      `json:"pinned"`
}

/****************
 * INTERACTIONS *
 ****************/
type Like struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	UserID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_topic_like" json:"-"`
	User   user.User      `gorm:"foreignKey:UserID" json:"user"`

	TopicID uuid.UUID  `gorm:"type:uuid;not null;uniqueIndex:idx_user_topic_like" json:"-"`
	Topic   Topic `gorm:"foreignKey:TopicID" json:"-"`

	CreatedAt time.Time `json:"created_at"`
}


type Comment struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	TopicID   uuid.UUID  `gorm:"type:uuid;not null;index" json:"-"`
	Topic     Topic `gorm:"foreignKey:TopicID" json:"-"`

	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"-"`
	User      user.User      `gorm:"foreignKey:UserID" json:"user"`

	Content   string `gorm:"type:text;not null" json:"content"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type View struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	UserID    *uuid.UUID `gorm:"type:uuid;index" json:"-"`
	User      *user.User      `gorm:"foreignKey:UserID" json:"user,omitempty"` // Nullable for guests

	TopicID   uuid.UUID  `gorm:"type:uuid;not null;index" json:"-"`
	Topic     Topic `gorm:"foreignKey:TopicID" json:"-"`

	IPAddress string    `gorm:"type:inet" json:"ip,omitempty"` // Optional
	CreatedAt time.Time `json:"created_at"`
}
