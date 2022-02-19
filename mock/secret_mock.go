package mock

import "github.com/stretchr/testify/mock"

type SecretMock struct {
	mock.Mock
}

func (m SecretMock) ClientSecret() string {
	args := m.Called()
	return args.String(0)
}
