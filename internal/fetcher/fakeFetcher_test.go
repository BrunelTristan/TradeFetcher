package fetcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFakeFetcher(t *testing.T) {
	fakeObject := NewFakeFetcher()

	assert.NotNil(t, fakeObject)
}

func TestFetchLastTrades(t *testing.T) {
	fakeObject := NewFakeFetcher()

	assert.NotNil(t, fakeObject)

	trades := fakeObject.FetchLastTrades()

	assert.NotNil(t, trades)
}
