package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kigawas/abchat/api/middleware"
	"github.com/kigawas/abchat/api/routers"
	"github.com/kigawas/abchat/app"
	"github.com/kigawas/abchat/websocket"
	"gorm.io/gorm"
)

func setupRouter(fApp *fiber.App) {
	var m = map[string]func(fiber.Router){
		"/users":         routers.CreateUserRouter,
		"/conversations": routers.CreateConversationRouter,
		"/messages":      routers.CreateMessageRouter,
	}
	for path, createRouter := range m {
		createRouter(fApp.Group(path))
	}
}

func CreateRouter(config app.Config) *fiber.App {
	db := app.SetupDB(config.DatabaseURL, &gorm.Config{})
	app.MigrateDB(db)

	router := fiber.New(createConfig())
	setupRouter(router)

	router.Get("/ws/:userId", websocket.New(middleware.HandleWebSocket))
	return router
}
