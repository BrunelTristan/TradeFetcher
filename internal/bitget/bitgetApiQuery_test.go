package bitget

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/url"
	"strconv"
	"testing"
	"time"
	"tradeFetcher/internal/generatedMocks"
	bitgetModel "tradeFetcher/model/bitget"
	"tradeFetcher/model/error"
)

func TestNewBitgetApiQuery(t *testing.T) {
	fakeObject := NewBitgetApiQuery(nil, nil)

	assert.NotNil(t, fakeObject)
}

func TestCallApiQueryWithNilParameters(t *testing.T) {
	api := NewBitgetApiQuery(nil, nil)

	output, err := api.Get(nil)

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Equal(t, 999, err.(*error.RestApiError).HttpCode)
}

func TestCallApiQueryWithUnsupportedChars(t *testing.T) {
	api := NewBitgetApiQuery(nil, nil)
	parameters := &bitgetModel.ApiQueryParameters{
		Route: "@^\\``||[{#~/public/time",
	}

	output, err := api.Get(parameters)

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*url.Error)))
}

func TestCallApiQueryWithUnkwownRoute(t *testing.T) {
	mockController := gomock.NewController(t)

	signatureBuilderMock := generatedMocks.NewMockISignatureBuilder(mockController)

	signatureBuilderMock.
		EXPECT().
		Sign(gomock.Any())

	accountCfg := &bitgetModel.AccountConfiguration{
		ApiKey:     "key",
		PassPhrase: "phrase",
		SecretKey:  "secret",
	}

	api := NewBitgetApiQuery(accountCfg, signatureBuilderMock)
	parameters := &bitgetModel.ApiQueryParameters{
		Route: ".apis/vXXXX/public/time",
	}

	output, err := api.Get(parameters)

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*url.Error)))
}

func TestCallApiQueryWithoutErrorWithoutQueryStringWithoutBody(t *testing.T) {
	mockController := gomock.NewController(t)

	signatureBuilderMock := generatedMocks.NewMockISignatureBuilder(mockController)

	accountCfg := &bitgetModel.AccountConfiguration{
		ApiKey:     "key",
		PassPhrase: "phrase",
		SecretKey:  "secret",
	}
	parameters := &bitgetModel.ApiQueryParameters{
		Route: "/api/v2/public/time",
	}

	signatureBuilderMock.
		EXPECT().
		Sign(gomock.Eq([]byte(strconv.FormatInt(int64(time.Now().UnixNano()/1000000), 10) + "GET/api/v2/public/time"))).
		Times(1)

	api := NewBitgetApiQuery(accountCfg, signatureBuilderMock)

	output, err := api.Get(parameters)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.NotEmpty(t, output)
}

func TestCallApiQueryWithoutErrorWithQueryStringWithoutBody(t *testing.T) {
	mockController := gomock.NewController(t)

	signatureBuilderMock := generatedMocks.NewMockISignatureBuilder(mockController)

	accountCfg := &bitgetModel.AccountConfiguration{
		ApiKey:     "key",
		PassPhrase: "phrase",
		SecretKey:  "secret",
	}
	parameters := &bitgetModel.ApiQueryParameters{
		Route: "/api/v2/public/time?param1=yesterday",
	}

	signatureBuilderMock.
		EXPECT().
		Sign(gomock.Eq([]byte(strconv.FormatInt(int64(time.Now().UnixNano()/1000000), 10) + "GET/api/v2/public/time?param1=yesterday"))).
		Times(1)

	api := NewBitgetApiQuery(accountCfg, signatureBuilderMock)

	output, err := api.Get(parameters)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.NotEmpty(t, output)
}

// Valid data from bitget are tested in Integration tests
