package controllers

import (
	"pokemon/database"
	"pokemon/models"
	"time"

	"github.com/google/uuid"
)

func CreateNotification(userID uuid.UUID, notifType string, message string) {
    notif := models.Notification{
        ID:        uuid.New(),
        UserID:    userID,
        Type:      notifType,
        Message:   message,
        CreatedAt: time.Now(),
    }
    _ = database.DB.Create(&notif)
}
