package bitget

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"tradeFetcher/internal/generatedMocks"
	bitgetModel "tradeFetcher/model/bitget"
)

func TestNewBitgetApiSignatureBuilder(t *testing.T) {
	fakeObject := NewBitgetApiSignatureBuilder(nil, nil, nil)

	assert.NotNil(t, fakeObject)
}

func TestSignMessage(t *testing.T) {
	mockController := gomock.NewController(t)

	message := "Simple message to sign"
	expectedSignature := []byte("Signed and encoded Message")
	accountCfg := &bitgetModel.AccountConfiguration{
		ApiKey:     "key",
		PassPhrase: "phrase",
		SecretKey:  "secret",
	}

	crypterMock := generatedMocks.NewMockICrypter(mockController)
	encoderMock := generatedMocks.NewMockIEncoder(mockController)

	crypterMock.EXPECT().
		Crypt(gomock.Eq(message), "secret").
		Times(1).
		Return([]byte("Signed"))

	encoderMock.EXPECT().
		Encode(gomock.Eq([]byte("Signed"))).
		Times(1).
		Return(expectedSignature)

	builder := NewBitgetApiSignatureBuilder(accountCfg, crypterMock, encoderMock)

	signature := builder.Sign([]byte(message))

	assert.NotEmpty(t, signature)
	assert.Equal(t, expectedSignature, signature)
}
