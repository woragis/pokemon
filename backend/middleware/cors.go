package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupCors() (fiber.Handler) {
	return cors.New(cors.Config{
			AllowOrigins:     "http://localhost:5173, https://yourdomain.com", // comma separated allowed origins
			AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
			AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
			AllowCredentials: true,
			// ExposeHeaders:   "Content-Length", // optional
			MaxAge: 3600,
		})
}