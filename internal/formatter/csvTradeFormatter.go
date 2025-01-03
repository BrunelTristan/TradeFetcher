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
	openClose := "C"

	if trade.Open {
		openClose = "O"
	}

	return fmt.Sprintf(
		"%d,%s,%s;%.8f;%.8f;%.8f",
		trade.ExecutedTimestamp,
		openClose,
		trade.Pair,
		trade.Price,
		trade.Quantity,
		trade.Fees)
}
