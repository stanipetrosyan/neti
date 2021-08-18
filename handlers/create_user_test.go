package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"neti/mocks"
	"neti/pkg/db"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var router *gin.Engine

func TestCreateUser(t *testing.T) {
	t.Run("should create a new user", func(t *testing.T) {
		var body = []byte(`{
			"username": "user",
			"password": "pass"
		}`)

		users := mocks.UsersMock{}
		users.On("Add", mock.IsType(db.User{})).Return(true)

		gin.SetMode(gin.TestMode)
		router = gin.Default()
		router.POST("/users", PostCreateUser(users))

		request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		request.Header.Set("content-type", "application/json")
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		users.AssertExpectations(t)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}
