package processUnit

import (
	"tradeFetcher/model/trading"
)

type IProcessUnit interface {
	ProcessTrades(trades []*trading.Trade) error
}
