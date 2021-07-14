package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type usersMock struct {
	mock.Mock
}

func (m usersMock) Add(username string, password string) bool {
	args := m.Called(username, password)
	return args.Bool(0)
}

func (m usersMock) FindUserBy(username string) (string, string) {
	args := m.Called(username)
	return args.String(0), args.String(1)
}

var router *gin.Engine

func TestCreateUser(t *testing.T) {
	t.Run("should create a new user", func(t *testing.T) {
		users := usersMock{}
		users.On("Add", "admin", "admin").Return(true)

		gin.SetMode(gin.TestMode)
		router = gin.Default()
		router.POST("/users", PostCreateUser(users))

		request, _ := http.NewRequest("POST", "/users", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)
		users.MethodCalled("Add", "admin", "admin")

		users.AssertNumberOfCalls(t, "Add", 1)
		assert.Equal(t, http.StatusOK, response.Code)
	})
}
