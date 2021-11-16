package handlers

import (
	"encoding/json"
	"neti/internals/domain"
	"neti/internals/repositories"
	"neti/pkg/crypto"

	"github.com/gin-gonic/gin"
)

func PostCreateUser(users repositories.Users, password crypto.Password) gin.HandlerFunc {

	return func(context *gin.Context) {
		body := json.NewDecoder(context.Request.Body)
		var user domain.User
		body.Decode(&user)

		hash := password.Salt(user.Password)

		if users.Add(domain.User{Username: user.Username, Password: string(hash)}) {
			context.JSON(200, "nil")
		}
	}
}
