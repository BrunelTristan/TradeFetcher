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
	"tradeFetcher/internal/testingTools"
	bitgetModel "tradeFetcher/model/bitget"
	"tradeFetcher/model/error"
)

func TestNewApiQuery(t *testing.T) {
	fakeObject := NewApiQuery(nil, nil)

	assert.NotNil(t, fakeObject)
}

func TestCallApiQueryWithNilParameters(t *testing.T) {
	api := NewApiQuery(nil, nil)

	output, err := api.Get(nil)

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Equal(t, 999, err.(*error.RestApiError).HttpCode)
}

func TestCallApiQueryWithUnsupportedChars(t *testing.T) {
	api := NewApiQuery(nil, nil)
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

	api := NewApiQuery(accountCfg, signatureBuilderMock)
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

	expectedRoute := strconv.FormatInt(int64(time.Now().UnixNano()/1000000), 10) + "GET/api/v2/public/time"

	signatureBuilderMock.
		EXPECT().
		Sign(testingTools.NewByteSliceMatcherWithException([]byte(expectedRoute), []int{7, 8, 9, 10, 11, 12})).
		Times(1)

	api := NewApiQuery(accountCfg, signatureBuilderMock)

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

	expectedRoute := strconv.FormatInt(int64(time.Now().UnixNano()/1000000), 10) + "GET/api/v2/public/time?param1=yesterday"

	signatureBuilderMock.
		EXPECT().
		Sign(testingTools.NewByteSliceMatcherWithException([]byte(expectedRoute), []int{7, 8, 9, 10, 11, 12})).
		Times(1)

	api := NewApiQuery(accountCfg, signatureBuilderMock)

	output, err := api.Get(parameters)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.NotEmpty(t, output)
}

// Valid data from bitget are tested in Integration tests
