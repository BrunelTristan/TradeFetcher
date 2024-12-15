package fetcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFakeFetcher(t *testing.T) {
	fakeObject := NewFakeFetcher()

	assert.NotNil(t, fakeObject)
}

func TestFakeFetcherFetchLastTrades(t *testing.T) {
	fakeObject := NewFakeFetcher()

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotNil(t, trades)
	assert.NotEmpty(t, trades)

	for _, trade := range trades {
		assert.NotNil(t, trade)
		assert.NotEmpty(t, trade)
	}
}
