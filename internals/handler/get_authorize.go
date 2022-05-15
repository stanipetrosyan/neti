package handler

import (
	"neti/internals/repository"

	"github.com/gin-gonic/gin"
)

type AuthorizeRequest struct {
	ResponseType string `json:"response_type"`
	ClientId     string `json:"client_id"`
}

type AuthorizeResponse struct {
	Code string `json:"code"`
}

func GetAuthorizeApi(clients repository.Clients, codes repository.Codes) gin.HandlerFunc {

	return func(context *gin.Context) {
		var request AuthorizeRequest
		context.BindJSON(&request)

		if clients.Exist(request.ClientId) {
			codes.Add(repository.AuthorizationCode{ClientId: request.ClientId, Code: "aCode"})
			context.JSON(200, AuthorizeResponse{Code: "aCode"})
		}
	}
}
