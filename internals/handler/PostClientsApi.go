package handler

import (
	repository "neti/internals/repository"

	"github.com/gin-gonic/gin"
)

func PostClientsApi(clients repository.Clients) gin.HandlerFunc {
	return func(context *gin.Context) {
		if clients.Add("aClient") {
			context.JSON(200, nil)
		}
	}
}
