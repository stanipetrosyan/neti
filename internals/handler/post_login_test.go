package handler

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"neti/internals/domain"
	"neti/mock"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginApi(t *testing.T) {
	t.Run("should return 200", func(t *testing.T) {
		password := mock.PasswordMock{}
		password.On("Compare", "hashPassword", []byte("admin")).Return(true)

		users := mock.UsersMock{}
		users.On("FindBy", "admin").Return(domain.User{Username: "admin", Password: "hashPassword"})

		var router = gin.Default()
		PostLoginApi(router, users, password)

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		fw, _ := writer.CreateFormField("username")
		io.Copy(fw, strings.NewReader("admin"))
		fw, _ = writer.CreateFormField("password")
		io.Copy(fw, strings.NewReader("admin"))

		writer.Close()

		request, _ := http.NewRequest("POST", "/login", bytes.NewReader(body.Bytes()))
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		users.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, response.Code)
	})

}

func TestUnauthorizedLogin(t *testing.T) {
	t.Run("should return 403 if login credentials are wrong", func(t *testing.T) {
		password := mock.PasswordMock{}
		password.On("Compare", "hashPassword", []byte("wrong")).Return(false)

		users := mock.UsersMock{}
		users.On("FindBy", "admin").Return(domain.User{Username: "admin", Password: "hashPassword"})

		var router = gin.Default()
		PostLoginApi(router, users, password)

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		fw, _ := writer.CreateFormField("username")
		io.Copy(fw, strings.NewReader("admin"))
		fw, _ = writer.CreateFormField("password")
		io.Copy(fw, strings.NewReader("wrong"))

		writer.Close()

		request, _ := http.NewRequest("POST", "/login", bytes.NewReader(body.Bytes()))
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		users.AssertExpectations(t)
		assert.Equal(t, http.StatusForbidden, response.Code)
	})
}
