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

func TestFormat(t *testing.T) {
	formatter := NewCsvTradeFormatter()

	trade := trading.Trade{
		Pair:     "testingToken",
		Price:    1.23,
		Quantity: 98.74,
		Fees:     0.00256,
	}

	assert.NotNil(t, formatter)

	output := formatter.Format(&trade)

	assert.Equal(t, "testingToken;1.23000000;98.74000000;0.00256000", output)
}
