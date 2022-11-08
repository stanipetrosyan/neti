package mock

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

func (m UsersMock) FindBy(username string) domain.User {
	args := m.Called(username)
	return args.Get(0).(domain.User)
}

func (m UsersMock) AddRole(username string, role string) bool {
	args := m.Called(username, role)
	return args.Bool(0)
}
