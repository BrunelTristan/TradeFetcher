package security

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSha256Crypter(t *testing.T) {
	fakeObject := NewSha256Crypter()

	assert.NotNil(t, fakeObject)
}

func TestCryptMessage(t *testing.T) {
	crypter := NewSha256Crypter()

	crypted := crypter.Crypt("Message", "key")
	expected, _ := hex.DecodeString("d3c76ce0cfad906c9264ae29bc7f82363b6ed5e63bd647437bc41aa6e46c1b49")

	assert.Equal(t, expected, crypted)
}
