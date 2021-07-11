package handlers

import "github.com/gin-gonic/gin"

func LoginApi(context *gin.Context) {
	var form UserForm
	context.ShouldBind(&form)

	/* 	println(form.Username)
	   	println(form.Password)

	   	context.Request.ParseMultipartForm(1000)
	   	println("cu")
	   	println(context.GetPostForm("username"))
	   	context.Request.PostForm.Add("username", "admin")
	   	username := context.PostForm("username")
	   	password := context.PostForm("password")
	   	println("username")
	   	println(context.Request.PostForm.Get("username")) */
	if form.Username == "admin" && form.Password == "admin" {
		context.JSON(200, nil)
	} else {
		context.JSON(403, nil)
	}
}

type UserForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
