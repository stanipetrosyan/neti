package handlers

import "github.com/gin-gonic/gin"

func LoginApi() gin.HandlerFunc {
	return func(context *gin.Context) {
		var form UserForm
		context.ShouldBind(&form)
		if form.Username == "admin" && form.Password == "admin" {
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
