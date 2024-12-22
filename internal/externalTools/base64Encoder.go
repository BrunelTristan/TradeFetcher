package externalTools

import (
	"encoding/base64"
)

type Base64Encoder struct {
}

func NewBase64Encoder() IEncoder {
	return &Base64Encoder{}
}

func (e *Base64Encoder) Encode(message []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(message))
}
