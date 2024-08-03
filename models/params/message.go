package params

type SendMessageParams struct {
	SenderID       string `json:"sender_id"`
	ConversationID string `json:"conversation_id"`
	Content        string `json:"content"`
}

type CreateReadReceiptParams struct {
	MessageID string `json:"message_id"`
	ReadBy    string `json:"read_by"`
}
