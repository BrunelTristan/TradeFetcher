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

func TestNewSpotFillsGetter(t *testing.T) {
	fakeObject := NewSpotFillsGetter(nil, nil)

	assert.NotNil(t, fakeObject)
}

func TestGetSpotFillsWithNilParameters(t *testing.T) {
	mockController := gomock.NewController(t)

	apiGetterMock := generatedMocks.NewMockIQuery[bitgetModel.ApiQueryParameters](mockController)
	routeBuilderMock := generatedMocks.NewMockIApiRouteBuilder(mockController)

	apiGetterMock.
		EXPECT().
		Get(gomock.Any()).
		Times(0)

	routeBuilderMock.
		EXPECT().
		BuildRoute(gomock.Any(), gomock.Any()).
		Times(0)

	fakeObject := NewSpotFillsGetter(apiGetterMock, routeBuilderMock)

	output, err := fakeObject.Get(nil)

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Equal(t, 999, err.(*error.RestApiError).HttpCode)
}

func TestGetSpotFillsWithBitgetError(t *testing.T) {
	parameters := &bitgetModel.SpotGetFillQueryParameters{Symbol: "ETHUSDT"}
	expectedRoute := []string{"/api/v2/spot", "/trade/fills"}
	expectedRouteParams := map[string]string{"symbol": "ETHUSDT"}
	expectedApiParameters := &bitgetModel.ApiQueryParameters{
		Route: "/api/v2/spot/trade/fills?symbol=ETHUSDT",
	}
	mockController := gomock.NewController(t)

	apiGetterMock := generatedMocks.NewMockIQuery[bitgetModel.ApiQueryParameters](mockController)
	routeBuilderMock := generatedMocks.NewMockIApiRouteBuilder(mockController)

	apiGetterMock.
		EXPECT().
		Get(gomock.Eq(expectedApiParameters)).
		Times(1).
		Return(nil, &error.RestApiError{HttpCode: 501})

	routeBuilderMock.
		EXPECT().
		BuildRoute(gomock.Eq(expectedRoute), gomock.Eq(expectedRouteParams)).
		Times(1).
		Return(expectedApiParameters.Route)

	fakeObject := NewSpotFillsGetter(apiGetterMock, routeBuilderMock)

	output, err := fakeObject.Get(parameters)

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Equal(t, 501, err.(*error.RestApiError).HttpCode)
}

func TestGetSpotFillsWithoutError(t *testing.T) {
	parameters := &bitgetModel.SpotGetFillQueryParameters{Symbol: "ETHUSDT"}
	expectedRoute := []string{"/api/v2/spot", "/trade/fills"}
	expectedRouteParams := map[string]string{"symbol": "ETHUSDT"}
	expectedApiParameters := &bitgetModel.ApiQueryParameters{
		Route: "/api/v2/spot/trade/fills?symbol=ETHUSDT",
	}
	mockController := gomock.NewController(t)

	apiGetterMock := generatedMocks.NewMockIQuery[bitgetModel.ApiQueryParameters](mockController)
	routeBuilderMock := generatedMocks.NewMockIApiRouteBuilder(mockController)

	apiGetterMock.
		EXPECT().
		Get(gomock.Eq(expectedApiParameters)).
		Times(1).
		Return("{\"msg\":\"sucess\"}", nil)

	routeBuilderMock.
		EXPECT().
		BuildRoute(gomock.Eq(expectedRoute), gomock.Eq(expectedRouteParams)).
		Times(1).
		Return(expectedApiParameters.Route)

	fakeObject := NewSpotFillsGetter(apiGetterMock, routeBuilderMock)

	output, err := fakeObject.Get(parameters)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output)
}
