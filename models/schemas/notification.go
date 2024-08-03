package schemas

import (
	"time"

	"github.com/kigawas/abchat/models/domains"
)

type NotificationSettingSchema struct {
	UserID    string    `json:"user_id"`
	Email     bool      `json:"email"`
	Push      bool      `json:"push"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromNotificationSetting(notificationSetting domains.NotificationSetting) NotificationSettingSchema {
	return NotificationSettingSchema{
		UserID:    notificationSetting.UserID,
		Email:     notificationSetting.Email,
		Push:      notificationSetting.Push,
		UpdatedAt: notificationSetting.UpdatedAt,
	}
}
