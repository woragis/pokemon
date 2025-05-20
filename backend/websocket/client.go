package websocket

import (
	"encoding/json"
	"pokemon/database"
	"pokemon/models"
	"sync"
	"time"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

type Client struct {
	UserID uuid.UUID
	Conn   *websocket.Conn
	Send   chan []byte
}

type Hub struct {
	Clients    map[uuid.UUID]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
	mu         sync.RWMutex
}

type Message struct {
	SenderID   uuid.UUID `json:"sender_id"`
	ReceiverID uuid.UUID `json:"receiver_id"`
	Content    string    `json:"content"`
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[uuid.UUID]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			h.Clients[client.UserID] = client
			h.mu.Unlock()

		case client := <-h.Unregister:
			h.mu.Lock()
			if _, ok := h.Clients[client.UserID]; ok {
				delete(h.Clients, client.UserID)
				close(client.Send)
			}
			h.mu.Unlock()

		case msg := <-h.Broadcast:
			h.mu.RLock()
			if receiver, ok := h.Clients[msg.ReceiverID]; ok {
				msgBytes, _ := json.Marshal(msg)
				receiver.Send <- msgBytes
			}
			h.mu.RUnlock()
		}
	}
}

func (c *Client) ReadPump(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var msg Message
		if err := c.Conn.ReadJSON(&msg); err != nil {
			break
		}

		// Save message
		db := database.DB
		db.Create(&models.ChatMessage{
			ID:         uuid.New(),
			SenderID:   msg.SenderID,
			ReceiverID: msg.ReceiverID,
			Content:    msg.Content,
			CreatedAt:  time.Now(),
		})

		hub.Broadcast <- msg
	}
}

func (c *Client) WritePump() {
	for {
		msg, ok := <-c.Send
		if !ok {
			c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
		c.Conn.WriteMessage(websocket.TextMessage, msg)
	}
}
