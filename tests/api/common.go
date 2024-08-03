package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertJsonResponse[T any](t *testing.T, resp *http.Response, expected T) {
	body, _ := io.ReadAll(resp.Body)

	var respBody T
	_ = json.Unmarshal(body, &respBody)

	assert.Equal(t, expected, respBody)
}
