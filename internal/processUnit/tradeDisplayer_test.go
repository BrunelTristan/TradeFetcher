package processUnit

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tradeFetcher/model/trading"
)

func TestNewTradeDisplayer(t *testing.T) {
	displayer := NewTradeDisplayer()

	assert.NotNil(t, displayer)
}

func TestProcessTradesOnEmptySlice(t *testing.T) {
	displayer := NewTradeDisplayer()

	assert.NotNil(t, displayer)

	trades := []trading.Trade{}

	displayer.ProcessTrades(trades)
}


func TestProcessTradesWithValues(t *testing.T) {
	displayer := NewTradeDisplayer()

	assert.NotNil(t, displayer)

	trades := []trading.Trade{
		trading.Trade{},
		trading.Trade{},
	}

	displayer.ProcessTrades(trades)
}
