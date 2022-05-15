package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"neti/internals/repository"
	"neti/mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthorizeApi(t *testing.T) {

	clientMock := mock.ClientsMock{}
	codesMock := mock.CodesMock{}
	clientMock.On("Exist", "aClientId").Return(true)
	codesMock.On("Add", repository.AuthorizationCode{ClientId: "aClientId", Code: "aCode"}).Return(true)

	body, _ := json.Marshal(AuthorizeRequest{ResponseType: "code", ClientId: "aClientId"})

	request, _ := http.NewRequest("GET", "/authorize", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	var router = gin.Default()
	router.GET("/authorize", GetAuthorizeApi(clientMock, codesMock))
	router.ServeHTTP(response, request)

	res, _ := json.Marshal(AuthorizeResponse{Code: "aCode"})
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, string(res), response.Body.String())
	clientMock.AssertExpectations(t)
	codesMock.AssertExpectations(t)
}
