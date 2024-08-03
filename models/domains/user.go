package domains

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        string `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
