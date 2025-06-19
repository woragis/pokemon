package shout

import (
	"errors"
	"pokemon/internal/domains/user"
	"pokemon/pkg/validation"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/********
 * MAIN *
 ********/

type Shout struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID       uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	User         user.User      `gorm:"foreignKey:UserID" json:"user"`
	Content      string         `gorm:"type:varchar(280);not null" json:"content"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	ReshoutOfID  *uuid.UUID     `gorm:"type:uuid;index" json:"reshout_of_id,omitempty"`
	ReshoutOf    *Shout         `gorm:"foreignKey:ReshoutOfID" json:"reshout_of,omitempty"`
	QuoteCommentID *uuid.UUID     `gorm:"type:uuid" json:"quote_comment_id,omitempty"`
	QuoteComment   *ShoutComment `gorm:"foreignKey:QuoteCommentID" json:"quote_comment,omitempty"`

	QuoteContent *string        `gorm:"type:varchar(280)" json:"quote_content,omitempty"`

	Likes        []ShoutLike    `gorm:"constraint:OnDelete:CASCADE" json:"likes"`
	Comments     []ShoutComment `gorm:"constraint:OnDelete:CASCADE" json:"comments"`
	LikeCount    int            `gorm:"default:0" json:"like_count"`
	CommentCount int            `gorm:"default:0" json:"comment_count"`
	IsFlagged    bool           `gorm:"default:false" json:"is_flagged"`
}

/****************
 * INTERACTIONS *
 ****************/

type ShoutComment struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	ShoutID   uuid.UUID `gorm:"type:uuid;not null;index" json:"shout_id"`
	User      user.User `gorm:"foreignKey:UserID" json:"user"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time `json:"created_at"`
	IsFlagged bool      `gorm:"default:false" json:"is_flagged"`

	ParentCommentID *uuid.UUID      `gorm:"type:uuid;index" json:"parent_comment_id,omitempty"`
	ParentComment   *ShoutComment   `gorm:"foreignKey:ParentCommentID" json:"-"`
	Replies         []ShoutComment  `gorm:"foreignKey:ParentCommentID" json:"replies,omitempty"`

	Likes     []ShoutCommentLike `gorm:"constraint:OnDelete:CASCADE" json:"likes"`
	LikeCount int                `gorm:"default:0" json:"like_count"`
}

type ShoutLike struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_shout_like" json:"user_id"`
	ShoutID   uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_shout_like" json:"shout_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ShoutCommentLike struct {
	ID        uuid.UUID    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID    `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_comment_like" json:"user_id"`
	CommentID uuid.UUID    `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_comment_like" json:"comment_id"`
	CreatedAt time.Time    `json:"created_at"`
}

type ShoutSave struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_shout_save" json:"user_id"`
	ShoutID   uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_shout_save" json:"shout_id"`
	CreatedAt time.Time `json:"created_at"`

	User  user.User `gorm:"foreignKey:UserID" json:"user"`
	Shout Shout     `gorm:"foreignKey:ShoutID" json:"shout"`
}

type ShoutView struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id,omitempty"`
	ShoutID   uuid.UUID `gorm:"type:uuid;not null;index" json:"shout_id"`
	CreatedAt time.Time `json:"created_at"`

	User  user.User `gorm:"foreignKey:UserID" json:"user"`
	Shout Shout     `gorm:"foreignKey:ShoutID" json:"shout"`
}

/***************
 * VALIDATIONS *
 ***************/

func (s *Shout) Validate() error {
	if err := validation.UUIDRequired(s.UserID, "user_id"); err != nil {
		return err
	}
	if err := validation.StringRequired(s.Content, "content"); err != nil {
		return err
	}
	if len(s.Content) > 280 {
		return errors.New("content must be 280 characters or less")
	}
	if s.QuoteContent != nil && len(*s.QuoteContent) > 280 {
		return errors.New("quote_content must be 280 characters or less")
	}
	return nil
}

func (c *ShoutComment) Validate() error {
	if err := validation.UUIDRequired(c.UserID, "user_id"); err != nil {
		return err
	}
	if err := validation.UUIDRequired(c.ShoutID, "shout_id"); err != nil {
		return err
	}
	if err := validation.StringRequired(c.Content, "content"); err != nil {
		return err
	}
	return nil
}

func (l *ShoutLike) Validate() error {
	if err := validation.UUIDRequired(l.UserID, "user_id"); err != nil {
		return err
	}
	if err := validation.UUIDRequired(l.ShoutID, "shout_id"); err != nil {
		return err
	}
	return nil
}

func (l *ShoutCommentLike) Validate() error {
	if err := validation.UUIDRequired(l.UserID, "user_id"); err != nil {
		return err
	}
	if err := validation.UUIDRequired(l.CommentID, "comment_id"); err != nil {
		return err
	}
	return nil
}

func (s *ShoutSave) Validate() error {
	if err := validation.UUIDRequired(s.UserID, "user_id"); err != nil {
		return err
	}
	if err := validation.UUIDRequired(s.ShoutID, "shout_id"); err != nil {
		return err
	}
	return nil
}

func (v *ShoutView) Validate() error {
	if err := validation.UUIDRequired(v.ShoutID, "shout_id"); err != nil {
		return err
	}
	// UserID is optional
	return nil
}
