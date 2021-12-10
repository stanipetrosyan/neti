package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessToken(t *testing.T) {
	auth := AuthService{}
	actual := auth.AccessToken()

	assert.NotEqual(t, "", actual)
}
