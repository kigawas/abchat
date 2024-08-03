package tests

import (
	"testing"

	"github.com/kigawas/abchat/app"
	"github.com/kigawas/abchat/app/persistence"
	"github.com/kigawas/abchat/models/params"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUser(t *testing.T) {
	db := app.SetupDB("sqlite://file::user:?cache=shared&mode=memory", &gorm.Config{})
	app.MigrateDB(db)

	user, _ := persistence.CreateUser(db, &params.CreateUserParams{Username: "test", Email: "test@abc.com"})
	assert.Equal(t, "test", user.Username)

	users, _ := persistence.ListUsers(db)
	assert.Equal(t, "test", users.Users[0].Username)
	assert.Equal(t, "test@abc.com", users.Users[0].Email)
}
