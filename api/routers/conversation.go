package routers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kigawas/abchat/app"
	"github.com/kigawas/abchat/app/persistence"
	"github.com/kigawas/abchat/models/params"
)

func CreateConversationRouter(router fiber.Router) {
	router.Get("/", conversationsGet).Post("/", conversationPost).
		Get("/:id", conversationGetId).Get("/:id/messages", conversationGetIdMessages)
}

func conversationPost(c fiber.Ctx) error {
	db := app.GetDB()

	p := &params.CreateConversationParams{}
	if err := c.Bind().JSON(p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if len(p.UserIDs) < 2 {
		return c.Status(400).JSON(fiber.Map{"error": "At least two users are required"})
	}

	conversation, err := persistence.CreateConversation(db, p)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create conversation"})
	}
	return c.JSON(conversation)
}

func conversationsGet(c fiber.Ctx) error {
	db := app.GetDB()

	conversations, err := persistence.ListConversations(db)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to list conversations"})
	}
	return c.JSON(conversations)
}

func conversationGetId(c fiber.Ctx) error {
	db := app.GetDB()
	conversationID := c.Params("id")

	conversation, err := persistence.GetConversation(db, conversationID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get conversation"})
	}
	return c.JSON(conversation)
}

func conversationGetIdMessages(c fiber.Ctx) error {
	db := app.GetDB()
	conversationID := c.Params("id")

	messages, err := persistence.GetMessages(db, conversationID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get messages"})
	}
	return c.JSON(messages)
}
