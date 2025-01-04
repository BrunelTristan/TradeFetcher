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

func TestNewBitgetFutureTransactionsGetter(t *testing.T) {
	fakeObject := NewBitgetFutureTransactionsGetter(nil, nil)

	assert.NotNil(t, fakeObject)
}

func TestGetFutureTransactionsWithNilParameters(t *testing.T) {
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

	fakeObject := NewBitgetFutureTransactionsGetter(apiGetterMock, routeBuilderMock)

	output, err := fakeObject.Get(nil)

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Equal(t, 999, err.(*error.RestApiError).HttpCode)
}

func TestGetFutureTransactionsWithBitgetError(t *testing.T) {
	parameters := &bitgetModel.FutureTransactionsQueryParameters{}
	expectedRoute := []string{"/api/v2/mix", "/order/fill-history"}
	expectedRouteParams := map[string]string{"productType": "USDT-FUTURES"}
	expectedApiParameters := &bitgetModel.ApiQueryParameters{
		Route: "/api/v2/mix/order/fill-history?productType=USDT-FUTURES",
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

	fakeObject := NewBitgetFutureTransactionsGetter(apiGetterMock, routeBuilderMock)

	output, err := fakeObject.Get(parameters)

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Equal(t, 501, err.(*error.RestApiError).HttpCode)
}

func TestGetFutureTransactionsWithoutError(t *testing.T) {
	parameters := &bitgetModel.FutureTransactionsQueryParameters{}
	expectedRoute := []string{"/api/v2/mix", "/order/fill-history"}
	expectedRouteParams := map[string]string{"productType": "USDT-FUTURES"}
	expectedApiParameters := &bitgetModel.ApiQueryParameters{
		Route: "/api/v2/mix/order/fill-history?productType=USDT-FUTURES",
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

	fakeObject := NewBitgetFutureTransactionsGetter(apiGetterMock, routeBuilderMock)

	output, err := fakeObject.Get(parameters)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output)
}
