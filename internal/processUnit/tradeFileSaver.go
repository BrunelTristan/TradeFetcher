package processUnit

import (
	"os"
	"path/filepath"
	"tradeFetcher/internal/formatter"
	"tradeFetcher/model/trading"
)

type TradeFileSaver struct {
	tradeFormatter formatter.ITradeFormatter
	filePath       string
}

func NewTradeFileSaver(tradeFormat formatter.ITradeFormatter, fPath string) IProcessUnit {
	return &TradeFileSaver{
		tradeFormatter: tradeFormat,
		filePath:       fPath,
	}
}

func (s TradeFileSaver) ProcessTrades(trades []*trading.Trade) error {
	if len(trades) == 0 {
		return nil
	}

	_ = os.MkdirAll(filepath.Dir(s.filePath), 0770)

	file, err := os.OpenFile(s.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, trade := range trades {
		_, _ = file.Write([]byte(s.tradeFormatter.Format(trade)))
		_, _ = file.Write([]byte("\n"))
	}

	return nil
}
