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
		Open:              true,
		Long:              false,
		Fees:              0.00256,
	}

	assert.NotNil(t, formatter)

	output := formatter.Format(&trade)

	assert.Equal(t, "172345687,O,S,testingToken;1.23000000;98.74000000;0.00256000", output)
}

func TestFormatCloseOrder(t *testing.T) {
	formatter := NewCsvTradeFormatter()

	trade := trading.Trade{
		Pair:              "testingToken",
		ExecutedTimestamp: 172345687,
		Price:             1.23,
		Quantity:          98.74,
		Open:              false,
		Long:              true,
		Fees:              0.00256,
	}

	assert.NotNil(t, formatter)

	output := formatter.Format(&trade)

	assert.Equal(t, "172345687,C,L,testingToken;1.23000000;98.74000000;0.00256000", output)
}
