package bitget

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	bitgetModel "tradeFetcher/model/bitget"
	"tradeFetcher/model/error"
	"tradeFetcher/model/trading"
)

func TestNewFutureTransactionToTradeConverter(t *testing.T) {
	fakeObject := NewFutureTransactionToTradeConverter()

	assert.NotNil(t, fakeObject)
}

func TestFutureTransactionsToTradeConverterWithNilInput(t *testing.T) {
	converter := NewFutureTransactionToTradeConverter()

	output, err := converter.Convert(nil)

	assert.Nil(t, err)
	assert.Nil(t, output)
}

func TestFutureTransactionsToTradeConverterWithPriceFloatingError(t *testing.T) {
	converter := NewFutureTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTransaction{
		Symbol:     "BTCUSDC",
		Side:       "sell",
		Price:      "abc",
		LastUpdate: "123456",
		Size:       "0.0054",
	}

	output, err := converter.Convert(input)

	assert.True(t, errors.As(err, new(*error.ConversionError)))
	assert.Nil(t, output)
}

func TestFutureTransactionsToTradeConverterWithQuantityFloatingError(t *testing.T) {
	converter := NewFutureTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTransaction{
		Symbol:     "BTCUSDC",
		Side:       "sell",
		Price:      "46",
		LastUpdate: "123456",
		Size:       "0..0054",
	}

	output, err := converter.Convert(input)

	assert.True(t, errors.As(err, new(*error.ConversionError)))
	assert.Nil(t, output)
}

func TestFutureTransactionsToTradeConverterWithExecTimeError(t *testing.T) {
	converter := NewFutureTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTransaction{
		Symbol:     "BTCUSDC",
		Side:       "sell",
		Price:      "46",
		LastUpdate: "abcsde",
		Size:       "0.0054",
		FeeDetail: []*bitgetModel.ApiFeeDetail{
			&bitgetModel.ApiFeeDetail{
				FeesValue: "0.00254",
			},
		},
	}

	output, err := converter.Convert(input)

	assert.True(t, errors.As(err, new(*error.ConversionError)))
	assert.Nil(t, output)
}

func TestFutureTransactionsToTradeConverterWithSideError(t *testing.T) {
	converter := NewFutureTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTransaction{
		Symbol:     "BTCUSDC",
		Side:       "bud",
		Price:      "46",
		LastUpdate: "1745698523",
		Size:       "0.0054",
		FeeDetail: []*bitgetModel.ApiFeeDetail{
			&bitgetModel.ApiFeeDetail{
				FeesValue: "0.00254",
			},
		},
	}

	output, err := converter.Convert(input)

	assert.True(t, errors.As(err, new(*error.ConversionError)))
	assert.Nil(t, output)
}

func TestFutureTransactionsToTradeConverterWithTradeSideError(t *testing.T) {
	converter := NewFutureTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTransaction{
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
	}

	output, err := converter.Convert(input)

	assert.True(t, errors.As(err, new(*error.ConversionError)))
	assert.Nil(t, output)
}

func TestFutureTransactionsToTradeConverterWithFeesFloatingError(t *testing.T) {
	converter := NewFutureTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTransaction{
		Symbol:     "BTCUSDC",
		Side:       "sell",
		Price:      "46",
		LastUpdate: "123456",
		Size:       "0.0054",
		FeeDetail: []*bitgetModel.ApiFeeDetail{
			&bitgetModel.ApiFeeDetail{
				FeesValue: "hundred",
			},
		},
	}

	output, err := converter.Convert(input)

	assert.True(t, errors.As(err, new(*error.ConversionError)))
	assert.Nil(t, output)
}

