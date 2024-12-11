package fetcher

import (
	"tradeFetcher/model/trading"
)

type IFetcher interface {
	FetchLastTrades() []trading.Trade
}
