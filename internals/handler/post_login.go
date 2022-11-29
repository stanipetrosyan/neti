package handler

import (
	"neti/internals/repository"
	services "neti/internals/service"

	"github.com/gin-gonic/gin"
)

func PostLoginApi(router *gin.Engine, users repository.Users, password services.Password) {
	router.POST("/login", func(context *gin.Context) {
		var form UserForm
		context.ShouldBind(&form)
		user := users.FindBy(form.Username)

		if password.Compare(user.Password, []byte(form.Password)) {
			context.JSON(200, nil)
		} else {
			context.JSON(403, nil)
		}
	})
}

type UserForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
