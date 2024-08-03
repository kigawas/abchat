package domains

import (
	"time"

	"gorm.io/gorm"
)

// Conversation represents a chat conversation (one-to-one or group)
type Conversation struct {
	ID        string `gorm:"primaryKey"`
	Name      string // Can be null for one-to-one chats
	IsGroup   bool   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// ConversationMember represents a user's membership in a conversation
type ConversationMember struct {
	ConversationID string `gorm:"primaryKey"`
	UserID         string `gorm:"primaryKey"`
	JoinedAt       time.Time
	LeftAt         *time.Time // Null if still a member
	CreatedAt      time.Time
	UpdatedAt      time.Time

	Conversation Conversation `gorm:"foreignKey:ConversationID"`
	User         User         `gorm:"foreignKey:UserID"`
}
