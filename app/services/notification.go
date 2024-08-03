package services

import (
	"log"
	"time"

	"github.com/kigawas/abchat/app"
	"github.com/kigawas/abchat/app/persistence"
)

func SendNotification(userID string, message string) {
	db := app.GetDB()
	setting, err := persistence.GetNotificationSetting(db, userID)
	// if not exists, create default
	if err != nil {
		_, _ = persistence.CreateDefaultNotificationSetting(db, userID)
		setting, _ = persistence.GetNotificationSetting(db, userID)
	}

	if !setting.Email && !setting.Push {
		return
	}

	if setting.Email {
		user, _ := persistence.GetUser(db, userID)
		notifyByEmail(userID, user.Email, message)
	}

	if setting.Push {
		notifyByPush(userID, "token-push", message)
	}
}

func notifyByEmail(userID string, email string, message string) {
	log.Printf("Notify user: %s, email: %s with message: %s\n", userID, email, message)
	time.Sleep(time.Second) // simulate latency
}

func notifyByPush(userID string, token string, message string) {
	log.Printf("Notify user: %s, token: %s with message: %s\n", userID, token, message)
	time.Sleep(time.Second)
}
