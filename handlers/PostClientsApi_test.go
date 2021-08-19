package handlers

import (
	"net/http"
	"net/http/httptest"
	"neti/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostClientsApi(t *testing.T) {
	t.Run("should create a new client", func(t *testing.T) {

		clients := mocks.ClientsMock{}
		clients.On("Add", "aClient").Return(true)

		gin.SetMode(gin.TestMode)
		var router = gin.Default()
		router.POST("/clients", PostClientsApi(clients))

		req, _ := http.NewRequest("POST", "/clients", nil)
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		clients.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, res.Code)
	})
}
