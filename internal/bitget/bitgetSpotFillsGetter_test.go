package bitget

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"tradeFetcher/internal/generatedMocks"
	bitgetModel "tradeFetcher/model/bitget"
	"tradeFetcher/model/error"
)

func TestNewBitgetSpotFillsGetter(t *testing.T) {
	fakeObject := NewBitgetSpotFillsGetter(nil)

	assert.NotNil(t, fakeObject)
}

func TestGetSpotFillsWithNilParameters(t *testing.T) {
	mockController := gomock.NewController(t)

	apiGetterMock := generatedMocks.NewMockICommand[bitgetModel.ApiCommandParameters](mockController)

	apiGetterMock.
		EXPECT().
		Get(gomock.Any()).
		Times(0)

	fakeObject := NewBitgetSpotFillsGetter(apiGetterMock)

	output, err := fakeObject.Get(nil)

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Equal(t, 999, err.(*error.RestApiError).HttpCode)
}

func TestGetSpotFillsWithBitgetError(t *testing.T) {
	parameters := &bitgetModel.SpotGetFillCommandParameters{Symbol: "ETHUSDT"}
	expectedApiParameters := &bitgetModel.ApiCommandParameters{
		Route: "/api/v2/spot/trade/fills?symbol=ETHUSDT",
	}
	mockController := gomock.NewController(t)

	apiGetterMock := generatedMocks.NewMockICommand[bitgetModel.ApiCommandParameters](mockController)

	apiGetterMock.
		EXPECT().
		Get(gomock.Eq(expectedApiParameters)).
		Times(1).
		Return(nil, &error.RestApiError{HttpCode: 501})

	fakeObject := NewBitgetSpotFillsGetter(apiGetterMock)

	output, err := fakeObject.Get(parameters)

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Equal(t, 501, err.(*error.RestApiError).HttpCode)
}
