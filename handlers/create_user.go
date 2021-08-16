package handlers

import (
	"log"
	"neti/pkg/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PostCreateUser(users db.Users) gin.HandlerFunc {

	// TODO use real value
	return func(context *gin.Context) {
		hash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.MinCost)
		if err != nil {
			log.Fatal("Error to salt password")
		}

		if users.Add(db.User{Username: "admin", Password: string(hash)}) {
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
