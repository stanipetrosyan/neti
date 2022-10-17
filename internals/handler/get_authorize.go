package handler

import (
	"neti/internals/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
			code := uuid.New().String()

			codes.Add(repository.AuthorizationCode{ClientId: request.ClientId, Code: code})
			context.JSON(200, AuthorizeResponse{Code: "aCode"})
		}
	}
}
