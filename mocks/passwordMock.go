package mocks

import "github.com/stretchr/testify/mock"

type PasswordMock struct {
	mock.Mock
}

func (m PasswordMock) Salt(password string) string {
	args := m.Called(password)
	return args.String(0)
}

func (m PasswordMock) Compare(hashedPwd string, plainPwd []byte) bool {
	args := m.Called(hashedPwd, plainPwd)
	return args.Bool(0)
}
