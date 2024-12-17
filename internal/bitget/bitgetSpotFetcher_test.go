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

func TestNewBitgetSpotFetcher(t *testing.T) {
	fakeObject := NewBitgetSpotFetcher(nil, nil)

	assert.NotNil(t, fakeObject)
}

func TestBitgetSpotFetcherFetchLastTradesWithGetError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIGetter(mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)

	externalGetterMock.
		EXPECT().
		Get(gomock.Any()).
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

	fakeObject := NewBitgetSpotFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWithJsonConvertError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIGetter(mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)

	externalGetterMock.
		EXPECT().
		Get(gomock.Any()).
		Times(1).
		Return("", nil)

	jsonConverterMock.
		EXPECT().
		Export(gomock.Any()).
		Times(0)
	jsonConverterMock.
		EXPECT().
		Import(gomock.Any()).
		Times(1).
		Return(nil, &error.JsonError{})

	fakeObject := NewBitgetSpotFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.JsonError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWithBitgetError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIGetter(mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)

	externalGetterMock.
		EXPECT().
		Get(gomock.Any()).
		Times(1).
		Return("", nil)

	jsonConverterMock.
		EXPECT().
		Export(gomock.Any()).
		Times(0)
	jsonConverterMock.
		EXPECT().
		Import(gomock.Any()).
		Times(1).
		Return(&bitgetModel.ApiSpotGetFills{
			ApiResponse: bitgetModel.ApiResponse{Code: "0999"},
		}, nil)

	fakeObject := NewBitgetSpotFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWithPriceFloatingError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIGetter(mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)

	externalGetterMock.
		EXPECT().
		Get(gomock.Any()).
		Times(1).
		Return("", nil)

	jsonConverterMock.
		EXPECT().
		Export(gomock.Any()).
		Times(0)
	jsonConverterMock.
		EXPECT().
		Import(gomock.Any()).
		Times(1).
		Return(&bitgetModel.ApiSpotGetFills{
			ApiResponse: bitgetModel.ApiResponse{Code: "000"},
			Data: []*bitgetModel.ApiSpotFill{
				&bitgetModel.ApiSpotFill{
					Symbol: "BTCUSDC",
					Side:   "sell",
					Price:  "abc",
					Size:   "0.0054",
				},
			},
		}, nil)

	fakeObject := NewBitgetSpotFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWithQuantityFloatingError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIGetter(mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)

	externalGetterMock.
		EXPECT().
		Get(gomock.Any()).
		Times(1).
		Return("", nil)

	jsonConverterMock.
		EXPECT().
		Export(gomock.Any()).
		Times(0)
	jsonConverterMock.
		EXPECT().
		Import(gomock.Any()).
		Times(1).
		Return(&bitgetModel.ApiSpotGetFills{
			ApiResponse: bitgetModel.ApiResponse{Code: "000"},
			Data: []*bitgetModel.ApiSpotFill{
				&bitgetModel.ApiSpotFill{
					Symbol: "BTCUSDC",
					Side:   "sell",
					Price:  "46",
					Size:   "0..0054",
				},
			},
		}, nil)

	fakeObject := NewBitgetSpotFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWithoutError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIGetter(mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)

	externalGetterMock.
		EXPECT().
		Get(gomock.Any()).
		Times(1).
		Return("", nil)

	jsonConverterMock.
		EXPECT().
		Export(gomock.Any()).
		Times(0)
	jsonConverterMock.
		EXPECT().
		Import(gomock.Any()).
		Times(1).
		Return(&bitgetModel.ApiSpotGetFills{
			ApiResponse: bitgetModel.ApiResponse{Code: "000"},
			Data: []*bitgetModel.ApiSpotFill{
				&bitgetModel.ApiSpotFill{
					Symbol: "ARUSDT",
					Side:   "buy",
					Price:  "0.3654",
					Size:   "1234785",
				},
				&bitgetModel.ApiSpotFill{
					Symbol: "BTCUSDC",
					Side:   "sell",
					Price:  "106452.12",
					Size:   "0.0054",
				},
			},
		}, nil)

	fakeObject := NewBitgetSpotFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 2)

	if 2 == len(trades) {
		assert.Equal(t, "ARUSDT", trades[0].Pair)
		assert.Equal(t, 0.3654, trades[0].Price)
		assert.Equal(t, 1234785.0, trades[0].Quantity)
		assert.Equal(t, "BTCUSDC", trades[1].Pair)
		assert.Equal(t, 106452.12, trades[1].Price)
		assert.Equal(t, 0.0054, trades[1].Quantity)
	}
}
