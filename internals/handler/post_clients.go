package handler

import (
	"neti/internals/domain"
	repository "neti/internals/repository"

	"github.com/gin-gonic/gin"
)

func PostClientsApi(clients repository.Clients) gin.HandlerFunc {
	return func(context *gin.Context) {
		var client domain.Client
		context.BindJSON(&client)

		if clients.Add(client) {
			context.JSON(200, nil)
		}
	}
}
