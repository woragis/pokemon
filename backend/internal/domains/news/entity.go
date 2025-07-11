package news

import (
	"errors"
	"pokemon/internal/domains/user"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type News struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;not null;default:gen_random_uuid()"`
	UserID    uuid.UUID      `json:"user_id" gorm:"type:uuid;not null"`
	User      user.User      `json:"user" gorm:"foreignKey:UserID"`
	Title     string         `json:"title" gorm:"type:text;not null"`
	SubTitle  string         `json:"subtitle" gorm:"type:text;not null"`
	Body      string         `json:"body" gorm:"type:text;not null"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type NewsView struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;not null;default:gen_random_uuid()"`
	NewsID    uuid.UUID `json:"news_id" gorm:"type:uuid;index;not null"`
	UserID    uuid.UUID `json:"user_id,omitempty" gorm:"type:uuid;index"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *News) Validate() error {
	if n.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	if strings.TrimSpace(n.Title) == "" {
		return errors.New("title is required")
	}
	if strings.TrimSpace(n.SubTitle) == "" {
		return errors.New("subtitle is required")
	}
	if strings.TrimSpace(n.Body) == "" {
		return errors.New("body is required")
	}
	return nil
}
