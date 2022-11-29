package handler

import (
	"neti/internals/domain"
	repository "neti/internals/repository"
	"neti/internals/service"

	"github.com/gin-gonic/gin"
)

func PostClientsApi(router *gin.Engine, clients repository.Clients, secret service.Secret) {
	router.POST("/clients", func(context *gin.Context) {
		var client domain.Client
		context.BindJSON(&client)

		client.ClientSecret = secret.ClientSecret()

		if clients.Add(client) {
			context.JSON(200, client)
		}
	})
}
