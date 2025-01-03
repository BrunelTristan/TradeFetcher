package fetcher

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"tradeFetcher/internal/generatedMocks"
	"tradeFetcher/model/error"
	"tradeFetcher/model/trading"
)

func TestNewSortByDateFetcherDecorator(t *testing.T) {
	fakeObject := NewSortByDateFetcherDecorator(nil)

	assert.NotNil(t, fakeObject)
}

func TestSortByDateFetcherDecoratorFetchLastTradesWithDecorateeError(t *testing.T) {
	mockController := gomock.NewController(t)

	fetcherMock := generatedMocks.NewMockIFetcher(mockController)

	fetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(1).
		Return(nil, &error.RestApiError{HttpCode: 404})

	fakeObject := NewSortByDateFetcherDecorator(fetcherMock)

	trades, err := fakeObject.FetchLastTrades()

	assert.NotNil(t, err)
	assert.Nil(t, trades)
}

func TestSortByDateFetcherDecoratorFetchLastTradesFromDecoratee(t *testing.T) {
	mockController := gomock.NewController(t)

	fetcherMock := generatedMocks.NewMockIFetcher(mockController)

	tradesCount := 51

	fetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(1).
		Return(make([]trading.Trade, tradesCount), nil)

	fakeObject := NewSortByDateFetcherDecorator(fetcherMock)

	trades, err := fakeObject.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotNil(t, trades)
	assert.NotEmpty(t, trades)

	assert.Equal(t, tradesCount, len(trades))
}

func TestSortByDateFetcherDecoratorFetchLastTradesSortByDate(t *testing.T) {
	mockController := gomock.NewController(t)

	fetcherMock := generatedMocks.NewMockIFetcher(mockController)

	fetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(1).
		Return([]trading.Trade{
			trading.Trade{ExecutedTimestamp: 456},
			trading.Trade{ExecutedTimestamp: 789},
			trading.Trade{ExecutedTimestamp: 123},
		}, nil)

	fakeObject := NewSortByDateFetcherDecorator(fetcherMock)

	trades, err := fakeObject.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotNil(t, trades)
	assert.NotEmpty(t, trades)

	assert.Equal(t, 3, len(trades))

	if 3 == len(trades) {
		assert.Equal(t, int64(123), trades[0].ExecutedTimestamp)
		assert.Equal(t, int64(456), trades[1].ExecutedTimestamp)
		assert.Equal(t, int64(789), trades[2].ExecutedTimestamp)
	}
}
