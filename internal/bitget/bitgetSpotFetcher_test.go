package bitget

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"tradeFetcher/internal/error"
	"tradeFetcher/internal/generatedMocks"
)

func TestNewBitgetSpotFetcher(t *testing.T) {
	fakeObject := NewBitgetSpotFetcher(nil)

	assert.NotNil(t, fakeObject)
}

func TestBitgetSpotFetcherFetchLastTradesWithGetError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIGetter(mockController)

	externalGetterMock.
		EXPECT().
		Get(gomock.Any()).
		Times(1).
		Return(nil, &error.RestApiError{HttpCode: 500})

	fakeObject := NewBitgetSpotFetcher(externalGetterMock)

	assert.NotNil(t, fakeObject)

	_, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.RestApiError)))
}
