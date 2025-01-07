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

func (f FakeFetcher) FetchLastTrades() ([]*trading.Trade, error) {
	tradeCount := MIN_TRADES_COUNT + rand.IntN(MAX_TRADES_COUNT-MIN_TRADES_COUNT)

	trades := make([]*trading.Trade, tradeCount)

	for index := 0; index < tradeCount; index++ {
		trades[index] = &trading.Trade{
			Pair:              "FakePair",
			Price:             rand.Float64(),
			Quantity:          rand.Float64(),
			Fees:              rand.Float64(),
			ExecutedTimestamp: rand.Int64(),
		}
	}

	return trades, nil
}
