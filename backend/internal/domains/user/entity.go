package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
    ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
    Email     string         `json:"email" gorm:"uniqueIndex;not null"`
    Username  string         `json:"username" gorm:"uniqueIndex;not null"`
    Password  string         `json:"-" gorm:"not null"`
    FirstName string         `json:"first_name"`
    LastName  string         `json:"last_name"`
    Active    bool           `json:"active" gorm:"default:true"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type CreateUserRequest struct {
    Email     string `json:"email" validate:"required,email"`
    Username  string `json:"username" validate:"required,min=3,max=20"`
    Password  string `json:"password" validate:"required,min=6"`
    FirstName string `json:"first_name" validate:"required"`
    LastName  string `json:"last_name" validate:"required"`
}

type UpdateUserRequest struct {
    FirstName *string `json:"first_name,omitempty"`
    LastName  *string `json:"last_name,omitempty"`
    Active    *bool   `json:"active,omitempty"`
}

type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}
