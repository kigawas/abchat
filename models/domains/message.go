package domains

import "time"

// Message represents a single message in a conversation
type Message struct {
	ID             string `gorm:"primaryKey"`
	ConversationID string `gorm:"not null"`
	SenderID       string `gorm:"not null"`
	Content        string `gorm:"type:text;not null"`
	IsDeleted      bool   `gorm:"default:false"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	Conversation Conversation `gorm:"foreignKey:ConversationID"`
	Sender       User         `gorm:"foreignKey:SenderID"`
	ReadReceipts []ReadReceipt
}

// ReadReceipt tracks whether a message has been read by a user
type ReadReceipt struct {
	MessageID string    `gorm:"primaryKey"`
	ReadBy    string    `gorm:"primaryKey"`
	ReadAt    time.Time `gorm:"not null"`

	Message Message `gorm:"foreignKey:MessageID"`
	Reader  User    `gorm:"foreignKey:ReadBy"`
}
