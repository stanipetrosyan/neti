package main

import (
	"neti/handlers"
	"neti/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	users := db.PostgresUsers{}
	var router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/auth", handlers.AuthApi)
	router.POST("/login", handlers.LoginApi)
	router.POST("/users", handlers.PostCreateUser(&users))

	router.Run()
}
