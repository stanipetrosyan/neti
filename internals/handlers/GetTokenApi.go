package handlers

import (
	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	GrantType string `json:"grant_type"`
	ClientId  string `json:"client_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	State       string `json:"state"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
}

func GetTokenApi() gin.HandlerFunc {

	return func(context *gin.Context) {

		var json TokenRequest

		context.BindJSON(&json)

		response := TokenResponse{AccessToken: "anAccessToken", State: "aState", TokenType: "aTokenType", ExpiresIn: "expired"}
		context.JSON(200, response)
	}
}
