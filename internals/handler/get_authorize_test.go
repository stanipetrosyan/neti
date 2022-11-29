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
	m "github.com/stretchr/testify/mock"
)

func TestAuthorizeApi(t *testing.T) {

	clientMock := mock.ClientsMock{}
	codesMock := mock.CodesMock{}
	clientMock.On("Exist", "aClientId").Return(true)
	codesMock.On("Add", m.Anything).Return(true)

	body, _ := json.Marshal(AuthorizeRequest{ResponseType: "code", ClientId: "aClientId"})

	request, _ := http.NewRequest("GET", "/authorize", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	var router = gin.Default()
	GetAuthorizeApi(router, clientMock, codesMock)
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	clientMock.AssertExpectations(t)
	codesMock.AssertExpectations(t)
}
