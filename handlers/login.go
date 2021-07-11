package handlers

import "github.com/gin-gonic/gin"

func LoginApi(context *gin.Context) {
	context.Request.ParseMultipartForm(1000)
	println("cu")
	println(context.GetPostForm("username"))
	context.Request.PostForm.Add("username", "admin")
	username := context.PostForm("username")
	password := context.PostForm("password")
	println("username")
	println(context.Request.PostForm.Get("username"))
	if username == "admin" && password == "admin" {
		context.JSON(200, nil)
	} else {
		context.JSON(403, nil)
	}
}
