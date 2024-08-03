package params

type CreateConversationParams struct {
	UserIDs []string `json:"user_ids"`
	Name    string   `json:"name,omitempty"`
}
