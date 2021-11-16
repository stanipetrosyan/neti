package mocks

import (
	"neti/internals/domain"

	"github.com/stretchr/testify/mock"
)

type UsersMock struct {
	mock.Mock
}

func (m UsersMock) Add(user domain.User) bool {
	args := m.Called(user)
	return args.Bool(0)
}

func (m UsersMock) FindBy(username string) (string, string) {
	args := m.Called(username)
	return args.String(0), args.String(1)
}
