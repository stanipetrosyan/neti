package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthApi(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", nil)
}
