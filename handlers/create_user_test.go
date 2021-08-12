package handlers

import (
	"net/http"
	"net/http/httptest"
	"neti/mocks"
	"neti/pkg/db"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestCreateUser(t *testing.T) {
	t.Run("should create a new user", func(t *testing.T) {
		users := mocks.UsersMock{}
		users.On("Add", db.User{Username: "admin", Password: "admin"}).Return(true)

		gin.SetMode(gin.TestMode)
		router = gin.Default()
		router.POST("/users", PostCreateUser(users))

		request, _ := http.NewRequest("POST", "/users", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		users.AssertExpectations(t)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}
