package fetcher

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"tradeFetcher/internal/generatedMocks"
	"tradeFetcher/model/error"
	"tradeFetcher/model/trading"
)

func TestNewBitgetFetcher(t *testing.T) {
	fakeObject := NewBitgetFetcher(nil)

	assert.NotNil(t, fakeObject)
}

func TestBitgetFetcherFetchLastTradesWithErrorOnSpotFetcher(t *testing.T) {
	mockController := gomock.NewController(t)

	spotFetcherMock := generatedMocks.NewMockIFetcher(mockController)
	futureFetcherMock := generatedMocks.NewMockIFetcher(mockController)

	spotFetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(1).
		Return(nil, &error.RestApiError{HttpCode: 404})
	futureFetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(0)

	fakeObject := NewBitgetFetcher([]IFetcher{spotFetcherMock, futureFetcherMock})

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Nil(t, trades)
}

func TestBitgetFetcherFetchLastTradesWithErrorOnFutureFetcher(t *testing.T) {
	mockController := gomock.NewController(t)

	spotFetcherMock := generatedMocks.NewMockIFetcher(mockController)
	futureFetcherMock := generatedMocks.NewMockIFetcher(mockController)

	spotFetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(1).
		Return(make([]*trading.Trade, 2), nil)
	futureFetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(1).
		Return(nil, &error.RestApiError{HttpCode: 502})

	fakeObject := NewBitgetFetcher([]IFetcher{spotFetcherMock, futureFetcherMock})

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Nil(t, trades)
}

func TestBitgetFetcherFetchLastTrades(t *testing.T) {
	mockController := gomock.NewController(t)

	spotFetcherMock := generatedMocks.NewMockIFetcher(mockController)
	futureFetcherMock := generatedMocks.NewMockIFetcher(mockController)

	spotFetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(2).
		Return(make([]*trading.Trade, 7), nil)
	futureFetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(3).
		Return(make([]*trading.Trade, 5), nil)

	fakeObject := NewBitgetFetcher([]IFetcher{futureFetcherMock, spotFetcherMock, futureFetcherMock, futureFetcherMock, spotFetcherMock})

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotNil(t, trades)
	assert.NotEmpty(t, trades)

	assert.Equal(t, 29, len(trades))
}
