package handlers

import (
	"neti/internals/repositories"

	"github.com/gin-gonic/gin"
)

func PostClientsApi(clients repositories.Clients) gin.HandlerFunc {
	return func(context *gin.Context) {
		if clients.Add("aClient") {
			context.JSON(200, nil)
		}
	}
}
