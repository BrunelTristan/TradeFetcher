package bitget

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"strconv"
	"testing"
	"time"
	"tradeFetcher/internal/generatedMocks"
	bitgetModel "tradeFetcher/model/bitget"
	"tradeFetcher/model/error"
)

func TestNewFutureTaxTransactionsGetter(t *testing.T) {
	fakeObject := NewFutureTaxTransactionsGetter(nil, nil)

	assert.NotNil(t, fakeObject)
}

func TestGetFutureTaxTransactionsWithBitgetError(t *testing.T) {
	parameters := &bitgetModel.FutureTaxTransactionsQueryParameters{}
	currentTimestamp := time.Now().Unix()
	previousTimestamp := currentTimestamp - (20 * 24 * 3600)
	currentTimestampString := strconv.FormatInt(currentTimestamp, 10)
	previousTimestampString := strconv.FormatInt(previousTimestamp, 10)
	expectedRoute := []string{"/api/v2/tax", "/future-record"}
	expectedRouteParams := map[string]string{"endTime": currentTimestampString, "startTime": previousTimestampString}
	expectedApiParameters := &bitgetModel.ApiQueryParameters{
		Route: "/api/v2/mix/future-record?endTime=" + currentTimestampString + "&startTime=" + previousTimestampString,
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

	fakeObject := NewFutureTaxTransactionsGetter(apiGetterMock, routeBuilderMock)

	output, err := fakeObject.Get(parameters)

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Equal(t, 501, err.(*error.RestApiError).HttpCode)
}

func TestGetFutureTaxTransactionsWithoutError(t *testing.T) {
	parameters := &bitgetModel.FutureTaxTransactionsQueryParameters{}
	currentTimestamp := time.Now().Unix()
	previousTimestamp := currentTimestamp - (20 * 24 * 3600)
	currentTimestampString := strconv.FormatInt(currentTimestamp, 10)
	previousTimestampString := strconv.FormatInt(previousTimestamp, 10)
	expectedRoute := []string{"/api/v2/tax", "/future-record"}
	expectedRouteParams := map[string]string{"endTime": currentTimestampString, "startTime": previousTimestampString}
	expectedApiParameters := &bitgetModel.ApiQueryParameters{
		Route: "/api/v2/mix/future-record?endTime=" + currentTimestampString + "&startTime=" + previousTimestampString,
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

	fakeObject := NewFutureTaxTransactionsGetter(apiGetterMock, routeBuilderMock)

	output, err := fakeObject.Get(parameters)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output)
}
