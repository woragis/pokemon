package main

import (
	"pokemon/database"
	"pokemon/middleware"
	"pokemon/routes"
	"pokemon/websocket"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    app := fiber.New()
    ws_hub := websocket.NewHub()

    database.ConnectDB()

    // Logging
    app.Use(logger.New())

    // API Group routes
    api := app.Group("/api")

    // Setup CORS for API routes
    api.Use(middleware.SetupCors())

    routes.SetupAuthRoutes(api)
    routes.RegisterShoutRoutes(api)
    routes.RegisterPokeFeedRoutes(api)
    routes.RegisterGameGuideRoutes(api)
    routes.RegisterChatRoutes(api, ws_hub)
    routes.RegisterPokedexRoutes(api)
    routes.BlogRoutes(api)
    routes.SetupGameRoutes(api)
    routes.SetupPermissionRoutes(api)
    routes.SetupRoleRoutes(api)
    routes.SetupUserRoutes(api)

    app.Listen(":3000")
}
