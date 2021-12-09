package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"neti/internals/domain"
	"neti/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTokenApi(t *testing.T) {

	auth := mocks.AuthMock{}
	auth.On("AccessToken").Return(domain.TokenResponse{AccessToken: "anAccessToken", State: "aState", TokenType: "aTokenType", ExpiresIn: "expired"})

	router := gin.Default()
	router.POST("/token", GetTokenApi(auth))

	body, _ := json.Marshal(TokenRequest{GrantType: "password", ClientId: "client_id", Username: "admin", Password: "admin"})

	request, _ := http.NewRequest("POST", "/token", bytes.NewBuffer(body))

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	auth.AssertExpectations(t)
	res, _ := json.Marshal(domain.TokenResponse{AccessToken: "anAccessToken", State: "aState", TokenType: "aTokenType", ExpiresIn: "expired"})
	assert.Contains(t, response.Body.String(), string(res))
}
