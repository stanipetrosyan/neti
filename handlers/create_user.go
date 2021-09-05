package handlers

import (
	"encoding/json"
	"neti/pkg/crypto"
	"neti/pkg/db"

	"github.com/gin-gonic/gin"
)

func PostCreateUser(users db.Users, password crypto.Password) gin.HandlerFunc {

	return func(context *gin.Context) {
		body := json.NewDecoder(context.Request.Body)
		var user db.User
		body.Decode(&user)

		hash := password.Salt(user.Password)

		if users.Add(db.User{Username: user.Username, Password: string(hash)}) {
			context.JSON(200, "nil")
		}
	}
}
