package services

import "neti/internals/domain"

type Auth interface {
	AccessToken() domain.TokenResponse
}

type AuthService struct{}

func (s *AuthService) AccessToken() domain.TokenResponse {
	return domain.TokenResponse{}
}
