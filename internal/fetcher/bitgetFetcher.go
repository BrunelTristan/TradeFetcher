package fetcher

import (
	"tradeFetcher/model/trading"
)

type BitgetFetcher struct {
	fetchers []IFetcher
}

func NewBitgetFetcher(
	fetchs []IFetcher,
) IFetcher {
	return &BitgetFetcher{
		fetchers: fetchs,
	}
}

func (f BitgetFetcher) FetchLastTrades() ([]trading.Trade, error) {
	trades := make([]trading.Trade, 0)

	for _, fetcher := range f.fetchers {
		list, err := fetcher.FetchLastTrades()

		if err != nil {
			return nil, err
		}

		trades = append(trades, list...)
	}

	return trades, nil
}
