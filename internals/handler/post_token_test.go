package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"neti/internals/domain"
	"neti/mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostTokenApi(t *testing.T) {
	var password = mock.PasswordMock{}
	var users = mock.UsersMock{}
	var auth = mock.AuthMock{}
	var clients = mock.ClientsMock{}
	clients.On("FindBy", "aClientId").Return(domain.Client{ClientId: "aClientId", ClientSecret: "aClientSecret"})

	var router = gin.Default()

	t.Run("should check if user credential are right when grant type is password", func(t *testing.T) {
		password.On("Compare", "hashPassword", []byte("admin")).Return(true)
		users.On("FindBy", "admin").Return("admin", "hashPassword")
		auth.On("AccessToken").Return(domain.TokenResponse{AccessToken: "anAccessToken", State: "aState", TokenType: "aTokenType", ExpiresIn: "expired"})

		router.POST("/token", PostTokenApi(auth, users, password, clients))

		body, _ := json.Marshal(TokenRequest{GrantType: "password", ClientId: "client_id", Username: "admin", Password: "admin"})
		request, _ := http.NewRequest("POST", "/token", bytes.NewBuffer(body))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		auth.AssertExpectations(t)
		users.AssertExpectations(t)
		password.AssertExpectations(t)

		res, _ := json.Marshal(domain.TokenResponse{AccessToken: "anAccessToken", State: "aState", TokenType: "aTokenType", ExpiresIn: "expired"})
		assert.Contains(t, response.Body.String(), string(res))
	})

	t.Run("should check if client credentials are right when grant type is credentials", func(t *testing.T) {
		auth.On("AccessToken").Return(domain.TokenResponse{AccessToken: "anAccessToken", State: "aState", TokenType: "aTokenType", ExpiresIn: "expired"})

		//router.POST("/token", PostTokenApi(auth, users, password, clients))

		body, _ := json.Marshal(TokenRequest{GrantType: "credentials", ClientId: "aClientId", ClientSecret: "aClientSecret"})
		request, _ := http.NewRequest("POST", "/token", bytes.NewBuffer(body))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		clients.AssertExpectations(t)

		res, _ := json.Marshal(domain.TokenResponse{AccessToken: "anAccessToken", State: "aState", TokenType: "aTokenType", ExpiresIn: "expired"})
		assert.Contains(t, response.Body.String(), string(res))

	})

}
