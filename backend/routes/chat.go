package routes

import (
	"pokemon/controllers"
	"pokemon/middleware"
	ws "pokemon/websocket"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func RegisterChatRoutes(api fiber.Router, hub *ws.Hub) {
	api.Get("/ws/chat", middleware.RequireAuth(), websocket.New(controllers.ChatWebSocket(hub)))
}
