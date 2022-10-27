package handler

import (
	repository "neti/internals/repository"
	services "neti/internals/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type TokenRequest struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Scope        string `json:"scope"`
}

func PostTokenApi(auth services.Auth, users repository.Users, password services.Password, clients repository.Clients, codes repository.Codes) gin.HandlerFunc {
	return func(context *gin.Context) {
		var request TokenRequest
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
			_, userPassword := users.FindBy(request.Username)
			if password.Compare(userPassword, []byte(request.Password)) {
				response := auth.AccessToken()
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
	}
}
