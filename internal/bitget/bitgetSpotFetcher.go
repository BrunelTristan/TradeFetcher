package bitget

import (
	"tradeFetcher/internal/externalTools"
	"tradeFetcher/internal/fetcher"
	"tradeFetcher/model/trading"
)

type BitgetSpotFetcher struct {
	tradeGetter externalTools.IGetter
}

func NewBitgetSpotFetcher(tGetter externalTools.IGetter) fetcher.IFetcher {
	return &BitgetSpotFetcher{
		tradeGetter: tGetter,
	}
}

func (f BitgetSpotFetcher) FetchLastTrades() ([]trading.Trade, error) {
	_, err := f.tradeGetter.Get(nil)

	if err != nil {
		return nil, err
	}

	trades := make([]trading.Trade, 0)

	return trades, nil
}
