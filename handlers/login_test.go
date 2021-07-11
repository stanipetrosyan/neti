package handlers

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginApi(t *testing.T) {
	t.Run("should return 200", func(t *testing.T) {
		var json = []byte(`{"username": "admin"}`)

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, _ := writer.CreatePart(textproto.MIMEHeader{"Content-Type": {"application/json"}})
		part.Write(json)

		/* fw, _ := writer.CreateFormField("username")
		io.Copy(fw, strings.NewReader("admin"))
		fw, _ = writer.CreateFormField("password")
		io.Copy(fw, strings.NewReader("admin")) */
		writer.Close()

		println(body.String())

		request, _ := http.NewRequest("POST", "/login", body)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()
		router := setupRouter()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	var router = gin.Default()
	router.POST("/login", LoginApi)
	return router
}
