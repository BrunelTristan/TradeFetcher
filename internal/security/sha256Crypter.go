package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"tradeFetcher/internal/externalTools"
)

type Sha256Crypter struct{}

func NewSha256Crypter() externalTools.ICrypter {
	return &Sha256Crypter{}
}

func (c *Sha256Crypter) Crypt(message string, key string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	return mac.Sum(nil)
}
