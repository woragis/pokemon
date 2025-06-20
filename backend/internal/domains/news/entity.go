package news

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type news struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;not null;default:gen_random_uuid()"`
	Title     string         `json:"title" gorm:"type:text;not null"`
	SubTitle  string         `json:"subtitle" gorm:"type:text;not null"`
	Body      string         `json:"body" gorm:"type:text;not null"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type newsView struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;not null;default:gen_random_uuid()"`
	NewsID    uuid.UUID `json:"news_id" gorm:"type:uuid;index;not null"`
	UserID    uuid.UUID `json:"user_id,omitempty" gorm:"type:uuid;index"`
	CreatedAt time.Time `json:"created_at"`
}
