package handlers

import (
	"neti/internals/repositories"
	services "neti/internals/services"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	GrantType string `json:"grant_type"`
	ClientId  string `json:"client_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func GetTokenApi(auth services.Auth, users repositories.Users, password services.Password) gin.HandlerFunc {
	return func(context *gin.Context) {
		var request TokenRequest
		context.BindJSON(&request)
		
		_, userPassword := users.FindBy(request.Username)
		if password.Compare(userPassword, []byte(request.Password)) {
			response := auth.AccessToken()
			context.JSON(200, response)
		} else {
			context.JSON(403, nil)
		}
		
	}
}
