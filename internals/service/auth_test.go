package service

import (
	"neti/internals/domain"
	"neti/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessToken(t *testing.T) {
	users := mock.UsersMock{}
	users.On("FindBy", "aUser").Return(domain.User{Username: "aUser", Password: "aPass", Role: "aRole"})
	auth := AuthService{users: users}

	actual := auth.UserAccessToken("aUser")

	users.AssertExpectations(t)
	assert.NotEqual(t, "", actual)
}
