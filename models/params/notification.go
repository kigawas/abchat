package params

type UpdateNotificationSettingParams struct {
	Email bool `json:"email"`
	Push  bool `json:"push"`
}
