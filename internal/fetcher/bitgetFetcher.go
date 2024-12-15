package fetcher

import (
	"tradeFetcher/model/trading"
)

type BitgetFetcher struct {
	spotFetcher IFetcher
}

func NewBitgetFetcher(sFetcher IFetcher) IFetcher {
	return &BitgetFetcher{
		spotFetcher: sFetcher,
	}
}

func (f BitgetFetcher) FetchLastTrades() ([]trading.Trade, error) {
	trades, err := f.spotFetcher.FetchLastTrades()

	return trades, err
}
