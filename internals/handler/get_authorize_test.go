package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"neti/mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthorizeApi(t *testing.T) {

	clientMock := mock.ClientsMock{}
	clientMock.On("Exist", "aClientId").Return(true)

	body, _ := json.Marshal(AuthorizeRequest{ResponseType: "code", ClientId: "aClientId"})

	request, _ := http.NewRequest("GET", "/authorize", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	var router = gin.Default()
	router.GET("/authorize", GetAuthorizeApi(clientMock))
	router.ServeHTTP(response, request)

	res, _ := json.Marshal(AuthorizeResponse{Code: "aCode"})
	clientMock.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, string(res), response.Body.String())
}
