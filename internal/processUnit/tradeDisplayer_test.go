package processUnit

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"tradeFetcher/internal/generatedMocks"
	"tradeFetcher/model/trading"
)

func TestNewTradeDisplayer(t *testing.T) {
	displayer := NewTradeDisplayer(nil)

	assert.NotNil(t, displayer)
}

func TestProcessTradesDisplayerOnEmptySlice(t *testing.T) {
	mockController := gomock.NewController(t)

	tradeFormatterMock := generatedMocks.NewMockITradeFormatter(mockController)

	tradeFormatterMock.
		EXPECT().
		Format(gomock.Any()).
		Times(0)

	displayer := NewTradeDisplayer(tradeFormatterMock)

	assert.NotNil(t, displayer)

	trades := []*trading.Trade{}

	err := displayer.ProcessTrades(trades)

	assert.Nil(t, err)
}

func TestProcessTradesDisplayerWithValues(t *testing.T) {
	mockController := gomock.NewController(t)

	tradeFormatterMock := generatedMocks.NewMockITradeFormatter(mockController)

	tradeFormatterMock.
		EXPECT().
		Format(gomock.Any()).
		Return("WINNING TRADE").
		Times(2)

	displayer := NewTradeDisplayer(tradeFormatterMock)

	assert.NotNil(t, displayer)

	trades := []*trading.Trade{
		&trading.Trade{},
		&trading.Trade{},
	}

	err := displayer.ProcessTrades(trades)

	assert.Nil(t, err)
}
