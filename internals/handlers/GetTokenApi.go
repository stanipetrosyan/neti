package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetTokenApi() gin.HandlerFunc {

	return func(context *gin.Context) {
		context.JSON(200, "nil")
	}
}
