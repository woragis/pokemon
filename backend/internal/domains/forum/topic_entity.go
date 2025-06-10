package forum

import (
	"pokemon/internal/domains/user"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Topic1 struct {
	ID           uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:generate_uuid_v4()"`
	UserID       uuid.UUID      `json:"user_id" gorm:"type:uuid;not null;index"`
	Title        string         `json:"title" gorm:"not null"`
	Content      string         `json:"content" gorm:"not null"`

	ViewCount    int64          `json:"view_count" gorm:"default:0"`
	CommentCount int64          `json:"comment_count" gorm:"default:0"`
	LikeCount    int64          `json:"like_count" gorm:"default:0"`
	DislikeCount int64          `json:"dislike_count" gorm:"default:0"`

	User         user.User      `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
    DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type TopicView struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:generate_uuid_v4()"`
	TopicID uuid.UUID `json:"topic_id" gorm:"type:uuid;index;not null"`
	UserID  uuid.UUID `json:"user_id" gorm:"type:uuid;index"`

	Topic Topic         `json:"topic" gorm:"foreignKey:TopicID"`
	User  user.User     `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
}

type TopicLike struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:generate_uuid_v4()"`
	TopicID uuid.UUID `json:"topic_id" gorm:"type:uuid;not null;index:idx_topic_like,unique"`
	UserID  uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index:idx_topic_like,unique"`

	Like    bool      `json:"like" gorm:"not null"`

	User      user.User `json:"-" gorm:"foreignKey:UserID"`
	Topic     Topic     `json:"-" gorm:"foreignKey:TopicID"`
	CreatedAt time.Time `json:"created_at"`
}

type Comment1 struct {
	ID           uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:generate_uuid_v4()"`
	TopicID      uuid.UUID      `json:"topic_id" gorm:"type:uuid;not null;index:idx_topic_like,unique"`
	UserID       uuid.UUID      `json:"user_id" gorm:"type:uuid;not null;index:idx_topic_like,unique"`
	ParentID     uuid.UUID      `json:"parent_id,omitempty" gorm:"type:uuid;index"`
	Content      string         `json:"content" gorm:"type:text;not null"`
	LikeCount    int64          `json:"like_count" gorm:"default:0"`
	DislikeCount int64          `json:"dislike_count" gorm:"default:0"`
	User         user.User      `json:"user" gorm:"foreignKey:UserID"`
	Topic        Topic          `json:"topic" gorm:"foreignKey:TopicID"`
	Parent       *Comment1      `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Replies      []Comment1     `json:"replies,omitempty" gorm:"foreignKey:ParentID"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
    DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type CommentLike struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:generate_uuid_v4()"`
	CommentID uuid.UUID `json:"comment_id" gorm:"type:uuid;not null;index:idx_topic_comment_like,unique"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index:idx_topic_comment_like,unique"`

	Like      bool      `json:"like" gorm:"not null"`

	User      user.User `json:"-" gorm:"foreignKey:UserID"`
	Comment   Comment1  `json:"-" gorm:"foreignKey:CommentID"`
	CreatedAt time.Time `json:"created_at"`
}