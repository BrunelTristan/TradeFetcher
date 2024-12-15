package formatter

import (
	"fmt"
	"tradeFetcher/model/trading"
)

type CsvTradeFormatter struct {
}

func NewCsvTradeFormatter() ITradeFormatter {
	return &CsvTradeFormatter{}
}

func (t CsvTradeFormatter) Format(trade *trading.Trade) string {
	return fmt.Sprintf(
		"%s;%.8f;%.8f;%.8f",
		trade.Pair,
		trade.Price,
		trade.Quantity,
		trade.Fees)
}
