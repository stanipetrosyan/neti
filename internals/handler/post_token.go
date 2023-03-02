package handler

import (
	"neti/internals/domain"
	repository "neti/internals/repository"
	services "neti/internals/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func PostTokenApi(router *gin.Engine, auth services.Auth, users repository.Users, password services.Password, clients repository.Clients, codes repository.Codes) {
	router.POST("/token", func(context *gin.Context) {
		var request domain.TokenRequest
		context.BindJSON(&request)

		switch request.GrantType {
		case "credentials":
			client := clients.FindBy(request.ClientId)
			log.Info("client found with id: ", client.ClientId)

			if client.ClientSecret == request.ClientSecret {
				response := auth.AccessToken()
				context.JSON(200, response)
			} else {
				context.JSON(401, nil)
			}
		case "password":
			user := users.FindBy(request.Username)
			if password.Compare(user.Password, []byte(request.Password)) {
				response := auth.UserAccessToken(request.Username)
				context.JSON(200, response)
			} else {
				context.JSON(401, nil)
			}
		case "code":
			if codes.FindBy(request.ClientId) == request.Code {
				response := auth.AccessToken()
				codes.DeleteBy(request.ClientId)
				context.JSON(200, response)
			} else {
				context.JSON(401, nil)
			}
		}
	})
}
