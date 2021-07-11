package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthApi(t *testing.T) {
	t.Run("should return 200", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/auth", nil)
		response := httptest.NewRecorder()
		router := tupRouter()
		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Contains(t, response.Body.String(), "<title>Login page</title>")
	})
}

func tupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	// da capire bene la questione del path
	html := template.Must(template.ParseFiles("../templates/login.html"))
	router.SetHTMLTemplate(html)
	router.POST("/auth", AuthApi)

	return router
}
