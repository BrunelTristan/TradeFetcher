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

func TestNewTradesFetcher(t *testing.T) {
	fakeObject := NewTradesFetcher[bitgetModel.SpotGetFillQueryParameters, bitgetModel.ApiSpotFill](nil, nil)

	assert.NotNil(t, fakeObject)
}

func TestTradesFetcherFetchLastTradesWithGetError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	tradeConverterMock := generatedMocks.NewMockIStructConverter[bitgetModel.ApiSpotFill, trading.Trade](mockController)

	externalGetterMock.
		EXPECT().
		Get(gomock.Nil()).
		Times(1).
		Return(nil, &error.RestApiError{HttpCode: 500})

	tradeConverterMock.
		EXPECT().
		Convert(gomock.Any()).
		Times(0)

	fakeObject := NewTradesFetcher[bitgetModel.SpotGetFillQueryParameters, bitgetModel.ApiSpotFill](externalGetterMock, tradeConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Nil(t, trades)
}

func TestTradesFetcherFetchLastTradesWithBitgetError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	tradeConverterMock := generatedMocks.NewMockIStructConverter[bitgetModel.ApiSpotFill, trading.Trade](mockController)

	externalGetterMock.
		EXPECT().
		Get(gomock.Nil()).
		Times(1).
		Return(&bitgetModel.ApiSpotGetFills{
			ApiResponse: bitgetModel.ApiResponse{Code: "0999"},
		}, nil)

	tradeConverterMock.
		EXPECT().
		Convert(gomock.Any()).
		Times(0)

	fakeObject := NewTradesFetcher[bitgetModel.SpotGetFillQueryParameters, bitgetModel.ApiSpotFill](externalGetterMock, tradeConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestTradesFetcherFetchLastTradesWithTradeConversionError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	tradeConverterMock := generatedMocks.NewMockIStructConverter[bitgetModel.ApiSpotFill, trading.Trade](mockController)

	apiSubObject := &bitgetModel.ApiSpotFill{
		Symbol: "ORCAUSDT",
	}
	apiResponse := &bitgetModel.ApiSpotGetFills{
		ApiResponse: bitgetModel.ApiResponse{Code: "00000"},
		Data: []*bitgetModel.ApiSpotFill{
			apiSubObject,
		},
	}

	externalGetterMock.
		EXPECT().
		Get(gomock.Nil()).
		Times(1).
		Return(apiResponse, nil)

	tradeConverterMock.
		EXPECT().
		Convert(gomock.Eq(apiSubObject)).
		Times(1).
		Return(nil, &error.ConversionError{Message: "Conversion error message"})

	fakeObject := NewTradesFetcher[bitgetModel.SpotGetFillQueryParameters, bitgetModel.ApiSpotFill](externalGetterMock, tradeConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)

	if errors.As(err, new(*error.BitgetError)) {
		assert.True(t, strings.Contains(err.(*error.BitgetError).Error(), "Conversion error message"))
	}
}

func TestTradesFetcherFetchLastTradesWithoutError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	tradeConverterMock := generatedMocks.NewMockIStructConverter[bitgetModel.ApiSpotFill, trading.Trade](mockController)

	builtBitgetTrades := &bitgetModel.ApiSpotGetFills{
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
		Get(gomock.Nil()).
		Times(1).
		Return(builtBitgetTrades, nil)

	for index, trade := range builtBitgetTrades.Data {
		tradeConverterMock.
			EXPECT().
			Convert(gomock.Eq(trade)).
			Times(1).
			Return(convertedTrades[index], nil)
	}

	fakeObject := NewTradesFetcher[bitgetModel.SpotGetFillQueryParameters, bitgetModel.ApiSpotFill](externalGetterMock, tradeConverterMock)

	assert.NotNil(t, fakeObject)

	trades, err := fakeObject.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 3)

	for index, trade := range trades {
		assert.Equal(t, convertedTrades[index], trade)
	}
}
