package utils

/*
Key Features:
Token Management:

✅ Access & Refresh Tokens - Complete token pair system
✅ Token Validation - Comprehensive validation with proper error handling
✅ Token Refresh - Secure token refresh mechanism
✅ Expiry Management - Flexible expiration times

Security Features:

✅ Token Blacklisting - Interface for revoking tokens
✅ Role-based Access - Built-in role checking
✅ Proper Error Handling - Specific error types
✅ Header Extraction - Safe token extraction from headers

Middleware Integration:

✅ Fiber Middleware - Ready-to-use authentication middleware
✅ Role Middleware - Permission-based access control
✅ Context Storage - User info stored in request context
*/

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
    ErrInvalidToken     = errors.New("invalid token")
    ErrExpiredToken     = errors.New("token has expired")
    ErrTokenNotFound    = errors.New("token not found")
    ErrInvalidClaims    = errors.New("invalid token claims")
    ErrTokenMalformed   = errors.New("token is malformed")
    ErrTokenUnverifiable = errors.New("token is unverifiable")
)

// JWTClaims represents the JWT token claims
type JWTClaims struct {
    UserID   uuid.UUID `json:"user_id"`
    Email    string `json:"email"`
    Username string `json:"username,omitempty"`
    Role     string `json:"role,omitempty"`
    TokenType string `json:"token_type"` // "access" or "refresh"
    jwt.RegisteredClaims
}

// JWTManager handles JWT token operations
type JWTManager struct {
    secretKey           []byte
    accessTokenExpiry   time.Duration
    refreshTokenExpiry  time.Duration
    issuer              string
}

// NewJWTManager creates a new JWT manager
func NewJWTManager(secretKey string, accessExpiry, refreshExpiry time.Duration, issuer string) *JWTManager {
    return &JWTManager{
        secretKey:          []byte(secretKey),
        accessTokenExpiry:  accessExpiry,
        refreshTokenExpiry: refreshExpiry,
        issuer:             issuer,
    }
}

// TokenPair represents access and refresh tokens
type TokenPair struct {
    AccessToken  string    `json:"access_token"`
    RefreshToken string    `json:"refresh_token"`
    TokenType    string    `json:"token_type"`
    ExpiresIn    int64     `json:"expires_in"`
    ExpiresAt    time.Time `json:"expires_at"`
}

// GenerateTokenPair generates both access and refresh tokens
func (j *JWTManager) GenerateTokenPair(userID uuid.UUID, email, username, role string) (*TokenPair, error) {
    // Generate access token
    accessToken, accessExp, err := j.generateToken(userID, email, username, role, "access", j.accessTokenExpiry)
    if err != nil {
        return nil, fmt.Errorf("failed to generate access token: %w", err)
    }

    // Generate refresh token
    refreshToken, _, err := j.generateToken(userID, email, username, role, "refresh", j.refreshTokenExpiry)
    if err != nil {
        return nil, fmt.Errorf("failed to generate refresh token: %w", err)
    }

    return &TokenPair{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
        TokenType:    "Bearer",
        ExpiresIn:    int64(j.accessTokenExpiry.Seconds()),
        ExpiresAt:    accessExp,
    }, nil
}

// GenerateAccessToken generates only an access token
func (j *JWTManager) GenerateAccessToken(userID uuid.UUID, email, username, role string) (string, time.Time, error) {
    return j.generateToken(userID, email, username, role, "access", j.accessTokenExpiry)
}

// GenerateRefreshToken generates only a refresh token
func (j *JWTManager) GenerateRefreshToken(userID uuid.UUID, email, username, role string) (string, time.Time, error) {
    return j.generateToken(userID, email, username, role, "refresh", j.refreshTokenExpiry)
}

// generateToken is the internal method to generate tokens
func (j *JWTManager) generateToken(userID uuid.UUID, email, username, role, tokenType string, expiry time.Duration) (string, time.Time, error) {
    now := time.Now()
    expiresAt := now.Add(expiry)

    claims := &JWTClaims{
        UserID:    userID,
        Email:     email,
        Username:  username,
        Role:      role,
        TokenType: tokenType,
        RegisteredClaims: jwt.RegisteredClaims{
            ID:        fmt.Sprintf("%s_%s_%d", userID.String(), tokenType, now.Unix()),
            Subject:   fmt.Sprintf("%d", userID),
            Audience:  jwt.ClaimStrings{j.issuer},
            Issuer:    j.issuer,
            IssuedAt:  jwt.NewNumericDate(now),
            NotBefore: jwt.NewNumericDate(now),
            ExpiresAt: jwt.NewNumericDate(expiresAt),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(j.secretKey)
    if err != nil {
        return "", time.Time{}, fmt.Errorf("failed to sign token: %w", err)
    }

    return tokenString, expiresAt, nil
}

// ValidateToken validates and parses a JWT token
func (j *JWTManager) ValidateToken(tokenString string) (*JWTClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        // Verify signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return j.secretKey, nil
    })

    if err != nil {
        // Handle specific JWT errors
        if errors.Is(err, jwt.ErrTokenMalformed) {
            return nil, ErrTokenMalformed
        } else if errors.Is(err, jwt.ErrTokenExpired) {
            return nil, ErrExpiredToken
        } else if errors.Is(err, jwt.ErrTokenNotValidYet) {
            return nil, ErrInvalidToken
        }
        return nil, fmt.Errorf("failed to parse token: %w", err)
    }

    claims, ok := token.Claims.(*JWTClaims)
    if !ok || !token.Valid {
        return nil, ErrInvalidClaims
    }

    return claims, nil
}

