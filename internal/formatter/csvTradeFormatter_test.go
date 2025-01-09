package formatter

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tradeFetcher/model/trading"
)

func TestNewCsvTradeFormatter(t *testing.T) {
	formatter := NewCsvTradeFormatter()

	assert.NotNil(t, formatter)
}

func TestFormatOpenOrder(t *testing.T) {
	formatter := NewCsvTradeFormatter()

	trade := trading.Trade{
		Pair:              "testingToken",
		ExecutedTimestamp: 172345687,
		Price:             1.23,
		Quantity:          98.74,
		TransactionType:   trading.OPENING,
		Long:              false,
		Fees:              0.00256,
	}

	assert.NotNil(t, formatter)

	output := formatter.Format(&trade)

	assert.Equal(t, "172345687;O;S;testingToken;98.74000000;1.23000000;0.00256000", output)
}

func TestFormatCloseOrder(t *testing.T) {
	formatter := NewCsvTradeFormatter()

	trade := trading.Trade{
		Pair:              "testingToken",
		ExecutedTimestamp: 172345687,
		Price:             1.23,
		Quantity:          98.74,
		TransactionType:   trading.CLOSE,
		Long:              true,
		Fees:              0.00256,
	}

	assert.NotNil(t, formatter)

	output := formatter.Format(&trade)

	assert.Equal(t, "172345687;C;L;testingToken;98.74000000;1.23000000;0.00256000", output)
}

func TestFormatFundingTransaction(t *testing.T) {
	formatter := NewCsvTradeFormatter()

	trade := trading.Trade{
		Pair:              "testingToken",
		ExecutedTimestamp: 172345687,
		TransactionType:   trading.FUNDING,
		Fees:              -0.00256,
	}

	assert.NotNil(t, formatter)

	output := formatter.Format(&trade)

	assert.Equal(t, "172345687;F;;testingToken;0.00000000;0.00000000;-0.00256000", output)
}
