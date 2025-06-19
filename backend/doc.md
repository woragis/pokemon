# Documentation

## PKG

### Utils

#### JWT

```go
// Initialize JWT Manager
jwtManager := NewJWTManager(
    "your-secret-key",
    24*time.Hour,    // access token expiry
    7*24*time.Hour,  // refresh token expiry
    "your-app-name"  // issuer
)

// Generate token pair
tokens, err := jwtManager.GenerateTokenPair(userID, email, username, "user")

// Validate token
claims, err := jwtManager.ValidateToken(tokenString)

// Refresh tokens
newTokens, err := jwtManager.RefreshTokens(refreshToken)

// Use in Fiber routes
app.Use("/api", JWTMiddleware(jwtManager, nil))
app.Use("/admin", RoleMiddleware("admin", "moderator"))
```

```go
// In your main.go or config setup
jwtManager := utils.NewJWTManager(
    cfg.JWTSecret,
    time.Duration(cfg.JWTExpireHours)*time.Hour,
    time.Duration(cfg.JWTRefreshExpireHours)*time.Hour,
    "your-app-name",
)
```

#### Validation

```go
// Basic struct validation
type User struct {
    Email    string `json:"email" validate:"required,email"`
    Username string `json:"username" validate:"required,username"`
    Password string `json:"password" validate:"required,strong_password"`
}

user := User{Email: "test@test.com", Username: "john_doe", Password: "SecurePass123!"}
if err := ValidateStruct(&user); err != nil {
    // Returns structured error with field-specific messages
    fmt.Println(err) // "email is required; username must be 3-20 characters..."
}

// Individual field validation
if err := ValidateEmail("test@example.com"); err != nil {
    fmt.Println("Invalid email")
}

// Password confirmation
if err := ValidatePasswordConfirmation("SecurePass123!", "SecurePass123!"); err != nil {
    fmt.Println("Passwords don't match")
}

// File upload validation
fileParams := FileUploadParams{
    FileName: "document.pdf",
    FileSize: 1024000,
    MimeType: "application/pdf",
    ContentType: "application/pdf",
}
if err := ValidateStruct(&fileParams); err != nil {
    fmt.Println("Invalid file parameters")
}
```

```go
// In your user domain
type CreateUserRequest struct {
    Email           string `json:"email" validate:"required,email"`
    Username        string `json:"username" validate:"required,username"`
    Password        string `json:"password" validate:"required,strong_password"`
    ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
    FirstName       string `json:"first_name" validate:"required,min=2,max=50,alpha"`
    LastName        string `json:"last_name" validate:"required,min=2,max=50,alpha"`
    Phone           string `json:"phone" validate:"omitempty,phone"`
    Website         string `json:"website" validate:"omitempty,url"`
}

// In your handler
func (h *Handler) CreateUser(c *fiber.Ctx) error {
    var req CreateUserRequest
    if err := c.BodyParser(&req); err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
    }

    // This will return detailed validation errors
    if err := utils.ValidateStruct(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(err)
    }

    // Continue with user creation...
}
```

```json
{
  "errors": [
    {
      "field": "email",
      "value": "invalid-email",
      "tag": "email",
      "message": "email must be a valid email address"
    },
    {
      "field": "password",
      "value": "123",
      "tag": "strong_password",
      "message": "password must contain at least 8 characters with uppercase, lowercase, number and special character"
    }
  ]
}
```

## Install packages

```sh
// Additional useful packages for go.mod
/*
go mod init your-project
go get github.com/gofiber/fiber/v2
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/go-redis/redis/v8
go get golang.org/x/crypto
go get github.com/golang-jwt/jwt/v5
go get github.com/go-playground/validator/v10
go get github.com/joho/godotenv
go get github.com/google/uuid
go get go.uber.org/zap          // Better logging
go get github.com/stretchr/testify // Testing
go get github.com/golang-migrate/migrate/v4 // Database migrations
go get github.com/swaggo/fiber-swagger // API documentation
go get github.com/goccy/go-json // Faster JSON
*/
```
