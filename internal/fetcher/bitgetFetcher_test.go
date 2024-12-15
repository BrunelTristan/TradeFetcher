package fetcher

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"tradeFetcher/internal/generatedMocks"
	"tradeFetcher/model/trading"
)

func TestNewBitgetFetcher(t *testing.T) {
	fakeObject := NewBitgetFetcher(nil)

	assert.NotNil(t, fakeObject)
}

func TestBitgetFetcherFetchLastTrades(t *testing.T) {
	mockController := gomock.NewController(t)

	spotFetcherMock := generatedMocks.NewMockIFetcher(mockController)

	tradesCount := 7

	spotFetcherMock.
		EXPECT().
		FetchLastTrades().
		Times(1).
		Return(make([]trading.Trade, tradesCount))

	fakeObject := NewBitgetFetcher(spotFetcherMock)

	assert.NotNil(t, fakeObject)

	trades := fakeObject.FetchLastTrades()

	assert.NotNil(t, trades)
	assert.NotEmpty(t, trades)

	assert.Equal(t, tradesCount, len(trades))
}
