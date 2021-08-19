package mocks

import "github.com/stretchr/testify/mock"

type ClientsMock struct {
	mock.Mock
}

func (m ClientsMock) Add(client string) bool {
	args := m.Called(client)
	return args.Bool(0)
}
