package handler

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthApi(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	html := template.Must(template.ParseFiles("../../templates/login.html"))
	router.SetHTMLTemplate(html)
	GetAuthApi(router)

	t.Run("should return 200", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/auth", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Contains(t, response.Body.String(), "<title>Login page</title>")
	})
}
