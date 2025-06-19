package forum

import (
	"errors"
	"pokemon/internal/domains/user"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*********
 * MAIN *
 *********/

type Topic struct {
	ID           uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID       uuid.UUID      `json:"user_id" gorm:"type:uuid;not null;index"`
	Title        string         `json:"title" gorm:"not null"`
	Content      string         `json:"content" gorm:"not null"`
	Pinned       bool           `json:"pinned" gorm:"default:false"`

	ViewCount    int64          `json:"view_count" gorm:"default:0"`
	CommentCount int64          `json:"comment_count" gorm:"default:0"`
	LikeCount    int64          `json:"like_count" gorm:"default:0"`
	DislikeCount int64          `json:"dislike_count" gorm:"default:0"`

	User         user.User      `json:"user" gorm:"foreignKey:UserID"`

	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
    DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

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

/***************
 * VALIDATIONS *
 ***************/

func (t *Topic) Validate() error {
	// Normalize input
	t.Title = strings.TrimSpace(t.Title)
	t.Content = strings.TrimSpace(t.Content)

	// Title validations
	if t.Title == "" {
		return errors.New("title is required")
	}
	if len(t.Title) > 200 {
		return errors.New("title cannot be longer than 200 characters")
	}

	// Content validations
	if t.Content == "" {
		return errors.New("content is required")
	}
	if len(t.Content) > 5000 {
		return errors.New("content cannot be longer than 5000 characters")
	}

	// UserID must be valid
	if t.UserID == uuid.Nil {
		return errors.New("user_id is required and must be valid")
	}

	return nil
}

/****************
 * INTERACTIONS *
 ****************/

type TopicView struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TopicID   uuid.UUID `json:"topic_id" gorm:"type:uuid;index;not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;index"`

	IPAddress string    `gorm:"type:inet" json:"ip,omitempty"` // Optional

	Topic     Topic     `json:"topic" gorm:"foreignKey:TopicID"`
	User      user.User `json:"user" gorm:"foreignKey:UserID"`

	CreatedAt time.Time `json:"created_at"`
}

type TopicLike struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TopicID   uuid.UUID `json:"topic_id" gorm:"type:uuid;not null;index:idx_topic_like,unique"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index:idx_topic_like,unique"`

	Like      bool      `json:"like" gorm:"not null"`

	User      user.User `json:"-" gorm:"foreignKey:UserID"`
	Topic     Topic     `json:"-" gorm:"foreignKey:TopicID"`

	CreatedAt time.Time `json:"created_at"`
}

type TopicComment struct {
	ID           uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TopicID      uuid.UUID      `json:"topic_id" gorm:"type:uuid;not null;index:idx_topic_like,unique"`
	UserID       uuid.UUID      `json:"user_id" gorm:"type:uuid;not null;index:idx_topic_like,unique"`
	ParentID     uuid.UUID      `json:"parent_id,omitempty" gorm:"type:uuid;index"`

	Content      string         `json:"content" gorm:"type:text;not null"`
	LikeCount    int64          `json:"like_count" gorm:"default:0"`
	DislikeCount int64          `json:"dislike_count" gorm:"default:0"`

	User         user.User      `json:"user" gorm:"foreignKey:UserID"`
	Topic        Topic          `json:"topic" gorm:"foreignKey:TopicID"`
	Parent       *TopicComment  `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Replies      []TopicComment `json:"replies,omitempty" gorm:"foreignKey:ParentID"`

	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
    DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type TopicCommentLike struct {
	ID        uuid.UUID    `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CommentID uuid.UUID    `json:"comment_id" gorm:"type:uuid;not null;index:idx_topic_comment_like,unique"`
	UserID    uuid.UUID    `json:"user_id" gorm:"type:uuid;not null;index:idx_topic_comment_like,unique"`

	Like      bool         `json:"like" gorm:"not null"`

	User      user.User    `json:"-" gorm:"foreignKey:UserID"`
	Comment   TopicComment `json:"-" gorm:"foreignKey:CommentID"`

	CreatedAt time.Time    `json:"created_at"`
}

type TopicCategory struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name       string    `gorm:"uniqueIndex;not null" json:"name"`
	Color      string    `gorm:"not null" json:"color"`
	Description string   `json:"description"`
}

/***************
 * VALIDATIONS *
 ***************/

func (l *TopicLike) Validate() error {
	if l.TopicID == uuid.Nil {
		return errors.New("topic_id is required")
	}
	if l.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	return nil
}

func (c *TopicComment) Validate() error {
	if c.TopicID == uuid.Nil {
		return errors.New("topic_id is required")
	}
	if c.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	if strings.TrimSpace(c.Content) == "" {
		return errors.New("content is required")
	}
	return nil
}

func (l *TopicCommentLike) Validate() error {
	if l.CommentID == uuid.Nil {
		return errors.New("comment_id is required")
	}
	if l.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	return nil
}

func (c *TopicCategory) Validate() error {
	if strings.TrimSpace(c.Name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(c.Color) == "" {
		return errors.New("color is required")
	}
	return nil
}
