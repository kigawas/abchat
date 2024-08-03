package schemas

import (
	"time"

	"github.com/kigawas/abchat/models/domains"
)

type ConversationSchema struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsGroup   bool   `json:"is_group"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ConversationMembersSchema struct {
	ConversationSchema
	Members []string `json:"members"`
}

func FromConversation(c domains.Conversation) ConversationSchema {
	return ConversationSchema{
		ID:        c.ID,
		Name:      c.Name,
		IsGroup:   c.IsGroup,
		CreatedAt: c.CreatedAt.Format(time.RFC3339),
		UpdatedAt: c.UpdatedAt.Format(time.RFC3339),
	}
}

func FromConversations(c []domains.Conversation) []ConversationSchema {
	conversations := make([]ConversationSchema, len(c))
	for i, conversation := range c {
		conversations[i] = FromConversation(conversation)
	}
	return conversations
}

func FromConversationMembers(c domains.Conversation, members []domains.ConversationMember) ConversationMembersSchema {
	membersIDs := make([]string, len(members))
	for i, member := range members {
		membersIDs[i] = member.UserID
	}
	return ConversationMembersSchema{
		ConversationSchema: FromConversation(c),
		Members:            membersIDs,
	}
}
