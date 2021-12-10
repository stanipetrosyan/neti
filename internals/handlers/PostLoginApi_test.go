package handlers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"neti/mocks"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginApi(t *testing.T) {
	t.Run("should return 200", func(t *testing.T) {
		password := mocks.PasswordMock{}
		password.On("Compare", "hashPassword", []byte("admin")).Return(true)

		users := mocks.UsersMock{}
		users.On("FindBy", "admin").Return("admin", "hashPassword")
		router := setupRouter(users, password)

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
		password := mocks.PasswordMock{}
		password.On("Compare", "hashPassword", []byte("wrong")).Return(false)

		users := mocks.UsersMock{}
		users.On("FindBy", "admin").Return("admin", "hashPassword")

		router := setupRouter(users, password)
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

func setupRouter(users mocks.UsersMock, password mocks.PasswordMock) *gin.Engine {
	gin.SetMode(gin.TestMode)
	var router = gin.Default()
	router.POST("/login", PostLoginApi(users, password))
	return router
}
