package persistence

import (
	"github.com/kigawas/abchat/models/domains"
	"github.com/kigawas/abchat/models/params"
	"github.com/kigawas/abchat/models/schemas"
	"gorm.io/gorm"
)

func GetNotificationSetting(db *gorm.DB, userID string) (schemas.NotificationSettingSchema, error) {
	var notificationSetting domains.NotificationSetting
	if err := db.Where("user_id = ?", userID).First(&notificationSetting).Error; err != nil {
		return schemas.NotificationSettingSchema{}, err
	}
	return schemas.FromNotificationSetting(notificationSetting), nil
}

func CreateDefaultNotificationSetting(db *gorm.DB, userID string) (schemas.NotificationSettingSchema, error) {
	notificationSetting := domains.NotificationSetting{
		UserID: userID,
		Email:  false,
		Push:   true,
	}
	result := db.Create(&notificationSetting)
	if result.Error != nil {
		return schemas.NotificationSettingSchema{}, result.Error
	}
	return schemas.FromNotificationSetting(notificationSetting), nil
}

func UpdateNotificationSetting(db *gorm.DB, userID string, params *params.UpdateNotificationSettingParams) (schemas.NotificationSettingSchema, error) {
	var notificationSetting domains.NotificationSetting
	if err := db.Where("user_id = ?", userID).First(&notificationSetting).Error; err != nil {
		return schemas.NotificationSettingSchema{}, err
	}

	notificationSetting.Email = params.Email
	notificationSetting.Push = params.Push

	result := db.Save(&notificationSetting)
	if result.Error != nil {
		return schemas.NotificationSettingSchema{}, result.Error
	}
	return schemas.FromNotificationSetting(notificationSetting), nil
}
