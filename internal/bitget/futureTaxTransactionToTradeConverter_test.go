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

func TestFutureTaxTransactionsToTradeConverterWithoutFunding(t *testing.T) {
	converter := NewFutureTaxTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTaxTransaction{
		Symbol:    "LINKBTC",
		Timestamp: "16549876877",
		Amount:    "-0.0012",
		Fee:       "0",
		TaxType:   "open_short",
	}

	output, err := converter.Convert(input)

	assert.Nil(t, err)
	assert.Nil(t, output)
}

func TestFutureTaxTransactionsToTradeConverterWithExecTimeError(t *testing.T) {
	converter := NewFutureTaxTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTaxTransaction{
		Symbol:    "BTCUSDC",
		Timestamp: "abcsde",
		Amount:    "-0.00459712",
		Fee:       "0",
		TaxType:   "contract_main_settle_fee",
	}

	output, err := converter.Convert(input)

	assert.True(t, errors.As(err, new(*error.ConversionError)))
	assert.Nil(t, output)
}

func TestFutureTaxTransactionsToTradeConverterWithAmountTaxFloatingError(t *testing.T) {
	converter := NewFutureTaxTransactionToTradeConverter()

	input := &bitgetModel.ApiFutureTaxTransaction{
		Symbol:    "BTCUSDC",
		Timestamp: "123456",
		Amount:    "number",
		Fee:       "0.000456",
		TaxType:   "contract_main_settle_fee",
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
		Amount:    "-0.0012",
		Fee:       "0",
		TaxType:   "contract_main_settle_fee",
	}

	output, err := converter.Convert(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.Equal(t, "LINKBTC", output.Pair)
	assert.Equal(t, 0.0012, output.Fees)
	assert.Equal(t, int64(16549876), output.ExecutedTimestamp)
	assert.Equal(t, trading.FUNDING, output.TransactionType)
}
