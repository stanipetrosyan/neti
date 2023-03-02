package service

import (
	"crypto/rand"
	"encoding/hex"
)

type Secret interface {
	ClientSecret() string
}

type CryptoSecret struct{}

func (s *CryptoSecret) ClientSecret() string {
	random, _ := rand.Prime(rand.Reader, 256)
	return hex.EncodeToString(random.Bytes())
}
