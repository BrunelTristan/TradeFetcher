package bitget

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	bitgetModel "tradeFetcher/model/bitget"
	"tradeFetcher/model/error"
	"tradeFetcher/model/trading"
)

func TestNewSpotFillToTradeConverter(t *testing.T) {
	fakeObject := NewSpotFillToTradeConverter()

	assert.NotNil(t, fakeObject)
}

func TestSpotFillToTradeConverterWithNilInput(t *testing.T) {
	converter := NewSpotFillToTradeConverter()

	output, err := converter.Convert(nil)

	assert.Nil(t, err)
	assert.Nil(t, output)
}

func TestSpotFillToTradeConverterWithPriceFloatingError(t *testing.T) {
	converter := NewSpotFillToTradeConverter()

	input := &bitgetModel.ApiSpotFill{
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

func TestSpotFillToTradeConverterWithQuantityFloatingError(t *testing.T) {
	converter := NewSpotFillToTradeConverter()

	input := &bitgetModel.ApiSpotFill{
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

func TestSpotFillToTradeConverterWithExecTimeError(t *testing.T) {
	converter := NewSpotFillToTradeConverter()

	input := &bitgetModel.ApiSpotFill{
		Symbol:     "BTCUSDC",
		Side:       "sell",
		Price:      "46",
		LastUpdate: "abcsde",
		Size:       "0.0054",
		FeeDetail: &bitgetModel.ApiFeeDetail{
			FeesValue: "0.00254",
		},
	}

	output, err := converter.Convert(input)

	assert.True(t, errors.As(err, new(*error.ConversionError)))
	assert.Nil(t, output)
}

func TestSpotFillToTradeConverterWithSideError(t *testing.T) {
	converter := NewSpotFillToTradeConverter()

	input := &bitgetModel.ApiSpotFill{
		Symbol:     "BTCUSDC",
		Side:       "bud",
		Price:      "46",
		LastUpdate: "1745698523",
		Size:       "0.0054",
		FeeDetail: &bitgetModel.ApiFeeDetail{
			FeesValue: "0.00254",
		},
	}

	output, err := converter.Convert(input)

	assert.True(t, errors.As(err, new(*error.ConversionError)))
	assert.Nil(t, output)
}

func TestSpotFillToTradeConverterWithFeesFloatingError(t *testing.T) {
	converter := NewSpotFillToTradeConverter()

	input := &bitgetModel.ApiSpotFill{
		Symbol:     "BTCUSDC",
		Side:       "sell",
		Price:      "46",
		LastUpdate: "123456",
		Size:       "0.0054",
		FeeDetail: &bitgetModel.ApiFeeDetail{
			FeesValue: "hundred",
		},
	}

	output, err := converter.Convert(input)

	assert.True(t, errors.As(err, new(*error.ConversionError)))
	assert.Nil(t, output)
}

func TestSpotFillToTradeConverterBuyOrder(t *testing.T) {
	converter := NewSpotFillToTradeConverter()

	input := &bitgetModel.ApiSpotFill{
		Symbol:     "LINKBTC",
		Side:       "buy",
		Price:      "0.03654",
		LastUpdate: "16549876877",
		Size:       "1234.785",
		FeeDetail: &bitgetModel.ApiFeeDetail{
			FeesValue: "0.0012",
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

func TestSpotFillToTradeConverterSellOrder(t *testing.T) {
	converter := NewSpotFillToTradeConverter()

	input := &bitgetModel.ApiSpotFill{
		Symbol:     "ADAUSDC",
		Side:       "sell",
		Price:      "0.4565",
		LastUpdate: "4565987546",
		Size:       "47.348",
		FeeDetail: &bitgetModel.ApiFeeDetail{
			FeesValue: "0.000489",
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
	assert.Equal(t, trading.CLOSE, output.TransactionType)
	assert.True(t, output.Long)
}
