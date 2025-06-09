package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// ErrorResponse represents the structure of error responses
type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}

// ErrorHandler is a custom error handler for Fiber
func ErrorHandler(c *fiber.Ctx, err error) error {
	// Default error code
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	// Check if it's a Fiber error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	// Log the error for debugging
	log.Printf("Error: %v, Code: %d, Path: %s, Method: %s", 
		err, code, c.Path(), c.Method())

	// Send custom error response
	return c.Status(code).JSON(&ErrorResponse{
		Error:   true,
		Message: message,
		Code:    code,
	})
}
