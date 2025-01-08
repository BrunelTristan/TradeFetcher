package bitget

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	bitgetModel "tradeFetcher/model/bitget"
	"tradeFetcher/model/error"
	"tradeFetcher/model/trading"
)

func TestNewFutureTaxTransactionToTradeConverter(t *testing.T) {
	fakeObject := NewFutureTaxTransactionToTradeConverter()

	assert.NotNil(t, fakeObject)
}

func TestFutureTaxTransactionsToTradeConverterWithNilInput(t *testing.T) {
	converter := NewFutureTaxTransactionToTradeConverter()

	output, err := converter.Convert(nil)

	assert.Nil(t, err)
	assert.Nil(t, output)
}

func TestFutureTaxTransactionsToTradeConverterWithExecTimeError(t *testing.T) {
	converter := NewFutureTaxTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTaxTransaction{
		Symbol:    "BTCUSDC",
		Timestamp: "abcsde",
		Fee:       "0.00254",
	}

	output, err := converter.Convert(input)

	assert.True(t, errors.As(err, new(*error.ConversionError)))
	assert.Nil(t, output)
}

func TestFutureTaxTransactionsToTradeConverterWithFeesFloatingError(t *testing.T) {
	converter := NewFutureTaxTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTaxTransaction{
		Symbol:    "BTCUSDC",
		Timestamp: "123456",
		Fee:       "hundred",
	}

	output, err := converter.Convert(input)

	assert.True(t, errors.As(err, new(*error.ConversionError)))
	assert.Nil(t, output)
}

func TestFutureTaxTransactionsToTradeConverterWithoutError(t *testing.T) {
	converter := NewFutureTaxTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTaxTransaction{
		Symbol:    "LINKBTC",
		Timestamp: "16549876877",
		Fee:       "-0.0012",
	}

	output, err := converter.Convert(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.Equal(t, "LINKBTC", output.Pair)
	assert.Equal(t, 0.0012, output.Fees)
	assert.Equal(t, int64(16549876), output.ExecutedTimestamp)
	assert.Equal(t, trading.FUNDING, output.TransactionType)
}
