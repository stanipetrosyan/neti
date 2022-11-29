package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthApi(router *gin.Engine) {
	router.GET("/auth", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", nil)
	})
}
