package bitget

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"strings"
	"testing"
	"tradeFetcher/internal/generatedMocks"
	bitgetModel "tradeFetcher/model/bitget"
	"tradeFetcher/model/error"
	"tradeFetcher/model/trading"
)

func TestNewBitgetSpotFetcher(t *testing.T) {
	fakeObject := NewBitgetSpotFetcher("", nil, nil, nil)

	assert.NotNil(t, fakeObject)
}

func TestBitgetSpotFetcherFetchLastTradesWithGetError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)
	tradeConverterMock := generatedMocks.NewMockIStructConverter[bitgetModel.ApiSpotFill, trading.Trade](mockController)

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

	tradeConverterMock.
		EXPECT().
		Convert(gomock.Any()).
		Times(0)

	fakeObject := NewBitgetSpotFetcher("", externalGetterMock, jsonConverterMock, tradeConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWithJsonConvertError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)
	tradeConverterMock := generatedMocks.NewMockIStructConverter[bitgetModel.ApiSpotFill, trading.Trade](mockController)

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

	tradeConverterMock.
		EXPECT().
		Convert(gomock.Any()).
		Times(0)

	fakeObject := NewBitgetSpotFetcher("", externalGetterMock, jsonConverterMock, tradeConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.JsonError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWithBitgetError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)
	tradeConverterMock := generatedMocks.NewMockIStructConverter[bitgetModel.ApiSpotFill, trading.Trade](mockController)

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

	tradeConverterMock.
		EXPECT().
		Convert(gomock.Any()).
		Times(0)

	fakeObject := NewBitgetSpotFetcher("", externalGetterMock, jsonConverterMock, tradeConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetSpotFetcherFetchLastTradesWithTradeConversionError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)
	tradeConverterMock := generatedMocks.NewMockIStructConverter[bitgetModel.ApiSpotFill, trading.Trade](mockController)

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

	tradeConverterMock.
		EXPECT().
		Convert(gomock.Any()).
		Times(1).
		Return(nil, &error.ConversionError{Message: "Conversion error message"})

	fakeObject := NewBitgetSpotFetcher("", externalGetterMock, jsonConverterMock, tradeConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)

	if errors.As(err, new(*error.BitgetError)) {
		assert.True(t, strings.Contains(err.(*error.BitgetError).Error(), "Conversion error message"))
	}
}

func TestBitgetSpotFetcherFetchLastTradesWithoutError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiSpotGetFills](mockController)
	tradeConverterMock := generatedMocks.NewMockIStructConverter[bitgetModel.ApiSpotFill, trading.Trade](mockController)

	buildedBitgetTrades := &bitgetModel.ApiSpotGetFills{
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
	}

	convertedTrades := []*trading.Trade{
		&trading.Trade{Pair: "A"},
		&trading.Trade{Pair: "B"},
		&trading.Trade{Pair: "CDE"},
	}

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
		Return(buildedBitgetTrades, nil)

	tradeConverterMock.
		EXPECT().
		Convert(gomock.Eq(buildedBitgetTrades.Data[0])).
		Times(1).
		Return(convertedTrades[0], nil)
	tradeConverterMock.
		EXPECT().
		Convert(gomock.Eq(buildedBitgetTrades.Data[1])).
		Times(1).
		Return(convertedTrades[1], nil)
	tradeConverterMock.
		EXPECT().
		Convert(gomock.Eq(buildedBitgetTrades.Data[2])).
		Times(1).
		Return(convertedTrades[2], nil)

	fakeObject := NewBitgetSpotFetcher("LINKBTC", externalGetterMock, jsonConverterMock, tradeConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 3)

	for index, trade := range trades {
		assert.Equal(t, *convertedTrades[index], trade)
	}
}
