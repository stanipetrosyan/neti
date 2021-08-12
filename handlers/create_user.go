package handlers

import (
	"neti/pkg/db"

	"github.com/gin-gonic/gin"
)

func PostCreateUser(users db.Users) gin.HandlerFunc {

	return func(context *gin.Context) {
		if users.Add(db.User{Username: "admin", Password: "admin"}) {
			context.JSON(200, "nil")
		}
	}
}
