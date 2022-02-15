package mock

import (
	"neti/internals/domain"

	"github.com/stretchr/testify/mock"
)

type ClientsMock struct {
	mock.Mock
}

func (m ClientsMock) Add(client domain.Client) bool {
	args := m.Called(client)
	return args.Bool(0)
}

func (m ClientsMock) FindBy(id string) domain.Client {
	args := m.Called(id)
	return args.Get(0).(domain.Client)
}
