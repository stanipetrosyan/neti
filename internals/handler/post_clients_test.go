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

func TestPostClientsApi(t *testing.T) {
	t.Run("should create a new client", func(t *testing.T) {

		clients := mock.ClientsMock{}
		clients.On("Add", domain.Client{ClientId: "aClientId", ClientSecret: "aClientSecret"}).Return(true)

		gin.SetMode(gin.TestMode)
		var router = gin.Default()
		router.POST("/clients", PostClientsApi(clients))

		body := []byte(`{"client_id": "aClientId", "client_secret": "aClientSecret"}`)

		req, _ := http.NewRequest("POST", "/clients", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		clients.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, res.Code)
	})
}