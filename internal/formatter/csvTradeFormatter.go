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
	longShort := "S"

	if trade.Open {
		openClose = "O"
	}

	if trade.Long {
		longShort = "L"
	}

	return fmt.Sprintf(
		"%d,%s,%s,%s;%.8f;%.8f;%.8f",
		trade.ExecutedTimestamp,
		openClose,
		longShort,
		trade.Pair,
		trade.Price,
		trade.Quantity,
		trade.Fees)
}
