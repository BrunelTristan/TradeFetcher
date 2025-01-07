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

func TestNewApiQueryToStructDecorator(t *testing.T) {
	fakeObject := NewApiQueryToStructDecorator[bitgetModel.SpotGetFillQueryParameters, bitgetModel.ApiSpotGetFills](nil, nil, nil)

	assert.NotNil(t, fakeObject)
}

func TestApiQueryToStructDecoratorGetWithErrorOnParamBuilder(t *testing.T) {
	mockController := gomock.NewController(t)

	paramBuilderMock := generatedMocks.NewMockIQueryParametersBuilder[bitgetModel.SpotGetFillQueryParameters](mockController)
	apiQueryMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)

	paramBuilderMock.
		EXPECT().
		BuildQueryParameters().
		Times(1).
		Return(nil, errors.New("an error"))

	apiQueryMock.
		EXPECT().
		Get(gomock.Any()).
		Times(0)

	jsonConverterMock.
		EXPECT().
		Export(gomock.Any()).
		Times(0)
	jsonConverterMock.
		EXPECT().
		Import(gomock.Any()).
		Times(0)

	getter := NewApiQueryToStructDecorator[bitgetModel.SpotGetFillQueryParameters, bitgetModel.ApiSpotGetFills](apiQueryMock, paramBuilderMock, jsonConverterMock)

	output, err := getter.Get(nil)

	assert.Nil(t, output)
	assert.NotNil(t, err)
}

func TestApiQueryToStructDecoratorGetWithErrorOnDecoratee(t *testing.T) {
	mockController := gomock.NewController(t)

	paramBuilderMock := generatedMocks.NewMockIQueryParametersBuilder[bitgetModel.SpotGetFillQueryParameters](mockController)
	apiQueryMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)

	builtParameters := &bitgetModel.SpotGetFillQueryParameters{Symbol: "NONE"}

	paramBuilderMock.
		EXPECT().
		BuildQueryParameters().
		Times(1).
		Return(builtParameters, nil)

	apiQueryMock.
		EXPECT().
		Get(gomock.Eq(builtParameters)).
		Times(1).
		Return(nil, &error.RestApiError{HttpCode: 500})

	jsonConverterMock.
		EXPECT().
		Export(gomock.Any()).
		Times(0)
	jsonConverterMock.
		EXPECT().
		Import(gomock.Any()).
		Times(0)

	getter := NewApiQueryToStructDecorator[bitgetModel.SpotGetFillQueryParameters, bitgetModel.ApiSpotGetFills](apiQueryMock, paramBuilderMock, jsonConverterMock)

	output, err := getter.Get(nil)

	assert.Nil(t, output)
	assert.NotNil(t, err)
	assert.True(t, errors.As(err, new(*error.RestApiError)))
}

func TestApiQueryToStructDecoratorGetWithErrorOnJsonConversion(t *testing.T) {
	mockController := gomock.NewController(t)

	paramBuilderMock := generatedMocks.NewMockIQueryParametersBuilder[bitgetModel.SpotGetFillQueryParameters](mockController)
	apiQueryMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)

	builtParameters := &bitgetModel.SpotGetFillQueryParameters{Symbol: "NONE"}

	paramBuilderMock.
		EXPECT().
		BuildQueryParameters().
		Times(1).
		Return(builtParameters, nil)

	apiQueryMock.
		EXPECT().
		Get(gomock.Eq(builtParameters)).
		Times(1).
		Return("json", nil)

	jsonConverterMock.
		EXPECT().
		Export(gomock.Any()).
		Times(0)
	jsonConverterMock.
		EXPECT().
		Import(gomock.Eq("json")).
		Times(1).
		Return(nil, &error.JsonError{})

	getter := NewApiQueryToStructDecorator[bitgetModel.SpotGetFillQueryParameters, bitgetModel.ApiSpotGetFills](apiQueryMock, paramBuilderMock, jsonConverterMock)

	output, err := getter.Get(nil)

	assert.Nil(t, output)
	assert.NotNil(t, err)
	assert.True(t, errors.As(err, new(*error.JsonError)))
}

func TestApiQueryToStructDecoratorGetWithoutError(t *testing.T) {
	mockController := gomock.NewController(t)

	paramBuilderMock := generatedMocks.NewMockIQueryParametersBuilder[bitgetModel.SpotGetFillQueryParameters](mockController)
	apiQueryMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)

	builtParameters := &bitgetModel.SpotGetFillQueryParameters{Symbol: "NONE"}
	expected := &bitgetModel.ApiSpotGetFills{ApiResponse: bitgetModel.ApiResponse{Code: "OK"}}

	paramBuilderMock.
		EXPECT().
		BuildQueryParameters().
		Times(1).
		Return(builtParameters, nil)

	apiQueryMock.
		EXPECT().
		Get(gomock.Eq(builtParameters)).
		Times(1).
		Return("json", nil)

	jsonConverterMock.
		EXPECT().
		Export(gomock.Any()).
		Times(0)
	jsonConverterMock.
		EXPECT().
		Import(gomock.Eq("json")).
		Times(1).
		Return(expected, nil)

	getter := NewApiQueryToStructDecorator[bitgetModel.SpotGetFillQueryParameters, bitgetModel.ApiSpotGetFills](apiQueryMock, paramBuilderMock, jsonConverterMock)

	output, err := getter.Get(nil)

	assert.Nil(t, err)
	assert.Equal(t, expected, output)
}
