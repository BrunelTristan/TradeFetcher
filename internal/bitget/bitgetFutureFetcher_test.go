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

func TestNewBitgetFutureFetcher(t *testing.T) {
	fakeObject := NewBitgetFutureFetcher(nil, nil, nil)

	assert.NotNil(t, fakeObject)
}

func TestBitgetFutureFetcherFetchLastTradesWithGetError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)
	tradeConverterMock := generatedMocks.NewMockIStructConverter[bitgetModel.ApiFutureTransaction, trading.Trade](mockController)

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

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock, tradeConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Nil(t, trades)
}

func TestBitgetFutureFetcherFetchLastTradesWithJsonConvertError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)
	tradeConverterMock := generatedMocks.NewMockIStructConverter[bitgetModel.ApiFutureTransaction, trading.Trade](mockController)

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

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock, tradeConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.JsonError)))
	assert.Nil(t, trades)
}

func TestBitgetFutureFetcherFetchLastTradesWithBitgetError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)
	tradeConverterMock := generatedMocks.NewMockIStructConverter[bitgetModel.ApiFutureTransaction, trading.Trade](mockController)

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
		Return(&bitgetModel.ApiFutureTransactions{
			ApiResponse: bitgetModel.ApiResponse{Code: "0999"},
		}, nil)

	tradeConverterMock.
		EXPECT().
		Convert(gomock.Any()).
		Times(0)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock, tradeConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetFutureFetcherFetchLastTradesWithTradeConversionError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)
	tradeConverterMock := generatedMocks.NewMockIStructConverter[bitgetModel.ApiFutureTransaction, trading.Trade](mockController)

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
		Return(&bitgetModel.ApiFutureTransactions{
			ApiResponse: bitgetModel.ApiResponse{Code: "000"},
			Data: &bitgetModel.ApiFutureTransactionsList{
				FillList: []*bitgetModel.ApiFutureTransaction{
					&bitgetModel.ApiFutureTransaction{},
				},
			},
		}, nil)

	tradeConverterMock.
		EXPECT().
		Convert(gomock.Any()).
		Times(1).
		Return(nil, &error.ConversionError{Message: "Conversion error message"})

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock, tradeConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)

	if errors.As(err, new(*error.BitgetError)) {
		assert.True(t, strings.Contains(err.(*error.BitgetError).Error(), "Conversion error message"))
	}
}

func TestBitgetFutureFetcherFetchLastTradesWithoutError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)
	tradeConverterMock := generatedMocks.NewMockIStructConverter[bitgetModel.ApiFutureTransaction, trading.Trade](mockController)

	buildedBitgetTrades := &bitgetModel.ApiFutureTransactions{
		ApiResponse: bitgetModel.ApiResponse{Code: "000"},
		Data: &bitgetModel.ApiFutureTransactionsList{
			FillList: []*bitgetModel.ApiFutureTransaction{
				&bitgetModel.ApiFutureTransaction{Symbol: "LINKBTC"},
				&bitgetModel.ApiFutureTransaction{Symbol: "AVAXBTC"},
				&bitgetModel.ApiFutureTransaction{Symbol: "ETHBTC"},
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
		Get(gomock.Any()).
		Times(1).
		Return("BTC", nil)

	jsonConverterMock.
		EXPECT().
		Export(gomock.Any()).
		Times(0)
	jsonConverterMock.
		EXPECT().
		Import(gomock.Eq("BTC")).
		Times(1).
		Return(buildedBitgetTrades, nil)

	tradeConverterMock.
		EXPECT().
		Convert(gomock.Eq(buildedBitgetTrades.Data.FillList[0])).
		Times(1).
		Return(convertedTrades[0], nil)
	tradeConverterMock.
		EXPECT().
		Convert(gomock.Eq(buildedBitgetTrades.Data.FillList[1])).
		Times(1).
		Return(convertedTrades[1], nil)
	tradeConverterMock.
		EXPECT().
		Convert(gomock.Eq(buildedBitgetTrades.Data.FillList[2])).
		Times(1).
		Return(convertedTrades[2], nil)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock, tradeConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 3)

	for index, trade := range trades {
		assert.Equal(t, *convertedTrades[index], trade)
	}
}
