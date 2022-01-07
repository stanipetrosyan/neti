package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var password = CryptoPassword{}

func TestPassword(t *testing.T) {
	t.Run("should compare plained pwd with salted pwd", func(t *testing.T) {
		hash := password.Salt("password")

		assert.True(t, password.Compare(hash, []byte("password")))
	})
}
