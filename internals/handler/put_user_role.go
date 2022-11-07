package handler

import (
	"net/http"
	"neti/internals/repository"

	"github.com/gin-gonic/gin"
)

type RoleRequest struct {
	Role string `json:"role"`
}

func PutUserRoleApi(users repository.Users, router *gin.Engine) {
	router.PUT("/users/:id", func(context *gin.Context) {
		username := context.Param("id")
		var request RoleRequest
		context.BindJSON(&request)

		if users.AddRole(username, request.Role) {
			context.JSON(http.StatusOK, username)
		} else {
			context.JSON(404, nil)
		}
	})
}
