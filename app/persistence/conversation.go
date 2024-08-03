package persistence

import (
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/kigawas/abchat/models/domains"
	"github.com/kigawas/abchat/models/params"
	"github.com/kigawas/abchat/models/schemas"
	"gorm.io/gorm"
)

func CreateConversation(db *gorm.DB, params *params.CreateConversationParams) (schemas.ConversationMembersSchema, error) {
	isGroup := len(params.UserIDs) > 2
	userIDs := params.UserIDs

	var convID string
	if isGroup {
		convID = uuid.New().String()
	} else {
		slices.Sort(userIDs)
		convID = fmt.Sprintf("Conv-%s-%s", userIDs[0], userIDs[1])
	}

	conversation := domains.Conversation{
		ID:      convID,
		IsGroup: isGroup,
		Name:    params.Name,
	}
	members := make([]domains.ConversationMember, len(userIDs))
	for i, userID := range userIDs {
		members[i] = domains.ConversationMember{
			ConversationID: convID,
			UserID:         userID,
			JoinedAt:       time.Now(),
		}
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&conversation).Error; err != nil {
			return err
		}
		if err := tx.Create(&members).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return schemas.ConversationMembersSchema{}, err
	}
	return schemas.FromConversationMembers(conversation, members), nil
}

func ListConversations(db *gorm.DB) ([]schemas.ConversationSchema, error) {
	var conversations []domains.Conversation
	if err := db.Find(&conversations).Error; err != nil {
		return nil, err
	}
	return schemas.FromConversations(conversations), nil
}

func GetConversation(db *gorm.DB, conversationID string) (schemas.ConversationMembersSchema, error) {
	var conversation domains.Conversation
	if err := db.First(&conversation, "id = ?", conversationID).Error; err != nil {
		return schemas.ConversationMembersSchema{}, err
	}
	var conversationMembers []domains.ConversationMember
	if err := db.Find(&conversationMembers, "conversation_id = ?", conversationID).Error; err != nil {
		return schemas.ConversationMembersSchema{}, err
	}
	return schemas.FromConversationMembers(conversation, conversationMembers), nil
}
