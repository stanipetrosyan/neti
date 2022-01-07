package handler

import (
	"neti/internals/repository"
	services "neti/internals/service"

	"github.com/gin-gonic/gin"
)

func PostLoginApi(users repository.Users, password services.Password) gin.HandlerFunc {
	return func(context *gin.Context) {
		var form UserForm
		context.ShouldBind(&form)
		_, userPassword := users.FindBy(form.Username)

		if password.Compare(userPassword, []byte(form.Password)) {
			context.JSON(200, nil)
		} else {
			context.JSON(403, nil)
		}
	}
}

type UserForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
