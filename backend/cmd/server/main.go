package main

import (
	"log"
	"pokemon/internal/config"
	"pokemon/internal/database"
	"pokemon/internal/migrations"
	"pokemon/pkg/utils"

	"pokemon/internal/domains/blog"
	favMon "pokemon/internal/domains/favorite-pokemon"
	"pokemon/internal/domains/forum"
	"pokemon/internal/domains/game"
	"pokemon/internal/domains/guide"
	"pokemon/internal/domains/news"
	"pokemon/internal/domains/shout"
	"pokemon/internal/domains/team"
	"pokemon/internal/domains/user"
	"pokemon/internal/domains/walkthrough"
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
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

    // Run database migrations
    if err := migrations.RunAll(db, migrations.GetAllMigrators()); err != nil {
		log.Fatal("migration failed:", err)
	}
    
    // Initialize cache with redis
    redis := database.NewRedis(cfg.RedisURL)
    
    // Initialize Fiber app
    app := fiber.New(fiber.Config{
        ErrorHandler: middleware.ErrorHandler,
    })

    c := config.Load()
    
    // Global middleware
    app.Use(logger.New())
    app.Use(utils.SetupCors(c))
    
    // API routes
    api := app.Group("/api/v1")
    
    // Initialize domains
    user.NewHandler(db, redis).RegisterRoutes(api)
    blog.NewHandler(db, redis).RegisterRoutes(api)
    favMon.NewHandler(db, redis).RegisterRoutes(api)
    forum.NewHandler(db, redis).RegisterRoutes(api)
    game.NewHandler(db, redis).RegisterRoutes(api)
    guide.NewHandler(db, redis).RegisterRoutes(api)
    news.NewHandler(db, redis).RegisterRoutes(api)
    shout.NewHandler(db, redis).RegisterRoutes(api)
    team.NewHandler(db, redis).RegisterRoutes(api)
    walkthrough.NewHandler(db, redis).RegisterRoutes(api)
    
    log.Fatal(app.Listen(":" + cfg.Port))
}
