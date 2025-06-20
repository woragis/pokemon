package walkthrough

import (
	"pokemon/internal/domains/user"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Walkthrough struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title     string         `gorm:"not null" json:"title"`                          // e.g. "Pokémon Scarlet Full Walkthrough"
	Game      string         `gorm:"not null" json:"game"`                          // e.g. "Scarlet", "Violet", etc.
	AuthorID  uuid.UUID      `gorm:"type:uuid;not null" json:"author_id"`           // FK to users
	Author    user.User      `gorm:"foreignKey:AuthorID" json:"author"`
	Tags      pq.StringArray `gorm:"type:text[]" json:"tags"`                       // Optional: ["Story", "Guide", "Exploration"]

	Steps     []WalkthroughStep `gorm:"foreignKey:WalkthroughID" json:"steps"`     // Walkthrough content sections

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type WalkthroughStep struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	WalkthroughID  uuid.UUID `gorm:"type:uuid;not null" json:"walkthrough_id"`         // FK to Walkthrough
	Title          string    `gorm:"not null" json:"title"`                            // e.g. "Route 1 and First Battle"
	Content        string    `gorm:"type:text;not null" json:"content"`                // Markdown or HTML
	StepNumber     int       `gorm:"not null" json:"step_number"`                      // Order in the walkthrough

	MediaURLs      pq.StringArray `gorm:"type:text[]" json:"media_urls"`              // Optional: image/video links
	Tags           pq.StringArray `gorm:"type:text[]" json:"tags"`                    // e.g. ["Wild Pokémon", "Catching", "Battle Tips"]

	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type WalkthroughComment struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	WalkthroughID uuid.UUID `gorm:"type:uuid;not null" json:"walkthrough_id"`
	UserID        uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User          user.User `gorm:"foreignKey:UserID" json:"user"`
	Content       string    `gorm:"type:text;not null" json:"content"`

	CreatedAt     time.Time `json:"created_at"`
}
