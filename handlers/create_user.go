package handlers

import "github.com/gin-gonic/gin"

func PostCreateUser(context *gin.Context) {
	//u.Add("User", "user")
	context.JSON(200, "nil")
}

type Users interface {
	Add(username string, password string) bool
}
