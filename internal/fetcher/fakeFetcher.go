package fetcher

import (
	"tradeFetcher/model/trading"
)

type FakeFetcher struct {
}

func NewFakeFetcher() IFetcher {
	return &FakeFetcher{}
}

func (f FakeFetcher) FetchLastTrades() []trading.Trade {
	trades := make([]trading.Trade, 0)

	return trades
}