// RefreshTokens validates a refresh token and generates new token pair
func (j *JWTManager) RefreshTokens(refreshToken string) (*TokenPair, error) {
    claims, err := j.ValidateToken(refreshToken)
    if err != nil {
        return nil, fmt.Errorf("invalid refresh token: %w", err)
    }

    // Ensure it's a refresh token
    if claims.TokenType != "refresh" {
        return nil, errors.New("token is not a refresh token")
    }

    // Generate new token pair
    return j.GenerateTokenPair(claims.UserID, claims.Email, claims.Username, claims.Role)
}

// ExtractTokenFromHeader extracts JWT token from Authorization header
func ExtractTokenFromHeader(authHeader string) (string, error) {
    if authHeader == "" {
        return "", ErrTokenNotFound
    }

    const bearerPrefix = "Bearer "
    if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
        return "", ErrTokenMalformed
    }

    return authHeader[len(bearerPrefix):], nil
}

// GetTokenClaims extracts claims from token string
func (j *JWTManager) GetTokenClaims(tokenString string) (*JWTClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        return j.secretKey, nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*JWTClaims)
    if !ok {
        return nil, ErrInvalidClaims
    }

    return claims, nil
}

// IsTokenExpired checks if a token is expired without validating signature
func IsTokenExpired(tokenString string) bool {
    token, _ := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte("dummy"), nil // We don't validate signature here
    })

    if claims, ok := token.Claims.(*JWTClaims); ok {
        return claims.ExpiresAt.Time.Before(time.Now())
    }

    return true
}

// GetTokenRemainingTime returns the remaining time until token expires
func (j *JWTManager) GetTokenRemainingTime(tokenString string) (time.Duration, error) {
    claims, err := j.GetTokenClaims(tokenString)
    if err != nil {
        return 0, err
    }

    remaining := time.Until(claims.ExpiresAt.Time)
    if remaining < 0 {
        return 0, ErrExpiredToken
    }

    return remaining, nil
}

// BlacklistManager interface for token blacklisting
type BlacklistManager interface {
    BlacklistToken(tokenID string, expiry time.Time) error
    IsTokenBlacklisted(tokenID string) bool
}

// RedisBlacklistManager implements BlacklistManager using Redis
type RedisBlacklistManager struct {
    // Implementation would use Redis client
    // This is just the interface definition
}

// ValidateAndBlacklistCheck validates token and checks blacklist
func (j *JWTManager) ValidateAndBlacklistCheck(tokenString string, blacklist BlacklistManager) (*JWTClaims, error) {
    claims, err := j.ValidateToken(tokenString)
    if err != nil {
        return nil, err
    }

    // Check if token is blacklisted
    if blacklist != nil && blacklist.IsTokenBlacklisted(claims.ID) {
        return nil, errors.New("token has been revoked")
    }

    return claims, nil
}

// Legacy functions for backward compatibility

// GenerateJWT generates a simple JWT token (legacy)
func GenerateJWT(userID uuid.UUID, email string) (string, error) {
    manager := NewJWTManager("default-secret", 24*time.Hour, 7*24*time.Hour, "default-issuer")
    token, _, err := manager.GenerateAccessToken(userID, email, "", "")
    return token, err
}

// ValidateJWT validates a JWT token (legacy)
func ValidateJWT(tokenString string) (*JWTClaims, error) {
    manager := NewJWTManager("default-secret", 24*time.Hour, 7*24*time.Hour, "default-issuer")
    return manager.ValidateToken(tokenString)
}

// Example usage and middleware integration

// JWTMiddleware creates a Fiber middleware for JWT authentication
func JWTMiddleware(jwtManager *JWTManager, blacklist BlacklistManager) func(c *fiber.Ctx) error {
    return func(c *fiber.Ctx) error {
        // Extract token from header
        authHeader := c.Get("Authorization")
        tokenString, err := ExtractTokenFromHeader(authHeader)
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Missing or malformed token",
            })
        }

        // Validate token
        claims, err := jwtManager.ValidateAndBlacklistCheck(tokenString, blacklist)
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        // Ensure it's an access token
        if claims.TokenType != "access" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Invalid token type",
            })
        }

        // Store claims in context
        c.Locals("userID", claims.UserID)
        c.Locals("email", claims.Email)
        c.Locals("username", claims.Username)
        c.Locals("role", claims.Role)
        c.Locals("tokenID", claims.ID)

        return c.Next()
    }
}

// RoleMiddleware creates a middleware to check user roles
func RoleMiddleware(allowedRoles ...string) func(c *fiber.Ctx) error {
    return func(c *fiber.Ctx) error {
        userRole, ok := c.Locals("role").(string)
        if !ok {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
                "error": "Role information not found",
            })
        }

        for _, role := range allowedRoles {
            if userRole == role {
                return c.Next()
            }
        }

        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error": "Insufficient permissions",
        })
    }
}