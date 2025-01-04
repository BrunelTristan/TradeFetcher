package fetcher

import (
	"tradeFetcher/model/trading"
)

type BitgetFetcher struct {
	spotFetcher   IFetcher
	futureFetcher IFetcher
}

func NewBitgetFetcher(
	sFetcher IFetcher,
	fFetcher IFetcher,
) IFetcher {
	return &BitgetFetcher{
		spotFetcher:   sFetcher,
		futureFetcher: fFetcher,
	}
}

func (f BitgetFetcher) FetchLastTrades() ([]trading.Trade, error) {
	trades, err := f.spotFetcher.FetchLastTrades()

	if err != nil {
		return nil, err
	}

	list, err := f.futureFetcher.FetchLastTrades()

	if err != nil {
		return nil, err
	}

	trades = append(trades, list...)

	return trades, err
}
