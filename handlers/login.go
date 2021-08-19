package handlers

import (
	"log"
	"neti/pkg/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginApi(users db.Users) gin.HandlerFunc {
	return func(context *gin.Context) {
		var form UserForm
		context.ShouldBind(&form)
		_, password := users.FindBy(form.Username)

		if comparePasswords(password, []byte(form.Password)) {
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

// Create interface to manage password for testing purpose
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
