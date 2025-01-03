package fetcher

import (
	"sort"
	"tradeFetcher/model/trading"
)

type SortByDateFetcherDecorator struct {
	decoratee IFetcher
}

func NewSortByDateFetcherDecorator(decorateeFetcher IFetcher) IFetcher {
	return &SortByDateFetcherDecorator{
		decoratee: decorateeFetcher,
	}
}

func (f SortByDateFetcherDecorator) FetchLastTrades() ([]trading.Trade, error) {
	trades, err := f.decoratee.FetchLastTrades()

	sort.SliceStable(trades, func(i, j int) bool {
		return trades[i].ExecutedTimestamp < trades[j].ExecutedTimestamp
	})

	return trades, err
}
