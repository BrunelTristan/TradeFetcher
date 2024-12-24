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
	fakeObject := NewBitgetSpotFetcher(nil, nil, nil)

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

	fakeObject := NewBitgetSpotFetcher([]string{""}, externalGetterMock, jsonConverterMock)

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

	fakeObject := NewBitgetSpotFetcher([]string{""}, externalGetterMock, jsonConverterMock)

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

	fakeObject := NewBitgetSpotFetcher([]string{""}, externalGetterMock, jsonConverterMock)

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
					Symbol: "BTCUSDC",
					Side:   "sell",
					Price:  "abc",
					Size:   "0.0054",
				},
			},
		}, nil)

	fakeObject := NewBitgetSpotFetcher([]string{""}, externalGetterMock, jsonConverterMock)

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
					Symbol: "BTCUSDC",
					Side:   "sell",
					Price:  "46",
					Size:   "0..0054",
				},
			},
		}, nil)

	fakeObject := NewBitgetSpotFetcher([]string{""}, externalGetterMock, jsonConverterMock)

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
					Symbol: "BTCUSDC",
					Side:   "sell",
					Price:  "46",
					Size:   "0.0054",
					Fees:   "hundred",
				},
			},
		}, nil)

	fakeObject := NewBitgetSpotFetcher([]string{""}, externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWithoutError(t *testing.T) {
	mockController := gomock.NewController(t)

	assetList := []string{"BTCUSDT", "ETHUSDC", "LINKBTC"}

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)

	externalGetterMock.
		EXPECT().
		Get(gomock.Eq(&bitgetModel.SpotGetFillQueryParameters{Symbol: "BTCUSDT"})).
		Times(1).
		Return("BTC", nil)
	externalGetterMock.
		EXPECT().
		Get(gomock.Eq(&bitgetModel.SpotGetFillQueryParameters{Symbol: "ETHUSDC"})).
		Times(1).
		Return("ETH", nil)
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
		Import(gomock.Eq("BTC")).
		Times(1).
		Return(&bitgetModel.ApiSpotGetFills{
			ApiResponse: bitgetModel.ApiResponse{Code: "000"},
			Data: []*bitgetModel.ApiSpotFill{
				&bitgetModel.ApiSpotFill{
					Symbol: "BTCUSDT",
					Side:   "sell",
					Price:  "106452.12",
					Size:   "0.0054",
					Fees:   "0.0007",
				},
				&bitgetModel.ApiSpotFill{
					Symbol: "BTCUSDT",
					Side:   "sell",
					Price:  "98456.74",
					Size:   "0.0012",
					Fees:   "0.000048",
				},
			},
		}, nil)
	jsonConverterMock.
		EXPECT().
		Import(gomock.Eq("ETH")).
		Times(1).
		Return(&bitgetModel.ApiSpotGetFills{
			ApiResponse: bitgetModel.ApiResponse{Code: "000"},
			Data: []*bitgetModel.ApiSpotFill{
				&bitgetModel.ApiSpotFill{
					Symbol: "ETHUSDC",
					Side:   "buy",
					Price:  "4000.01",
					Size:   "0.004",
					Fees:   "0.00017",
				},
			},
		}, nil)
	jsonConverterMock.
		EXPECT().
		Import(gomock.Eq("LINK")).
		Times(1).
		Return(&bitgetModel.ApiSpotGetFills{
			ApiResponse: bitgetModel.ApiResponse{Code: "000"},
			Data: []*bitgetModel.ApiSpotFill{
				&bitgetModel.ApiSpotFill{
					Symbol: "LINKBTC",
					Side:   "buy",
					Price:  "0.03654",
					Size:   "1234.785",
					Fees:   "0.0012",
				},
				&bitgetModel.ApiSpotFill{
					Symbol: "LINKBTC",
					Side:   "buy",
					Price:  "0.03654",
					Size:   "6547.13",
					Fees:   "0.0048",
				},
				&bitgetModel.ApiSpotFill{
					Symbol: "LINKBTC",
					Side:   "sell",
					Price:  "0.04012",
					Size:   "5555.55",
					Fees:   "0.0037",
				},
			},
		}, nil)

	fakeObject := NewBitgetSpotFetcher(assetList, externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 6)

	if 6 == len(trades) {
		assert.Equal(t, "BTCUSDT", trades[0].Pair)
		assert.Equal(t, 106452.12, trades[0].Price)
		assert.Equal(t, 0.0054, trades[0].Quantity)
		assert.Equal(t, 0.0007, trades[0].Fees)
		assert.Equal(t, "BTCUSDT", trades[1].Pair)
		assert.Equal(t, 98456.74, trades[1].Price)
		assert.Equal(t, 0.0012, trades[1].Quantity)
		assert.Equal(t, 0.000048, trades[1].Fees)
		assert.Equal(t, "ETHUSDC", trades[2].Pair)
		assert.Equal(t, 4000.01, trades[2].Price)
		assert.Equal(t, 0.004, trades[2].Quantity)
		assert.Equal(t, 0.00017, trades[2].Fees)
		assert.Equal(t, "LINKBTC", trades[3].Pair)
		assert.Equal(t, 0.03654, trades[3].Price)
		assert.Equal(t, 1234.785, trades[3].Quantity)
		assert.Equal(t, 0.0012, trades[3].Fees)
		assert.Equal(t, "LINKBTC", trades[4].Pair)
		assert.Equal(t, 0.03654, trades[4].Price)
		assert.Equal(t, 6547.13, trades[4].Quantity)
		assert.Equal(t, 0.0048, trades[4].Fees)
		assert.Equal(t, "LINKBTC", trades[5].Pair)
		assert.Equal(t, 0.04012, trades[5].Price)
		assert.Equal(t, 5555.55, trades[5].Quantity)
		assert.Equal(t, 0.0037, trades[5].Fees)
	}
}
