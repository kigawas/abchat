package schemas

import "github.com/kigawas/abchat/models/domains"

type UserSchema struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserListSchema struct {
	Users []UserSchema `json:"users"`
}

func FromUser(u domains.User) UserSchema {
	return UserSchema{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}

func FromUsers(u []domains.User) UserListSchema {
	users := make([]UserSchema, len(u))
	for i, user := range u {
		users[i] = FromUser(user)
	}
	return UserListSchema{
		Users: users,
	}
}
