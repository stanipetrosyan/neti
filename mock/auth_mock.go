package mock

import (
	"neti/internals/domain"

	"github.com/stretchr/testify/mock"
)

type AuthMock struct {
	mock.Mock
}

func (m AuthMock) AccessToken() domain.TokenResponse {
	args := m.Called()
	return args.Get(0).(domain.TokenResponse)
}

func (m AuthMock) UserAccessToken(username string) domain.TokenResponse {
	args := m.Called(username)
	return args.Get(0).(domain.TokenResponse)
}
