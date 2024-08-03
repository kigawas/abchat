package middleware

import (
	"context"
	"encoding/json"
	"log"

	"github.com/kigawas/abchat/app"
	"github.com/kigawas/abchat/app/persistence"
	"github.com/kigawas/abchat/app/services"
	"github.com/kigawas/abchat/models/params"
	"github.com/kigawas/abchat/models/schemas"
	"github.com/kigawas/abchat/websocket"
	"github.com/redis/go-redis/v9"
)

func HandleWebSocket(c *websocket.Conn) {
	db := app.GetDB()
	userID := c.Params("userId")

	exist, err := persistence.DoesUserExist(db, userID)
	if !exist || err != nil {
		c.WriteMessage(websocket.TextMessage, []byte("User not found"))
		c.Close()
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	online, err := services.CheckOnlineUser(ctx, userID)
	if online || err != nil {
		c.WriteMessage(websocket.TextMessage, []byte("Already online"))
		c.Close()
		return
	}

	services.AddOnlineUser(ctx, userID)
	// services.SaveUserConnection(c, userID)
	pullOfflineMessages(ctx, c, userID)
	go listenForOnlineMessages(ctx, c, userID)

	defer func() {
		log.Println("Exiting websocket handler for user:", userID)
		services.RemoveOnlineUser(ctx, userID)
		// services.DeleteUserConnection(userID)
		c.Close()
	}()

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			break
		}
	}
}

func pullOfflineMessages(ctx context.Context, c *websocket.Conn, userID string) {
	for {
		// Get queued message from Redis
		result, err := services.PopOfflineMessages(ctx, userID)
		if err == redis.Nil {
			// No more messages in queue
			break
		} else if err != nil {
			log.Printf("Error retrieving queued message: %v", err)
			break
		}

		// Send message via WebSocket
		var message schemas.MessageSchema
		json.Unmarshal([]byte(result), &message)
		deliverMessage(c, userID, &message)
	}
}

func deliverMessage(conn *websocket.Conn, receiverUserID string, message *schemas.MessageSchema) {
	messageJSON, _ := json.Marshal(message)
	err := conn.WriteMessage(websocket.TextMessage, messageJSON)
	if err != nil {
		log.Println("Failed to deliver message to user:", receiverUserID)
	} else {
		go services.SendNotification(receiverUserID, message.Content)
		persistence.CreateReadReceipt(app.GetDB(), params.CreateReadReceiptParams{
			MessageID: message.ID,
			ReadBy:    receiverUserID,
		})
	}
}

func listenForOnlineMessages(ctx context.Context, conn *websocket.Conn, userID string) {
	pubsub := services.SubscribeOnlineMessages(ctx, userID)
	defer pubsub.Close()

	for {
		select {
		case <-ctx.Done():
			log.Println("Stop listening for user:", userID)
			return

		default:
			redisMsg, err := pubsub.ReceiveMessage(ctx)

			if err != nil {
				if err == context.Canceled {
					log.Println("Redis receive canceled for user:", userID)
					return
				}
				log.Println("Redis receive error:", err)
				continue
			}

			var message schemas.MessageSchema
			err = json.Unmarshal([]byte(redisMsg.Payload), &message)
			if err != nil {
				log.Println("JSON unmarshal error:", err)
				continue
			}

			deliverMessage(conn, userID, &message)
		}
	}
}
