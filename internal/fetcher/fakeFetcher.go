package fetcher

import (
	"math/rand/v2"
	"tradeFetcher/model/trading"
)

const MIN_TRADES_COUNT int = 1
const MAX_TRADES_COUNT int = 5

type FakeFetcher struct {
}

func NewFakeFetcher() IFetcher {
	return &FakeFetcher{}
}

func (f FakeFetcher) FetchLastTrades() []trading.Trade {
	tradeCount := MIN_TRADES_COUNT + rand.IntN(MAX_TRADES_COUNT-MIN_TRADES_COUNT)

	trades := make([]trading.Trade, tradeCount)

	for index := 0; index < tradeCount; index++ {
		trades[index].Pair = "FakePair"
		trades[index].Price = rand.Float64()
		trades[index].Quantity = rand.Float64()
		trades[index].Fees = rand.Float64()
	}

	return trades
}
