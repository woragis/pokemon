package main

import (
	"log"
	"pokemon/internal/config"
	"pokemon/internal/database"

	"pokemon/internal/domains/user"
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    // Load configuration
    cfg := config.Load()
    
    // Initialize database with postgres
    db, err := database.NewPostgres(cfg.DatabaseURL)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    
    // Initialize cache with redis
    redis := database.NewRedis(cfg.RedisURL)
    
    // Initialize Fiber app
    app := fiber.New(fiber.Config{
        ErrorHandler: middleware.ErrorHandler,
    })
    
    // Global middleware
    app.Use(logger.New())
    app.Use(cors.New(cors.Config{
        AllowMethods: "GET,POST,PUT,DELETE",
        AllowHeaders: "Content-Type,Authorization",
        AllowOrigins: "http://localhost:5173",
    }))
    
    // API routes
    api := app.Group("/api/v1")
    
    // Initialize domains
    user.NewHandler(db, redis).RegisterRoutes(api)
    
    log.Fatal(app.Listen(":" + cfg.Port))
}
