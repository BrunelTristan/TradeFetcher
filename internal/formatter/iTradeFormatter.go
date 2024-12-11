package formatter

import (
	"tradeFetcher/model/trading"
)

type ITradeFormatter interface {
	Format(trade *trading.Trade) string
}
