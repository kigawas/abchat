package persistence

import (
	"github.com/google/uuid"
	"github.com/kigawas/abchat/models/domains"
	"github.com/kigawas/abchat/models/params"
	"github.com/kigawas/abchat/models/schemas"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, params *params.CreateUserParams) (schemas.UserSchema, error) {
	user := domains.User{
		ID:       uuid.New().String(),
		Username: params.Username,
		Email:    params.Email,
	}

	result := db.Create(&user)
	if result.Error != nil {
		return schemas.UserSchema{}, result.Error
	}
	return schemas.FromUser(user), nil
}

func GetUser(db *gorm.DB, ID string) (schemas.UserSchema, error) {
	var user domains.User
	if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
		return schemas.UserSchema{}, err
	}
	return schemas.FromUser(user), nil
}

func ListUsers(db *gorm.DB) (schemas.UserListSchema, error) {
	var users []domains.User
	if err := db.Find(&users).Error; err != nil {
		return schemas.UserListSchema{}, err
	}
	return schemas.FromUsers(users), nil
}

func DoesUserExist(db *gorm.DB, ID string) (bool, error) {
	var user domains.User
	if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}
