package controllers

import (
	"pokemon/websocket"

	ws "github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

func ChatWebSocket(hub *websocket.Hub) func(*ws.Conn) {
	return func(conn *ws.Conn) {
		userID, ok := conn.Locals("user_id").(uuid.UUID)
		if !ok {
			conn.WriteMessage(ws.TextMessage, []byte("Unauthorized"))
			conn.Close()
			return
		}

		client := &websocket.Client{
			UserID: userID,
			Conn:   conn,
			Send:   make(chan []byte),
		}

		hub.Register <- client

		go client.ReadPump(hub)
		go client.WritePump()
	}
}
