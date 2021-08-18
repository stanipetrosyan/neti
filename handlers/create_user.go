package handlers

import (
	"encoding/json"
	"log"
	"neti/pkg/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PostCreateUser(users db.Users) gin.HandlerFunc {

	return func(context *gin.Context) {
		body := json.NewDecoder(context.Request.Body)
		var user db.User
		body.Decode(&user)

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		if err != nil {
			log.Fatal("Error to salt password")
		}

		if users.Add(db.User{Username: user.Username, Password: string(hash)}) {
			context.JSON(200, "nil")
		}
	}
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
