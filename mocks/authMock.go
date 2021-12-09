package mocks

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
