package utils

import (
	"pokemon/internal/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupCors(config *config.Config) fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     strings.Join(config.CORSAllowedOrigins, ","),
		AllowMethods:     strings.Join(config.CORSAllowedMethods, ","),
		AllowHeaders:     strings.Join(config.CORSAllowedHeaders, ","),
		// AllowOrigins:     "*",
		// AllowMethods:     "*",
		// AllowHeaders:     "*",
		AllowCredentials: true,
		MaxAge:           3600,
	})
}
