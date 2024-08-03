package schemas

import (
	"time"

	"github.com/kigawas/abchat/models/domains"
)

type ReadReceiptSchema struct {
	ReadBy string `json:"read_by"`
	ReadAt string `json:"read_at"`
}
type MessageSchema struct {
	ID             string              `json:"id"`
	ConversationID string              `json:"conversation_id"`
	SenderID       string              `json:"sender_id"`
	Content        string              `json:"content"`
	SentAt         string              `json:"sent_at"`
	UpdatedAt      string              `json:"updated_at"`
	ReadReceipts   []ReadReceiptSchema `json:"read_receipts"`
}

type MessageListSchema struct {
	Messages []MessageSchema `json:"messages"`
}

func FromMessage(m domains.Message) MessageSchema {
	receipts := make([]ReadReceiptSchema, len(m.ReadReceipts))

	for i, receipt := range m.ReadReceipts {
		receipts[i] = ReadReceiptSchema{
			ReadBy: receipt.ReadBy,
			ReadAt: receipt.ReadAt.Format(time.RFC3339),
		}
	}

	return MessageSchema{
		ID:             m.ID,
		ConversationID: m.ConversationID,
		SenderID:       m.SenderID,
		Content:        m.Content,
		SentAt:         m.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      m.UpdatedAt.Format(time.RFC3339),
		ReadReceipts:   receipts,
	}
}

func FromMessages(m []domains.Message) MessageListSchema {
	messages := make([]MessageSchema, len(m))

	for i, message := range m {
		messages[i] = FromMessage(message)
	}

	return MessageListSchema{
		Messages: messages,
	}
}
