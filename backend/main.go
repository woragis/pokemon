package main

import (
	"pokemon/database"
	"pokemon/routes"
	"pokemon/websocket"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    db := database.ConnectDB()
    ws_hub := websocket.NewHub()
    routes.SetupAuthRoutes(app, db)
    routes.RegisterShoutRoutes(app)
    routes.RegisterPokeFeedRoutes(app)
    routes.RegisterGameGuideRoutes(app)
    routes.RegisterChatRoutes(app, ws_hub)
    routes.RegisterAdminRoutes(app)

    app.Listen(":3000")
}
