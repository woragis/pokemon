package main

import (
	"pokemon/database"
	"pokemon/middleware"
	"pokemon/routes"
	"pokemon/websocket"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    ws_hub := websocket.NewHub()

    database.ConnectDB()

    // API Group routes
    api := app.Group("/api")

    // Public routes
    public := api.Group("/")
    routes.SetupAuthRoutes(public)

    // Logged in routes
    auth := api.Group("/", middleware.RequireAuth())
    routes.RegisterShoutRoutes(auth)
    routes.RegisterPokeFeedRoutes(auth)
    routes.RegisterGameGuideRoutes(auth)
    routes.RegisterChatRoutes(auth, ws_hub)
    routes.RegisterPokedexRoutes(auth)

    // Admin routes
    admin := api.Group("/admin", middleware.RequireAuth(), middleware.RequireRole())
    routes.RegisterAdminRoutes(admin)

    app.Listen(":3000")
}