func TestFutureTransactionsToTradeConverterBuyOrder(t *testing.T) {
	converter := NewFutureTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTransaction{
		Symbol:     "LINKBTC",
		Side:       "buy",
		Price:      "0.03654",
		LastUpdate: "16549876877",
		Size:       "1234.785",
		TradeSide:  "buy",
		FeeDetail: []*bitgetModel.ApiFeeDetail{
			&bitgetModel.ApiFeeDetail{
				FeesValue: "-0.0012",
			},
		},
	}

	output, err := converter.Convert(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.Equal(t, "LINKBTC", output.Pair)
	assert.Equal(t, 0.03654, output.Price)
	assert.Equal(t, 1234.785, output.Quantity)
	assert.Equal(t, 0.0012, output.Fees)
	assert.Equal(t, int64(16549876), output.ExecutedTimestamp)
	assert.Equal(t, trading.OPENING, output.TransactionType)
	assert.True(t, output.Long)
}

func TestFutureTransactionsToTradeConverterSellOrder(t *testing.T) {
	converter := NewFutureTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTransaction{
		Symbol:     "ADAUSDC",
		Side:       "sell",
		TradeSide:  "sell",
		Price:      "0.4565",
		LastUpdate: "4565987546",
		Size:       "47.348",
		FeeDetail: []*bitgetModel.ApiFeeDetail{
			&bitgetModel.ApiFeeDetail{
				FeesValue: "-0.000489",
			},
		},
	}

	output, err := converter.Convert(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.Equal(t, "ADAUSDC", output.Pair)
	assert.Equal(t, 0.4565, output.Price)
	assert.Equal(t, 47.348, output.Quantity)
	assert.Equal(t, 0.000489, output.Fees)
	assert.Equal(t, int64(4565987), output.ExecutedTimestamp)
	assert.Equal(t, trading.OPENING, output.TransactionType)
	assert.False(t, output.Long)
}

func TestFutureTransactionsToTradeConverterOpenTradeSide(t *testing.T) {
	converter := NewFutureTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTransaction{
		Symbol:     "LINKBTC",
		Side:       "buy",
		Price:      "0.03654",
		LastUpdate: "16549876877",
		Size:       "1234.785",
		TradeSide:  "open",
		FeeDetail: []*bitgetModel.ApiFeeDetail{
			&bitgetModel.ApiFeeDetail{
				FeesValue: "-0.0012",
			},
		},
	}

	output, err := converter.Convert(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.Equal(t, trading.OPENING, output.TransactionType)
	assert.True(t, output.Long)
}

func TestFutureTransactionsToTradeConverterCloseTradeSide(t *testing.T) {
	converter := NewFutureTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTransaction{
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
	}

	output, err := converter.Convert(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.Equal(t, trading.CLOSE, output.TransactionType)
	assert.True(t, output.Long)
}

func TestFutureTransactionsToTradeConverterNearlyOpenTradeSide(t *testing.T) {
	converter := NewFutureTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTransaction{
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
	}

	output, err := converter.Convert(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.Equal(t, trading.OPENING, output.TransactionType)
	assert.True(t, output.Long)
}

func TestFutureTransactionsToTradeConverterNearlyCloseTradeSide(t *testing.T) {
	converter := NewFutureTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTransaction{
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
	}

	output, err := converter.Convert(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.Equal(t, trading.CLOSE, output.TransactionType)
	assert.True(t, output.Long)
}

func TestFutureTransactionsToTradeConverterBuyShortOrder(t *testing.T) {
	converter := NewFutureTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTransaction{
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
	}

	output, err := converter.Convert(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.Equal(t, trading.CLOSE, output.TransactionType)
	assert.False(t, output.Long)
}

func TestFutureTransactionsToTradeConverterSellLongOrder(t *testing.T) {
	converter := NewFutureTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTransaction{
		Symbol:     "LINKBTC",
		Side:       "buy",
		Price:      "0.03654",
		LastUpdate: "16549876877",
		Size:       "1234.785",
		TradeSide:  "ffgdgrgsella",
		FeeDetail: []*bitgetModel.ApiFeeDetail{
			&bitgetModel.ApiFeeDetail{
				FeesValue: "-0.0012",
			},
		},
	}

	output, err := converter.Convert(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.Equal(t, "LINKBTC", output.Pair)
	assert.Equal(t, 0.03654, output.Price)
	assert.Equal(t, 1234.785, output.Quantity)
	assert.Equal(t, 0.0012, output.Fees)
	assert.Equal(t, int64(16549876), output.ExecutedTimestamp)
	assert.Equal(t, trading.CLOSE, output.TransactionType)
	assert.True(t, output.Long)
}
