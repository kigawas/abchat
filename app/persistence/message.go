package persistence

import (
	"time"

	"github.com/google/uuid"
	"github.com/kigawas/abchat/models/domains"
	"github.com/kigawas/abchat/models/params"
	"github.com/kigawas/abchat/models/schemas"
	"gorm.io/gorm"
)

func CreateReadReceipt(db *gorm.DB, params params.CreateReadReceiptParams) error {
	readReceipt := domains.ReadReceipt{
		MessageID: params.MessageID,
		ReadBy:    params.ReadBy,
		ReadAt:    time.Now(),
	}
	return db.Create(&readReceipt).Error
}

func CreateMessage(db *gorm.DB, params *params.SendMessageParams) (schemas.MessageSchema, error) {
	message := domains.Message{
		ID:             uuid.New().String(),
		ConversationID: params.ConversationID,
		SenderID:       params.SenderID,
		Content:        params.Content,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	result := db.Create(&message)
	if result.Error != nil {
		return schemas.MessageSchema{}, result.Error
	}
	return schemas.FromMessage(message), nil
}

func GetMessages(db *gorm.DB, conversationID string) (schemas.MessageListSchema, error) {
	var messages []domains.Message
	if err := db.Model(&domains.Message{}).Preload("ReadReceipts").Find(&messages, "messages.is_deleted = false and messages.conversation_id = ?", conversationID).Error; err != nil {
		return schemas.MessageListSchema{}, err
	}
	return schemas.FromMessages(messages), nil
}

func DeleteMessage(db *gorm.DB, messageID string) error {
	// logic delete
	return db.Model(&domains.Message{}).Where("id = ?", messageID).Update("is_deleted", true).Error
}
