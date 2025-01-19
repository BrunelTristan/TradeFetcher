package fetcher

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"tradeFetcher/internal/generatedMocks"
	"tradeFetcher/model/error"
	"tradeFetcher/model/trading"
)

func TestNewFilterByDateFetcherDecorator(t *testing.T) {
	fakeObject := NewFilterByDateFetcherDecorator(nil, nil)

	assert.NotNil(t, fakeObject)
}

func TestFilterByDateFetcherDecoratorFetchLastTradesWithDecorateeError(t *testing.T) {
	mockController := gomock.NewController(t)

	fetcherMock := generatedMocks.NewMockIFetcher(mockController)
	lastProceedMock := generatedMocks.NewMockILastProceedRetriever(mockController)

	fetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(1).
		Return(nil, &error.RestApiError{HttpCode: 404})

	lastProceedMock.
		EXPECT().
		GetLastProceedTimestamp().
		Times(0)

	fakeObject := NewFilterByDateFetcherDecorator(fetcherMock, lastProceedMock)

	trades, err := fakeObject.FetchLastTrades()

	assert.NotNil(t, err)
	assert.Nil(t, trades)
}

func TestFilterByDateFetcherDecoratorFetchLastTradesFromDecoratee(t *testing.T) {
	mockController := gomock.NewController(t)

	fetcherMock := generatedMocks.NewMockIFetcher(mockController)
	lastProceedMock := generatedMocks.NewMockILastProceedRetriever(mockController)

	tradesCount := 5
	fetchedTrades := []*trading.Trade{
		&trading.Trade{ExecutedTimestamp: 456},
		&trading.Trade{ExecutedTimestamp: 789},
		&trading.Trade{ExecutedTimestamp: 123},
		&trading.Trade{ExecutedTimestamp: 852},
		&trading.Trade{ExecutedTimestamp: 369},
	}

	fetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(1).
		Return(fetchedTrades, nil)

	lastProceedMock.
		EXPECT().
		GetLastProceedTimestamp().
		Times(1).
		Return(int64(12))

	fakeObject := NewFilterByDateFetcherDecorator(fetcherMock, lastProceedMock)

	trades, err := fakeObject.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotNil(t, trades)
	assert.NotEmpty(t, trades)

	assert.Equal(t, tradesCount, len(trades))
}

func TestFilterByDateFetcherDecoratorFetchLastTradesFilterByDate(t *testing.T) {
	mockController := gomock.NewController(t)

	fetcherMock := generatedMocks.NewMockIFetcher(mockController)
	lastProceedMock := generatedMocks.NewMockILastProceedRetriever(mockController)

	fetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(1).
		Return([]*trading.Trade{
			&trading.Trade{ExecutedTimestamp: 123},
			&trading.Trade{ExecutedTimestamp: 456},
			&trading.Trade{ExecutedTimestamp: 789},
		}, nil)

	lastProceedMock.
		EXPECT().
		GetLastProceedTimestamp().
		Times(1).
		Return(int64(345))

	fakeObject := NewFilterByDateFetcherDecorator(fetcherMock, lastProceedMock)

	trades, err := fakeObject.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotNil(t, trades)
	assert.NotEmpty(t, trades)

	assert.Equal(t, 2, len(trades))

	if 2 == len(trades) {
		assert.Equal(t, int64(456), trades[0].ExecutedTimestamp)
		assert.Equal(t, int64(789), trades[1].ExecutedTimestamp)
	}
}
