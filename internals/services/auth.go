package services

import (
	"fmt"
	"neti/internals/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Auth interface {
	AccessToken() domain.TokenResponse
}

type AuthService struct{}

func (s *AuthService) AccessToken() domain.TokenResponse {
	expiresIn := time.Now().Add(time.Minute * 15).Unix()
	claims := jwt.MapClaims{}
	claims["sub"] = "aSubject"
	claims["exp"] = expiresIn
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("ACCESS_SECRET"))
	if err != nil {
		return domain.TokenResponse{}
	}

	return domain.TokenResponse{AccessToken: signedToken, State: "valid", TokenType: "jwt", ExpiresIn: fmt.Sprint(expiresIn)}
}
