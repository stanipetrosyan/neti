package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"neti/mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPutUserRole(t *testing.T) {

	users := mock.UsersMock{}
	users.On("AddRole", "aUsername", "aRole").Return(true)

	body := []byte(`{"role": "aRole"}`)
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("PUT", "/users/aUsername", bytes.NewBuffer(body))

	router := gin.Default()
	PutUserRoleApi(users, router)
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	users.AssertExpectations(t)
}
