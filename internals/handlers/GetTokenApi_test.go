package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTokenApi(t *testing.T) {
	router := gin.Default()
	router.POST("/token", GetTokenApi())

	body, _ := json.Marshal(TokenRequest{GrantType: "password", ClientId: "client_id", Username: "admin", Password: "admin"})

	request, _ := http.NewRequest("POST", "/token", bytes.NewBuffer(body))

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	res, _ := json.Marshal(TokenResponse{AccessToken: "anAccessToken", State: "aState", TokenType: "aTokenType", ExpiresIn: "expired"})
	assert.Contains(t, response.Body.String(), string(res))
}
