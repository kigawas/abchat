package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/kigawas/abchat/api"
	"github.com/kigawas/abchat/app"
	"github.com/kigawas/abchat/app/services"
)

func main() {
	app.LoadEnv()
	config := app.FromEnv()

	services.InitRedis(config.RedisURL)
	services.ClearOnlineUsers(context.Background())

	router := api.CreateRouter(config)

	go func() { // Graceful shutdown
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		router.Shutdown()
	}()

	if err := router.Listen(config.URL(), fiber.ListenConfig{
		EnablePrefork: config.Prefork,
	}); err != nil {
		log.Fatal(err)
	}

}
