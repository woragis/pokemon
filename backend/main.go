package main

import (
	"pokemon/database"
	"pokemon/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    db := database.ConnectDB()
    routes.SetupAuthRoutes(app, db)
    routes.RegisterShoutRoutes(app)
    routes.RegisterPokeFeedRoutes(app)
    routes.RegisterGameGuideRoutes(app)

    app.Listen(":3000")
}
