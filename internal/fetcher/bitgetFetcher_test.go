package fetcher

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"tradeFetcher/internal/error"
	"tradeFetcher/internal/generatedMocks"
	"tradeFetcher/model/trading"
)

func TestNewBitgetFetcher(t *testing.T) {
	fakeObject := NewBitgetFetcher(nil)

	assert.NotNil(t, fakeObject)
}

func TestBitgetFetcherFetchLastTradesWithErrorOnSpotFetcher(t *testing.T) {
	mockController := gomock.NewController(t)

	spotFetcherMock := generatedMocks.NewMockIFetcher(mockController)

	spotFetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(1).
		Return(nil, &error.RestApiError{HttpCode: 404})

	fakeObject := NewBitgetFetcher(spotFetcherMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Nil(t, trades)
}

func TestBitgetFetcherFetchLastTrades(t *testing.T) {
	mockController := gomock.NewController(t)

	spotFetcherMock := generatedMocks.NewMockIFetcher(mockController)

	tradesCount := 7

	spotFetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(1).
		Return(make([]trading.Trade, tradesCount), nil)

	fakeObject := NewBitgetFetcher(spotFetcherMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotNil(t, trades)
	assert.NotEmpty(t, trades)

	assert.Equal(t, tradesCount, len(trades))
}
