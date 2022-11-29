package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"neti/internals/domain"
	"neti/mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestCreateUser(t *testing.T) {
	t.Run("should create a new user", func(t *testing.T) {
		var body = []byte(`{
			"username": "user",
			"password": "pass"
		}`)

		password := mock.PasswordMock{}
		password.On("Salt", "pass").Return("passwordSalt")

		users := mock.UsersMock{}
		users.On("Add", domain.User{Username: "user", Password: "passwordSalt"}).Return(true)

		gin.SetMode(gin.TestMode)
		router = gin.Default()
		PostUsersApi(router, users, password)

		request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		request.Header.Set("content-type", "application/json")
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		users.AssertExpectations(t)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}
