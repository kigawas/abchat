package routers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kigawas/abchat/app"
	"github.com/kigawas/abchat/app/persistence"
	"github.com/kigawas/abchat/models/params"
)

func CreateUserRouter(router fiber.Router) {
	router.Get("/", usersGet).Get("/:id", userGet).Post("/", userPost).
		Get("/:id/setting", userSettingGet).Put("/:id/setting", userSettingPut)
}

func userPost(c fiber.Ctx) error {
	db := app.GetDB()
	p := &params.CreateUserParams{}
	if err := c.Bind().JSON(p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user data"})
	}

	user, err := persistence.CreateUser(db, p)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(201).JSON(user)
}

func usersGet(c fiber.Ctx) error {
	db := app.GetDB()

	users, err := persistence.ListUsers(db)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to list users"})
	}
	return c.JSON(users)
}

func userGet(c fiber.Ctx) error {
	db := app.GetDB()

	userID := c.Params("id")
	user, err := persistence.GetUser(db, userID)
	if err != nil {
		return handleError(c, err)
	}
	return c.JSON(user)
}

func userSettingGet(c fiber.Ctx) error {
	db := app.GetDB()

	userID := c.Params("id")
	setting, err := persistence.GetNotificationSetting(db, userID)
	if err != nil {
		return handleError(c, err)
	}
	return c.JSON(setting)
}

func userSettingPut(c fiber.Ctx) error {
	db := app.GetDB()

	userID := c.Params("id")
	p := &params.UpdateNotificationSettingParams{}
	if err := c.Bind().JSON(p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user data"})
	}

	setting, err := persistence.UpdateNotificationSetting(db, userID, p)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update notification setting"})
	}
	return c.JSON(setting)
}
