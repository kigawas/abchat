package tests

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/kigawas/abchat/api"
	"github.com/kigawas/abchat/app"
	"github.com/kigawas/abchat/models/schemas"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	config := app.Config{
		DatabaseURL: "sqlite://file::api-user:?cache=shared&mode=memory",
	}
	router := api.CreateRouter(config)

	get(t, router, schemas.UserListSchema{Users: []schemas.UserSchema{}})
	post(t, router, "test", "test@abc.com")
}

func get(t *testing.T, router *fiber.App, expected schemas.UserListSchema) {
	req := httptest.NewRequest("GET", "/users", nil)
	resp, _ := router.Test(req)
	assert.Equal(t, 200, resp.StatusCode)

	AssertJsonResponse(t, resp, expected)
}

func post(t *testing.T, router *fiber.App, username string, email string) {
	req := httptest.NewRequest("POST", "/users", strings.NewReader(`{"username": "`+username+`", "email": "`+email+`"}`))
	resp, _ := router.Test(req)

	assert.Equal(t, 201, resp.StatusCode)
	var users schemas.UserSchema
	body, _ := io.ReadAll(resp.Body)

	json.Unmarshal(body, &users)
	assert.Equal(t, username, users.Username)
	assert.Equal(t, email, users.Email)
}
