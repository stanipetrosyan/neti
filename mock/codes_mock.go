package mock

import (
	"neti/internals/repository"

	"github.com/stretchr/testify/mock"
)

type CodesMock struct {
	mock.Mock
}

func (m CodesMock) Add(codes repository.AuthorizationCode) bool {
	args := m.Called(codes)
	return args.Bool(0)
}

func (m CodesMock) FindBy(clientId string) string {
	args := m.Called(clientId)
	return args.String(0)
}

func (m CodesMock) DeleteBy(clientId string) bool {
	args := m.Called(clientId)
	return args.Bool(0)
}
