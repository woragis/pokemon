package routes

import (
	"pokemon/controllers"
	ws "pokemon/websocket"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func RegisterChatRoutes(app *fiber.App, hub *ws.Hub) {
	app.Get("/ws/chat", websocket.New(controllers.ChatWebSocket(hub)))
}
