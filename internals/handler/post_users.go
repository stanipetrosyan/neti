package handler

import (
	"encoding/json"
	"neti/internals/domain"
	"neti/internals/repository"
	service "neti/internals/service"

	"github.com/gin-gonic/gin"
)

func PostUsersApi(router *gin.Engine, users repository.Users, password service.Password) {
	router.POST("/users", func(context *gin.Context) {
		body := json.NewDecoder(context.Request.Body)
		var user domain.User
		body.Decode(&user)

		hash := password.Salt(user.Password)

		if users.Add(domain.User{Username: user.Username, Password: string(hash)}) {
			context.JSON(200, "nil")
		}
	})
}
