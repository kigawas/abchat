package routers

import (
	"context"
	"encoding/json"
	"log"
	"slices"

	"github.com/gofiber/fiber/v3"
	"github.com/kigawas/abchat/app"
	"github.com/kigawas/abchat/app/persistence"
	"github.com/kigawas/abchat/app/services"
	"github.com/kigawas/abchat/models/params"
	"github.com/kigawas/abchat/models/schemas"
)

func CreateMessageRouter(router fiber.Router) {
	router.Post("/", messagePost).Delete("/:id", messageDelete)
}

func messagePost(c fiber.Ctx) error {
	db := app.GetDB()
	p := &params.SendMessageParams{}
	if err := c.Bind().JSON(p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Broadcast message to conversation participants
	conversation, err := persistence.GetConversation(db, p.ConversationID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get conversation"})
	}

	// check sender id is in conversation
	if !slices.Contains(conversation.Members, p.SenderID) {
		return c.Status(400).JSON(fiber.Map{"error": "Sender is not in conversation"})
	}

	message, err := persistence.CreateMessage(db, p)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create message"})
	}

	for _, member := range conversation.Members {
		if member != p.SenderID {
			go sendMessageToUser(member, &message)
		}
	}

	return c.JSON(message)
}

func messageDelete(c fiber.Ctx) error {
	db := app.GetDB()
	messageID := c.Params("id")

	err := persistence.DeleteMessage(db, messageID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete message"})
	}
	return c.SendStatus(204)
}

func sendMessageToUser(userID string, message *schemas.MessageSchema) {
	ctx := context.Background()
	isOnline, err := services.CheckOnlineUser(ctx, userID)
	if err != nil {
		log.Printf("Error checking online status: %v", err)
		return
	}
	log.Println("Sending message to user:", userID, "is online", isOnline)

	messageJSON, _ := json.Marshal(message)

	if isOnline {
		err := services.PublishOnlineMessage(ctx, userID, messageJSON)
		if err != nil {
			log.Printf("Error sending online message: %v", err)
		}
	} else {
		err := services.PushOfflineMessage(ctx, userID, messageJSON)
		if err != nil {
			log.Printf("Error sending offline message: %v", err)
		}
	}
}
