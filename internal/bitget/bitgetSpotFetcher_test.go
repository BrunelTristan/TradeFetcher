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
	fakeObject := NewBitgetSpotFetcher("", nil, nil)

	assert.NotNil(t, fakeObject)
}

func TestBitgetSpotFetcherFetchLastTradesWithGetError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
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

	fakeObject := NewBitgetSpotFetcher("", externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWithJsonConvertError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
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

	fakeObject := NewBitgetSpotFetcher("", externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.JsonError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWithBitgetError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
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

	fakeObject := NewBitgetSpotFetcher("", externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWithPriceFloatingError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
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
					Symbol:     "BTCUSDC",
					Side:       "sell",
					Price:      "abc",
					LastUpdate: "123456",
					Size:       "0.0054",
				},
			},
		}, nil)

	fakeObject := NewBitgetSpotFetcher("", externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWithQuantityFloatingError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
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
					Symbol:     "BTCUSDC",
					Side:       "sell",
					Price:      "46",
					LastUpdate: "123456",
					Size:       "0..0054",
				},
			},
		}, nil)

	fakeObject := NewBitgetSpotFetcher("", externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWithExecutingTimeError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
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
					Symbol:     "BTCUSDC",
					Side:       "sell",
					Price:      "46",
					LastUpdate: "abcsde",
					Size:       "0.0054",
					FeeDetail: &bitgetModel.ApiFeeDetail{
						FeesValue: "0.00254",
					},
				},
			},
		}, nil)

	fakeObject := NewBitgetSpotFetcher("", externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWithSideError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
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
					Symbol:     "BTCUSDC",
					Side:       "bud",
					Price:      "46",
					LastUpdate: "1745698523",
					Size:       "0.0054",
					FeeDetail: &bitgetModel.ApiFeeDetail{
						FeesValue: "0.00254",
					},
				},
			},
		}, nil)

	fakeObject := NewBitgetSpotFetcher("", externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWitFeesFloatingError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
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
					Symbol:     "BTCUSDC",
					Side:       "sell",
					Price:      "46",
					LastUpdate: "123456",
					Size:       "0.0054",
					FeeDetail: &bitgetModel.ApiFeeDetail{
						FeesValue: "hundred",
					},
				},
			},
		}, nil)

	fakeObject := NewBitgetSpotFetcher("", externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWithoutError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)

	externalGetterMock.
		EXPECT().
		Get(gomock.Eq(&bitgetModel.SpotGetFillQueryParameters{Symbol: "LINKBTC"})).
		Times(1).
		Return("LINK", nil)

	jsonConverterMock.
		EXPECT().
		Export(gomock.Any()).
		Times(0)
	jsonConverterMock.
		EXPECT().
		Import(gomock.Eq("LINK")).
		Times(1).
		Return(&bitgetModel.ApiSpotGetFills{
			ApiResponse: bitgetModel.ApiResponse{Code: "000"},
			Data: []*bitgetModel.ApiSpotFill{
				&bitgetModel.ApiSpotFill{
					Symbol:     "LINKBTC",
					Side:       "buy",
					Price:      "0.03654",
					LastUpdate: "16549876877",
					Size:       "1234.785",
					FeeDetail: &bitgetModel.ApiFeeDetail{
						FeesValue: "0.0012",
					},
				},
				&bitgetModel.ApiSpotFill{
					Symbol:     "LINKBTC",
					Side:       "buy",
					Price:      "0.03654",
					LastUpdate: "16549976789",
					Size:       "6547.13",
					FeeDetail: &bitgetModel.ApiFeeDetail{
						FeesValue: "0.0048",
					},
				},
				&bitgetModel.ApiSpotFill{
					Symbol:     "LINKBTC",
					Side:       "sell",
					Price:      "0.04012",
					LastUpdate: "16550876654",
					Size:       "5555.55",
					FeeDetail: &bitgetModel.ApiFeeDetail{
						FeesValue: "0.0037",
					},
				},
			},
		}, nil)

	fakeObject := NewBitgetSpotFetcher("LINKBTC", externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 3)

	if 3 == len(trades) {
		assert.Equal(t, "LINKBTC", trades[0].Pair)
		assert.Equal(t, 0.03654, trades[0].Price)
		assert.Equal(t, 1234.785, trades[0].Quantity)
		assert.Equal(t, 0.0012, trades[0].Fees)
		assert.Equal(t, int64(16549876), trades[0].ExecutedTimestamp)
		assert.True(t, trades[0].Open)
		assert.True(t, trades[0].Long)
		assert.Equal(t, "LINKBTC", trades[1].Pair)
		assert.Equal(t, 0.03654, trades[1].Price)
		assert.Equal(t, 6547.13, trades[1].Quantity)
		assert.Equal(t, 0.0048, trades[1].Fees)
		assert.Equal(t, int64(16549976), trades[1].ExecutedTimestamp)
		assert.True(t, trades[1].Open)
		assert.True(t, trades[1].Long)
		assert.Equal(t, "LINKBTC", trades[2].Pair)
		assert.Equal(t, 0.04012, trades[2].Price)
		assert.Equal(t, 5555.55, trades[2].Quantity)
		assert.Equal(t, 0.0037, trades[2].Fees)
		assert.Equal(t, int64(16550876), trades[2].ExecutedTimestamp)
		assert.False(t, trades[2].Open)
		assert.True(t, trades[2].Long)
	}
}
