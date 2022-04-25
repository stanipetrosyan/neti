package handler

import (
	repository "neti/internals/repository"
	services "neti/internals/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type TokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Scope        string `json:"scope"`
}

func PostTokenApi(auth services.Auth, users repository.Users, password services.Password, clients repository.Clients) gin.HandlerFunc {
	return func(context *gin.Context) {
		var request TokenRequest
		context.BindJSON(&request)

		if request.GrantType == "credentials" {
			client := clients.FindBy(request.ClientId)
			log.Info("client found with id: ", client.ClientId)

			if client.ClientSecret == request.ClientSecret {
				response := auth.AccessToken()
				context.JSON(200, response)
			} else {
				context.JSON(401, nil)
			}
		} else {
			_, userPassword := users.FindBy(request.Username)
			if password.Compare(userPassword, []byte(request.Password)) {
				response := auth.AccessToken()
				context.JSON(200, response)
			} else {
				context.JSON(401, nil)
			}
		}
	}
}
