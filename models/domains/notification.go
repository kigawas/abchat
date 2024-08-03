package domains

import "time"

type NotificationSetting struct {
	UserID    string `gorm:"primaryKey"`
	Email     bool
	Push      bool
	User      User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
