package main

import (
	"neti/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/auth", handlers.AuthApi)
	router.POST("/login", handlers.LoginApi)
	router.POST("/users", handlers.PostCreateUser)

	router.Run()
}
