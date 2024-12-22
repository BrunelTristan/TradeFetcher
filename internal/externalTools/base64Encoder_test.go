package externalTools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBase64Encoder(t *testing.T) {
	object := NewBase64Encoder()

	assert.NotNil(t, object)
}

func TestEncode(t *testing.T) {
	encoder := NewBase64Encoder()

	encoded := encoder.Encode([]byte("Th@s \nis a long and Compl€x \"text' to (encode) in &&~ b@se§%£¤\\ ^^"))

	assert.NotNil(t, encoded)
	assert.Equal(t, []byte("VGhAcyAKaXMgYSBsb25nIGFuZCBDb21wbOKCrHggInRleHQnIHRvIChlbmNvZGUpIGluICYmfiBiQHNlwqclwqPCpFwgXl4="), encoded)
}
