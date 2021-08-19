package handlers

import (
	"neti/pkg/db"

	"github.com/gin-gonic/gin"
)

func PostClientsApi(clients db.Clients) gin.HandlerFunc {
	return func(context *gin.Context) {
		if clients.Add("aClient") {
			context.JSON(200, nil)
		}
	}
}
