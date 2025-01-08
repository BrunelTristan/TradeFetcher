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

	if trade.Long {
		longShort = "L"
	}

	if trade.TransactionType == trading.OPENING {
		openClose = "O"
	} else if trade.TransactionType == trading.FUNDING {
		openClose = "F"
		longShort = ""
	}

	return fmt.Sprintf(
		"%d;%s;%s;%s;%.8f;%.8f;%.8f",
		trade.ExecutedTimestamp,
		openClose,
		longShort,
		trade.Pair,
		trade.Quantity,
		trade.Price,
		trade.Fees)
}
