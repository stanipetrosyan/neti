package service

import (
	"fmt"
	"neti/internals/domain"
	"neti/internals/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

type Auth interface {
	AccessToken() domain.TokenResponse
	UserAccessToken(username string) domain.TokenResponse
}

type AuthService struct {
	Users repository.Users
}

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

	return domain.TokenResponse{
		AccessToken: signedToken,
		State:       "valid",
		TokenType:   "jwt",
		ExpiresIn:   fmt.Sprint(expiresIn),
	}
}

func (s *AuthService) UserAccessToken(username string) domain.TokenResponse {
	user := s.Users.FindBy(username)

	expiresIn := time.Now().Add(time.Minute * 15).Unix()
	claims := jwt.MapClaims{}
	claims["sub"] = "aSubject"
	claims["exp"] = expiresIn
	claims["roles"] = []string{user.Role}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("ACCESS_SECRET"))
	if err != nil {
		return domain.TokenResponse{}
	}

	return domain.TokenResponse{
		AccessToken: signedToken,
		State:       "valid",
		TokenType:   "jwt",
		ExpiresIn:   fmt.Sprint(expiresIn),
		Roles:       []string{user.Role},
	}
}
