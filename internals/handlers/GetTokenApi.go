package handlers

import (
	services "neti/internals/services"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	GrantType string `json:"grant_type"`
	ClientId  string `json:"client_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func GetTokenApi(auth services.Auth) gin.HandlerFunc {

	return func(context *gin.Context) {

		var json TokenRequest

		context.BindJSON(&json)

		response := auth.AccessToken()
		context.JSON(200, response)
	}
}
