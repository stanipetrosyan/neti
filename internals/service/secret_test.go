package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientSecret(t *testing.T) {
	secret := CryptoSecret{}

	var actual = secret.ClientSecret()

	assert.Len(t, actual, 64)
}
