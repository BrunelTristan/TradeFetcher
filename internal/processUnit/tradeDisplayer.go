package processUnit

import (
	"fmt"
	"tradeFetcher/internal/formatter"
	"tradeFetcher/model/trading"
)

type TradeDisplayer struct {
	tradeFormatter formatter.ITradeFormatter
}

func NewTradeDisplayer(tradeFormat formatter.ITradeFormatter) IProcessUnit {
	return &TradeDisplayer{
		tradeFormatter: tradeFormat,
	}
}

func (t TradeDisplayer) ProcessTrades(trades []*trading.Trade) {
	for _, trade := range trades {
		fmt.Println(t.tradeFormatter.Format(trade))
	}
}
