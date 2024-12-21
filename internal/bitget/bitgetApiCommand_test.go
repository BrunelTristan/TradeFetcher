package bitget

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/url"
	"testing"
	"tradeFetcher/internal/generatedMocks"
	bitgetModel "tradeFetcher/model/bitget"
	"tradeFetcher/model/error"
)

func TestNewBitgetApiCommand(t *testing.T) {
	fakeObject := NewBitgetApiCommand(nil, nil)

	assert.NotNil(t, fakeObject)
}

func TestCallApiCommandWithNilParameters(t *testing.T) {
	api := NewBitgetApiCommand(nil, nil)

	output, err := api.Get(nil)

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Equal(t, 999, err.(*error.RestApiError).HttpCode)
}

func TestCallApiCommandWithUnsupportedChars(t *testing.T) {
	api := NewBitgetApiCommand(nil, nil)
	parameters := &bitgetModel.ApiCommandParameters{
		Route: "@^\\``||[{#~/public/time",
	}

	output, err := api.Get(parameters)

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*url.Error)))
}

func TestCallApiCommandWithUnkwownRoute(t *testing.T) {
	mockController := gomock.NewController(t)

	signatureBuilderMock := generatedMocks.NewMockISignatureBuilder(mockController)

	signatureBuilderMock.
		EXPECT().
		Sign(gomock.Any())

	api := NewBitgetApiCommand(nil, signatureBuilderMock)
	parameters := &bitgetModel.ApiCommandParameters{
		Route: ".apis/vXXXX/public/time",
	}

	output, err := api.Get(parameters)

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*url.Error)))
}

func TestCallApiCommandWithoutErrorWithoutQueryStringWithoutBody(t *testing.T) {
	mockController := gomock.NewController(t)

	signatureBuilderMock := generatedMocks.NewMockISignatureBuilder(mockController)

	api := NewBitgetApiCommand(nil, signatureBuilderMock)
	parameters := &bitgetModel.ApiCommandParameters{
		Route: "/api/v2/public/time",
	}

	signatureBuilderMock.
		EXPECT().
		Sign(gomock.Eq([]byte("GET/api/v2/public/time"))).
		Times(1)

	output, err := api.Get(parameters)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.NotEmpty(t, output)
}

func TestCallApiCommandWithoutErrorWithQueryStringWithoutBody(t *testing.T) {
	mockController := gomock.NewController(t)

	signatureBuilderMock := generatedMocks.NewMockISignatureBuilder(mockController)

	api := NewBitgetApiCommand(nil, signatureBuilderMock)
	parameters := &bitgetModel.ApiCommandParameters{
		Route: "/api/v2/public/time?param1=yesterday",
	}

	signatureBuilderMock.
		EXPECT().
		Sign(gomock.Eq([]byte("GET/api/v2/public/time?param1=yesterday"))).
		Times(1)

	output, err := api.Get(parameters)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.NotEmpty(t, output)
}

// Valid data from bitget are tested in Integration tests
