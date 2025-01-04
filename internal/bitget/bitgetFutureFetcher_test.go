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

func TestNewBitgetFutureFetcher(t *testing.T) {
	fakeObject := NewBitgetFutureFetcher(nil, nil)

	assert.NotNil(t, fakeObject)
}

func TestBitgetFutureFetcherFetchLastTradesWithGetError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.RestApiError)))
	assert.Nil(t, trades)
}

func TestBitgetFutureFetcherFetchLastTradesWithJsonConvertError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.JsonError)))
	assert.Nil(t, trades)
}

func TestBitgetFutureFetcherFetchLastTradesWithBitgetError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetFutureFetcherFetchLastTradesWithPriceFloatingError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "BTCUSDC",
						Side:       "sell",
						Price:      "abc",
						LastUpdate: "123456",
						Size:       "0.0054",
						TradeSide:  "open",
					},
				},
			},
		}, nil)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetFutureFetcherFetchLastTradesWithQuantityFloatingError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "BTCUSDC",
						Side:       "sell",
						Price:      "46",
						LastUpdate: "123456",
						Size:       "0..0054",
						TradeSide:  "open",
					},
				},
			},
		}, nil)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetFutureFetcherFetchLastTradesWithExecutingTimeError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "BTCUSDC",
						Side:       "sell",
						Price:      "46",
						LastUpdate: "abcsde",
						Size:       "0.0054",
						TradeSide:  "open",
						FeeDetail: []*bitgetModel.ApiFeeDetail{
							&bitgetModel.ApiFeeDetail{
								FeesValue: "0.00254",
							},
						},
					},
				},
			},
		}, nil)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetFutureFetcherFetchLastTradesWithSideError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "BTCUSDC",
						Side:       "bud",
						Price:      "46",
						LastUpdate: "1745698523",
						Size:       "0.0054",
						TradeSide:  "open",
						FeeDetail: []*bitgetModel.ApiFeeDetail{
							&bitgetModel.ApiFeeDetail{
								FeesValue: "0.00254",
							},
						},
					},
				},
			},
		}, nil)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetFutureFetcherFetchLastTradesWithTradeSideError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "BTCUSDC",
						Side:       "buy",
						Price:      "46",
						LastUpdate: "1745698523",
						Size:       "0.0054",
						TradeSide:  "noSide",
						FeeDetail: []*bitgetModel.ApiFeeDetail{
							&bitgetModel.ApiFeeDetail{
								FeesValue: "0.00254",
							},
						},
					},
				},
			},
		}, nil)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetFutureFetcherFetchLastTradesWitFeesFloatingError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "BTCUSDC",
						Side:       "sell",
						Price:      "46",
						LastUpdate: "123456",
						Size:       "0.0054",
						TradeSide:  "open",
						FeeDetail: []*bitgetModel.ApiFeeDetail{
							&bitgetModel.ApiFeeDetail{
								FeesValue: "something",
							},
						},
					},
				},
			},
		}, nil)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.True(t, errors.As(err, new(*error.BitgetError)))
	assert.Nil(t, trades)
}

func TestBitgetFutureFetcherFetchLastTradesWithOpenTradeSide(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "LINKBTC",
						Side:       "buy",
						Price:      "0.03654",
						LastUpdate: "16549876877",
						Size:       "1234.785",
						TradeSide:  "open",
						FeeDetail: []*bitgetModel.ApiFeeDetail{
							&bitgetModel.ApiFeeDetail{
								FeesValue: "0.0012",
							},
						},
					},
				},
			},
		}, nil)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 1)

	if 1 == len(trades) {
		assert.True(t, trades[0].Open)
	}
}

func TestBitgetFutureFetcherFetchLastTradesWithCloseTradeSide(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "LINKBTC",
						Side:       "buy",
						Price:      "0.03654",
						LastUpdate: "16549876877",
						Size:       "1234.785",
						TradeSide:  "close",
						FeeDetail: []*bitgetModel.ApiFeeDetail{
							&bitgetModel.ApiFeeDetail{
								FeesValue: "0.0012",
							},
						},
					},
				},
			},
		}, nil)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 1)

	if 1 == len(trades) {
		assert.False(t, trades[0].Open)
	}
}

func TestBitgetFutureFetcherFetchLastTradesWithNearlyOpenTradeSide(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "LINKBTC",
						Side:       "buy",
						Price:      "0.03654",
						LastUpdate: "16549876877",
						Size:       "1234.785",
						TradeSide:  "fakeopendata",
						FeeDetail: []*bitgetModel.ApiFeeDetail{
							&bitgetModel.ApiFeeDetail{
								FeesValue: "0.0012",
							},
						},
					},
				},
			},
		}, nil)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 1)

	if 1 == len(trades) {
		assert.True(t, trades[0].Open)
	}
}

func TestBitgetFutureFetcherFetchLastTradesWithBuyBuyTrade(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "LINKBTC",
						Side:       "buy",
						Price:      "0.03654",
						LastUpdate: "16549876877",
						Size:       "1234.785",
						TradeSide:  "ffgdgbuyrga",
						FeeDetail: []*bitgetModel.ApiFeeDetail{
							&bitgetModel.ApiFeeDetail{
								FeesValue: "0.0012",
							},
						},
					},
				},
			},
		}, nil)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 1)

	if 1 == len(trades) {
		assert.True(t, trades[0].Open)
	}
}

