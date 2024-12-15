package bitget

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"tradeFetcher/internal/generatedMocks"
)

func TestNewBitgetSpotFetcher(t *testing.T) {
	fakeObject := NewBitgetSpotFetcher(nil)

	assert.NotNil(t, fakeObject)
}

func TestBitgetSpotFetcherFetchLastTradesWithoutStrindGet(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIGetter(mockController)

	externalGetterMock.
		EXPECT().
		Get(gomock.Any()).
		Times(1).
		Return(52, nil)

	fakeObject := NewBitgetSpotFetcher(externalGetterMock)

	assert.NotNil(t, fakeObject)

	_ = fakeObject.FetchLastTrades()
}
