package fetcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFakeFetcher(t *testing.T) {
	fakeObject := NewFakeFetcher()
	
	assert.NotNil(t, fakeObject)
}