func TestBitgetFutureFetcherFetchLastTradesWithBuySellTrade(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "LINKBTC",
						Side:       "sell",
						Price:      "0.03654",
						LastUpdate: "16549876877",
						Size:       "1234.785",
						TradeSide:  "ffgdgbuyrga",
						FeeDetail: []*bitgetModel.ApiFeeDetail{
							&bitgetModel.ApiFeeDetail{
								FeesValue: "0.0012",
							},
						},
					},
				},
			},
		}, nil)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 1)

	if 1 == len(trades) {
		assert.False(t, trades[0].Open)
	}
}

func TestBitgetFutureFetcherFetchLastTradesWithSellSellTrade(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "LINKBTC",
						Side:       "sell",
						Price:      "0.03654",
						LastUpdate: "16549876877",
						Size:       "1234.785",
						TradeSide:  "ffgdsellgbrga",
						FeeDetail: []*bitgetModel.ApiFeeDetail{
							&bitgetModel.ApiFeeDetail{
								FeesValue: "0.0012",
							},
						},
					},
				},
			},
		}, nil)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 1)

	if 1 == len(trades) {
		assert.True(t, trades[0].Open)
	}
}

func TestBitgetFutureFetcherFetchLastTradesWithSellBuyTrade(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "LINKBTC",
						Side:       "buy",
						Price:      "0.03654",
						LastUpdate: "16549876877",
						Size:       "1234.785",
						TradeSide:  "ffgdgrgsella",
						FeeDetail: []*bitgetModel.ApiFeeDetail{
							&bitgetModel.ApiFeeDetail{
								FeesValue: "0.0012",
							},
						},
					},
				},
			},
		}, nil)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 1)

	if 1 == len(trades) {
		assert.False(t, trades[0].Open)
	}
}

func TestBitgetFutureFetcherFetchLastTradesWithNearlyCloseTradeSide(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "LINKBTC",
						Side:       "buy",
						Price:      "0.03654",
						LastUpdate: "16549876877",
						Size:       "1234.785",
						TradeSide:  "someclosetext",
						FeeDetail: []*bitgetModel.ApiFeeDetail{
							&bitgetModel.ApiFeeDetail{
								FeesValue: "0.0012",
							},
						},
					},
				},
			},
		}, nil)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 1)

	if 1 == len(trades) {
		assert.False(t, trades[0].Open)
	}
}

func TestBitgetFutureFetcherFetchLastTradesWithoutError(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverterMock := generatedMocks.NewMockIJsonConverter[bitgetModel.ApiFutureTransactions](mockController)

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
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "LINKBTC",
						Side:       "buy",
						Price:      "0.03654",
						LastUpdate: "16549876877",
						Size:       "1234.785",
						TradeSide:  "sell",
						FeeDetail: []*bitgetModel.ApiFeeDetail{
							&bitgetModel.ApiFeeDetail{
								FeesValue: "0.0012",
							},
						},
					},
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "LINKBTC",
						Side:       "buy",
						Price:      "0.03654",
						LastUpdate: "16549976789",
						Size:       "6547.13",
						TradeSide:  "close",
						FeeDetail: []*bitgetModel.ApiFeeDetail{
							&bitgetModel.ApiFeeDetail{
								FeesValue: "0.0048",
							},
						},
					},
					&bitgetModel.ApiFutureTransaction{
						Symbol:     "LINKBTC",
						Side:       "sell",
						Price:      "0.04012",
						LastUpdate: "16550876654",
						Size:       "5555.55",
						TradeSide:  "ffgdgrgsella",
						FeeDetail: []*bitgetModel.ApiFeeDetail{
							&bitgetModel.ApiFeeDetail{
								FeesValue: "0.0037",
							},
						},
					},
				},
			},
		}, nil)

	fetcher := NewBitgetFutureFetcher(externalGetterMock, jsonConverterMock)

	assert.NotNil(t, fetcher)

	trades, err := fetcher.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 3)

	if 3 == len(trades) {
		assert.Equal(t, "LINKBTC", trades[0].Pair)
		assert.Equal(t, 0.03654, trades[0].Price)
		assert.Equal(t, 1234.785, trades[0].Quantity)
		assert.Equal(t, 0.0012, trades[0].Fees)
		assert.Equal(t, int64(16549876), trades[0].ExecutedTimestamp)
		assert.False(t, trades[0].Open)
		assert.True(t, trades[0].Long)
		assert.Equal(t, "LINKBTC", trades[1].Pair)
		assert.Equal(t, 0.03654, trades[1].Price)
		assert.Equal(t, 6547.13, trades[1].Quantity)
		assert.Equal(t, 0.0048, trades[1].Fees)
		assert.Equal(t, int64(16549976), trades[1].ExecutedTimestamp)
		assert.False(t, trades[1].Open)
		assert.True(t, trades[1].Long)
		assert.Equal(t, "LINKBTC", trades[2].Pair)
		assert.Equal(t, 0.04012, trades[2].Price)
		assert.Equal(t, 5555.55, trades[2].Quantity)
		assert.Equal(t, 0.0037, trades[2].Fees)
		assert.Equal(t, int64(16550876), trades[2].ExecutedTimestamp)
		assert.True(t, trades[2].Open)
		assert.False(t, trades[2].Long)
	}
}
