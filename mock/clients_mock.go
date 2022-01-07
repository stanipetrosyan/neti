package mock

import "github.com/stretchr/testify/mock"

type ClientsMock struct {
	mock.Mock
}

func (m ClientsMock) Add(client string) bool {
	args := m.Called(client)
	return args.Bool(0)
}

func (m ClientsMock) Find(id string) string {
	args := m.Called(id)
	return args.String(0)
}
