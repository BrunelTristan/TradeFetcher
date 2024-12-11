package processUnit

import (
	"fmt"
	"tradeFetcher/model/trading"
)

type TradeDisplayer struct {
}

func NewTradeDisplayer() IProcessUnit {
	return &TradeDisplayer{}
}

func (t TradeDisplayer) ProcessTrades(trades []trading.Trade) {
	for _, trade := range trades {
		fmt.Println(trade)
	}
}
