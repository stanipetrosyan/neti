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

func TestPasswordGrantType(t *testing.T) {
	t.Run("should check if user credential are right when grant type is password", func(t *testing.T) {
		var password = mock.PasswordMock{}
		var users = mock.UsersMock{}
		var auth = mock.AuthMock{}
		var clients = mock.ClientsMock{}
		var codes = mock.CodesMock{}

		t.Run("should check if user credential are right when grant type is password", func(t *testing.T) {
			clients.On("FindBy", "aClientId").Return(domain.Client{ClientId: "aClientId", ClientSecret: "aClientSecret"})
			password.On("Compare", "hashPassword", []byte("admin")).Return(true)
			users.On("FindBy", "admin").Return(domain.User{Username: "admin", Password: "hashPassword"})
			auth.On("UserAccessToken", "admin").Return(domain.TokenResponse{AccessToken: "anAccessToken", State: "aState", TokenType: "aTokenType", ExpiresIn: "expired"})

			var router = gin.Default()
			PostTokenApi(router, auth, users, password, clients, codes)

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
	})
}

func TestCredentialsGranType(t *testing.T) {
	var password = mock.PasswordMock{}
	var users = mock.UsersMock{}
	var auth = mock.AuthMock{}
	var clients = mock.ClientsMock{}
	var codes = mock.CodesMock{}

	t.Run("", func(t *testing.T) {
		auth.On("AccessToken").Return(domain.TokenResponse{AccessToken: "anAccessToken", State: "aState", TokenType: "aTokenType", ExpiresIn: "expired"})
		clients.On("FindBy", "aClientId").Return(domain.Client{ClientId: "aClientId", ClientSecret: "aClientSecret"})

		var router = gin.Default()
		PostTokenApi(router, auth, users, password, clients, codes)

		body, _ := json.Marshal(TokenRequest{GrantType: "credentials", ClientId: "aClientId", ClientSecret: "aClientSecret"})
		request, _ := http.NewRequest("POST", "/token", bytes.NewBuffer(body))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		clients.AssertExpectations(t)
		auth.AssertExpectations(t)

		res, _ := json.Marshal(domain.TokenResponse{AccessToken: "anAccessToken", State: "aState", TokenType: "aTokenType", ExpiresIn: "expired"})
		assert.Contains(t, response.Body.String(), string(res))
	})
}

func TestCodeGrantType(t *testing.T) {
	var password = mock.PasswordMock{}
	var users = mock.UsersMock{}
	var auth = mock.AuthMock{}
	var clients = mock.ClientsMock{}
	var codes = mock.CodesMock{}

	t.Run("should check if authorization code is right when grant type is code", func(t *testing.T) {
		auth.On("AccessToken").Return(domain.TokenResponse{AccessToken: "anAccessToken", State: "aState", TokenType: "aTokenType", ExpiresIn: "expired"})
		codes.On("FindBy", "aClientId").Return("aCode")
		codes.On("DeleteBy", "aClientId").Return(true)

		var router = gin.Default()
		PostTokenApi(router, auth, users, password, clients, codes)

		body, _ := json.Marshal(TokenRequest{GrantType: "code", Code: "aCode", ClientId: "aClientId", ClientSecret: "aClientSecret"})
		request, _ := http.NewRequest("POST", "/token", bytes.NewBuffer(body))

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		auth.AssertExpectations(t)
		codes.AssertExpectations(t)

		res, _ := json.Marshal(domain.TokenResponse{AccessToken: "anAccessToken", State: "aState", TokenType: "aTokenType", ExpiresIn: "expired"})
		assert.Contains(t, response.Body.String(), string(res))
	})
}
